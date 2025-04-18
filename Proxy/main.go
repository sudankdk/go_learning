package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)



func HandleRequestAndRedirect(w http.ResponseWriter, r *http.Request){
	fmt.Printf("proxying request: %s %s",r.Method,r.URL.String() )

	//naya request banaune  url lai target garna
	outReq,err:=http.NewRequest(r.Method,r.URL.String(),r.Body)
	if  err != nil{
		http.Error(w,"Bad Request",http.StatusBadRequest)
		return 
	}

	outReq.Header= r.Header



	/// ya neri default transport  use garna sakinxa
	resp, err :=http.DefaultTransport.RoundTrip(outReq)
	if err != nil{
		http.Error(w, "Error forwarding request", http.StatusBadGateway)
		return
	}

	defer resp.Body.Close()

	for key,values := range resp.Header{
		for _,value:= range values{
			w.Header().Add(key,value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)

}


func main(){
	handler := http.HandlerFunc(HandleRequestAndRedirect)
	log.Println("Starting proxy server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}