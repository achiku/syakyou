package main

import (
	"io"

	"github.com/pkg/errors"
)

func write(w io.Writer, src []byte) error {
	if _, err := w.Write(src); err != nil {
		return errors.Wrap(err, "failed to write file")
	}
	return nil
}
