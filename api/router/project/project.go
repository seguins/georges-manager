package project

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/seguins/georges-manager/georges/project"
	"net/http"
)

func startService(response http.ResponseWriter, req *http.Request) {
	p, err := project.CreateNewProject("test", req.Body)
	if err != nil {
		log.Errorln(err)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := p.Start(); err != nil {
		log.Errorln(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Write([]byte("DONE"))
}

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/start-service", startService).Methods("POST")
}
