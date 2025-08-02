package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Hazrat Ali. I'm software engineer.")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	handelCors(w)
	handlePreflightReq(w, r)

	if r.Method != "GET" {
		http.Error(w, "Plz Give me Get Request", 400)
		return
	}

	sendData(w, productList, 200)

}

func creatProduct(w http.ResponseWriter, r *http.Request) {
	handelCors(w)
	handlePreflightReq(w, r)

	if r.Method != "POST" {
		http.Error(w, "Plz Give me Post Request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "pls give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func handelCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE, OPTIONS")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/products", getProducts)

	mux.HandleFunc("/creat-products", creatProduct)

	fmt.Println("Server is runing at port :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Fresh Apple form Bangladesh.",
		Price:       140,
		ImgURL:      "https://5.imimg.com/data5/AK/RA/MY-68428614/apple.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "Fresh orange form Bangladesh.",
		Price:       100,
		ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	// prd4 := Product{
	// 	ID:          4,
	// 	Title:       "Orange",
	// 	Description: "Fresh orange form Bangladesh.",
	// 	Price:       100,
	// 	ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	// }
	// prd5 := Product{
	// 	ID:          5,
	// 	Title:       "Orange",
	// 	Description: "Fresh orange form Bangladesh.",
	// 	Price:       100,
	// 	ImgURL:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	// }
	// prd6 := Product{
	// 	ID:          6,
	// 	Title:       "Mango",
	// 	Description: "Fresh Mango form Bangladesh.",
	// 	Price:       10000000,
	// 	ImgURL:      "https://img.freepik.com/free-photo/mango-still-life_23-2151542114.jpg?w=360",
	// }
	productList = append(productList, prd1, prd2, prd3)
	// productList = append(productList, prd4, prd5, prd6)
}
