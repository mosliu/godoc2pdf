package main

import (
    "os"
    "path/filepath"
)

func isFileExist(path string) bool {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        return false
    }
    return true
}


func getMainExePath() string {
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    return dir
}