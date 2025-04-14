package main

import (
	"fmt"
	compositions "oops/Compositions"
)

func main(){
	//object banako
	// p := structs.Person{}
	// p.SetName("ram")

	// fmt.Println("object of person is created: ",p)


	// var c poly.Shape= poly.Circle{}
	// c.Render()

	c := compositions.NewCar("hahaha",1,2)
	fmt.Println(c.HP())
	
	
}