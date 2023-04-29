package mysql

import (
	mysql2 "RedNews_Server/internal/adapters/db/mysql"
	"RedNews_Server/internal/domain/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type UserInterface interface {
	GetUser(id string) entity.User
	Registration(user entity.User) bool
	Authorization(user entity.User) bool
}

func GetUser(id string) (entity.User, error) {
	db := mysql2.OpenDB()
	defer db.Close()

	query := "SELECT * FROM users where Email = ?"
	if !strings.Contains(id, "@") {
		query = "SELECT * FROM users where UsersID = ?"
	}

	row := db.QueryRow(query, id)

	var user entity.User
	if err := row.Scan(&user.UserId, &user.Login, &user.Password, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("albumsById %d: no such album", id)
		}
		return user, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return user, nil
}

func Register(user *entity.User) error {
	db := mysql2.OpenDB()
	defer db.Close()

	if rows := db.QueryRow("SELECT * FROM users where Login = ? AND Email = ?", user.Login, user.Email, user.Password); rows != nil {
		if _, err := db.Exec("INSERT INTO users(Login, Password, Email) VALUE (?,?,?)", user.Login, user.Password, user.Email); err != nil {
			return errors.New("user has been Register")
		}
		log.Println("Пользователь зарегистрирован")
		return nil
	}
	return errors.New("user already")
}

func Authorization(user *entity.User) (*entity.User, error) {
	db := mysql2.OpenDB()
	defer db.Close()
	log.Println(user)
	var users entity.User
	if rows := db.QueryRow("SELECT * FROM users where Email = ? AND Password = ?", user.Email, user.Password); rows != nil {
		log.Println("Rows")
		if err := rows.Scan(&users.UserId, &users.Login, &users.Password, &users.Email); err != nil {
			log.Println("Err")
			return nil, errors.New("User not scan")
		}
		return &users, nil
	}
	return nil, errors.New("User has been Register")
}
