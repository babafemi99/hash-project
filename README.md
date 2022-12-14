![alt text](https://github.com/babafemi99/hash-project/blob/MAIN/CSV_TASK.png)

# **CSV TASK**
This is a simple CLI app that takes in a .csv file and returns a JSON and a .csv file as outputs.

## **Description**
* A tool that takes .csv file name and generate Chip-0007 JSON schema for each row in the csv file.

* It generates a JSON file from the JSON schema and create a SHA-256 checksum for the file.

* It then creates a file containing each line of the CSV input file and append the checksum to each row o.


# USAGE

* Make sure you have [golang](https://go.dev/dl/) installed on your system, check by typing `go version` in your terminal


* Clone repository to local machine


* Download csv file


* specify csv path while running the application with the flag -d
    eg : `go run main.go -d="/path-to-csv-file"` or
         `go run main.go -d "path-csv-file"` 


* JSON files are generated in a folder called "json-data"


* .csv file is generated in a folder called "csv-data"


---
## DATA SAMPLES

* Sample json file below
```
{
    "format": "CHIP-0007",
    "name": "Pikachu",
    "description": "Electric-type Pokémon with stretchy cheeks",
    "minting_tool": "SuperMinter/2.5.2",
    "sensitive_content": false,
    "series_number": 22,
    "series_total": 1000,
    "attributes": [
        {
            "trait_type": "Species",
            "value": "Mouse"
        },
        {
            "trait_type": "Color",
            "value": "Yellow"
        },
        {
            "trait_type": "Friendship",
            "value": 50,
            "min_value": 0,
            "max_value": 255
        }
    ],
    "collection": {
        "name": "Example Pokémon Collection",
        "id": "e43fcfe6-1d5c-4d6e-82da-5de3aa8b3b57",
        "attributes": [
            {
                "type": "description",
                "value": "Example Pokémon Collection is the best Pokémon collection. Get yours today!"
            },
            {
                "type": "icon",
                "value": "https://examplepokemoncollection.com/image/icon.png"
            },
            {
                "type": "banner",
                "value": "https://examplepokemoncollection.com/image/banner.png"
            },
            {
                "type": "twitter",
                "value": "ExamplePokemonCollection"
            },
            {
                "type": "website",
                "value": "https://examplepokemoncollection.com/"
            }
        ]
    },
    "data": {
        "example_data": "VGhpcyBpcyBhbiBleGFtcGxlIG9mIGRhdGEgdGhhdCB5b3UgbWlnaHQgd2FudCB0byBzdG9yZSBpbiB0aGUgZGF0YSBvYmplY3QuIE5GVCBhdHRyaWJ1dGVzIHdoaWNoIGFyZSBub3QgaHVtYW4gcmVhZGFibGUgc2hvdWxkIGJlIHBsYWNlZCB3aXRoaW4gdGhpcyBvYmplY3QsIGFuZCB0aGUgYXR0cmlidXRlcyBhcnJheSB1c2VkIG9ubHkgZm9yIGluZm9ybWF0aW9uIHdoaWNoIGlzIGludGVuZGVkIHRvIGJlIHJlYWQgYnkgdGhlIHVzZXIu"
    }
}
```

* Sample .csv file 

| Team Names | Series Number | FileName           | Name              | Description                                 | Gender | Attributes                                                                                                                        | UUID                                 |
|------------|---------------|--------------------|-------------------|---------------------------------------------|--------|-----------------------------------------------------------------------------------------------------------------------------------|--------------------------------------|
| TEAM BEVEL | 1             | adewale-the-amebo  | adewale the amebo | Adewale likes to be in everyone's business. | Male   | hair: bald; eyes: black; teeth: none; clothing: red; accessories: mask; expression: none; strength: powerful; weakness: curiosity | cad316c3-37f8-4b27-9f53-9d803bfcfee7 |


