//go:build windows
// +build windows

package cmd

import (
	"errors"
	"io"
	"net/http"
	"net/url"
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

func downloadDLL(f, link string) error {
	transport := http.DefaultTransport

	httpProxy := os.Getenv("HTTP_PROXY")
	if len(httpProxy) > 0 {
		proxyUrl, err := url.Parse(os.Getenv("HTTP_PROXY"))
		if proxyUrl != nil && err == nil {
			transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		}
	}

	// create a new HTTP client
	client := &http.Client{
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}

	// create a new GET request
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9,zh;q=0.8,zh-CN;q=0.7,zh-TW;q=0.6")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	// send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
