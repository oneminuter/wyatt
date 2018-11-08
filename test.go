package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

func main() {
	//getToken()
	//getCurrentPath()
	addMessage()
}

func getToken() {
	//token := util.NewToken(456, 2, "1232465aasdf")
	//util.Logger(token)

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJkYzBlZjdlNDkyZGJlMjA2Nzc5MDI4NzMxOGU3MzI5ZCIsImV4cCI6MTUzODE1MjgxNSwianRpIjoiMjgifQ.RTzUWq77BKPHmt19c5vKwadcBBwTKWnMA-R8HZP5-w8"

	t, err := util.ParseToken(token)
	if err != nil {
		if strings.Contains(err.Error(), constant.TokenExpired) {
			log.Println("expired")
		}
		log.Println(err)
		return
	}

	log.Println("uuid: ", t.UUID)
	log.Println("userId: ", t.UserId)
	log.Println("status: ", t.Status)
}

func getCurrentPath() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

func addMessage() {
	var m = model.Message{
		UserId:  1,
		MsgType: "system",
		Content: "这是一个测试消息",
	}

	m.Add()
}
