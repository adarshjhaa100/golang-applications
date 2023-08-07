package webbasics

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"nhooyr.io/websocket"
)

// For cross origin (Note: In prod, use a list instead of *)
var (
	wsAcceptOptions = &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	}
)

func readFromWS(conn *websocket.Conn, ctx *context.Context, buff *[]byte) (int, *websocket.MessageType) {

	typ, rdr, err := conn.Reader(*ctx) // reads from a connection (using child context)
	check(err)

	fmt.Printf("\nMessage type: %v\n", typ)

	n, err := rdr.Read(*buff)
	check(err)

	fmt.Printf("\n%v bytes read, data: %v\n", n, string(*buff))

	if strings.Contains(string(*buff), "close") {
		conn.Close(websocket.StatusNormalClosure, "")
	}

	// Clean the buffer post reading messages
	*buff = make([]byte, 4096)

	return n, &typ

}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Accept handshake from a client and returns connection object
	conn, err := websocket.Accept(w, r, wsAcceptOptions)
	check(err)

	// Close connection with error (In case something goes bad before end)
	defer conn.Close(websocket.StatusInternalError,
		"WS connection closed, error")

	//  Takes a context (here request context) and returns a child context
	//  A context in go is used to pass data
	//  Why did we have to create a child context here?
	ctx, cancel := context.WithTimeout(r.Context(), time.Minute*10)
	defer cancel() // release resources associated with the ctx

	// var res interface{}

	/*---------------Read messages---------------*/
	// // Receive JSON messages
	// for {
	// 	var v interface{}
	// 	err = wsjson.Read(ctx, conn, &v)
	// 	check(err)
	// 	fmt.Printf("\nRead message: %#v\n", v)
	// }

	// Reader or close read?
	// Keep reading until "close message"
	var buff = make([]byte, 4096)

	for {
		n, typ := readFromWS(conn, &ctx, &buff)

		/*---------------Write back messages---------------*/
		resp := fmt.Sprintf("\nYour message has %#v bytes\n", n)
		conn.Write(ctx, *typ, []byte(resp))

	}

}

func WSImpl() {

	mux := http.NewServeMux()

	mux.HandleFunc("/websocket", wsHandler)

	http.ListenAndServe(":8081", mux)
}
