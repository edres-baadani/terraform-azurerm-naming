package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// --- Existing Structs ---
// Resource definition for the package
type Resource struct {
	Name    string  `json:"name"`
	Length  *Length `json:"length,omitempty"`
	Regex   *string `json:"regex,omitempty"`
	Scope   *string `json:"scope,omitempty"`
	Slug    *string `json:"slug,omitempty"`
	Dashes  bool    `json:"dashes"`
}

// Length allowed for that resorce
type Length struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// --- New Structs for Environments and Locations ---
type Environment struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	SortOrder int    `json:"sortOrder"`
}

type Location struct {
	Name                string `json:"name"`
	DisplayName         string `json:"displayName"`
	ShortName           string `json:"shortName"`
	RegionalDisplayName string `json:"regionalDisplayName"`
	RegionCategory    string `json:"regionCategory"`
	PairedRegionNames    string `json:"pairedRegionNames"`
}

// --- Combined Data for Template Execution ---
type TemplateData struct {
	Resources   []Resource
	Environments []Environment
	Locations   []Location
}

func main() {
	files, err := ioutil.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}
	var fileNames = make([]string, len(files))
	for i, file := range files {
		fileNames[i] = "templates/" + file.Name()
	}
	caser := cases.Title(language.AmericanEnglish)
	parsedTemplate, err := template.New("templates").Funcs(template.FuncMap{
		"cleanRegex": func(dirtyString string) string {
			var re = regexp.MustCompile(`(?m)\(\?=.{\d+,\d+}\$\)|\(\?!\.\*--\)`)
			return re.ReplaceAllString(dirtyString, "")
		},
		"replace": strings.ReplaceAll,
		"title":   caser.String,
	}).ParseFiles(fileNames...)
	if err != nil {
		log.Fatal(err)
	}

	// Load and process Resources
	sourceDefinitions, err := ioutil.ReadFile("resourceDefinition.json")
	if err != nil {
		log.Fatal(err)
	}
	var resourcesData []Resource
	err = json.Unmarshal(sourceDefinitions, &resourcesData)
	if err != nil {
		log.Fatal(err)
	}
	sourceDefinitionsUndocumented, err := ioutil.ReadFile("resourceDefinition_out_of_docs.json")
	if err != nil {
		log.Fatal(err)
	}
	var resourcesUndocumentedData []Resource
	err = json.Unmarshal(sourceDefinitionsUndocumented, &resourcesUndocumentedData)
	if err != nil {
		log.Fatal(err)
	}
	resourcesData = append(resourcesData, resourcesUndocumentedData...)
	sort.Slice(resourcesData, func(i, j int) bool {
		return resourcesData[i].Name < resourcesData[j].Name
	})

	// --- New Code for Environments and Locations ---
	// Load and process Environments
	sourceEnvironments, err := ioutil.ReadFile("resourceenvironments.json")
	if err != nil {
		log.Fatal(err)
	}
	var environmentsData []Environment
	err = json.Unmarshal(sourceEnvironments, &environmentsData)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(environmentsData, func(i, j int) bool {
		return environmentsData[i].SortOrder < environmentsData[j].SortOrder
	})

	// Load and process Locations
	sourceLocations, err := ioutil.ReadFile("resourcelocations.json")
	if err != nil {
		log.Fatal(err)
	}
	var locationsData []Location
	err = json.Unmarshal(sourceLocations, &locationsData)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(locationsData, func(i, j int) bool {
		return locationsData[i].Name < locationsData[j].Name
	})
	
	// Prepare the final data structure to pass to the templates
	data := TemplateData{
		Resources:   resourcesData,
		Environments: environmentsData,
		Locations:   locationsData,
	}

	mainFile, err := os.OpenFile("main.tf", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Pass the combined data to the "main" template
	parsedTemplate.ExecuteTemplate(mainFile, "main", data)

	outputsFile, err := os.OpenFile("outputs.tf", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Pass the combined data to the "outputs" template
	parsedTemplate.ExecuteTemplate(outputsFile, "outputs", data)
}