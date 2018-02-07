//Demo
//interface inherit interface: combinatorial
package main

import (
	"log"
)

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type StringReader struct {
}

func (sr *StringReader) Read() {
	log.Println("read from stringreader")
}

type StringWriter struct {
}

func (sw *StringWriter) Write() {
	log.Println("write into stringreader")
}

type StringReaderWriter struct {
	*StringReader
	*StringWriter
}

func main() {
	rw := &StringReaderWriter{}
	Read(rw)
	Write(rw)
	ReadWrite(rw)
}

func Read(r Reader) {
	r.Read()
}

func Write(w Writer) {
	w.Write()
}

func ReadWrite(rw ReaderWriter) {
	Read(rw)
	Write(rw)
}
