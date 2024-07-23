package image

import "io"

type ImageRepository interface {
	AddOne(filename string, fileSize int64, reader io.Reader) (string, error)
}
