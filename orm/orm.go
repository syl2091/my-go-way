package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "", "127.0.0.1:3306", "blog"))
	//db.Ping()
	defer db.Close()
	if err != nil {
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}
	//创建数据库
	//_, err2 := db.Exec("CREATE TABLE user(id INT NOT NULL , name VARCHAR(20), PRIMARY KEY(ID));")
	//插入数据
	//_, err2 := db.Query("INSERT INTO user VALUES(1, 'Wade')")
	result, err2 := db.Query("select * from user")
	if err2 != nil {
		log.Fatal(err2)
	}

	for result.Next() {
		var id int
		var name string
		err = result.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d, Name: %s\n", id, name)
	}
}
