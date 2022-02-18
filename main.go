package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

//type loginReq struct {
//	username string `json:"username"`
//	password string `json:"password"`
//}

func main() {
	flag.Parse()

	wsServer := NewWebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		//var req loginReq

		if r.Method == "GET" {
			http.ServeFile(w, r, "./public/login.html")
			return
		}
		var loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&loginReq)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// Read body
		//b, err := ioutil.ReadAll(r.Body)
		//defer r.Body.Close()
		//if err != nil {
		//	http.Error(w, err.Error(), 500)
		//	return
		//}
		//
		//// Unmarshal
		//var req loginReq
		//err = json.Unmarshal(b, &req)
		//if err != nil {
		//	http.Error(w, err.Error(), 500)
		//	return
		//}
		if loginReq.Username == "test" && loginReq.Password == "admin" {
			//cookie := &http.Cookie{
			//	Name:   "username",
			//	Value:  "test",
			//	MaxAge: 300,
			//}
			w.WriteHeader(200)
			write, err := w.Write([]byte("oke"))
			if err != nil {
				return
			}
			fmt.Printf("%v", write)
		} else {
			w.WriteHeader(400)
			w.Write([]byte("fail"))
		}
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cookies := request.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "name" {
				http.FileServer(http.Dir("./public")).ServeHTTP(writer, request)
				return
			}
		}
		writer.Header().Set("Location", "/login")
		writer.WriteHeader(http.StatusSeeOther)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
