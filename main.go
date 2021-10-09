package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"task1/controllers"
)

func main() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kaustubh:kaustubh542002@cluster0.vjebv.mongodb.net/instadbAPI?retryWrites=true&w=majority"))

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// DB := client.Database("instadbAPI")
	// usercollection := DB.Collection("users")
	// fmt.Println(usercollection)
	fmt.Print("serverstarted")
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb+srv://kaustubh:kaustubh542002@cluster0.vjebv.mongodb.net")
	if err != nil {
		panic(err)
	}
	return s
}
