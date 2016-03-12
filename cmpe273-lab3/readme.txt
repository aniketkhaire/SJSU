To execute the program :

1. go run lab3server.go			//now the server listens at ports 3000, 3001, 3002
2. go run lab3client.go			//the client program automatically PUT's 10 key/value pairs and displays the response codes. Also, the client program then fetches all keys one by one and displays their responses

3. curl localhost:3000/keys 	//returns all keys stored at server localhost:3000
4. curl localhost:3001/keys 	//returns all keys stored at server localhost:3001
5. curl localhost:3002/keys 	//returns all keys stored at server localhost:3002
6. curl localhost:3000/keys/id  //if present, returns key-value from server else returns NULL