package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed secret.json
var Env embed.FS

type EnvJson struct {
	Host   string `json:"host"`
	DbPort string `json:"psql_port"`
	DbName string `json:"psql_dbname"`
	DbUser string `json:"psql_user"`
	DbPw   string `json:"psql_pw"`
}

// DBに接続できるか確認

type User struct {
	User_Id       string
	User_Password string
}

func main() {
	var env EnvJson
	envData, err := Env.ReadFile("secret.json")
	if err != nil {
		fmt.Errorf("env file read error: %w", err)
	}
	json.Unmarshal(envData, &env)

	host := env.Host
	dbPort := env.DbPort
	dbName := env.DbName
	dbUser := env.DbUser
	dbPw := env.DbPw
	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, dbUser, dbPw, dbName))
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
	err = db.QueryRow("INSERT INTO test_user(user_id, user_password) VALUES($1, $2) RETURNING user_id", 999, "yesyes").Scan(&userId)
	if err != nil {
		fmt.Errorf("insert error: %w", err)
	}
	fmt.Printf("new user id: %d\n", userId)
}
