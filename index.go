package goweb

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Users struct {
	Name  string `json:"name"`
	Uname string `json:"uname"`
	Pass  string `json:"pass"`
}

func init() {

	http.Handle("/", http.FileServer(http.Dir("site")))
	http.HandleFunc("/login/", login)
	http.HandleFunc("/user/", userPage)

}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]
	p := sha256.New()
	p.Write([]byte(password))
	p.Sum(nil)
	password = hex.EncodeToString(p.Sum(nil))
	user := getUser(username)

	if user.Pass == password {
		w.Write([]byte("True"))
	} else {
		w.Write([]byte("False"))
	}

}

func userPage(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	var uname string
	if len(url) >= 3 {
		uname = url[2]
	}
	user := getUser(uname)
	profile, _ := template.ParseFiles("site/profile.html")

	err := profile.Execute(w, user)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func getUser(username string) Users {
	raw, err := ioutil.ReadFile("cred.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var allUsers []Users
	json.Unmarshal(raw, &allUsers)
	for _, p := range allUsers {
		if p.Uname == username {
			return p
		}

	}
	return Users{}
}
