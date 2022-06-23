package main

import (
	"b64headerdecoder/pkg/lineprocessor"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputReader *os.File
	var outputWriter *os.File

	outputWriter, inputReader = defineIO()
	df := func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
	defer df(outputWriter)
	defer df(inputReader)

	scanner := bufio.NewScanner(inputReader)
	lp := lineprocessor.NewProcessor()
	for scanner.Scan() {
		str := scanner.Text()
		res, ready := lp.ProcessLine(str)
		if ready {
			_, err := fmt.Fprintln(outputWriter, res)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err := scanner.Err()
	if err == nil {
		log.Println("reached end of file")
	} else {
		log.Fatal("")
	}
}

func defineIO() (*os.File, *os.File) {
	var inputReader *os.File
	var outputWriter *os.File

	if len(os.Args) >= 2 && len(os.Args[1]) != 0 {
		inputFileName := os.Args[1]
		log.Println("input file name: " + inputFileName)

		var err error
		inputReader, err = os.Open(inputFileName)
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) >= 3 && len(os.Args[2]) != 0 {
			outputFilename := os.Args[2]
			log.Println("output file name: ", outputFilename)
			outputWriter, err = os.Create(outputFilename)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			outputWriter = os.Stdout
		}
		return outputWriter, inputReader
	} else {
		inputReader = os.Stdin
		outputWriter = os.Stdout
	}
	return outputWriter, inputReader
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
