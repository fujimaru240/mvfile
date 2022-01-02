package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type FileInfo struct {
	Path               string
	FileName           string
	FileNameWithoutExt string
	Ext                string
	Index              int
}

var (
	fileDetails []FileInfo
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("ERROR - invalid args.")
		return
	}
	dirwalk(args[0])
	for _, f := range fileDetails {
		srcFileName := f.FileName
		if f.Index > 0 {
			srcFileName = f.FileNameWithoutExt + "_" + strconv.Itoa(f.Index) + f.Ext
		}
		dstFilePath := args[1] + "/" + srcFileName
		fileCopy(f.Path, dstFilePath)
	}
}

func dirwalk(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirwalk(filepath.Join(dir, file.Name()))
			continue
		}
		fmt.Printf("dir name : %s\n", filepath.Dir(file.Name()))
		fileName := filepath.Base(file.Name())
		fi := FileInfo{
			Path:               filepath.Join(dir, file.Name()),
			FileName:           fileName,
			FileNameWithoutExt: getFileNameWithoutExt(fileName),
			Ext:                filepath.Ext(fileName),
			Index:              0,
		}
		for _, f := range fileDetails {
			if f.FileName == fi.FileName {
				fi.Index++
			}
		}
		fileDetails = append(fileDetails, fi)
	}
}

func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func fileCopy(srcPath, dstPath string) {
	src, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}
