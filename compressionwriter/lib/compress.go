package lib

import (
	"compress/gzip"
	"io"
	"log"
)

type CompressReader struct {
	underlying io.Reader
	compressor *gzip.Writer
	exposed    io.Reader
}

func NewCompressReader(r io.Reader) CompressReader {
	log.Print("creating a loopless Compressor!")
	retval := CompressReader{
		underlying: r,
	}
	pr, pw := io.Pipe()
	retval.compressor = gzip.NewWriter(pw)
	retval.exposed = pr
	go func() {
		log.Print("launching copy")
		n, err := io.Copy(retval.compressor, retval.underlying)
		log.Printf("copied %d bytes (err = %v)", n, err)
		log.Printf("closing: gzip.Writer.Close()")
		retval.compressor.Close()
		pw.Close()
	}()
	return retval
}

func (c *CompressReader) Read(b []byte) (int, error) {
	n, err := c.exposed.Read(b)
	log.Printf("Read with %d bytes and error: %v", n, err)
	return n, err
}
