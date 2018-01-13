package main

import (
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"text/template"
)

var (
	// startTime time.Time
	appName  string
	username string

	files = []string{}
	model = map[string]interface{}{
		"Name":     "",
		"Username": "",
	}
)

func init() {
	flag.StringVar(&appName, "name", "gpt", "your app name")
	flag.StringVar(&username, "username", "yangwenmai", "your Github username")
}

// GenFiles 根据模板生成内容
func GenFiles(files []string, startTime time.Time) {
	var wg sync.WaitGroup
	for _, n := range files {
		wg.Add(1)
		go func(n string) {
			genFile(n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	log.Printf("Completed and cost %v.", time.Now().Sub(startTime))
}

func genFile(n string) {
	log.Printf("start gen file %s\n", n)
	bs, err := Asset(n)
	if err != nil {
		log.Fatal(err)
	}
	templateExecute(bs, n)
}

func templateExecute(bs []byte, n string) {
	t := template.New("gpt")
	t, err := t.Parse(string(bs))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(n)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(f, model)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()
}

func AllFiles() []string {
	files := []string{
		// "docs/wiki.md",
		"README.md",
		"example.go",
		"example_test.go",
		".gitignore",
		".travis.yml",
		"LICENSE",
	}
	return files
}

func main() {
	startTime := time.Now()

	flag.Parse()
	model["Name"] = appName
	model["Username"] = username

	GenFiles(AllFiles(), startTime)
}
