package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("uuid2base64 uuid_string")
		os.Exit(1)
	}
	uuidStr := os.Args[1]
	uid, err := uuid.Parse(uuidStr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	data, err := uid.MarshalBinary()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
}
