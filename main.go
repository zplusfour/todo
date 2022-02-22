package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

func checkTodos(dir string) {
	var toret []string
	comments := []string{"//", "/*", "#", "- -", "<!--", "%"}
	files, err := ioutil.ReadDir(dir)
  if err != nil {
    log.Fatal(err)
  }
 
  for _, f := range files {
    if f.IsDir() {
			checkTodos(dir)
		} else {
			__data, err := ioutil.ReadFile(dir+"/"+f.Name())
			if err != nil {
				log.Fatal(err)
			}
			data := string(__data)

			lines := strings.Split(data, "\n")
			for _, b := range lines {
				for _, c := range comments {
					if strings.HasPrefix(b, c+" TODO") {
						toret = append(toret, f.Name()+"\n"+b)
					} else if strings.HasPrefix(b, c+"TODO") {
						toret = append(toret, f.Name()+"\n"+b)
					}
				}
			}

			for _, k := range toret {
				fmt.Println(k)
			}
		}
  }
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		checkTodos(".")
	} else {
		checkTodos(args[0])
	}
}