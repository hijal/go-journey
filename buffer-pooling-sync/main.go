package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferpool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func renderResponse(message string) string {
	buf := bufferpool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferpool.Put(buf)

	buf.WriteString(`{"message":"`)
	buf.WriteString(message)
	buf.WriteString(`"}`)
	return buf.String()
}

func main() {
	for i := 0; i < 3; i++ {
		response := renderResponse(fmt.Sprintf("request #%d handled", i))
		fmt.Println(response)
	}
}
