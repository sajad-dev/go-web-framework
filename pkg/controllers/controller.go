package controllers

import (
	"time"

	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"github.com/sajad-dev/go-web-framework/pkg/model"
	"github.com/sajad-dev/go-web-framework/pkg/utils"
)
var StructRegistry = map[string]interface{}{
}

func GetToken(selection []string, email string, password string) (string,map[string]string, bool) {
	getuser := model.Get(selection, "users", []model.Where_st{
		model.Where_st{Key: "email", Value: email, After: "", Operator: "="},
	}, "", false)
	token := ""
	ok := false
	if getuser[0]["password"] == password {
		ok = true
		token = utils.GenerateToken()
		now := time.Now().Add(time.Hour * 24)
		tokendata := map[string]string{
			"token":      token,
			"expiration": now.Format("2006/01/02 15:04:05"),
			"user_id":    getuser[0]["id"],
		}
		model.Insert(tokendata, "token")
	}
	go RemoveToken()
	return token,getuser[0], ok
}

func RemoveToken() {
	out := model.Get([]string{"expiration", "id"}, "token", []model.Where_st{}, "", false)
	for _, v := range out {
		parsstime, err := time.Parse("2006-01-02 15:04:05", v["expiration"])
		exception.Log(err)

		timenow := time.Now()
		if parsstime.Before(timenow) {
			model.Delete("token", [2]string{"id", v["id"]})
		}
	}

}


