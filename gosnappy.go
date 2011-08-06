package gosnappy

//#include<snappy-c.h>
//#include<stdlib.h>
import "C"
import "unsafe"

import "io"
import "io/ioutil"
import "os"
import "bytes"

func NewReader(r io.Reader) (io.Reader, os.Error) {
	bs, e := ioutil.ReadAll(r)
	if e != nil {
		return nil, e
	}
	bs, e = Decompress(nil, bs)
	return bytes.NewBuffer(bs), e
}

type writer struct {
	buf  bytes.Buffer
	orig io.Writer
}

func NewWriter(w io.Writer) io.WriteCloser {
	return &writer{orig: w}

}
func (w *writer) Write(p []byte) (n int, err os.Error) {
	n, err = w.buf.Write(p)
	return n, err
}
func (w *writer) Close() os.Error {
	bs := Compress(nil, w.buf.Bytes())
	_, e := w.orig.Write(bs)
	return e
}
func c_pchar(p *byte) *C.char {
	return (*C.char)(unsafe.Pointer(p))
}
func Compress(output []byte, input []byte) []byte {
	size := C.snappy_max_compressed_length(C.size_t(len(input)))
	if C.size_t(cap(output)) < size {
		output = make([]byte, size)
	}
	if C.snappy_compress(c_pchar(&input[0]),
		C.size_t(len(input)),
		c_pchar(&output[0]), &size) != C.SNAPPY_OK {
		panic("Unexpected Error")
	}
	return output[:size]
}
func Decompress(output []byte, input []byte) (out []byte, err os.Error) {
	var size C.size_t
	if C.snappy_uncompressed_length(c_pchar(&input[0]),
		C.size_t(len(input)), &size) != C.SNAPPY_OK {
		return nil, os.NewError("Invalid Input")
	}
	if C.size_t(cap(output)) < size {
		output = make([]byte, size)
	}
	if C.snappy_uncompress(c_pchar(&input[0]),
		C.size_t(len(input)),
		c_pchar(&output[0]), &size) != C.SNAPPY_OK {
		return nil, os.NewError("Invalid Input")
	}
	return output[:size], nil
}
