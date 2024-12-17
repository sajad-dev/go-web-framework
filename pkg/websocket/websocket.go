package websocket

import (
	"context"
	"fmt"
	"net/http"

	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sajad-dev/go-web-framework/pkg/exception"
	"github.com/sajad-dev/go-web-framework/pkg/model"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var ConnactionMap = map[string]*websocket.Conn{}

func Connaction(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		exception.Log(err)
		auth, ok := r.Context().Value("auth").(model.GetOutput)
		if !ok {
			http.Error(w, "User ID not found in context", http.StatusInternalServerError)
			return
		}
		fmt.Println(1)
		ConnactionMap[auth[0]["id"]] = conn
		defer conn.Close()
		next.ServeHTTP(w, r)

	})
}
func MessageHandel(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth, ok := r.Context().Value("auth").(model.GetOutput)
		if !ok {
			return
		}

		conn := ConnactionMap[auth[0]["id"]]
		for {
			_, p, err := conn.ReadMessage()
			exception.Log(err)
			text := strings.Split(string(p)[0:4], ",")

			if strings.TrimSpace(text[1]) == "1" {
				out := model.Get([]string{"user_id", "group_id"}, "linkgroups", []model.Where_st{
					model.Where_st{Key: "group_id", Value: text[0], After: "", Operator: "="},
				}, "", false)

				if len(out) == 0 {
					conn.Close()

				}

				for i, val := range out {
					if val["user_id"] == auth[0]["id"] {
						group := model.Get([]string{"id", "name", "discription", "slug", "profile"}, "groups", []model.Where_st{
							model.Where_st{Key: "id", Value: val["group_id"], After: "", Operator: "="},
						}, "", false)
						ctx := context.WithValue(r.Context(), "group_data", group)
						fmt.Println(group)

						next.ServeHTTP(w, r.WithContext(ctx))
					} else {
						fmt.Println(i)
						if i == len(out)-1 {
							conn.Close()
						}
					}
				}
			} else {
				user := model.Get([]string{"id", "first_name", "last_name", "email"}, "users", []model.Where_st{
					model.Where_st{Key: "id", Value: text[0], After: "", Operator: "="},
				}, "", false)
				user[0]["message"] = string(p)
				ctx := context.WithValue(r.Context(), "send_data", user)
				next.ServeHTTP(w, r.WithContext(ctx))

			}

		}
	})
}

func Controller(w http.ResponseWriter, r *http.Request) {
	auth, ok := r.Context().Value("send_data").(model.GetOutput)
	if !ok {
		return
	}
	conn := ConnactionMap[auth[0]["id"]]
	fmt.Println(auth[0],ConnactionMap)
	conn.WriteMessage(websocket.TextMessage, []byte(auth[0]["message"]))

}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1))
		where := []model.Where_st{
			model.Where_st{Key: "token", Value: token, After: "", Operator: "="},
		}
		output := model.Get([]string{"token", "expiration", "user_id"}, "token", where, "", false)
		if len(output) == 0 {
			return
		}
		parsstime, err := time.Parse("2006-01-02 15:04:05", output[0]["expiration"])
		exception.Log(err)
		timenow := time.Now()
		if len(output) > 0 && !parsstime.Before(timenow) {
			user := model.Get([]string{"id", "first_name", "last_name", "email"}, "users", []model.Where_st{
				model.Where_st{Key: "id", Value: output[0]["user_id"], After: "", Operator: "="},
			}, "", false)
			ctx := context.WithValue(r.Context(), "auth", user)
			fmt.Println(user, "--1")
			next.ServeHTTP(w, r.WithContext(ctx))

		}
	})

}

func Handel() {

	err := http.ListenAndServe(":3000", nil)
	exception.Log(err)
}
