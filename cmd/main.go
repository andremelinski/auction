package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andremelinski/auction/config/database/mongodb"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("error trying to load env")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("conectado")

}
