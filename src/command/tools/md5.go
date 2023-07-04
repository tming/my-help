package tools

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

// Info describe the os.FileInfo and handle some actions
type Info struct {
	filePath   string
	LinkTarget string

	// info and err are return from os.Stat
	info os.FileInfo
	err  error
}

// Stat do the os.Stat and return an Info
func Stat(fp string) *Info {
	info, err := os.Stat(fp)
	return &Info{
		filePath: fp,
		info:     info,
		err:      err,
	}
}

// Lstat do the os.Lstat and return an Info
func Lstat(fp string) *Info {
	info, err := os.Lstat(fp)
	return &Info{
		filePath: fp,
		info:     info,
		err:      err,
	}
}

// Md5 return the md5 of this file
func (i *Info) Md5() (string, error) {
	if i.err != nil {
		return "", i.err
	}

	f, err := os.Open(i.filePath)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = f.Close()
	}()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}

	md5string := fmt.Sprintf("%x", md5hash.Sum(nil))
	return md5string, nil
}

func Md5sum(filepath string) (string, error) {
	return Stat(filepath).Md5()
}
