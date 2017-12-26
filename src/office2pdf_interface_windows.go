// +build windows

// interface_windows
package main

type Exporter interface {
    Export(inFile, outDir string) (outFile string, err error)
}
