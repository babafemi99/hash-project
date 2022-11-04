CSV TASK

Description

A tool that takes .csv file name and generate Chip-0007 JSON schema for each line.

It generates a JSON file from the JSON schema and create a SHA-256 checksum for the file.

It then creates a file containing each line of the CSV input file and append a the checksum to each row or line.


STEPS TO FOLLOW

1) clone repository to local machine
2) download csv file
3) specify csv path while running the application with the flag -d
    eg : go run main.go -d="/path-to-csv-file" or
         go run main.go -d "path-csv-file"
4) json files is generated in a folder called "json-data"
5) csv file is generated in a folder called "csv-data"