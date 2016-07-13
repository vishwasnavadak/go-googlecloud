package hello

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type users struct {
	Name  string `json:"name"`
	Uname string `json:"uname"`
	Pass  string `json:"pass"`
}

func init() {

	http.Handle("/", http.FileServer(http.Dir("site")))
	http.HandleFunc("/login/", login)
	http.HandleFunc("/user/", user)

}
func login(w http.ResponseWriter, r *http.Request) {
	flag := false
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	p := sha256.New()
	p.Write([]byte(password))
	p.Sum(nil)
	password = hex.EncodeToString(p.Sum(nil))
	user := getUsers()
	for _, p := range user {
		if p.Uname == username && p.Pass == password {
			flag = true
			break
		}

	}
	if flag {
		w.Write([]byte("True"))
	} else {
		w.Write([]byte("False"))
	}

}
func user(w http.ResponseWriter, r *http.Request) {

}
func getUsers() []users {
	raw, err := ioutil.ReadFile("cred.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []users
	json.Unmarshal(raw, &c)
	return c
}
