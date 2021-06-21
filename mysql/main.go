package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type Post struct {
	k string
	v string
}

func main() {
	var err error
	Db, err = sql.Open("mysql", "root@/go?charset=utf8mb4&parseTime=true")
	if err != nil {
		fmt.Println("errer")
		panic(err.Error())
	}

	var posts []Post

	rows, err := Db.Query("select * from map")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.k, &post.v)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}

	fmt.Print(posts)

}
