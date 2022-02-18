package main

import (
	json2 "encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"

type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Target  Target `json:"target"`
	Sender  Client `json:"sender"`
}

type Target struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
	rooms    map[*Room]bool
	Name     string `json:"name"` // dùng để đặt tên cho client
}

//type Cookie struct {
//	Name     string
//	Value    string
//	Path     string
//	Domain   string
//	Expires  time.Time
//	MaxAge   int
//	Secure   bool
//	HttpOnly bool
//	Raw      string
//	UnParse  []string
//}

// {"action": "asdasd", "message": "asdasd",  "sender": {}}

//phương thức mã hóa có thể được gọi để tạo một đối tượng byte json [] đã sẵn sàng để gửi lại cho các máy khách
func (message *Message) encode() []byte {
	json, err := json2.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return json
}
