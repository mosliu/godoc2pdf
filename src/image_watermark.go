package main

import (
    "os"
    "fmt"
    "image"
    "image/png"
    "image/jpeg"
    //"io"
    //"github.com/sirupsen/logrus"
    "image/draw"
    "image/color"
    //"github.com/anthonynsimon/bild/transform"
    "github.com/anthonynsimon/bild/blend"
    "github.com/golang/freetype"
    "io/ioutil"
    "path/filepath"
)

//var log = logrus.New()

func getImageDimensions(imagePath string) (int, int) {
    file, err := os.Open(imagePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
    defer file.Close()
    imageConfig, _, err := image.DecodeConfig(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
    }
    return imageConfig.Width, imageConfig.Height
}

func imageWatermark(imgOriginPath string, imgWatermarkPath string, imgOutPath string, offsetX int, offsetY int, opacity uint8) {
    imgOriginFile, _ := os.Open(imgOriginPath)
    imgOrigin, imgOriginType, _ := image.Decode(imgOriginFile)
    fmt.Println("imgOrigin image type is:", imgOriginType)
    defer imgOriginFile.Close()
    fmt.Println(imgOrigin.At(imgOrigin.Bounds().Dx(), imgOrigin.Bounds().Dy()).RGBA())
    fmt.Println(imgOrigin.At(99999, 99999).RGBA())
    imgWatermarkFile, _ := os.Open(imgWatermarkPath)
    imgWaterMark, imgWatermarkType, _ := image.Decode(imgWatermarkFile)
    fmt.Println("imgWaterMark image type is:", imgWatermarkType)
    defer imgWatermarkFile.Close()

    dx := imgWaterMark.Bounds().Dx()
    dy := imgWaterMark.Bounds().Dy()

    newRgba := image.NewRGBA(imgWaterMark.Bounds())
    for x := 0; x < dx; x++ {
        for y := 0; y < dy; y++ {
            r, b, g, a := imgWaterMark.At(x, y).RGBA()
            a2 := uint8(a >> 8)
            if a2 > opacity {
                a2 = opacity
            }
            newRgba.Set(x, y, color.NRGBA{
                uint8(r >> 8),
                uint8(b >> 8),
                uint8(g >> 8),
                uint8(a2),
            }) //设定alpha图片的透明度
        }
    }
    offset := image.Pt(offsetX, offsetY)
    bounds := imgOrigin.Bounds()
    imgOut := image.NewRGBA(bounds)
    draw.Draw(imgOut, bounds, imgOrigin, image.ZP, draw.Src)
    draw.Draw(imgOut, newRgba.Bounds().Add(offset), newRgba, image.ZP, draw.Over)
    fmt.Println("imgOrigin:", imgOrigin.Bounds().Dx(), imgOrigin.Bounds().Dy())
    imgFileWaterMarked, _ := os.Create(imgOutPath)
    defer imgFileWaterMarked.Close()
    switch imgOriginType {
    case "jpeg":
        jpeg.Encode(imgFileWaterMarked, imgOut, &jpeg.Options{jpeg.DefaultQuality})
    case "png":
        png.Encode(imgFileWaterMarked, imgOut)
    }
}

// this method use bild,but the size of the pic is wrong.
// fix the size problem,but the method is not clever
func imageWatermark2(imgOriginPath string, imgWatermarkPath string, imgOutPath string, offsetX int, offsetY int, percent float64) {
    imgOriginFile, _ := os.Open(imgOriginPath)
    imgOrigin, imgOriginType, _ := image.Decode(imgOriginFile)
    fmt.Println("imgOrigin image type is:", imgOriginType)
    defer imgOriginFile.Close()

    imgWatermarkFile, _ := os.Open(imgWatermarkPath)
    imgWatermark, imgWatermarkType, _ := image.Decode(imgWatermarkFile)
    fmt.Println("imgWatermark image type is:", imgWatermarkType)
    defer imgWatermarkFile.Close()

    originBounds := imgOrigin.Bounds()
    WatermarkBounds := imgOrigin.Bounds()
    dx := originBounds.Dx()
    dy := originBounds.Dy()

    if WatermarkBounds.Dx() > dx {
        dx = WatermarkBounds.Dx()
    }
    if WatermarkBounds.Dy() > dy {
        dy = WatermarkBounds.Dy()
    }

    newRgba := image.NewRGBA(image.Rect(0, 0, dx, dy))
    for x := offsetX; x < dx; x++ {
        for y := offsetY; y < dy; y++ {
            r, b, g, a := imgWatermark.At(x-offsetX, y-offsetY).RGBA()
            newRgba.Set(x, y, color.RGBA{
                uint8(r >> 8),
                uint8(b >> 8),
                uint8(g >> 8),
                uint8(a >> 8),
            })
        }
    }

    result := blend.Opacity(imgOrigin, newRgba, percent)

    imgOutFile, _ := os.Create(imgOutPath)

    jpeg.Encode(imgOutFile, result, &jpeg.Options{jpeg.DefaultQuality})
    defer imgOutFile.Close()
    switch imgOriginType {
    case "jpeg":
        jpeg.Encode(imgOutFile, result, &jpeg.Options{jpeg.DefaultQuality})
    case "png":
        png.Encode(imgOutFile, result)
    }

}

func createTextImage(width int, height int, str string, fontPath string, fontsize float64) draw.Image {
    //图片的宽度
    dx := width
    //图片的高度
    dy := height
    img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

    //设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            //设置一块 白色(255,255,255)透明的背景
            img.Set(x, y, color.RGBA{255, 255, 255, 0})
        }
    }
    //读取字体数据
    fontBytes, err := ioutil.ReadFile(fontPath)
    if err != nil {
        log.Println(err)
    }
    //载入字体数据
    font, err := freetype.ParseFont(fontBytes)
    if err != nil {
        log.Println("load front fail", err)
    }
    f := freetype.NewContext()
    //设置分辨率
    f.SetDPI(72)
    //设置字体
    f.SetFont(font)
    //设置尺寸
    f.SetFontSize(fontsize)
    f.SetClip(img.Bounds())
    //设置输出的图片
    f.SetDst(img)
    //设置字体颜色(红色)
    f.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))

    //设置字体的位置
    pt := freetype.Pt(40, 40+int(f.PointToFixed(fontsize))>>8)

    _, err = f.DrawString(str, pt)
    if err != nil {
        log.Fatal(err)
    }

    return img
}

func addImageWaterMarkByConf(inputfile string) {
    outDir, outFilename := filepath.Split(inputfile)
    outputPath := filepath.Join(outDir, "Done_"+outFilename)

    if !config.ImageWatermark.Enable {
        return
    }

    imageWatermark2(inputfile, config.ImageWatermark.Path, outputPath, config.ImageWatermark.OffsetX, config.ImageWatermark.OffsetY, config.ImageWatermark.Opacity)
}
//
//func main() {
//    imageWatermark("toadd.jpg", "aaa.png", "watermarked.jpg", 0, 0, 25)
//    imageWatermark2("toadd.jpg", "aaa.png", "watermarked2.jpg", 50, 50, 0.1)
//    img := createTextImage(200, 20, "hello,世界,hello,My lord", "STXINWEI.ttf", 26)
//
//    imgfile, err := os.Create("test.png")
//    if err != nil {
//        fmt.Println(err)
//    }
//    defer imgfile.Close()
//    //以png 格式写入文件
//    err = png.Encode(imgfile, img)
//    if err != nil {
//        log.Fatal(err)
//    }
//}
