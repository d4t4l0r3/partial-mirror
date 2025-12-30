package dbparser

import "io"

type DbParser interface {
	ReadFile(reader io.Reader) error
	GetNewestPackageVersion(packageName string) (string, error)
}
