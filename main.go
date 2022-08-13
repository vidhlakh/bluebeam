package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vidhlakh/bluebeam/interleave"
)

func main() {
	var file1, file2, delimiter string
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("err getiing current dir", err)
	}
	fmt.Println("Get pwd", pwd)
	flag.StringVar(&file1, "f1", pwd+"/inputs/file1.txt", "Enter first file name")
	flag.StringVar(&file2, "f2", pwd+"/inputs/file2.txt", "Enter second file name")
	flag.StringVar(&delimiter, "d", " ", "Enter delimiter for the words")

	flag.Parse()

	res, err := interleave.InterleaveFiles(file1, file2, delimiter)
	if err != nil {
		log.Fatalln("Err in interleaving", err)
	}
	log.Println("Interleaved strings from both files\n", res)
}
