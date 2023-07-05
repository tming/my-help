package tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func DownloadFile(f string, urlstr string) error {

	if f == "" {
		dir, _ := os.Getwd()
		f = filepath.Join(dir, path.Base(urlstr))
		fmt.Printf("save target file:%s\n", f)
	}

	// Create the file
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	// Get the data
	resp, err := http.Get(urlstr)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
