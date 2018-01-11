// +build windows

// interface_windows
package main

import (
    "github.com/urfave/cli"
    "path/filepath"
)

type Exporter interface {
    Export(inFile, outDir string) (outFile string, err error)
}

func exporterMap() (m map[string]interface{}) {
    m = map[string]interface{}{
        ".doc":  new(Word),
        ".docx": new(Word),
        ".txt": new(Word),
        ".htm": new(Word),
        ".html": new(Word),
        ".mhtml": new(Word),
        ".xls":  new(Excel),
        ".xlsx": new(Excel),
        ".ppt":  new(PowerPoint),
        ".pptx": new(PowerPoint),
    }
    return
}


func cliOffice2pdf() cli.Command {
    command := cli.Command{
        Name:        "office2pdf",
        Aliases:     []string{"o2p"},
        Category:    "Tools",
        Usage:       "Convert a  office word/excel/ppt to a Pdf file",
        UsageText:   "Example: doc2pdf office2pdf ./xxxx.docx ",
        Description: "Convert a  office word/excel/ppt to a Pdf file",
        ArgsUsage:   "<inFilename> <outFilePath>",
        //Flags: []cli.Flag{
        //    cli.BoolFlag{
        //        Name:   "show,s",
        //        Usage:  "show current password",
        //        Hidden: true,
        //    },
        //},
        Action: func(c *cli.Context) error {
            if (c.NArg() < 2) {
                log.Fatal("Args' number less than 2")
            }

            inFile, outDir := c.Args().First(), c.Args().Get(1)

            log.Info("input file: " + inFile + "\noutput dir: " + outDir)

            if fileIsExist(inFile) && fileIsExist(outDir) {
                log.Info("Processing...")
                inFile, _ = filepath.Abs(inFile)
                outDir, _ = filepath.Abs(outDir)
                exporter := exporterMap()[filepath.Ext(inFile)]
                if _, ok := exporter.(Exporter); ok {
                    outFile, err := exporter.(Exporter).Export(inFile, outDir)
                    if err != nil {
                        log.Fatal(err)
                    }
                    log.Info("output file: " + outFile)
                }
            } else {
                log.Error("inputFile or outputPath path is error ")
            }

            return nil
        },
    }
    return command
}
