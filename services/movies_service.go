package services

import (
	"GoSession/helpers"
	"GoSession/models"
	"strconv"

	"net/http"
)

type MoviesService struct {
	ResponseHelper *helpers.ResponseHelper
}

// TODO to implement interface Register_router_service and include the code below in the function
//service.ResponseHelper = &helpers.ResponseHelper{}
//router.HandleFunc("/movies", service.AllMoviesEndPoint).Methods("GET")

// GET list of movies
func (service MoviesService) AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	page, err := strconv.Atoi(queryValues.Get("page"))
	if err != nil {
		service.ResponseHelper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: if page number is not in range 1,2,3 then defaut to 3

	// TODO: Create a buffered channel and read and merge data from channel by calling GoRoutine readJson

	// TODO: Append data to result variable

	service.ResponseHelper.RespondWithJSON(w, http.StatusOK, result)
}

func readJson(fileNumber int, ch chan []models.Movie) {
	fileName := "json\\movies" + strconv.Itoa(fileNumber) + ".json"

	// TODO: read data from file and then parse it then pass it back to channel
}
