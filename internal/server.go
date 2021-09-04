package internal

import (
	"fmt"
	"net/http"
)


func Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
	return
}

func Run() {
	PORT := ":7000"

	http.HandleFunc("/", Auth)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
	}

}
