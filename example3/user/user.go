package user

import (
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

type Server struct {
	DB     *gorm.DB
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	return &Message{Body : "Hello"}, nil
}

func (s *Server) GetUser(ctx context.Context, useridx *UserIdx) (*User, error) {
	
	//u := User{}
	u := User{}
	
	rows, err := u.GetUser(s.DB, int(useridx.Idx))

	if err != nil {
		log.Fatal(err)
	}
	 return rows, nil
}

func (u *User) GetUser(db *gorm.DB, useridx int) (*User, error) {

	err := db.Debug().Model(User{}).Where("idx = ?", useridx).Find(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}