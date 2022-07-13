package main

import(
	"fmt"
)

func Print(){
	fmt.Println("Hello-world-from-pkg-print")
}

func main(){
	fmt.Println("Hello-world-from-pkg-main")
	Print()	
}

