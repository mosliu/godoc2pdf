package main

import (
    "fmt"
    "os"
    "time"

    "path/filepath"

    "github.com/fatih/color"
    "github.com/urfave/cli"
    "strings"
)

var done = make(chan bool, 1)
// SUPPORTOFFICETYPE: 支持的office文件类型
var SUPPORTOFFICETYPE = []string{".docx", ".doc",".txt",".htm",".html",".mhtml", ".xls", "xlsx", ".ppt", ".pptx",}
// SUPPORTIMAGETYPE: 支持的image文件类型
var SUPPORTIMAGETYPE = []string{".jpg", ".jpeg",".png",}
var compileDate = "2018/1/10"
func main() {

    app := cli.NewApp()office2pdf_excel_windows.go
    app.Name = "Doc2Pdf Tool"
    app.Version = "1.0.0.0"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        {
            Name:  "liuxuan",
            Email: "liuxuan@liuxuan.net",
        },
    }
    app.Copyright = "(c) 2018 Labthink Support Group."
    app.Usage = "Convert A doc file to pdf, add watermark and use password to protect it."
    app.UsageText = "Example: doc2pdf ./aaa.doc"
    app.ArgsUsage = "<filename>"
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "verbose",
            //Hidden: true,
            Usage: "show detail convert info (not work)",
        },
    }

    app.Action = func(c *cli.Context) error {

        //log.Debug("Opening Serial Port...0")
        //log.Info("Opening Serial Port...1")
        //log.Warn("Opening Serial Port...2")
        //log.Error("Opening Serial Port...3")
        //log.Panic("Opening Serial Port...4")
        //log.Fatal("Opening Serial Port...5")
        debugflag := c.Bool("debug")
        if debugflag {
            initLogger("debug")
        } else {
            initLogger("info")
        }
        log.Info(c.NArg(), " args [", c.Args().First(), "]")
        logF.Info(c.NArg(), " args [", c.Args().First(), "]")
        log.Info("Author:Liu Xuan,last modified at "+compileDate)

        if c.NArg() > 1 || c.NArg() < 0 {
            log.Error("Wrong usage.Try to drag one file on this tool.")
            os.Exit(2)
        } else {
            inputOfficeFilePath := c.Args().First()
            if isFileExist(inputOfficeFilePath) {
                startConvert(inputOfficeFilePath)
            } else {
                logF.Error("the input file is not existed.")
            }
        }

        //time.AfterFunc(time.Minute*1, func() {
        //    done <- true
        //})
        //wait for terminal signal
        start()

        return nil
    }

    app.Commands = []cli.Command{
        //生成conf.yaml模板
        createTemplate(),
        cliWaterMarkAndEncrypt(),
        //office2pdf(),
        cliOffice2pdf(),
    }

    defer func() {
        if e := recover(); e != nil {
            log.WithField("error", e).Error("Panicing,error occured")
            //color.Red("Panicing,error occured: %s\r\n", e)
        }
    }()

    app.Run(os.Args)
}
func startConvert(officeFile string) {
    var outFile string
    var err error
    log.Info("Start Processing...")
    rootPath := filepath.Dir(officeFile)
    inFile, _ := filepath.Abs(officeFile)
    outDir, _ := filepath.Abs(rootPath)
    inFileExt := filepath.Ext(inFile)
    inFileExt = strings.ToLower(inFileExt)
    var isOfficeFileExtFlag = false
    for _, item := range SUPPORTOFFICETYPE {
        if strings.Compare(inFileExt, item) == 0 {
            isOfficeFileExtFlag = true
        }
    }
    var isImageFileExtFlag = false
    for _,item := range SUPPORTIMAGETYPE {
        if strings.Compare(inFileExt, item) == 0 {
            isImageFileExtFlag = true
        }
    }
    //convert to pdf
    if isOfficeFileExtFlag {
        exporter := exporterMap()[filepath.Ext(inFile)]
        if _, ok := exporter.(Exporter); ok {
            outFile, err = exporter.(Exporter).Export(inFile, outDir)
            if err != nil {
                logF.Fatal(err)
            }
            log.Info("output pdf file: " + outFile)
        }
        addWaterMarkAndEncryptByConf(outFile)
    } else if isImageFileExtFlag{
        addImageWaterMarkByConf(inFile)
    } else {
        if strings.Compare(inFileExt, ".pdf") == 0 {
            ok, err := testEncrypt(inFile)
            if err != nil {
                log.Fatal(err)
            }
            if ok {
                log.Info(color.RedString("The file is Encrypted!Can not operate"))
                printAccessInfo(inFile, "")
            }else{
                addWaterMarkAndEncryptByConf(inFile)
            }
        }
    }
}

func start() {

    //<-done
    color.Blue("Operate finished.Press Enter to Close This Program!")
    var out string
    fmt.Scanln(&out)

    //fmt.Println("buf",buf)
    //log.Printf("%q", buf[:n])
}
