package db

import (
	"RedNews_Server/internal/domain/entity"
)

type Database struct {
	User     string
	Password string
	Host     string
	Database string
}

type UserInterface interface {
	GetUser(int32) entity.User

	GetID(string) int32
	GetName(int32) string
	GetEmail(int32) string
}
