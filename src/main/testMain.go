package main

import (
	bomgger "bomz.co.kr/bomgger"
	"fmt"
)

func main() {
	fmt.Printf("bomz logger start")
	fmt.Println("logger=", bomgger.GetInstance("aa"))
}
