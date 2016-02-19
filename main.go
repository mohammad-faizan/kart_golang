package main

import(
	"simple-server/server"
	"simple-server/db"
)

func main(){
	db, err := db.NewDbConnection()

	if err != nil {
		panic(err)
		return
	}

	server := server.NewServer(db)
	server.Run(":3000")
}