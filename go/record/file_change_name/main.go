package main

import (
	"fmt"
	"github.com/mikkyang/id3-go"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk("./", func(path string, info fs.FileInfo, err error) error {

		//fmt.Println("path:",path)
		//fmt.Println("infoName:",info.Name())
		if info.IsDir() || path == "./" {
			return nil
		}
		//fmt.Println(info.Name())

		li := strings.LastIndex(info.Name(), ".")
		var extensionName = ""
		if li != -1 {
			extensionName = info.Name()[li+1:]
			//fmt.Println(extensionName)
		}

		if strings.ToLower(extensionName) == "mp3" {
			ChangeMp3FileTag("./"+path, info.Name())
		}

		return nil
	})

}

func ChangeMp3FileTag(filePath, newName string) {
	mp3File, err := id3.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer mp3File.Close()

	//fmt.Println(mp3File.Artist())

	mp3File.SetTitle(newName)

	fmt.Println(fmt.Sprintf(`Change mp3File %s title to %s Success!`, filePath, newName))

}
