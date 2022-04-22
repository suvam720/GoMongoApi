package main

import (
	"fmt"
	"net/http"

	"github.com/suvam720/mongoapi/pkg/router"
)

func main() {
	fmt.Println("server started...")
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("on port 4000...")
}
