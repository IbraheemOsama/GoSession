package services

import (
	"GoSession/helpers"
	"GoSession/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"net/http"
)

type MoviesService struct {
	ResponseHelper *helpers.ResponseHelper
}

func (service MoviesService) InitAndRegisterService(router *mux.Router) {
	service.ResponseHelper = &helpers.ResponseHelper{}
	router.HandleFunc("/movies", service.AllMoviesEndPoint).Methods("GET")
}

// GET list of movies
func (service MoviesService) AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	page, err := strconv.Atoi(queryValues.Get("page"))
	if err != nil {
		service.ResponseHelper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if page > 3 || page < 1 {
		page = 3
	}

	ch := make(chan []models.Movie, page)

	for i := 1; i <= page; i++ {
		go readJson(i, ch)
	}

	var result []models.Movie
	for i := 1; i <= page; i++ {
		result = append(result, <-ch...)
	}

	service.ResponseHelper.RespondWithJSON(w, http.StatusOK, result)
}

func readJson(fileNumber int, ch chan []models.Movie) {
	fileName := "json\\movies" + strconv.Itoa(fileNumber) + ".json"
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened " + fileName)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var movies []models.Movie
	json.Unmarshal(byteValue, &movies)

	ch <- movies
}
