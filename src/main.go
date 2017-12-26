package main

import (
    "fmt"
    "os"
    "time"

    "github.com/fatih/color"
    "github.com/urfave/cli"
)

var done = make(chan bool, 1)

func main() {
    initLogger("info")

    app := cli.NewApp()
    app.Name = "RPC test Tool"
    app.Version = "1.0.0.0"
    app.Compiled = time.Now()
    app.Authors = []cli.Author{
        {
            Name:  "liuxuan",
            Email: "liuxuan@liuxuan.net",
        },
    }
    app.Copyright = "(c) 2017 Labthink Support Group."
    app.Usage = "CZY-6S device time correction tool"
    app.UsageText = "Example: Tool.exe COM1"
    app.ArgsUsage = "<COMPORT>"
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "debug,d",
            //Hidden: true,
            Usage: "language for the greeting",
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
        log.Info(c.NArg(),c.Args().First())
        log.Info("Author:Liu Xuan,last modified at 2017/11/30")

        log.Info("Opening Serial Port...")
        time.AfterFunc(time.Minute*1, func() {
            done <- true
        })
        //wait for terminal signal
        start()



        return nil
    }

    app.Commands = []cli.Command{
        pdf1(),
        pdf2(),
        //office2pdf(),
        office2pdfcli(),
    }

    defer func() {
        if e := recover(); e != nil {
            log.WithField("error", e).Error("Panicing,error occured")
            //color.Red("Panicing,error occured: %s\r\n", e)
        }
    }()



    app.Run(os.Args)
}

func start() {

    <-done
    color.Blue("本次校正结束,可以关闭了")
    var out string
    fmt.Scanln(&out)

    //fmt.Println("buf",buf)
    //log.Printf("%q", buf[:n])
}
