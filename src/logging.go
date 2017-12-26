package main

//Use github.com/sirupsen/logrus

import (
    "os"

    "github.com/shiena/ansicolor"
    "github.com/sirupsen/logrus"
)

// 你可以创建很多instance
//log to stdout.
var log = logrus.New()

// log to File.
var logF = logrus.New()

func init() {
    initLogger("info")
}

func initLogger(level string) {

    lvl, err := logrus.ParseLevel(level)
    if err != nil {
        log.Fatal(err)
        log.SetLevel(logrus.InfoLevel)
    } else {
        log.SetLevel(lvl)
    }

    // force colors on for TextFormatter
    log.Formatter = &logrus.TextFormatter{ForceColors: true,}
    // then wrap the log output with it
    // 用于解决windows的terminal中彩色不正确的问题

    log.Out = ansicolor.NewAnsiColorWriter(os.Stdout)

    //init logF
    file, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        logF.Out = file
        logF.SetLevel(logrus.InfoLevel)
    } else {
        log.Info("Failed to log to file, using default stderr")
    }
    //log.WithFields(logrus.Fields{
    //    "filename": "123.txt",
    //}).Info("打开文件失败")
}
