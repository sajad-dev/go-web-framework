package checksystem

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"math"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	testutils "github.com/sajad-dev/go-web-framework/internal/test-utils"
	"github.com/sajad-dev/go-web-framework/pkg/api"
	"github.com/sajad-dev/go-web-framework/pkg/controllers"
)

func TestApi(t *testing.T) {


	api.RouteRun()
	client := http.Client{Timeout: time.Second * 2}
	for _, route := range api.RouteList {

		funcname := runtime.FuncForPC(reflect.ValueOf(route.Controller).Pointer()).Name()
		funcname = strings.Split(funcname, ".")[2]
		structName := funcname + "Struct"
		structT := controllers.StructRegistry[structName]

		v := reflect.ValueOf(structT)

		arr := []map[string]string{}
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := v.Type().Field(i)
				tags := field.Tag.Get("validation")
				valid := strings.Split(tags, "|")
				for _, v := range valid {
					if strings.Contains(v, ":") {
						spl := strings.Split(v, ":")
						i, _ := strconv.Atoi(spl[1])
						good, bad := testutils.TestMapWithArgs[spl[0]](i)
						arr = append(arr, map[string]string{"good": good, "bad": bad, "fieldname": field.Name})
					} else {
						good, bad := testutils.TestMap[v]()
						arr = append(arr, map[string]string{"good": good, "bad": bad, "fieldname": field.Name})
					}
				}
			}
		}

		x := 0
		ou_bodys := []map[string]string{}
		for x < 2*len(arr)+1 {
			x++
			body := map[string]string{}
			for i, v := range arr {
				if math.Floor(float64((x+1)/2)) == float64(i+1) {
					if (x+1)%2 == 1 {
						body[strings.ToLower(v["fieldname"])] = v["bad"]
					} else {
						body[strings.ToLower(v["fieldname"])] = v["good"]
					}

				}
			}
			ou_bodys = append(ou_bodys, body)
		}
		reqnum := []int{}
		for i, v := range ou_bodys {
			reqnum = append(reqnum, i)
			jsonData, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			req, err := http.NewRequest(string(route.Method), "http://127.0.0.1:8000"+route.Pattern, bytes.NewBuffer(jsonData))
			if err != nil {
				t.Fatal("You Have Error In Routing:(Add Req)", err, "-> In route ::", route.Pattern)
			}
			req.Header.Set("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				t.Fatal("You Have Error In Routing :(Send Request)", err, "-> In route ::", route.Pattern)
			}

			body, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return
			}

			status := res.Status
			if i != 0 && !strings.Contains(status, "409") {
				t.Fatal("You have problemm in validation", v, string(jsonData), string(body), string(route.Method))
			}
			if i == 0 && !strings.Contains(status, "200") {
				t.Fatal("You have problemm in validation (2)")

			}
		}

	}
	

}
