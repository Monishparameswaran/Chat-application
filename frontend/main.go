package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Message struct {
	Name string
	Msg  string
}

var (
	usrname  string
	messages []Message
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("/home/jagan/Development/Chat-application/frontend")))
	http.HandleFunc("/home", getnamehandler)
	http.HandleFunc("/post", posthandler)
	http.HandleFunc("/first.html", servefirstpage)

	fmt.Println("server started running http/localhost:8080 ..!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server cannot start8080")
	}
	fmt.Println("server started running in port 8080")
}
func getnamehandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		fmt.Println(name)
		usrname = name
	}

	http.ServeFile(w, r, "first.html")
}
func posthandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		msg := r.FormValue("msg")

		fmt.Println(usrname + " " + msg)
		if usrname != "" && msg != "" {
			chatmessage := Message{usrname, msg}

			messages = append(messages, chatmessage)
		}

	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "error in parsing the html")
	}

	err = tmpl.Execute(w, messages)
	if err != nil {
		fmt.Println("cannot complete the parse")
	}
}
func servefirstpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "first.html")
}
