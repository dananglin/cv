package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	cvDataPath       string = "data/cv.json"
	cvTemplateDir    string = "template/"
	cvOutputDir      string = "__output/"
	cvOutputFileName string = "cv.tex"
)

func main() {

	var cv Cv
	cvOutputPath := cvOutputDir + cvOutputFileName

	fmap := template.FuncMap{
		"notLastElement":   notLastElement,
		"join":             join,
		"durationToString": durationToString,
	}

	// Read the JSON data
	log.Printf("INFO: Attempting to read data from %s...", cvDataPath)
	data, err := ioutil.ReadFile(cvDataPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to read data from file. %s", err.Error())
	}
	log.Print("INFO: Successfully read data.")

	log.Printf("INFO: Attempting to unmarshal JSON data...")
	if err = json.Unmarshal(data, &cv); err != nil {
		log.Fatalf("ERROR: Unable to unmarshal JSON data. %s", err.Error())
	}
	log.Println("INFO: JSON unmarshalling was successful.")

	// if CV_CONTACT_PHONE is set then add it to the CV
	phone := os.Getenv("CV_CONTACT_PHONE")
	if len(phone) > 0 {
		cv.Contact.Phone = phone
	}

	// Create the Output tex file
	log.Printf("INFO: Attempting to create output file %s...", cvOutputPath)
	if err = os.MkdirAll(cvOutputDir, 0750); err != nil {
		log.Fatalf("ERROR: Unable to create output directory %s. %s", cvOutputDir, err.Error())
	}
	output, err := os.Create(cvOutputPath)
	if err != nil {
		log.Fatalf("ERROR: Unable to create output file %s. %s.", cvOutputPath, err.Error())
	}
	defer output.Close()
	log.Printf("INFO: Successfully created output file %s.", cvOutputPath)

	// Execute template engine and produce the resulting TEX file
	log.Println("INFO: Attempting template execution...")
	t := template.Must(template.New("cv.tex.tmpl").Funcs(fmap).Delims("<<", ">>").ParseGlob(cvTemplateDir + "*.tex.tmpl"))

	if err = t.Execute(output, cv); err != nil {
		log.Fatalf("ERROR: Unable to execute the CV template. %s", err.Error())
	}
	log.Println("INFO: Template execution successful.")
}

// notLastElement returns true if an element of an array
// or a slice is not the last.
func notLastElement(pos, length int) bool {
	return pos < length-1
}

// join uses strings.Join to join all string elements into
// a single string.
func join(s []string) string {
	return strings.Join(s, " ")
}

// durationToString outputs the employment/education
// duration as a formatted string.
func durationToString(d Duration) string {
	start := fmt.Sprintf("%s, %s", d.Start.Month, d.Start.Year)
	present := strings.ToLower(d.Present)
	end := ""

	if present == "yes" || present == "true" {
		end = "Present"
	} else {
		end = fmt.Sprintf("%s, %s", d.End.Month, d.End.Year)
	}

	return start + " - " + end
}
