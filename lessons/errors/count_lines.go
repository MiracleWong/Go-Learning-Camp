package main

import (
	"bufio"
	"io"
)

func CountLines(reader io.Reader) (int, error) {
	var (
		br   = bufio.NewReader(reader)
		line = 0
		err  error
	)

	for {
		_, err = br.ReadString('\n')
		// 这里注意：要先line++再处理err，否则会少加行数
		line++
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		return 0, err
	}
	return line, nil
}

func ScannerLines(reader io.Reader) (int, error) {
	line := 0
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line++
	}

	return line, scanner.Err()
}

