package main

import (
	"fmt"
	"log"
	"simpleCrudMongoDB/models"
)

func main() {
	var user models.IUser
	user = &models.User{}

	userData := models.User{
		Username: "fiqri",
		Email: "fiqrikm8@gmail.com",
		Address: "bandung barat",
	}
	res, err := user.CreateUser(userData)
	if err != nil {
		log.Fatal(err.Error())
	}
	if res.InsertedID != nil {
		fmt.Println("Insert Success")
	}

	userFilter := models.User{
		Username: "fiqri",
	}
	res2, err := user.FindUser(userFilter)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Name:", res2.Username)
	fmt.Println("Email:", res2.Email)
	fmt.Println("Address:", res2.Address)

	userData2 := models.User{
		Username: "alya",
	}
	res3, err := user.UpdateUser(userFilter, userData2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Name:", res3.Username)
	fmt.Println("Email:", res3.Email)
	fmt.Println("Address:", res3.Address)

	userFilter = models.User{
		Username: "alya",
	}
	err = user.Delete(userFilter)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("User data deleted")
}
