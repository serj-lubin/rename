package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
		fmt.Println(file.Name(), file.IsDir())
		Original_Path :=  file.Name()
		New_Path := strings.ToLower(Original_Path)
		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			log.Fatal(e)
		}
	}

}
