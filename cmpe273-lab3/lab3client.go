package main

import (
	"fmt"
	"net/http"
	"bytes"
	"strconv"
	"io/ioutil"
//	"crypto/md5"
//	"encoding/binary"
	"hash/crc32"
)

type KeyValuePair struct{
	Key int `json:"key"`
	Value string `json:"value"`
}

var InputKeys[] KeyValuePair

type Server struct{
	name string
	hash uint32
}

type HashingRing struct{
	circle []Server
}

func (obj *HashingRing) add(keyvaluepair KeyValuePair) {
	for i:=0; i<3;i++ {
		key:=strconv.Itoa(keyvaluepair.Key)
		if obj.circle[i].hash > getHashValue(key) {
			callPost(obj.circle[i].name, keyvaluepair)
			return
		}
	}
	callPost(obj.circle[0].name, keyvaluepair)
}

func (obj *HashingRing) get(keyvaluepair KeyValuePair) {
	for i:=0; i<3;i++ {
		key:=strconv.Itoa(keyvaluepair.Key)
		if obj.circle[i].hash > getHashValue(key) {
			callGet(obj.circle[i].name, keyvaluepair)
			return
		}
	}
	callGet(obj.circle[0].name, keyvaluepair)
}

func (obj *HashingRing) addServer(portNumber string) {
	var myserver Server
	myserver.hash = getHashValue(portNumber)
	myserver.name = portNumber
	obj.circle = append(obj.circle, myserver)
}

func getHashValue(hashInput string) uint32 {
	hash := crc32.ChecksumIEEE
	hashing := hash([]byte(hashInput))
	return hashing
}

func callPost(myServer string, keyvaluepair KeyValuePair){
	
	client:=&http.Client{}
	
	hashValue := strconv.Itoa(keyvaluepair.Key)
	
	urlStr:="http://localhost:"+myServer+"/keys/"+hashValue+"/"+keyvaluepair.Value

	jsonStr := []byte(`{}`)
	req, _ := http.NewRequest("PUT", urlStr, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    res, err := client.Do(req)
    if err!=nil{
    	panic(err)
    }
    fmt.Println("Response Code: ", res.StatusCode)
   	defer res.Body.Close()
}

func callGet(myServer string, keyvaluepair KeyValuePair) {
	key:=strconv.Itoa(keyvaluepair.Key)
	//generate url
	urlStr:="http://localhost:"+myServer+"/keys/"+key
	//fmt.Printf("GET %s\n", urlStr)
	res, _:=http.Get(urlStr)
		defer res.Body.Close()

	data,_:=ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", string(data))
}

func main(){

	var obj HashingRing

	obj.addServer("3000")
	obj.addServer("3001")
	obj.addServer("3002")

	tempArray := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	//initializing the KeyValue array
	for index:=0;index<10;index++{
		InputKeys = append(InputKeys, KeyValuePair{})
		InputKeys[index].Key = index+1
		InputKeys[index].Value = tempArray[index]
	}

	//fmt.Println(InputKeys)

	//adding all key-value pairs 
	for index:=0;index<10;index++{
		obj.add(InputKeys[index])
	}

	//fetching all key-value pairs through key
	for index:=0;index<10;index++{
		obj.get(InputKeys[index])
	}

	//fmt.Println(obj.circle)
}