package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vidhlakh/bluebeam/interleave"
)

func main() {
	var file1, file2 string
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("err getting current dir", err)
	}
	flag.StringVar(&file1, "f1", pwd+"/inputs/file1.txt", "Enter first file name")
	flag.StringVar(&file2, "f2", pwd+"/inputs/file2.txt", "Enter second file name")

	flag.Parse()

	res, err := interleave.InterleaveFiles(file1, file2)
	if err != nil {
		log.Fatalln("Err in interleaving", err)
	}
	fmt.Println("Interleaved strings from both files")
	for _, out := range res {
		fmt.Printf("%v ", out)
	}
	fmt.Println("")
}
