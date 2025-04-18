package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){

	// http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w,"This is a server")
	// })

	// if err:= http.ListenAndServe(":8080",nil); err!=nil{
	// 	fmt.Println("server error.",err)
	// }

	resp, err:=http.Get("https://medium.com/@emonemrulhasan35/net-http-package-in-go-e178c67d87f1")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	body, err :=io.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(body))
}