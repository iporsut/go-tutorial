package main

import (
	"log"

	"github.com/globalsign/mgo"
)

type Employee struct {
	ID        int    `bson:"employee_id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
}

func main() {
	// Connect Server
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Select DB and Collection
	c := session.DB("store").C("employees")

	// Insert Document
	// emp := &Employee{
	// 	ID:        1002,
	// 	FirstName: "Fay",
	// 	LastName:  "Van Damme",
	// 	Email:     "Fay.Van Damme@testing.com",
	// }

	// err = c.Insert(emp)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Find All
	// var emps []Employee

	// err = c.Find(nil).All(&emps)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(emps)

	// Find By ID
	// var emp Employee
	// err = c.Find(bson.M{
	// 	"employee_id": 1002,
	// }).One(&emp)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(emp)

	// Update by ID
	// err = c.Update(bson.M{
	// 	"employee_id": 1002,
	// }, bson.M{
	// 	"$set": bson.M{
	// 		"first_name": "John",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Delete by ID
	// err = c.Remove(bson.M{
	// 	"employee_id": 1002,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
