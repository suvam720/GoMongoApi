package main

import (
	"fmt"
	"net/http"

	"github.com/suvam720/mongoapi/pkg/router"
)

func main() {
	fmt.Println("server started...")
	fmt.Println("on port 4000...")

	r := router.Router()
	http.ListenAndServe(":4000", r)

}
