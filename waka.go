// waka nano game
//  aq@okaq.com
// 2021-07-30
package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	WAKA = "waka.html"
)

func motd() {
	fmt.Println(time.Now().String())
	fmt.Println("okaq web localhost:8080")
}

func WakaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	http.ServeFile(w,r,WAKA)
}

func main() {
	motd()
	http.HandleFunc("/", WakaHandler)
	http.ListenAndServe(":8080", nil)
}



