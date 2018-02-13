package main

import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/unidoc/unidoc/pdf/creator"
    pdf "github.com/unidoc/unidoc/pdf/model"
    "strings"
    "strconv"
    "time"
)

// Watermark pdf file based on an image.
func addWaterMarkImage(c *creator.Creator, f *os.File) error {
    pdfReader, err := pdf.NewPdfReader(f)
    if err != nil {
        return err
    }

    numPages, err := pdfReader.GetNumPages()
    if err != nil {
        return err
    }

    //水印图片方法
    if config.Pdfs.Watermark.Enable {
        watermarkImgPath := config.Pdfs.Watermark.Path
        if !isFileExist(watermarkImgPath) {
            watermarkImgPath = filepath.Join(getMainExePath(), watermarkImgPath)
        }
        watermarkImg, err := creator.NewImageFromFile(watermarkImgPath)
        if err != nil {
            return err
        }

        for i := 0; i < numPages; i++ {
            pageNum := i + 1

            // Read the page.
            page, err := pdfReader.GetPage(pageNum)
            if err != nil {
                return err
            }

            //_ = processPage(page)

            // Add to creator.
            c.AddPage(page)

            if config.Pdfs.Watermark.ScaleToHeight{
                watermarkImg.ScaleToHeight(c.Context().PageHeight)
            }
            if config.Pdfs.Watermark.ScaleToWidth{
                watermarkImg.ScaleToWidth(c.Context().PageWidth)
            }
            x := config.Pdfs.Watermark.WidthPos
            y := config.Pdfs.Watermark.HeightPos
            watermarkImg.SetPos(x, y)
            watermarkImg.SetOpacity(config.Pdfs.Watermark.Opacity)

            _ = c.Draw(watermarkImg)

            //if !addDateFlag {
            //    p := creator.NewParagraph("hahahahaha")
            //
            //    p.SetFont(fonts.NewFontTimesBoldItalic())
            //
            //    p.SetPos(20.0, 20.0)
            //    _ = c.Draw(p)
            //}
        }
    }
    return nil
}

func addHeaderAndFooter(c *creator.Creator) error {
    // if both set to false,skip
    if (config.Pdfs.Textmark.HeadArea.Enable || config.Pdfs.Textmark.FootArea.Enable) == false {
        return nil
    }

    var footerFont, headerFont *pdf.PdfFont
    var err error

    headerFontPath := config.Pdfs.Textmark.HeadArea.FontPath
    if (len(headerFontPath) == 0) || (!config.Pdfs.Textmark.HeadArea.Enable) {
        headerFont = nil
    } else {
        if !isFileExist(headerFontPath) {
            headerFontPath = filepath.Join(getMainExePath(), headerFontPath)
        }
        headerFont, err = pdf.NewCompositePdfFontFromTTFFile(headerFontPath)
        if err != nil {
            return err
        }
    }
    if !config.Pdfs.Textmark.FootArea.Enable {
        footerFont = nil
    } else if (headerFont != nil) && strings.EqualFold(config.Pdfs.Textmark.HeadArea.FontPath, config.Pdfs.Textmark.FootArea.FontPath) {
        footerFont = headerFont
    } else {
        footerFontPath := config.Pdfs.Textmark.FootArea.FontPath
        if !isFileExist(footerFontPath) {
            footerFontPath = filepath.Join(getMainExePath(), footerFontPath)
        }
        footerFont, err = pdf.NewCompositePdfFontFromTTFFile(footerFontPath)
        if err != nil {
            return err
        }
    }

    // 写页面头部
    if config.Pdfs.Textmark.HeadArea.Enable {
        c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
            // Draw the on a block for each page.
            //使用了unidoc的compositefonts 分支
            contents := config.Pdfs.Textmark.HeadArea.Contents
            farg := footerAndHeaderArgs{args.PageNum, args.TotalPages}
            //block.ScaleToWidth(c.Context().Width)
            footerAndHeaderDrawler(contents, farg, headerFont, block)
        })
    }

    // 写页面脚部
    if config.Pdfs.Textmark.FootArea.Enable {
        c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
            // Draw the on a block for each page.
            //使用了unidoc的compositefonts 分支
            contents := config.Pdfs.Textmark.FootArea.Contents
            //block.ScaleToWidth(c.Context().Width)
            farg := footerAndHeaderArgs{args.PageNum, args.TotalPages}
            footerAndHeaderDrawler(contents, farg, footerFont, block)
        })
    }

    //err = addFooter(c, footerFont)
    //if err != nil {
    //    return err
    //}
    return nil
}

type footerAndHeaderArgs struct {
    PageNum    int
    TotalPages int
}

func footerAndHeaderDrawler(contents []Content, args footerAndHeaderArgs, footFont *pdf.PdfFont, block *creator.Block) {
    for i := 0; i < len(contents); i++ {
        if len(contents[i].Text) == 0 {
            continue
        }
        text := contents[i].Text
        if strings.Contains(contents[i].Text, "${PageNum}") {
            text = strings.Replace(text, "${PageNum}", strconv.Itoa(args.PageNum), -1)
        }
        if strings.Contains(contents[i].Text, "${TotalPages}") {
            text = strings.Replace(text, "${TotalPages}", strconv.Itoa(args.TotalPages), -1)
        }
        if strings.Contains(contents[i].Text, "${Date}") {
            Datestr := time.Now().Format(contents[i].DateFormat)
            text = strings.Replace(text, "${Date}", Datestr, -1)
        }

        p := creator.NewParagraph(text)
        if (footFont != nil) && (contents[i].UseFont) {
            p.SetFont(footFont)
        }
        fontSize := contents[i].FontSize
        //default 8
        if fontSize <= 0 {
            fontSize = 8
        }
        p.SetFontSize(fontSize)
        //p.SetPos(50, 20)
        x := contents[i].PosX0
        y := contents[i].PosY0

        p.SetPos(x, y)
        p.SetColor(creator.ColorRGBFrom8bit(contents[i].RGB[0], contents[i].RGB[1], contents[i].RGB[2]))

        block.Draw(p)
    }
}

func addWatermarkImageAndDateMark(inputPath string, outputPath string) error {
    //unicommon.Log.Debug("Input PDF: %v", inputPath)
    //unicommon.Log.Debug("Watermark image: %s", watermarkImgPath)

    //dateMarkFont, err := pdf.NewPdfFontFromTTFFile("./STXINWEI.TTF")

    c := creator.New()

    // Read the input pdf file.
    f, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer f.Close()

    err = addWaterMarkImage(c, f)
    if err != nil {
        return err
    }
    err = addHeaderAndFooter(c)
    if err != nil {
        return err
    }
    c.SetPageMargins(config.Pdfs.Textmark.Margins.Left, config.Pdfs.Textmark.Margins.Right, config.Pdfs.Textmark.Margins.Top, config.Pdfs.Textmark.Margins.Bottom)
    //c.SetPageMargins(0, 0, 0, 0)

    err = c.WriteToFile(outputPath)
    return err
}

// 获取页面信息，并输出
func processPage(page *pdf.PdfPage) error {
    mBox, err := page.GetMediaBox()
    if err != nil {
        return err
    }
    pageWidth := mBox.Urx - mBox.Llx
    pageHeight := mBox.Ury - mBox.Lly

    fmt.Printf(" Page: %+v\n", page)
    fmt.Printf(" Page mediabox: %+v\n", page.MediaBox)
    fmt.Printf(" Page height: %f\n", pageHeight)
    fmt.Printf(" Page width: %f\n", pageWidth)

    return nil
}
