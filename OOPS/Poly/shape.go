package poly

import "fmt"

//interfacae

type Shape interface{
	Render()
}

type Circle struct{}
type Square struct{}

func (c Circle) Render(){
	fmt.Println("Circle print hudai xa")
}

func(s Square) Render(){
	fmt.Print("Sq print hudai xa")
}