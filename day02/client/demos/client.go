package demos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func DealPostHandler() {
	apiUrl := "http://192.168.56.11:8081/req/post"
	contentType := "application/json"
	data := `{"name":"天才","password":"654321"}`
	resp, _ := http.Post(apiUrl, contentType, strings.NewReader(data))
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}

func DealGetHandler() {
	apiUrl := "http://192.168.56.11:8081/req/get"
	data := url.Values{}
	data.Set("name", "天才")
	u, _ := url.ParseRequestURI(apiUrl)
	u.RawQuery = data.Encode()
	resp, _ := http.Get(u.String())
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("请求路由为：", u.String())
	fmt.Println("返回数据为：", string(b))
}

func DealGetHandler1() {
	apiUrl := "http://192.168.56.11:8081/req/get?name=天才"
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
