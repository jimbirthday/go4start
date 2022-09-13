package httpx

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Wrold!\r\n的奋斗奋斗") //这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "123")                   //这个写入到w的是输出到客户端的
}

func TestClient(t *testing.T) {
	base, err := url.Parse("http://0.0.0.0:7000")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	to := 10

	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
		Timeout: time.Second * time.Duration(to),
	}

	requ, err := http.NewRequestWithContext(context.Background(), "post", base.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	requ.Header.Set("transfer-encoding", "chunked")
	resp, err := client.Do(requ)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println(string(all))
}
