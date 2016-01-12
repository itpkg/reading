package epub

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

type WalkFunc func(name string) error

func Walk(root string, fn WalkFunc) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			if filepath.Ext(info.Name()) == ".epub" {
				return fn(path)
			}
		}
		return nil
	})
}

func Open(archive string) (*Book, error) {
	var book Book
	target := archive[0 : len(archive)-5]
	if err := Unzip(archive, target); err != nil {
		return nil, err
	}

	if mt, err := readMimeType(target); err == nil {
		book.MimeType = mt
	} else {
		return nil, err
	}
	if ct, err := readContainer(target); err == nil {
		book.Container = ct
	} else {
		return nil, err
	}

	for i, rf := range book.Container.RootFiles {
		if opf, err := readRootFile(target, rf); err == nil {
			book.Container.RootFiles[i].Opf = opf
		} else {
			return nil, err
		}
	}
	return &book, nil
}

func readMimeType(target string) (string, error) {
	mt, err := ioutil.ReadFile(filepath.Join(target, "mimetype"))
	return string(mt), err
}

func Unzip(archive, target string) error {

	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0700); err != nil {
		return err
	}
	for _, file := range reader.File {
		name := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(name, file.Mode())
			continue
		}
		fr, err := file.Open()
		if err != nil {
			return err
		}
		defer fr.Close()

		os.MkdirAll(path.Dir(name), 0700)

		tw, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer tw.Close()

		if _, err := io.Copy(tw, fr); err != nil {
			return err
		}
	}

	return nil
}
