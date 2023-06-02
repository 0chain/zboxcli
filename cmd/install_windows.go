//go:build windows
// +build windows

package cmd

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func InstallDLLs() {
	pwd, _ := os.Getwd()

	libstdc := filepath.Join(pwd, "libstdc++-6.dll")

	if _, err := os.Stat(libstdc); errors.Is(err, os.ErrNotExist) {
		err = downloadDLL(libstdc, "https://github.com/0chain/blobber/wiki/windows/libstdc++-6.dll")
		if err != nil {
			panic("install: " + libstdc + " " + err.Error())
		}
	}

	libgcc_s_seh := filepath.Join(pwd, "libgcc_s_seh-1.dll")
	if _, err := os.Stat(libgcc_s_seh); errors.Is(err, os.ErrNotExist) {
		err = downloadDLL(libgcc_s_seh, "https://github.com/0chain/blobber/wiki/windows/libgcc_s_seh-1.dll")
		if err != nil {
			panic("install: " + libstdc + " " + err.Error())
		}
	}

	libwinpthread := filepath.Join(pwd, "libwinpthread-1.dll")
	if _, err := os.Stat(libwinpthread); errors.Is(err, os.ErrNotExist) {
		err = downloadDLL(libwinpthread, "https://github.com/0chain/blobber/wiki/windows/libwinpthread-1.dll")
		if err != nil {
			panic("install: " + libstdc + " " + err.Error())
		}
	}

}

func downloadDLL(f, url string) error {

	// create a new HTTP client
	client := &http.Client{}

	// create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check if the response was a redirect
	if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		redirectUrl, err := resp.Location()
		if err != nil {

			return err
		}

		// create a new GET request to follow the redirect
		req.URL = redirectUrl
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
	}

	// Create the file
	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err

}
