package main

import (
	"fmt"
	"log"

	"github.com/ratmir9/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("hello.txt", []byte("das"))
	if err != nil {
		log.Fatal(err)
	}
	f, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
