package main

import (
	"fmt"
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type Person struct {
	Name  string
	Phone string
}

//export MONGOHQ_URL=mongodb://user:pass@server.compose.io/db_name

func main() {
	url := os.Getenv("MONGOHQ_URL")
	fmt.Printf("%s ===== \n", url)
	if url == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}
	maxWait := time.Duration(5 * time.Second)

	session, err := mgo.DialWithTimeout(url,maxWait)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer session.Close()

//	session.SetSafe(&mgo.Safe{})
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("fingled_sb").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
