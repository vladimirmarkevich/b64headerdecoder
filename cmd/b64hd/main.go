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

	if checkForHelpArg() {
		printUsage()
		os.Exit(0)
	}
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
	if err != nil {
		log.Fatal("")
	}
}

func checkForHelpArg() bool {
	for _, v := range os.Args {
		if v == "-h" || v == "--help" {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("Usage: ", os.Args[0], " [-h, --help] [inputfile] [outputfile]")
	fmt.Println("Reads input stream, finds 'X-ImunifyEmail-Filter-Info:' following base64 encoded block. " +
		"Decode the block to output stream and print other data without changes ")
}

func defineIO() (*os.File, *os.File) {
	var inputReader *os.File
	var outputWriter *os.File
	var err error

	if len(os.Args) >= 2 && len(os.Args[1]) != 0 {
		inputFileName := os.Args[1]
		//log.Println("input file name: " + inputFileName)

		inputReader, err = os.Open(inputFileName)
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) >= 3 && len(os.Args[2]) != 0 {
			outputFilename := os.Args[2]
			//log.Println("output file name: ", outputFilename)
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
