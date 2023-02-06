package tool

import (
	"bytes"
	"github.com/klauspost/compress/gzip"

	"io"
)

// GZipBytes 压缩
func GZipBytes(data []byte) ([]byte, error) {
	var input bytes.Buffer
	g := gzip.NewWriter(&input)
	_, err := g.Write(data)
	if err != nil {
		return nil, err
	}
	err = g.Close()
	if err != nil {
		return nil, err
	}
	return input.Bytes(), nil
}

// UGZipBytes 解压
func UGZipBytes(data []byte) ([]byte, error) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	r, _ := gzip.NewReader(&in)
	err := r.Close()
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
