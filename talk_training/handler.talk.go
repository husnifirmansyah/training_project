package talk_training

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.api+json")
	response := GetTalks(ps.ByName("product_id"))
	json.NewEncoder(w).Encode(response)
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	log.Printf(ps.ByName("product_id"))
	log.Printf(ps.ByName("user_id"))
	log.Printf(ps.ByName("message"))
	//InsertTalks(ps.ByName("product_id"), ps.ByName("user_id"), ps.ByName("message"))
	//json.NewEncoder(w).Encode(response)
}
