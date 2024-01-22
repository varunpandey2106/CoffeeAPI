package main

import "fmt"

func main(){
	fmt.Println(("API is runnning"))

}

type Config struct{
	Port string
}

type Application struct{
	Config Config
}