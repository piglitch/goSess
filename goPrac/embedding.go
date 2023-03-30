package main
import (
	"fmt"
)

type Device struct{
	name string
	wireless bool
}

type mobilePhone struct{
	Device
	company string
	screen float32
	processor string
}

func main(){
	p:= mobilePhone{}
	p.name = "smartphone"
	p.wireless = true
	p.company = "OnePlus"
	p.screen = 6.62
	p.processor = "Snapdragon 8Gen1"
	fmt.Println(p)
}