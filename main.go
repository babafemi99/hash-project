package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"task-2.01/models"
)

// var csvPath = "./HNGi9 CSV FILE - Sheet1.csv"
var dataBytes bytes.Buffer
var nfts []models.NFT
var tAttributes []string

const collectionName = "Zuri NFT Tickets for Free Lunch"
const collectionID = "b774f676-c1d5-422e-beed-00ef5510c64d"

func main() {

	directory := flag.String("d", "", "default location for csv file")
	flag.Parse()
	csvPath := fmt.Sprintf("%v", *directory)
	log.Println("directory is ", csvPath)

	fileCsv, err := os.Open(csvPath)
	if err != nil {
		if csvPath == "" {
			log.Println("specify a download path using -d flag")
			return
		}
		log.Println("error opening csv file, did you specify a path to your csv file ?")
		log.Printf("error details: %v \n", err)
		return
	}
	defer fileCsv.Close()

	os.Mkdir("json-data", 0777)
	os.Mkdir("csv-data", 0777)
	os.Chdir("json-data")

	fileReader := csv.NewReader(fileCsv)
	fileReader.FieldsPerRecord = -1

	all, err := fileReader.ReadAll()
	if err != nil {
		return
	}

	var nftJson models.NFT
	var teamName string

	for i, data := range all {
		if i == 0 {
			continue
		}

		if strings.Contains(data[0], "TEAM") {
			teamName = data[0]
		}

		nftJson.FileName = data[2]
		nftJson.Format = "CHIP-0007"
		nftJson.Name = data[3]
		nftJson.SensitiveContent = false
		nftJson.MintingTools = teamName
		nftJson.Description = data[4]
		nftJson.Gender = data[5]
		nftJson.ID = data[7]
		nftJson.SeriesTotal = 420
		count, _ := strconv.Atoi(data[1])
		nftJson.SeriesNumber = count
		nftJson.Collections.Name = collectionName
		nftJson.Collections.Id = collectionID
		var attributes []models.Attributes
		var attribute models.Attributes
		split := strings.Split(data[6], ";")

		for _, s := range split {
			ss := strings.Split(s, ":")
			attribute.TraitType = ss[0]
			attribute.Value = strings.Join(ss[1:], " ")
			attributes = append(attributes, attribute)
		}
		nftJson.Attributes = attributes
		tAttributes = append(tAttributes, data[6])

		//hash the json data

		h := sha256.New()
		enc := gob.NewEncoder(&dataBytes)

		err = enc.Encode(&nftJson)
		if err != nil {
			return
		}

		h.Write(dataBytes.Bytes())
		sum := h.Sum(nil)
		nftJson.Data.Hash = fmt.Sprintf("%x", sum)

		marshal, err := json.MarshalIndent(nftJson, " ", "\t")
		if err != nil {
			log.Printf("something went wrong while marshalling data: %v", err)
			os.Exit(1)
		}
		jsonFile, err := os.Create(fmt.Sprintf("./%s.json", nftJson.FileName))
		if err != nil {
			log.Println(err)
			return
		}
		jsonFile.Write(marshal)
		jsonFile.Close()

		nfts = append(nfts, nftJson)
	}

	err = os.Chdir("../csv-data")
	if err != nil {
		log.Println("err3", err)
	}

	// create new csv file
	csvFile, err := os.Create("filename.output.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	err = csvwriter.Write([]string{"series number", "name", "description", "attributes", "gender", "uuid", "hash"})
	if err != nil {
		log.Println(err)
	}

	for i, nft := range nfts {
		var res []string
		res = append(res, strconv.Itoa(nft.SeriesNumber), nft.Name, nft.Description, tAttributes[i], nft.Gender, nft.ID, nft.Data.Hash)
		_ = csvwriter.Write(res)
	}
	csvwriter.Flush()
	csvFile.Close()
}
