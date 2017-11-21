package chatroom

"github.com/gorilla/websocket"

type ChatRoom struct{
  Clients [*websocket.Conn]bool
  Broadcast chan
  Upgrader websocket.Upgrader
}
