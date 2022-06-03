package controller

import (
	"bufio"
	"os"
)

func deleteLastLine(f *os.File) (err error) {
	delim := "<"
	var offset int64
	r := bufio.NewReader(f)
	for {
		offset--
		_, err := f.Seek(offset, os.SEEK_END)
		if err != nil {
			return err
			break
		}
		r.Reset(f)
		b, err := r.Peek(1)
		if err != nil {
			return err
			break
		}
		if string(b) == delim {
			break
		}
	}
	info, err := f.Stat()
	if err != nil {
		return err
	}
	err = f.Truncate(info.Size() + offset)
	if err != nil {
		return err
	}
	return nil
}

func addLine(f *os.File, str []string) (err error) {
	for _, s := range str {
		_, err := f.WriteString(s + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
