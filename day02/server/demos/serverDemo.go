package demos

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//处理Get请求
func DealGetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	type Info struct {
		Name     string
		Password string
		Age      int
	}
	user := Info{
		Name:     name,
		Password: "12345",
		Age:      12,
	}
	json.NewEncoder(w).Encode(user)
}

//处理Post请求
func DealPostHandler(w http.ResponseWriter, r *http.Request) {
	bodyContent, _ := ioutil.ReadAll(r.Body)
	type Info struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var d Info
	json.Unmarshal(bodyContent, &d)
	json.NewEncoder(w).Encode(d)
}

func DealGetHandler2(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	type Info struct {
		Name     string
		Password string
		Age      int
	}
	user := Info{
		Name:     name,
		Password: "12345",
		Age:      12,
	}
	json.NewEncoder(w).Encode(user)
}

func DealGetHandler1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var names string
	if len(query["name"]) > 0 {
		names = strings.Join(query["name"], "")
		if names == "" {
			names = "陌生人"
		}
	}
	var build strings.Builder
	build.WriteString("Hello ")
	build.WriteString(names)
	build.WriteString(" !")
	w.Write([]byte(build.String()))
}
