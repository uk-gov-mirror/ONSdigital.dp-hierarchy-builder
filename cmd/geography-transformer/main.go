package main

/**

This generator takes a v4 file and infers a hierarchy from the code in the label.

 */

import (
	"flag"
	"os"

	"github.com/ONSdigital/go-ns/log"
	"strings"
	"encoding/csv"
	"github.com/ONSdigital/dp-hierarchy-builder/hierarchy"
	"github.com/ONSdigital/dp-hierarchy-builder/cypher"
)

var filepath = flag.String("file", "cmd/geography-transformer/WD16_LAD16_CTY16_OTH_UK_LU.csv", "The path to the import filepath")
var csvFile = flag.String("csv", "cmd/geography-transformer/output/hierarchy.csv", "")
var cypherFile = flag.String("cypher", "cmd/geography-transformer/output/hierarchy.cypher", "")
var cypherDelFile = flag.String("cypher-delete", "cmd/geography-transformer/output/hierarchy-delete.cypher", "")

func main() {

	flag.Parse()

	f, err := os.Open(*filepath)
	checkErr(err)

	csvReader := csv.NewReader(f)
	defer f.Close()

	// identify the code lists used in this file from the header row.
	headerRow, err := csvReader.Read()
	checkErr(err)

	// maintain a slice of the code lists used in this file to reference later.
	codelists := createCodeListSlice(headerRow)

	// Create a map of code:node
	nodeMap := createNodeMap(csvReader, codelists)

	// populate a slice of top level nodes, i.e. the root elements
	rootNodes := hierarchy.IdentifyRootNodes(nodeMap)

	// populate the children of each node using the map to look up parents
	hierarchy.PopulateChildNodes(nodeMap)

	log.Debug("Generating csv", nil)
	err = hierarchy.CreateCSVFile(rootNodes, *csvFile)
	logErr(err)

	log.Debug("Generating cypher script", nil)
	err = cypher.CreateCypherFile(rootNodes, *cypherFile)
	logErr(err)

	log.Debug("Generating cypher deletion script", nil)
	err = cypher.CreateCypherDeleteFile(rootNodes, *cypherDelFile)
	logErr(err)
}

func createNodeMap(csvr *csv.Reader, codelists []string) *map[string]*hierarchy.Node {

	var nodeMap = make(map[string]*hierarchy.Node, 0)
	var err error = nil

	// read the rest of the CSV rows and populate a map of node objects with the code as the key.
	for err == nil {

		csvRow, err := csvr.Read()
		if err != nil {
			break
		}

		// each row has a number of codes, so iterate through the row and pick out each code / label
		for columnOffset, code := range csvRow {

			isCode := columnOffset%2 == 0
			codelistIndex := columnOffset / 2
			codelist := codelists[codelistIndex]
			codelist = "123" // todo : remove this override of the code list multiple codelists in a single geog - what to do???

			// create a node for each code in the row
			// the code will be empty where gaps exist between an area and its parent so ignore them
			if isCode && len(code) > 0 {

				// skip it if its already been added.
				if _, ok := nodeMap[code]; !ok {

					node := createNode(columnOffset, csvRow, codelist, code)
					nodeMap[code] = node
				}
			}
		}
	}

	return &nodeMap
}

func createNode(columnOffset int, csvRow []string, codelist string, code string) *hierarchy.Node {

	parentCode := ""
	parentCodeIndex := columnOffset + 2

	// only get the parent code if its not the last column
	if columnOffset+2 < len(csvRow) {
		parentCode = csvRow[parentCodeIndex]

		// if a parent code is not found next to this code, try looking in the next columns
		for parentCode == "" {
			parentCodeIndex = parentCodeIndex + 2
			parentCode = csvRow[parentCodeIndex]
		}
	}

	trimmedLabel := strings.Trim(csvRow[columnOffset+1], " ")
	escapedLabel := strings.Replace(trimmedLabel, "'", "\\'", -1)

	option := &hierarchy.Node{
		CodeList:   codelist,
		Code:       code,
		Label:      escapedLabel,
		ParentCode: parentCode,
	}

	return option

}

func createCodeListSlice(headerRow []string) []string {
	codelists := make([]string, len(headerRow)/2)
	// read the header row and populate an array of code list ID's as they are spread across the header row.
	for i, cell := range headerRow {

		isCode := i%2 == 0
		codelistIndex := i / 2

		if isCode {
			codelists[codelistIndex] = strings.TrimSuffix(cell, "CD")
		}
	}
	return codelists
}

func checkErr(err error) {
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

func logErr(err error) {
	if err != nil {
		log.Error(err, nil)
	}
}
