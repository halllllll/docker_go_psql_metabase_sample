package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DBに接続できるか確認

type User struct {
	User_Id       string
	User_Password string
}

func main() {
	host := "xxx.xxx.xxx.xxx"
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=5432 user=postgres password=postgres dbname=metabase sslmode=disable", host))
	defer db.Close()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM test_user")
	if err != nil {
		fmt.Errorf("query error: %w", err)
	}

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.User_Id, &u.User_Password)
		users = append(users, u)
	}
	fmt.Printf("%v\n", users)

	// insert

	var userId int
	err = db.QueryRow("INSERT INTO test_user(user_id, user_password) VALUES($1, $2) RETURNING user_id", 666, "newuserpassword").Scan(&userId)
	if err != nil {
		fmt.Errorf("insert error: %w", err)
	}
	fmt.Printf("new user id: %d\n", userId)
}
