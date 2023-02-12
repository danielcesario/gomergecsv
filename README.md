# Go Merge CSV
Simple Go program to merge many files, with same stricture, into one.

## Requirements
- Go 1.19 or later

## Preparing
You can put every files on folder **files**, in project root, the name of files doesn't matter but I recommending that the structure are the same.

## Running
You can execute program using follow command on the project root:

    go run cmd/mergecsv/main.go 

The program will generate a **result.csv** file, in project root, with rows of every file. On the header the program will use of the last file, it will ignore header of other files.