package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	// "encoding/json"

// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// func queryBuyers() string {
// 	resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return string(body)
// }

// func queryProducts() string {
// 	resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return string(body)
// }

// func queryTransactions() string {
// 	resp, err := http.Get("https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return string(body)
// }

// func main() {
// 	fmt.Println("Service start")

// 	fmt.Println(queryBuyers())
// 	fmt.Println(queryProducts())
// 	fmt.Println(queryTransactions())
// }
