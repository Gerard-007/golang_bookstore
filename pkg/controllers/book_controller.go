package controllers

import (
	"encoding/json"
	"github.com/Gerard-007/golang_bookstore/pkg/models"
	"github.com/Gerard-007/golang_bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(writer http.ResponseWriter, request *http.Request) {

}

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetBookByID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookid := vars["bookid"]
	ID, err := strconv.ParseInt(bookid, 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
	}
	bookDetails, _ := models.GetBookById(ID)
	response, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {

}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {

}
