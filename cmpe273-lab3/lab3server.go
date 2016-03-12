package main

import (
	"fmt"
	"net/http"
//	"bytes"
	"strconv"
//	"io/ioutil"
//	"crypto/md5"
//	"encoding/binary"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type KeyValuePair struct{
	Key int `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	KeyValuePairArray []KeyValuePair `json:"response"`
}

var Server1Response Response
var Server2Response Response
var Server3Response Response

func getAll3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, _ := json.MarshalIndent(Server1Response, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func getAll3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, _ := json.MarshalIndent(Server2Response, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func getAll3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, _ := json.MarshalIndent(Server3Response, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func get3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, _ := strconv.Atoi(p.ByName("id"))
	keyvaluepair := KeyValuePair{}
	i := 0
	for i = 0; i < len(Server1Response.KeyValuePairArray); i++ {
		if Server1Response.KeyValuePairArray[i].Key == id {
			keyvaluepair = Server1Response.KeyValuePairArray[i]
			break
		}
	}
	var getResponse Response
	if i != len(Server1Response.KeyValuePairArray) {
		getResponse.KeyValuePairArray = append(getResponse.KeyValuePairArray, keyvaluepair)
	}
	
	data, _ := json.MarshalIndent(getResponse, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func get3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, _ := strconv.Atoi(p.ByName("id"))
	keyvaluepair := KeyValuePair{}
	i := 0
	for i = 0; i < len(Server2Response.KeyValuePairArray); i++ {
		if Server2Response.KeyValuePairArray[i].Key == id {
			keyvaluepair = Server2Response.KeyValuePairArray[i]
			break
		}
	}
	var getResponse Response
	if i != len(Server2Response.KeyValuePairArray) {
		getResponse.KeyValuePairArray = append(getResponse.KeyValuePairArray, keyvaluepair)
	}
	
	data, _ := json.MarshalIndent(getResponse, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func get3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, _ := strconv.Atoi(p.ByName("id"))
	keyvaluepair := KeyValuePair{}
	i := 0
	for i = 0; i < len(Server3Response.KeyValuePairArray); i++ {
		if Server3Response.KeyValuePairArray[i].Key == id {
			keyvaluepair = Server3Response.KeyValuePairArray[i]
			break
		}
	}
	var getResponse Response
	if i != len(Server3Response.KeyValuePairArray) {
		getResponse.KeyValuePairArray = append(getResponse.KeyValuePairArray, keyvaluepair)
	}
	
	data, _ := json.MarshalIndent(getResponse, "", "\t")
	
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func put3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	
	temp := KeyValuePair{}
	temp.Key = Id
	temp.Value = value
	
	Server1Response.KeyValuePairArray = append(Server1Response.KeyValuePairArray, temp)
	
	data, _ := json.MarshalIndent(Server1Response, "", "\t")
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func put3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	
	temp := KeyValuePair{}
	temp.Key = Id
	temp.Value = value
	
	Server2Response.KeyValuePairArray = append(Server2Response.KeyValuePairArray, temp)
	
	data, _ := json.MarshalIndent(Server2Response, "", "\t")
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func put3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	
	temp := KeyValuePair{}
	temp.Key = Id
	temp.Value = value
	
	Server3Response.KeyValuePairArray = append(Server3Response.KeyValuePairArray, temp)
	
	data, _ := json.MarshalIndent(Server3Response, "", "\t")
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", data)
}

func main() {
	mux1 := httprouter.New()
	mux2 := httprouter.New()
	mux3 := httprouter.New()
	
	mux1.GET("/keys", getAll3000)
	mux2.GET("/keys", getAll3001)
	mux3.GET("/keys", getAll3002)

	mux1.GET("/keys/:id", get3000)
	mux2.GET("/keys/:id", get3001)
	mux3.GET("/keys/:id", get3002)
	
	mux1.PUT("/keys/:id/:value", put3000)
	mux2.PUT("/keys/:id/:value", put3001)
	mux3.PUT("/keys/:id/:value", put3002)
	
	go func() {
		server1 := http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: mux1,
		}
		server1.ListenAndServe()
	}()
	go func() {
		server2 := http.Server{
			Addr:    "0.0.0.0:3001",
			Handler: mux2,
		}
		server2.ListenAndServe()
	}()
		server3 := http.Server{
			Addr:    "0.0.0.0:3002",
			Handler: mux3,
		}
		server3.ListenAndServe()
}
