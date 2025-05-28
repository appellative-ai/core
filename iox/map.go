package iox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	//token = byte(' ')
	eol       = byte('\n')
	comment   = "//"
	delimiter = ":"
	newline   = "\r\n"
)

func ReadMap(uri any) (map[string]string, error) {
	buf, err := ReadFile(uri)
	if err != nil {
		return nil, err
	}
	return ParseMap(buf)
}

func ParseMap(buf []byte) (map[string]string, error) {
	m := make(map[string]string)
	if len(buf) == 0 {
		return m, nil
	}
	r := bytes.NewReader(buf)
	reader := bufio.NewReader(r)
	var line string
	var err error
	for {
		line, err = reader.ReadString('\n')
		k, v, err0 := parseLine(line)
		if err0 != nil {
			return m, err0
		}
		if len(k) > 0 {
			m[k] = v
		}
		if err == io.EOF {
			break
		} else {
			if err != nil {
				break
			}
		}
	}
	return m, nil
}

func parseLine(line string) (string, string, error) {
	if len(line) == 0 {
		return "", "", nil
	}
	line = strings.TrimLeft(line, " ")
	if strings.HasSuffix(line, newline) {
		i := len(line) - len(newline)
		line = line[:i]
	}
	if !valid(line) {
		return "", "", nil
	}
	i := strings.Index(line, delimiter)
	if i == -1 {
		return "", "", fmt.Errorf("invalid argument : line does not contain the ':' delimeter : [%v]", line)
	}
	key := line[:i]
	val := line[i+1:]
	return strings.TrimSpace(key), strings.TrimSpace(val), nil
}

func valid(line string) bool {
	if len(line) == 0 {
		//fmt.Printf("line: %v\n", line)
		return false
	}
	if strings.HasPrefix(line, comment) {
		//fmt.Printf("line: %v\n", "<empty>")
		return false
	}
	return true
}
