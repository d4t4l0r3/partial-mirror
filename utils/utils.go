package utils

import (
	"bytes"
	"fmt"
	"io"
)

func SeekAhead(reader io.Reader, numBytes int) error {
	buffer := make([]byte, numBytes)

	n, err := reader.Read(buffer)
	if err != nil {
		return err
	}
	if n != numBytes {
		return fmt.Errorf("Tried to seek %v bytes, could only read %v", numBytes, n)
	}
	return nil
}

func ReadToString(reader io.Reader, numBytes int) (string, error) {
	buffer := make([]byte, numBytes)

	n, err := reader.Read(buffer)
	if err != nil {
		return "", err
	}
	if n != numBytes {
		return "", fmt.Errorf("Tried to read %v bytes, could only read %v", numBytes, n)
	}
	eol := bytes.IndexByte(buffer, 0x00)
	if eol != -1 {
		return string(buffer[:eol]), nil
	} else {
		return string(buffer), nil
	}
}
