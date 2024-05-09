package main

import (
	"fmt"
	"io"

	//"github.com/gorilla/websocket"
	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func newServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}
func (s *Server) reedLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error : ", err)
			continue
		}

		msg := buf[:n]
		fmt.Println(string(msg))

		ws.Write([]byte("Message received successfully :) "))
	}
}
func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New Incoming Connection from client : ", ws.RemoteAddr())

	s.conns[ws] = true

	s.reedLoop(ws)

}
func main() {

}
