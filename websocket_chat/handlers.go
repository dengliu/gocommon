package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	TYPE_WS_JOIN              = "TYPE_WS_JOIN"
	TYPE_WS_LEAVE             = "TYPE_WS_LEAVE"
	TYPE_WS_MSG               = "TYPE_WS_MSG"
	TYPE_WS_CHECK_ONLINE_JOIN = "TYPE_WS_CHECK_ONLINE_JOIN"
)

type Message struct {
	Type         string `json:"type"`
	Username     string `json:"username"`
	Message      string `json:"message"`
	AmountOnline int    `json:"amount_online"`
	AmountJoin   int    `json:"amount_join"`
}

type Application struct {
	Upgrader   websocket.Upgrader
	Clients    map[*websocket.Conn]string
	Broadcast  chan Message
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	FileSystem http.FileSystem
}

func NewApplicatoin(infLog, errLog *log.Logger, filedir string) *Application {
	// Makes sure you put folder name correctly and exists
	fsStatic, err := fs.Sub(staticfiles, filedir)
	if err != nil {
		errLog.Fatal(err)
	}

	return &Application{
		Upgrader:   websocket.Upgrader{},
		Clients:    make(map[*websocket.Conn]string),
		Broadcast:  make(chan Message),
		InfoLog:    infLog,
		ErrorLog:   errLog,
		FileSystem: http.FS(fsStatic),
	}

}

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	// Websockets
	mux.HandleFunc("/ws", app.HandleConnections)

	// Check Username
	mux.HandleFunc("/check-username", app.HandleCheckUsername)

	// Static Files
	mux.Handle("/", app.HandleStaticFiles())

	return mux
}

func (app *Application) Run() {
	for {
		msg := <-app.Broadcast
		// app.InfoLog.Printf("msg: %+v\n", msg)
		for client := range app.Clients {
			if err := client.WriteJSON(msg); err != nil {
				msg := app.LeaveChat(client)
				go func() {
					app.Broadcast <- msg
				}()
			}
		}
	}
}

func (app *Application) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := app.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
	defer ws.Close()

	app.Clients[ws] = ""
	msg := app.CheckOnlineAndJoin()
	app.Broadcast <- msg

	for {
		msg := Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			msg = app.LeaveChat(ws)
			app.Broadcast <- msg
			break
		}
		switch msg.Type {
		case TYPE_WS_JOIN:
			app.JoinChat(ws, &msg)
			app.Broadcast <- msg
		case TYPE_WS_MSG:
			app.Broadcast <- msg
		}
	}
}

func (app *Application) HandleStaticFiles() http.Handler {
	return http.FileServer(app.FileSystem)
}

func (app *Application) AmountOnlineAndJoin() (int, int) {
	join := 0
	for _, username := range app.Clients {
		if username != "" {
			join++
		}
	}
	return len(app.Clients), join
}

func (app *Application) JoinChat(ws *websocket.Conn, msg *Message) {
	app.Clients[ws] = msg.Username
	cOn, cJoin := app.AmountOnlineAndJoin()
	msg.AmountOnline = cOn
	msg.AmountJoin = cJoin
	app.InfoLog.Output(2, fmt.Sprintf("%s join chat\n", app.Clients[ws]))
}

func (app *Application) LeaveChat(ws *websocket.Conn) Message {
	msg := Message{}
	username := app.Clients[ws]
	delete(app.Clients, ws)

	defer func() {
		app.InfoLog.Output(3, fmt.Sprintf("%s leave chat\n", username))
		ws.Close()
	}()

	cOn, cJoin := app.AmountOnlineAndJoin()
	if username != "" {
		msg = Message{
			Type:         TYPE_WS_LEAVE,
			Username:     username,
			AmountOnline: cOn,
			AmountJoin:   cJoin,
		}
	} else {
		username = ws.RemoteAddr().String()
		msg = Message{
			Type:         TYPE_WS_CHECK_ONLINE_JOIN,
			AmountOnline: cOn,
			AmountJoin:   cJoin,
		}
	}
	return msg
}

func (app *Application) CheckOnlineAndJoin() Message {
	cOn, cJoin := app.AmountOnlineAndJoin()
	msg := Message{
		Type:         TYPE_WS_CHECK_ONLINE_JOIN,
		AmountOnline: cOn,
		AmountJoin:   cJoin,
	}
	return msg
}

func (app *Application) IsUsernameExists(username string) bool {
	for _, uname := range app.Clients {
		if uname == username {
			return true
		}
	}
	return false
}

func (app *Application) HandleCheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		m := map[string]interface{}{
			"status_code": http.StatusBadRequest,
			"username":    username,
			"is_allow":    false,
			"reason":      "Bad Request",
		}
		app.SendJSON(w, http.StatusOK, m)
		return
	}

	isExists := app.IsUsernameExists(username)
	if isExists {
		m := map[string]interface{}{
			"status_code": http.StatusOK,
			"username":    username,
			"is_allow":    false,
			"reason":      "Username already in use",
		}
		app.SendJSON(w, http.StatusOK, m)
		return
	}

	m := map[string]interface{}{
		"status_code": http.StatusOK,
		"username":    username,
		"is_allow":    true,
		"reason":      "",
	}
	app.SendJSON(w, http.StatusOK, m)
}

func (app *Application) SendJSON(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	vJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(vJson)
}
