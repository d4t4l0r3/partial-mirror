package dbparser

import (
	"compress/gzip"
	"fmt"
	"io"
	"strings"

	"github.com/d4t4l0r3/partial-mirror/utils"
	"github.com/charmbracelet/log"
)

type AlpmParser struct {
	Packages map[string]string
}

func NewAlpmParser() (parser AlpmParser) {
	parser.Packages = make(map[string]string)
	return
}

func (parser *AlpmParser) ReadFile(reader io.Reader) error {
	decompressor, err := gzip.NewReader(reader)
	if err != nil {
		log.Error("Failed to decompress ALPM-DB", "error", err)
		return err
	}

	for {
		err = utils.SeekAhead(decompressor, 0x400)
		if err != nil {
			log.Error("Failed to seek ahead decompressor", "error", err)
			return err
		}

		info, err := utils.ReadToString(decompressor, 0x400)
		if err != nil {
			log.Error("Failed to read from decompressor", "error", err)
		}
		packageName, version, err := ParseInfo(info)
		if err != nil {
			log.Error("Failed to parse package info", "info", info, "error", err)
			return err
		}
		log.Debug("Found package", "name", packageName, "version", version)
		parser.Packages[packageName] = version
	}
}

func (parser AlpmParser) GetNewestPackageVersion(packageName string) (string, error) {
	version, exists := parser.Packages[packageName]
	if exists {
		return version, nil
	} else {
		return "", fmt.Errorf("Package %v does not exist", packageName)
	}
}

func ParseInfo(info string) (packageName, version string, err error) {
	// find %NAME% field
	err = parseSeek(&info, "%NAME%")
	if err != nil {
		return
	}
	// skip to next line
	info = info[7:]

	// find end of name
	i := strings.IndexByte(info, 0x0a)
	if i == -1 {
		err = fmt.Errorf("Failed to find line break")
		return
	}
	packageName = info[:i]
	info = info[i:]

	// find %VERSION% field
	err = parseSeek(&info, "%VERSION%")
	if err != nil {
		return
	}
	// skip to next line
	info = info[10:]

	// find end of version
	i = strings.IndexByte(info, 0x0a)
	if i == -1 {
		err = fmt.Errorf("Failed to find line break")
		return
	}
	version = info[:i]
	return
}

func parseSeek(str *string, seek string) error {
	i := strings.Index(*str, seek)
	if i == -1 {
		return fmt.Errorf("Failed to find \"%v\"", seek)
	}
	*str = (*str)[i:]
	return nil
}
