package main

import (
	"fmt"
	"log"
	"net/http"

	db "golang_service/database"
	"golang_service/handler"
	"golang_service/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	DB := db.Init()
	h := handler.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/create-pair", middlewares.CheckJwt(h.CreatePair)).Methods(http.MethodPost)
	router.HandleFunc("/create-exam", middlewares.CheckJwt(h.CreateExams)).Methods(http.MethodPost)
	router.HandleFunc("/submit-test", middlewares.CheckJwt(h.SubmitTest)).Methods(http.MethodPost)
	router.HandleFunc("/end-submit", middlewares.CheckJwt(h.EndSubmit)).Methods(http.MethodPost)
	router.HandleFunc("/submit-answer", middlewares.CheckJwt(h.SubmitAnswer)).Methods(http.MethodPost)
	router.HandleFunc("/caculate-winner", middlewares.CheckJwt(h.CaculateWinner)).Methods(http.MethodPost)
	router.HandleFunc("/withdraw", middlewares.CheckJwt(h.Withdraw)).Methods(http.MethodPost)
	router.HandleFunc("/check-account-balance", middlewares.CheckJwt(h.CheckAccountBalance)).Methods(http.MethodPost)
	router.HandleFunc("/check-contract-balance", middlewares.CheckJwt(h.CheckContractBalance)).Methods(http.MethodGet)
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
