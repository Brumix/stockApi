package models

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"stockApi/services/core"
)

type Hello struct {
	Msg string `json:"msg"`
}

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}

type Grpc struct {
	Connection *grpc.ClientConn
	Server     core.AuthServiceClient
}

type UserRequest struct {
	Email string `json:"email"`
}
