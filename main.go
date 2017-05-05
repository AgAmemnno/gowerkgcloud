package main

import (
	//"fmt"
	"log"
	"net/http"
       "github.com/AgAmemnno/gowerkdoci/soio"
	
       

)




func main() {

    

	server, err := soio.Soio()

        if err != nil {

		log.Fatal(err)
	}
    
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:3641...")
	log.Fatal(http.ListenAndServe(":3641", nil))
}
