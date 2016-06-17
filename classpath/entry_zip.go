package classpath

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"io/ioutil"
)

type ZipEntry struct {
	absPath   string
	zipReader *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if e.zipReader == nil {
		err := e.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := e.findClass(className)
	if classFile == nil {
		return nil, nil, fmt.Errorf("class not found %s", className)
	}

	data, err := readClass(classFile)
	return data, e, err
}

func (e *ZipEntry) String() string {
	return e.absPath
}

func (e *ZipEntry) openJar() error {
	reader, err := zip.OpenReader(e.absPath)
	if err != nil {
		return err
	}
	e.zipReader = reader
	return nil
}

func (e *ZipEntry) findClass(className string) *zip.File {
	for _, f := range e.zipReader.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}
