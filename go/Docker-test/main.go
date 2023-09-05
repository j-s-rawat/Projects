package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

type post struct {
	HelloUserId int `json:"userId"`
	Id          int
	Title       string
	Body        string
}

type reply struct {
	Title string `json:"Modidied Title"`
	Body  string `json:"Modified Body"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts/{id}", myHandler)
	http.ListenAndServe(":8080", r)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	//query variables, you don't do that with mux. Just grab them straight from the request.
	//r.URL.Query().Get("varName")
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p post
	json.Unmarshal(b, &p)
	spew.Dump(p)

	rep := reply{p.Title, p.Body}
	rb, err := json.Marshal(rep)
	if err != nil {
		panic(err)
	}
	w.Write(rb)
}
