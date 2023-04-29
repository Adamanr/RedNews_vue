package mysql

import (
	"RedNews_Server/internal/adapters/db/mysql"
	"RedNews_Server/internal/domain/entity"
	"errors"
	"strconv"
	"strings"
)

type ArticlesInterface interface {
	GetArticles(id string) entity.Articles
}

func GetArticles(selects string) (*entity.Articles, error) {
	var query interface{} = "SELECT * FROM articles WHERE ArticlesID = ?"

	if _, err := strconv.Atoi(selects); err != nil {
		selects = strings.Replace(selects, "_", " ", -1)
		query = "SELECT * FROM articles WHERE Title LIKE concat(?,'%')"
	}

	db := mysql.OpenDB()
	defer db.Close()

	if rows := db.QueryRow(query.(string), selects); rows != nil {
		var articles entity.Articles
		if err := rows.Scan(&articles.ArticlesID, &articles.Title, &articles.Body, &articles.AuthorID, &articles.Image, &articles.Created_at, &articles.Updated_at); err != nil {
			return nil, err
		}
		return &articles, nil
	}
	return nil, errors.New("not find this articles")
}

func NewArticles(articles entity.Articles) (err error) {
	db := mysql.OpenDB()
	defer db.Close()

	if _, err := db.Exec("INSERT INTO articles(Title, Body, Author_id, Image) VALUE(?,?,?,?)", articles.Title, articles.Body, articles.AuthorID, articles.Image); err != nil {
		return err
	}

	return nil
}
