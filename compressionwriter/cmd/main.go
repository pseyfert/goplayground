package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	mygzip "github.com/pseyfert/playground/compressionwriter/lib"
)

func main() {
	r := strings.NewReader("hai there, this text should be compressed")

	// b := make([]byte, 128)
	// n, err := r.Read(b)
	// fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
	// fmt.Printf("b[:n] = %q\n", b[:n])
	// n, err = r.Read(b)
	// fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
	// fmt.Printf("b[:n] = %q\n", b[:n])

	// var buf bytes.Buffer
	// zw := gzip.NewWriter(&buf)
	//
	// n, err := zw.Write([]byte("A long time ago in a galaxy far, far away ..."))
	// fmt.Printf("n = %v, err = %v, b = %v\n", n, err, buf.Bytes())
	// fmt.Printf("b[:n] = %q\n", buf.Bytes()[:n])
	// n, err = zw.Write([]byte("you are our only hope"))
	// fmt.Printf("n = %v, err = %v, b = %v\n", n, err, buf.Bytes())
	// fmt.Printf("b[:n] = %q\n", buf.Bytes()[:n])
	//
	// zr, err := gzip.NewReader(&buf)
	// if err != nil {
	// 	log.Fatalf("can't create reader %v", err)
	// }
	// if n, err := io.Copy(os.Stdout, zr); err != nil {
	// 	log.Fatalf("can't copy %v", err)
	// } else {
	// 	fmt.Printf("these have been n = %v\n", n)
	// }

	compressed := mygzip.NewCompressReader(r)

	b := make([]byte, 128)
	var bigbuffer bytes.Buffer
	for {
		log.Printf("trying to read")
		n, err := compressed.Read(b)

		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

		bigbuffer.Write(b[:n])
		hmpf := bytes.NewReader(bigbuffer.Bytes())
		validator, err := gzip.NewReader(hmpf)
		if err != nil {
			panic(err)
		}
		fmt.Println("can read back the following:")
		io.Copy(os.Stdout, validator)
		fmt.Println("EOM")
	}
}
