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

	router.Post("/", func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("HELLO WORD"))
	})

	router.Get("/{id}", func(response http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")

		response.WriteHeader(http.StatusOK)
		response.Write([]byte(id))
	})

	serverPort := viper.GetString("server.http.port")
	fmt.Println(serverPort)
	http.ListenAndServe(":"+serverPort, router)
}
