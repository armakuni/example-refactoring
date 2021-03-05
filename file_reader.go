package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type FileReader struct {
	filename string
}

func NewFileReader(filename string) *FileReader {
	return &FileReader{filename}
}

func (f *FileReader) Paginate(handler func(int, string)) (Paginated, error) {
	paginated := Paginated{}
	count := 0

	reader, err := os.Open(f.filename)
	if err != nil {
		return paginated, err
	}

	defer reader.Close()

	breader := bufio.NewReader(reader)

	for paginated.Total = 0; ; paginated.Total++ {
		line, err := breader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return paginated, err
		}
		if count < paginated.PageSize {
			handler(paginated.Total, strings.TrimSpace(line))
			count++
		}
	}

	return paginated, nil
}
