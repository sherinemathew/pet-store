package service

  import (
	"github.com/sherinemathew/petStore/petService/dbClient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
  )

  var DBClient dbclient.IBoltClient

func GetPet(w http.ResponseWriter, r *http.Request) {

	// Read the 'petId' path parameter from the route
	var petId = mux.Vars(r)["petId"]

    // Read the pet struct BoltDB
	pet, err := DBClient.QueryPet(petId)

        // If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

    // If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(pet)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}