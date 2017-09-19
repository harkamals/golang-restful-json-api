package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	CreatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	fmt.Println("Init: Post")

	var err error
	Db, err = gorm.Open("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&Post{}, &Comment{})

}

func main() {
	fmt.Println("Main: Post")

	// Create post
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	Db.Create(&post)
	fmt.Println(post)

	// Create comment
	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment)
	fmt.Println(comment)

	// Get comments from a post
	var readPost Post
	Db.Where("Id = $1", post.Id).First(&readPost)

	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments)

}
