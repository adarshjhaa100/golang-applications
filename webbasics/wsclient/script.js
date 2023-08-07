console.log("Hello From WS Client!")


const form = document.getElementById("form");
const msg = document.getElementById("msg");


const conn = new WebSocket("ws://localhost:8081/websocket")

conn.addEventListener("error", (e) => {
    console.error(e)
})

conn.addEventListener("close", (e) => {
    console.log("Connection CLosed")
})

conn.addEventListener("message", (e) => {
    console.log("Message")
    console.log(e.data) // data from connection
})

conn.addEventListener("open", (e) => {
    console.log("Open")
    // conn.send("Hello from client!!")
})


form.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("Message: ", msg.value);
    conn.send(msg.value)
    // conn.send(JSON.stringify({
    //     "name": "Hello",
    //     "age": 12
    // }))
})