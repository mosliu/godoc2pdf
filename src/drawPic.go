package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "image/png"
    "io/ioutil"
    "os"

    "github.com/golang/freetype"
    //"github.com/sirupsen/logrus"
)

//var log = logrus.New()
func imageAddWaterMark() {

    //原始图片是sam.jpg
    imgb, _ := os.Open("sam.jpg")
    img, _ := jpeg.Decode(imgb)
    defer imgb.Close()

    wmb, _ := os.Open("text.png")
    watermark, _ := png.Decode(wmb)
    defer wmb.Close()

    //把水印写到右下角，并向0坐标各偏移10个像素
    offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
    b := img.Bounds()
    m := image.NewNRGBA(b)

    draw.Draw(m, b, img, image.ZP, draw.Src)
    draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

    //生成新图片new.jpg，并设置图片质量..
    imgw, _ := os.Create("new.jpg")
    jpeg.Encode(imgw, m, &jpeg.Options{100})

    defer imgw.Close()

    fmt.Println("水印添加结束,请查看new.jpg图片...")
}

func createImage() {
    //图片的宽度
    dx := 200
    //图片的高度
    dy := 200
    imgfile, err := os.Create("test.png")
    if err != nil {
        fmt.Println(err)
    }
    defer imgfile.Close()
    img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

    //设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
    for y := 0; y < dy; y++ {
        for x := 0; x < dx; x++ {
            //设置一块 白色(255,255,255)不透明的背景
            img.Set(x, y, color.RGBA{255, 255, 255, 0})
        }
    }
    //读取字体数据
    fontBytes, err := ioutil.ReadFile("STXINWEI.ttf")
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
    f.SetFontSize(26)
    f.SetClip(img.Bounds())
    //设置输出的图片
    f.SetDst(img)
    //设置字体颜色(红色)
    f.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))

    //设置字体的位置
    pt := freetype.Pt(40, 40+int(f.PointToFixed(26))>>8)

    _, err = f.DrawString("hello,世界", pt)
    if err != nil {
        log.Fatal(err)
    }

    //以png 格式写入文件
    err = png.Encode(imgfile, img)
    if err != nil {
        log.Fatal(err)
    }
}

//
//func main()  {
//    createImage()
//}


