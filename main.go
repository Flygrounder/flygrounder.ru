package main

import (
	_ "embed"
	"encoding/json"
	"html/template"
	"log"
	"os"
)

type content struct {
	Socials  []social
	Sections []section
}

type social struct {
	Title string
	Link  string
}

type section struct {
	Title   string
	Entries []entry
}

type entry struct {
	Title  string
	Points []template.HTML
	Period *string
	Tags   []string
	Image  string
}

//go:embed content.json
var pageContentBytes []byte

//go:embed page.tmpl
var pageTemplate string

func main() {
	t, err := template.New("page").Parse(pageTemplate)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var pageContent content
	err = json.Unmarshal(pageContentBytes, &pageContent)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = t.Execute(os.Stdout, pageContent)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
