package main

import (
	"golang.org/x/exp/errors/fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please give root folder folderName as argument")
		return
	}

	folderName := os.Args[1]
	folderPath := "/Users/akshay/Tekion/goWorkspace/src/github.com/akshay1192/nomercy/Algos/" + folderName

	helperPath := filepath.Join(folderPath, "helperFunc")

	err := os.MkdirAll(helperPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	createFiles(folderPath, []string{"main.go", "problem.jpg", "README.md", "/helperFunc/helper.go", "/helperFunc/helper_test.go"})

}

func createFiles(currentDir string, paths []string) {

	err := os.Chdir(currentDir)
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		newFile, err := os.Create(currentDir + "/" + path)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newFile)
		newFile.Close()
	}
}
