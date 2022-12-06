package main

func weird()(ret string{
	defer func(){
		recover()
		ret = "hi"
	}()
	panic("omg")
}

func main