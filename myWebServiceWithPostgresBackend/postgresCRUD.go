package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=webservice dbname=webservice sslmode=disable")
	if err != nil {
		fmt.Println("cannot connect to the database")
		panic(err)
	}
}


// Create an entry in the Posts table
func (post *Post) create() (err error){
	statement := "INSERT into posts (content,author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)

	if err != nil{
		fmt.Println("cannot INSERT into the database")
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}




// retrieve an entry from the posts table in Postgres DB
func retrieve(id int) (post Post, err error){
	post = Post{}
	err =  Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// update an entry in the Postgres
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2 author = $3 where id = $1",post.Id, post.Content, post.Author)
	return
}

//delete an entry from posts table in the Postgres DB
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}


