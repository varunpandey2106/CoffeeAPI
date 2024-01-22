package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	fmt.Println(("API is runnning"))

}

type Config struct{
	Port string
}

func (app*Application) Serve() error{

	err:=godotenv.Load()
	if err != nil{
		log.Fatal("error loading .env file")
	}
	port:= os.Getenv("PORT")
	fmt.Println("API is listening on port: ",port)

	srv:= &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}
	return srv.ListenAndServe()

}

type Application struct{
	Config Config
}