package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
    "strconv"
    "net/url"
    "math/rand"
    "time"
)

//Structure for request json
type Request struct{
    Name   string `json:"name" bson:"name"`
    Address string `json:"address" bson:"address"`
    City string `json:"city" bson:"city"`   
    State string `json:"state" bson:"state"`
    Zip   string    `json:"zip" bson:"zip"`
}
//Structure for response json
type Response struct{
    Id int `json:"id" "bson":"id"`
    Name   string `json:"name" bson:"name"`
    Address string `json:"address" bson:"address"`
    City string `json:"city" bson:"city"`   
    State string `json:"state" bson:"state"`
    Zip    string    `json:"zip" bson:"zip"`
    Coordinate interface{} `json:"coordinate" bson:"coordinate"`
}

var mgoSession *mgo.Session 

func myGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    
    id := p.ByName("id")
    num, _:=strconv.Atoi(id)

    mgoSession, err := mgo.Dial("mongodb://user:user@ds043962.mongolab.com:43962/gorest");
    if err!=nil{
        panic(err)
    }

    res := Response{}

    if err := mgoSession.DB("gorest").C("user").Find(bson.M{"id":num}).One(&res); err!=nil{
        w.WriteHeader(404)
        return
    }

    data, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", data)    

}

func myPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {  
    req := Request{}
    res := Response{}

    json.NewDecoder(r.Body).Decode(&req)

    //the magic happens here
    address :=req.Address+" "+req.City+" "+req.State+" "+req.Zip
    responseCoordinate := getLatLng(address)
    
    //fmt.Println(responseCoordinate)
    mgoSession, err := mgo.Dial("mongodb://user:user@ds043962.mongolab.com:43962/gorest");
    if err!=nil{
        panic(err)
    }
    rand.Seed( time.Now().UTC().UnixNano())
    myRnd := rand.Intn(10000)
    //idCounter += 1
    res.Id = rand.Intn(myRnd)
    res.Name = req.Name
    res.Address = req.Address
    res.City = req.City
    res.State = req.State
    res.Zip = req.Zip
    res.Coordinate = responseCoordinate
    
    mgoSession.DB("gorest").C("user").Insert(res)

    data, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", data)
    
}

func myDelete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    id := p.ByName("id")
    num, _:= strconv.Atoi(id)
    //fmt.Println(num)

    mgoSession, err := mgo.Dial("mongodb://user:user@ds043962.mongolab.com:43962/gorest");
    if err!=nil{
        panic(err)
    }
    
    if err := mgoSession.DB("gorest").C("user").Remove(bson.M{"id":num}); err != nil {
        w.WriteHeader(404)
        return
    }
    w.WriteHeader(200)
}

func myPut(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    id := p.ByName("id")
    num, _:=strconv.Atoi(id)

    mgoSession, err := mgo.Dial("mongodb://user:user@ds043962.mongolab.com:43962/gorest");
    if err!=nil{
        panic(err)
    }

    res := Response{}
    req := Request{}

    json.NewDecoder(r.Body).Decode(&req)

    if err := mgoSession.DB("gorest").C("user").Find(bson.M{"id":num}).One(&res); err!=nil{
        w.WriteHeader(404)
        return
    }
    
    res.Id = num
    
    if req.Name != ""{
        res.Name = req.Name
    }
    if req.Address != ""{
        res.Address = req.Address
    }
    if req.City != ""{
        res.City = req.City
    }
    if req.State != ""{
        res.State = req.State
    }
    if req.Zip != ""{
        res.Zip = req.Zip
    }

    address :=res.Address+" "+res.City+" "+res.State+" "+res.Zip
    res.Coordinate = getLatLng(address)

    if err := mgoSession.DB("gorest").C("user").Update(bson.M{"id":num}, res); err != nil {
        w.WriteHeader(404)
        return
    }

    data, _ := json.Marshal(res)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", data)
}

func getLatLng(address string) (interface{}){
    var Url *url.URL
    Url, err := url.Parse("http://maps.google.com")
    if err != nil {
      panic("Error Panic")
    }

    Url.Path += "/maps/api/geocode/json"
    //fmt.Println(address)
    parameters := url.Values{}
    parameters.Add("address", address)
    Url.RawQuery = parameters.Encode()
    Url.RawQuery += "&sensor=false"

    res, err := http.Get(Url.String())
    
    if err != nil {
      panic("Error Panic")
    }
  
    defer res.Body.Close()
    var v map[string] interface{}
    dec:= json.NewDecoder(res.Body);
    if err := dec.Decode(&v); err != nil {
      fmt.Println("ERROR: " + err.Error())
    }   

    myRes := v["results"].([]interface{})[0].(map[string] interface{})["geometry"].(map[string] interface{})["location"]
    
    return myRes
}


func main() {
    mux := httprouter.New()
    mux.GET("/locations/:id", myGet)
    mux.POST("/locations", myPost)
    mux.DELETE("/locations/:id", myDelete)
    mux.PUT("/locations/:id", myPut)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}

