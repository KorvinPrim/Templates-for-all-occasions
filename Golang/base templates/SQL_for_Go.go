package main

//The template is made to interact with a Mysql database deployed on the basis of an Open Server.
import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	//handleRequest()
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Установка данных

	//insert, err := db.Query("INSERT INTO users (name,age) VALUES ('Anton',22)")
	//res, err := db.Query("SELECT name,age FROM users")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()
	//defer res.Close()

	res, err := db.Query("SELECT name,age FROM users")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))

	}

}
