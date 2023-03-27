package main
import ("fmt")

const(
	_ = iota
	KB = 1 << (10*iota)
	MB
	GB

)

func main(){

fileSize := 4000000000.
fmt.Printf("%.2f GB\n", fileSize/GB)
fmt.Println(KB,GB)


}