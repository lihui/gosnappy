package snappy

import "testing"
import "bytes"
import "io/ioutil"
import "io"

var s = "==hello world,hello world,hello world=="

func TestSnappyIO(t *testing.T) {
	var wbytes bytes.Buffer
	w := NewWriter(&wbytes)

	if _, err := io.Copy(w, bytes.NewBufferString(s)); err != nil {
		t.Error(err)
	}
	if err := w.Close(); err != nil {
		t.Error(err)
	}

	r, err := NewReader(&wbytes)
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		t.Error(err)
	}

	if string(bs) != s {
		t.Error("compress-decompress cycle not matched")
	}

}
func TestSnappBytes(t *testing.T) {

	bytes := Compress(nil, []byte(s))
	bytes, err := Decompress(nil, bytes)
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != s {
		t.Error("compress-decompress cycle not matched")
	}

	bytes, err = Decompress(nil, []byte(s))
	if err == nil {
		t.Error(bytes, err)
	}

}
