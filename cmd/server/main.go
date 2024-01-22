package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println(("API is runnning"))

}

type Config struct{
	Port string
}

func (app*Application) Serve() error{

	port:="8080"

	fmt.Println("API is listening on Port:", port)

	srv:= &http.Server{
		Addr:	fmt.Sprint(":%s", port),
		Handler: router.Routes(),
	}
	return srv.ListenAndServer()
}

type Application struct{
	Config Config
}