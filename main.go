package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

)

const (
	USER_GENERAL = "general"
	USER_PATIENT = "patient"
)
func main() {
	router := mux.NewRouter()
	router.Handle("/member/{type}", NewMemberHandler())
	router.Handle("/record/{type}", NewRecordHandler())
	router.Handle("/evaluation/{type}", NewEvaluationHandler())

	fmt.Println("Thais-mood Api on port :5000")
	if err := http.ListenAndServe(":5000", router); err != nil {
		log.Fatal(err)
	}
}

func NewMemberHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := mux.NewRouter()
		router.Use(setContentTypeMiddleWare)
		params := mux.Vars(r)
		fmt.Println(params["type"])

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPost:
			break
		case http.MethodPut:
			break
		case http.MethodDelete:
			break
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}


	})
}

func NewRecordHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := mux.NewRouter()
		router.Use(setContentTypeMiddleWare)
		params := mux.Vars(r)
		fmt.Println(params["type"])

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPost:
			break
		case http.MethodPut:
			break
		case http.MethodDelete:
			break
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}


	})
}

func NewEvaluationHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := mux.NewRouter()
		router.Use(setContentTypeMiddleWare)
		params := mux.Vars(r)
		fmt.Println(params["type"])

		switch r.Method {
		case http.MethodGet:
			break
		case http.MethodPost:
			break
		case http.MethodPut:
			break
		case http.MethodDelete:
			break
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}


	})
}

func setContentTypeMiddleWare(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}



