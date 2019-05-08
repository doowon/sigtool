package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/doowon/sigtool"
)

func main() {
	inParam := flag.String("in", "filename", "This specifies the input Signed PE filename to read from")
	outParam := flag.String("out", "filename", "This specifies the output PKCS#7 filename to write to")

	flag.Parse()
	if *inParam == "filename" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	buf, err := sigtool.ExtractDigitalSignature(*inParam)
	if err != nil {
		log.Fatal(err)
	}

	if *outParam != "filename" {
		ioutil.WriteFile(*outParam, buf, 0644)
	} else {
		fileName := filepath.Base(*inParam)
		if fileName == "." {
			fmt.Println("Input file path is not correct.")
			os.Exit(1)
		}
		ioutil.WriteFile(fileName+".pkcs7", buf, 0644)
	}
}
