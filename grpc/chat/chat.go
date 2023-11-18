package chat

import (
	context "context"
	"log"
)

type Server struct {

}

func (s *Server) SayHello (ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}