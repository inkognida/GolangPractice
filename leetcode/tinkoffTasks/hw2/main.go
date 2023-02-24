package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type introduction interface {
	hello() string
	goodbye() string
}

type human struct {
	Name string
	Msg  string
}

type animal struct {
	Name string
	Msg  string
}

func (human human) hello() string {
	return "hello from human"
}

func (human human) goodbye() string {
	return "goodbye from human"
}

func (animal animal) hello() string {
	return "hello from animal"
}

func (animal animal) goodbye() string {
	return "goodbye from animal"
}

type OperationObject struct {
	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Id    int         `json:"id,omitempty"`
	Date  interface{} `json:"created_at"`
}
type Bill struct {
	Company   string          `json:"company,omitempty"`
	Operation OperationObject `json:"operation,omitempty"`

	Type  string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Id    int         `json:"id,omitempty"`
	Date  interface{} `json:"created_at"`
}

const EnvVarFile = "file"

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Failed to open file")
	}
	return file
}

func handleInput() *os.File {
	var flagFile = flag.String("file", "", ".json filename expected")
	flag.Parse()

	if *flagFile != "" {
		return openFile(*flagFile)
	}

	if envFile, exist := os.LookupEnv(EnvVarFile); exist != false {
		return openFile(envFile)
	}

	var userFile string
	if _, err := fmt.Scan(&userFile); err != nil {
		return openFile(userFile)
	}

	return nil
}

func main() {
	intros := []introduction{human{}, animal{}}
	for _, creature := range intros {
		fmt.Println(creature.hello())
	}

	for _, creature := range intros {
		fmt.Println(creature.goodbye())
	}
	if file := handleInput(); file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println("Failed to close file")
			}
		}(file)

	} else {
		log.Fatalln("No file to handle")
	}

}
