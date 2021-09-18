package backend

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var dict map[string]string = map[string]string{
	"admin": "1234",
}


func Run() {
	PORT := ":7002"
	tmpl,_ := template.ParseFiles("./frontend/home.html")
	data, err := ioutil.ReadFile("users.json")

	if err != nil {
		log.Println(err)
		return
	}
	var users []Users
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(users)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	//fileContents, err := ioutil.ReadFile("./frontend/user.html")
	//	//if err != nil {
	//	//	log.Println(err)
	//	//	w.WriteHeader(http.StatusNotFound)
	//	//	return
	//	//}
	//	// w.Write(fileContents)
	//	err := tmpl.Execute(w, users)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//})

	http.HandleFunc("/auth/", func(w http.ResponseWriter, r *http.Request) {
		fileContents, err := ioutil.ReadFile("./frontend/index.html")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method == "GET" {
			w.Write(fileContents)
		}
		if r.Method == "POST" {
			username, password := "", ""
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
			}
			username = r.FormValue("username")
			password = r.FormValue("password")
			fmt.Println(username, password)
			if _, ok := dict[username]; ok && dict[username] == password{
				fmt.Println("OK")
				http.Redirect(w, r, "home", http.StatusSeeOther)
			} else {
				w.Write(fileContents)
			}
		}
	})

	http.HandleFunc("/auth/home", func(w http.ResponseWriter, r *http.Request) {
		//fileContents, err := ioutil.ReadFile("./frontend/home.html")

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method == "GET" {
			err := tmpl.Execute(w, users)
			if err != nil {
				log.Println(err)
				return
			}
			// w.Write(fileContents)
		}
		if r.Method == "POST" {
				password, confirmPassword := "", ""

				err := r.ParseForm()
				if err != nil {
					log.Println(err)
				}
				password = r.FormValue("password")
				confirmPassword = r.FormValue("confirm-password")
				fmt.Println(password, confirmPassword)

			if password == confirmPassword {
				fmt.Println("OK")
				http.Redirect(w, r, "home", http.StatusSeeOther)
			}
			} else {
				// w.Write(fileContents)
				err := tmpl.Execute(w, users)
				if err != nil {
					log.Println(err)
					return
				}
			}
	})

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}

}
