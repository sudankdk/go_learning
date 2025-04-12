package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",handleForm)
	http.HandleFunc("/hello",handleHello)

	fmt.Println("server running at port 8080")

	if err := http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}

}


func handleHello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method !="GET"{
		http.Error(w,"This method is not supported",http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Helloe")
}
func handleForm(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Println(w,"eroor aayo hai: %v",err)
		return
	}
	fmt.Fprintf(w, "form added nice")
	name := r.FormValue("name")
	fmt.Fprintf(w,"ya aayo name: %s",name)
}