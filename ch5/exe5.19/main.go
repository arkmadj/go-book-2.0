package main

import "fmt"

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hi there!!"
	}()
	panic("omg!!")
}

func main() {
	fmt.Println(weird())
}
