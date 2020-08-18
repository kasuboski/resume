package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseFiles("hack/resume.html.tmpl"))
	bs, err := ioutil.ReadFile("resume.json")
	if err != nil {
		log.Fatalf("couldn't read resume.json: %v", err)
	}

	f, err := os.Create("resume.html")
	if err != nil {
		log.Fatalf("couldn't open out.html: %v", err)
	}
	defer f.Close()

	var params map[string]interface{}
	err = json.Unmarshal(bs, &params)
	if err != nil {
		log.Fatalf("unable to unmarshal json: %v", err)
	}

	tmpl.Execute(f, params)
}
