package services

import (
	"fmt"
	"log"
	"user-server/models"
	"user-server/utils"
)

func GetUser() []models.User {
	conn := utils.DbConnect()
	userList := make([]models.User, 0)
	rows, err := conn.Query("select * from user where email = ?", "lujun@example.com")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Status)
		if err != nil {
			log.Fatal(err)
		}
		userList = append(userList, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return userList
}

func CreateUser(user models.User) {
	conn := utils.DbConnect()
	sql := "INSERT INTO user (`username`, `password`, `email`, `status`) VALUES (?, ?, ?, ?)"
	res, err := conn.Exec(sql, user.Username, user.Password, user.Email, user.Status)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
