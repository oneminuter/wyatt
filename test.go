package main

import (
	"log"
	"strings"
	"wyatt/api/constant"
	"wyatt/util"
)

func main() {
	getToken()
	//modelTest()
}

func modelTest() {
	util.GetIpInfo("49.4.136.242")
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
