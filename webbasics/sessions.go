package webbasics

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

/* Session cookies: allows the server to re-identify the user

To maintain state between server and client

here: get session cooking using /login ,to access /secret
can revoke using /logout
*/

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key) // this is the store which has the key
)

/*
Returns a session cookie that can be used by the browser for reconnect

Browser will automatically store the cookie and sent it again in subsequent requests
*/
func login(w http.ResponseWriter, r *http.Request) {
	// Return session with the given name, creates new if doesn't exist
	session, _ := store.Get(r, "cookie-name") // This cookie store was created above
	fmt.Printf("\n Session values: %#v \n", session.IsNew)

	// Add to server that this session is authenticated
	session.Values["authenticated"] = true

	// Save the session state before returning the response
	session.Save(r, w)

}

func logout(w http.ResponseWriter, r *http.Request) {
	// Return session with the given name, creates new if doesn't exist
	// here, the cookie would generally exist
	session, _ := store.Get(r, "cookie-name")

	fmt.Printf("\n Session values: %#v \n", session.IsNew)

	// Add to server that this session is authenticated
	session.Values["authenticated"] = false

	// Save the session state before returning the response
	session.Save(r, w)

	// Expire cookie/ or send an empty cookie
	w.Header().Add("Set-Cookie", "cookie-name= ; Path=/; Expires=Tue, 05 Sep 2023 18:28:01 GMT; Max-Age=0")

}

/*
$ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
*/
func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	fmt.Println("Cookie", r.Header.Get("Cookie"))

	fmt.Printf("\n Session values: %#v \n", session.IsNew)

	// Check if authenticated, return false if not
	if auth, ok := session.Values["authenticated"].(bool); !(auth) || !(ok) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// Print secret message
	fmt.Fprintln(w, "The cake is a lie!")
}

func SessionImpl() {

	mtx := http.NewServeMux()

	mtx.HandleFunc("/login", login)
	mtx.HandleFunc("/logout", logout)
	mtx.HandleFunc("/secret", secret)

	http.ListenAndServe(":8081", mtx)

}
