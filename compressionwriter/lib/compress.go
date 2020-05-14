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
	log.Print("creating a new Compressor!")
	retval := CompressReader{
		underlying: r,
	}
	pr, pw := io.Pipe()
	retval.compressor = gzip.NewWriter(pw)
	retval.exposed = pr
	go func() {
		log.Print("launching copy")
		for {
			log.Printf("trying to copy")
			n, err := io.Copy(retval.compressor, retval.underlying)
			log.Printf("copied %d bytes", n)
			// this is wrong ... Copy does not return EOF
			if err == io.EOF || n == 0 {
				break
			}
		}
		log.Printf("closing: gzip.Writer.Close()")
		retval.compressor.Close()
		pw.Close()
	}()
	return retval
}

func (c *CompressReader) Read(b []byte) (int, error) {
	return c.exposed.Read(b)
}
