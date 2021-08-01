package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	confFile     = "config"
	pathConfFile = "."
	typeConfFile = "yml"
)

func main() {
	config, err := loadEnv(confFile, pathConfFile, typeConfFile)

	fs := http.FileServer(http.Dir("./static"))

	if err != nil {
		fmt.Println(err)
	}
	mux := http.NewServeMux()
	handlers(mux, fs)

	if err := http.ListenAndServe(":"+config.Server.Port, mux); err != nil {
		log.Fatal(err)
	}

}
