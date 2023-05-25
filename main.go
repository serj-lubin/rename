package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// 	const ShellToUse = "sh"
	// 	cmd := exec.Command(ShellToUse, "-c", "  echo 1 >/tmp/TeSt")
	// 	err := cmd.Run()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		Original_Path := file.Name()
		New_Path := strings.ToLower(Original_Path)
		if Original_Path != New_Path {
			fmt.Println(file.Name(), file.IsDir())
			e := os.Rename(Original_Path, New_Path)
			if e != nil {
				log.Fatal(e)
			}
		}
	}

}
