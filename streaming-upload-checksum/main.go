package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type CountingWriter struct {
	W io.Writer
	N int64
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.W.Write(p)
	c.N += int64(n)
	return n, err
}

const maxUploadBytes = 1 << 20 // 1MB

var ErrTooLarge = errors.New("upload exceeds size limit")

func SaveUpload(dst io.Writer, src io.Reader) (checksum string, size int64, err error) {
	hasher := sha256.New()
	limited := io.LimitReader(src, maxUploadBytes+1)

	tee := io.TeeReader(limited, hasher)

	counter := &CountingWriter{W: dst}
	if _, err := io.Copy(counter, tee); err != nil {
		return "", 0, fmt.Errorf("copy upload : %w", err)
	}

	if counter.N > maxUploadBytes {
		return "", 0, ErrTooLarge
	}

	return hex.EncodeToString(hasher.Sum(nil)), counter.N, nil
}

func main() {
	body := strings.NewReader("invoice-2026-09.pdf contents...")
	var storage bytes.Buffer

	sum, n, err := SaveUpload(&storage, body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("stored %d bytes, sha256=%s\n", n, sum)
	fmt.Printf("storage now holds: %q\n", storage.String())

	huge := bytes.NewReader(make([]byte, maxUploadBytes+10))

	_, _, err = SaveUpload(io.MultiWriter(&bytes.Buffer{}), huge)
	fmt.Println("large upload:", err)
}
