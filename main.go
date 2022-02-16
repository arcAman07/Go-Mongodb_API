package main

import (
	"MongoDb_SampleCRUD/Users"
	"MongoDb_SampleCRUD/config"
)

func main() {
	coll := config.Configure()
	Users.Insert(coll)
	Users.Find(coll)
	Users.Update(coll)
	Users.Delete(coll) 

	// IMPORTANT TO CLOSE THE CONNECTION AFTER TASKS HAVE BEEN EXECUTED

	// After doing all the CRUD operations create another function to close the mongodb connection.
}
