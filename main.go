package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := chi.NewRouter()

	serverPort := viper.GetString("server.http.port")
	fmt.Println(serverPort)
	http.ListenAndServe(":"+serverPort, router)
}
