package main

import (
	"b64headerdecoder/pkg/lineprocessor"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	if len(fileName) == 0 {
		log.Fatal("First arg has to be file name")
	} else {
		log.Println("filename: " + fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	lp := lineprocessor.NewProcessor()
	for scanner.Scan() {
		str := scanner.Text()
		res, ready := lp.ProcessLine(str)
		if ready {
			fmt.Println(res)
		}
	}
	err = scanner.Err()
	if err == nil {
		log.Println("reached end of file")
	} else {
		log.Fatal("")
	}
}

/*
file, err := os.OpenFile(
        "test.txt",
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )

    // Open file and create a buffered reader on top
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    bufferedReader := bufio.NewReader(file)
	// Read up to and including delimiter
    // Returns byte slice
    dataBytes, err := bufferedReader.ReadBytes('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read bytes: %s\n", dataBytes)

    // Read up to and including delimiter
    // Returns string
    dataString, err := bufferedReader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read string: %s\n", dataString)

*/
