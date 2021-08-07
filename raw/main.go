package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type config struct {
  API struct {
    AuthToken string
    Port      string
  }

  Server struct {
    IP       string
    Username string
    Password string
  }
}

var C config

type jsonStruct struct {
	Key    string
	Target string
	Method string
	Port   int
	Time   int
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	w.Header().Set("Content-Type", "text")
	fmt.Fprintf(w, "Access Forbidden")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Armour API\n")
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		var t jsonStruct
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Println(err)
		}
		t.Key = CheckStr(t.Key) // Security reasons
		if t.Key != C.API.AuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized\n")
			return
		} else {
			t.Method = CheckStr(t.Method) // Security reasons
			res := CheckInput(t.Target, t.Method, t.Port, t.Time) // Security reasons
			if res != true {
				w.WriteHeader(http.StatusNotAcceptable)
				fmt.Fprintf(w, "Input not accepted\n")
				return
			}
                        // Call python script to connect to server via SSH
                        go ConnectServer(t.Method, t.Target, t.Port, t.Time)
		}
	}
}

func main() {
        ParseJson("./config.json")
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/api", apiHandler)
        fmt.Println("Started the API - Coded by M0ba")
        fmt.Println(http.ListenAndServe(C.API.Port, nil))
}
