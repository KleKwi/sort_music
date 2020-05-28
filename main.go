package main

import (
	"log"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

func main() {
	src := os.Args[1]+"/"
	out := os.Args[2]+"/"
	list := getMusic(os.Args[1])
	for _, file := range list {
		title, artist := getTag(src+file)
		fmt.Println(title, artist)
		os.Mkdir(out+artist, os.ModePerm)
		err := os.Rename(src+file, out+artist+"/"+file)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getTag(f string) (string, string) {
	file, err := os.OpenFile(f, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err)
	}
	m, err := tag.ReadFrom(file)
	if err != nil {
		fmt.Println(err)
	}
	return m.Title(), m.Artist()
}

func getMusic(rootpath string) []string {

	baseName := []string{}

	musicExt := []string{".flac", ".mp4", "m4a"}

	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		for _, ext := range musicExt {
			if filepath.Ext(path) == ext {
				baseName = append(baseName, info.Name())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return baseName
}
