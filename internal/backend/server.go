package backend

import (
	"fmt"
	"log"
	"net/http"
)


func Auth(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println(username, password)
	} else {
		http.ServeFile(w, r, "./internal/frontend/index.html")
		return
	}
}

func Run() {
	PORT := ":7000"

	http.HandleFunc("/", Auth)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}

}
