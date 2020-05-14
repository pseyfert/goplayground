package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	mygzip "github.com/pseyfert/goplayground/compressionwriter/lib"
)

func main() {
	r := strings.NewReader("hai there, this text should be compressed")

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
