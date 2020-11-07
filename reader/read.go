package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

var replacer = strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func createFile() {
	content := "Martin Luther\nAngelina Jolie\nPablo Escobar\nJack Sparrow"
	e := ioutil.WriteFile("outfile", []byte(content), 0755)
	if e != nil {
		fmt.Println("Error creating file.", e)
	}
}

func procFile(pth string) []name {
	names := make([]name, 0) // performance penalty for this

	file, err := os.Open(pth)
	checkErr(err)
	defer file.Close()

	scn := bufio.NewScanner(file)
	for scn.Scan() {
		v := strings.Split(scn.Text(), " ")
		names = append(names, name{fname: replacer.Replace(v[0]), lname: replacer.Replace(v[1])})
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	return names
}

func main() {
	createFile() // Create file called outfile already populated with names
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Names reader")
	fmt.Println("======================")
	fmt.Print("Provide file path: ")
	fileP, _ := reader.ReadString('\n')
	fileP = replacer.Replace(fileP)

	var names []name

	if fileP != "" {
		fmt.Println("Content")
		names = procFile(fileP)
	} else {
		fmt.Println("Using default file. Content:")
		names = procFile("outfile")
	}

	for _, nm := range names {
		fmt.Println("Name: " + nm.fname + " " + nm.lname)
	}
}
