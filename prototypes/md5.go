package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	/* Args control */
	if len(os.Args) != 2 {
		fmt.Println("Usage : md5 FILE")
		os.Exit(1)
	}

	/* Open the file */
	file	:= os.Args[1]
	f, err  := os.Open(file)

	/* Test error ? */
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}

	/* Computing the md5 */
	hash := md5.New()
	io.Copy(hash, f)

	/* Display the md5 sum */
	fmt.Printf("%x  %s\n", hash.Sum(nil),file)

	os.Exit(0)
}
