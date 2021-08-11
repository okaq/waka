// waka nano game
//  aq@okaq.com
// 2021-08-12
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	RAKA = "raka.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("okaq web localhost:8080")
}

func RakaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,RAKA)
}

func main() {
	motd()
	http.HandleFunc("/", RakaHandler)
	http.ListenAndServe(":8080", nil)
}



