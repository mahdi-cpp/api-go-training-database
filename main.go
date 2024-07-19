package main

import "github.com/mahdi-cpp/api-go-gallery/repository"

func main() {
	//m := model.Music{Name: "name", Width: 58}
	//fmt.Println(m)
	repository.DatabaseInit()

	//repository.UserStart()
	repository.PostStart()

	Run()

}
