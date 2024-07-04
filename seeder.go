package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)
// should be run with make from root - parent project dir
const rootDir = "db/seeds"

func loadSQLFiles() []string {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func loadSQLFiles() []string {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}
