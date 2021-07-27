package user

import "golang.org/x/net/context"

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	return &Message{Body : "Hello"}, nil
}

func (s *Server) GetUser(ctx context.Context, useridx *UserIdx) (*User, error) {
	u := User{}
	u.Idx = 1
	u.Name = "김현지"
	u.UserId = "abasdasdaa"
	u.Password = "dslkadlkldskldsl"

	return &u, nil 
}