package main

import "C"
import (
	"archive/zip"
	"io"
	"log"
	"os"
)

//export Zip
func Zip(zipname *C.char, filenames []*C.char) int32 {
	f, err := os.Create(C.GoString(zipname))
	if err != nil {
		log.Println(err)
		return -1
	}
	defer f.Close()

	w := zip.NewWriter(f)
	defer w.Close()

	for _, name := range filenames {
		ww, err := w.Create(C.GoString(name))
		if err != nil {
			log.Println(err)
			return -1
		}

		fi, err := os.Open(C.GoString(name))
		if err != nil {
			log.Println(err)
			return -1
		}
		defer fi.Close()

		if _, err := io.Copy(ww, fi); err != nil {
			log.Println(err)
			return -1
		}
	}

	return 0
}

//export Unzip
func Unzip(zipname *C.char) int32 {
	r, err := zip.OpenReader(C.GoString(zipname))
	if err != nil {
		log.Println(err)
		return -1
	}
	defer r.Close()

	// Iterate through the files in the archive,
	for _, file := range r.File {
		ri, err := file.Open()
		if err != nil {
			log.Println(err)
			return -1
		}
		defer ri.Close()

		wi, err := os.Create(file.Name)
		if err != nil {
			log.Println(err)
			return -1
		}
		defer wi.Close()

		_, err = io.Copy(wi, ri)
		if err != nil {
			log.Println(err)
			return -1
		}
	}

	return 0
}

func main() {}
