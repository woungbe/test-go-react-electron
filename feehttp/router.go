package feeHttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sample struct {
	UserID string `json:userid`
	Pwd    string `json:pwd`
	Key    string `json:key`
}

type ReturnSample struct {
	Name     string `json:name`
	Assets   string `json:assets`
	Balanace string `json:balanace`
}

func Init() {
	Router()
	http.ListenAndServe(":5000", nil)
}

func Router() {
	defer DeferTest("Router")
	http.HandleFunc("/hello", HelloWorld)
	http.HandleFunc("/post/exmple01", PostSample)
	http.HandleFunc("/post/bodyData", PostBodyData)

	http.HandleFunc("/start", Start) // 스타트
	http.HandleFunc("/stop", Stop)   // 정지
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

func PostSample(w http.ResponseWriter, req *http.Request) {
	// w.Write([]byte("Hello World"))
	fmt.Println("Method : ", req.Method) // GET, POST, PUT, DELETE,
	fmt.Println("URL : ", req.URL)
	fmt.Println("Header : ", req.Header)

	if req.Method != "POST" {
		return
	}

	b, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	fmt.Println("Body : ", string(b))

	var data Sample
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(data.Key, data.Pwd, data.UserID)

	send := &ReturnSample{
		Name:     "박경령",
		Assets:   "BTCUSDT",
		Balanace: "2000000",
	}

	byteSend, err := json.Marshal(send)
	if err != nil {
		fmt.Println("err : ", err)
	}

	w.Write(byteSend)
}

func PostBodyData(w http.ResponseWriter, req *http.Request) {
	// w.Write([]byte("Hello World"))
	fmt.Println("Method : ", req.Method)
	fmt.Println("URL : ", req.URL)
	fmt.Println("Header : ", req.Header)

	b, _ := ioutil.ReadAll(req.Body)
	// defer req.Body.Close()
	fmt.Println("Body : ", string(b))
}

func Start(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

func Stop(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}

func DeferTest(loc string) {
	fmt.Println(loc)
	r := recover()
	if r != nil {
		fmt.Println("error : ", "loc: ", r)
	}
}
