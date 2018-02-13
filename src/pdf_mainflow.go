package main

import (
    "path/filepath"

    unicommon "github.com/unidoc/unidoc/common"
    "github.com/urfave/cli"
)

func init() {
    // Set debug-level logging via console.
    unicommon.SetLogger(unicommon.NewConsoleLogger(unicommon.LogLevelDebug))
}

func cliWaterMarkAndEncrypt() cli.Command {
    command := cli.Command{
        Name:        "pdfLock",
        Aliases:     []string{"p"},
        Category:    "Tools",
        Usage:       "Watermark and Lock a pdf file",
        UsageText:   "Example: doc2pdf pdfLock ./xxxx.pdf",
        Description: "Watermark and Lock a pdf file",
        ArgsUsage:   "<filename> <password>",
        //Flags: []cli.Flag{
        //    cli.BoolFlag{
        //        Name:   "show,s",
        //        Usage:  "show current password",
        //        Hidden: true,
        //    },
        //},
        Action: func(c *cli.Context) error {
            if c.NArg() < 2 {
                log.Fatal("Args' number less than 2")
            }
            addWaterMarkAndEncryptByConf(c.Args().First())
            //calcBase(c.Args().First(), c.Bool("debase"))
            return nil
        },
    }
    return command
}

func addWaterMarkAndEncryptByConf(inputfile string) {
    outDir, outFilename := filepath.Split(inputfile)
    outputPath := filepath.Join(outDir, "Done_"+outFilename)
    //watermarkFile := config.Watermark.Path
    //userPass := config.Security.UserPass.Password2Add
    //ownerPass := config.Security.OwnerPass.Password2Add
    addWaterMarkAndEncrypt(inputfile, outputPath)

}
func addWaterMarkAndEncrypt(inputfile string, outputPath string) {
    err := addWatermarkImageAndDateMark(inputfile, outputPath)
    if err != nil {
        log.Error(err)
    }
    if config.Pdfs.Security.UserPass.Enable == false {
        config.Pdfs.Security.UserPass.Password2Add = ""
        //userPass = ""
    }
    if config.Pdfs.Security.OwnerPass.Enable == false {
        config.Pdfs.Security.OwnerPass.Password2Add = ""
        //ownerPass = ""
    }
    //如果有一个需要加密则执行
    if config.Pdfs.Security.UserPass.Enable || config.Pdfs.Security.OwnerPass.Enable {
        addPassword(outputPath, outputPath, config.Pdfs.Security.UserPass.Password2Add, config.Pdfs.Security.OwnerPass.Password2Add)
    }
    err = printAccessInfo(inputfile, "")
    if err != nil {
        log.Errorf("Error: %v\n", err)
    }
}
