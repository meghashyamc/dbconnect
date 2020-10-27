package dbapi

import (
	"encoding/json"
	"net/http"

	redis "github.com/meghashyamc/dbconnect/redisadapter"
)

func Add(w http.ResponseWriter, req *http.Request) {

	dbpayload := Dbpayload{}
	if err := json.NewDecoder(req.Body).Decode(&dbpayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorStatus + ": " + err.Error()))
		return
	}

	if _, err := redis.HSet(dbpayload.Key, dbpayload.Field, dbpayload.Value); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errorStatus + ": " + err.Error()))
		return
	}

	w.Write([]byte(successStatus))

}

func Get(w http.ResponseWriter, req *http.Request) {

	dbpayload := Dbpayload{}
	if err := json.NewDecoder(req.Body).Decode(&dbpayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errorStatus + ": " + err.Error()))
		return
	}
	val, err := redis.HGet(dbpayload.Key, dbpayload.Field)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errorStatus + ": required value not found"))
		return
	}

	w.Write([]byte(val))
}
