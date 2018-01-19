# doc2pdf

## 功能 [Function]
1. Use OLE to convert a word/excel/powerpoint file to pdf
2. Add  watermark and owner password protect to the pdf.
3. Add user defined page header and footer。

the three funcs  can be used separately.

---

1. 可以OLE使用OFFICE等将word、excel、powerpoint等文件转换为pdf
2. 对PDF增加水印和修改限制等功能，增加常用于发布一些不希望别人修改的内容
3. 可对PDF增加自定义的页眉页脚内容。

注：三个功能可以独立使用



## 使用 [Usage]

pre-condition: Ms Office  2013+ is needed for function 1 .

1. Put the file to convert in the .exe directory（Maybe needn't,For My test, file in other Directories is OK,at win10.）

2. drag the file to doc2pdf.exe（Easier than knock the keyboard to input the filepath）

3. if the file drag in is XXX.doc，then it will produce an unlock file name XXX.doc.pdf and a locked file Done_XXX.doc.pdf

4. if a pdf is dragged in, A locked/watermarked/HeaderAdded/FooterAdded  pdf will generate then. 

---

功能1使用条件为本机安装Ms Office 2013+

1. 将需要转换的文件放到doc2pdf目录中来。 （不必须，我在win10测试可不用）

2. 将需要转换的文件拖到doc2pdf.exe之上

3. 假设拖入的文件名为XXX.doc，则生成的文件中，不加密不加水印的为XXX.doc.pdf，加密加水印的为Done_XXX.doc.pdf

4. 拖入未加密的pdf文件，会生成加密加水印加页眉页脚的pdf文件

## 配置

Use the conf.yaml to stroe the configuration info.The Default one is below.

---

使用conf.yaml保存配置，修改配置需要遵循YAML语法，默认配置如下

```
loglevel: INFO
compress: false
#是否开启转换
convert:
  enable: true
  #该配置暂无效
  suffixallow: [doc, docx, xls, xlsx, ppt, pptx, txt, html, htm]
watermark:
#是否开启水印
  enable: true
  #图片目录，如非全路径则默认以程序位置为原位置
  path: "aaa.png"
  #图片是否水平居中，true/false
  widthCenter: true
  #如果设定了水平居中则忽略该项，否则设定为水印图片起始水平位置
  widthPos: 0
  #图片是否垂直居中，true/false
  heightCenter: true
  #如果设定了垂直居中则忽略该项，否则设定为水印图片起始垂直位置
  heightPos: 0
  #图片透明度
  opacity: 0.2
  #图片是否水平拉伸，启用则拉伸至与页面等宽
  scaleToWidth: true
  #图片是否垂直拉伸，启用则拉伸至与页面等长
  scaleToHeight: false
textmark:
  margin:
    #页面属性 四边边距
    left: 0
    right: 0
    top: 0
    bottom: 0
  headArea:
    #是否开启页眉文字添加
    enable: false
    #页眉文字字体，如非全路径则默认以程序位置为原位置
    fontPath: "STXINWEI.TTF"
    contents:
      #配置为列表，每个列表项目以- 开始
        #text是文字内容
      - text: "a啊哈开奖号萨福建省的烦恼了速度快烦死豆奶粉蓝思科技我IE如果屏幕的看病难的sss"
        #字号
        fontSize: 0.1
        #是否启用字体
        useFont: true
        #水平起始位置
        posX0: 50
        #竖直起始位置
        posY0: 20
        #颜色
        rgb: [63,68,76]
        # 特殊类型 ${PageNum} 表示当前页号，${TotalPages} 表示总页数
      - text: "Page ${PageNum} of ${TotalPages}"
        fontSize: 8
        useFont: false
        posX0: 300
        posY0: 20
        rgb: [63,68,76]
      - text: "Convert at ${Date}"
        fontSize: 8
        useFont: true
        posX0: 500
        posY0: 40
        rgb: [63,68,76]
        dateFormat: "2006年01月02日15时"
  footArea:
    enable: true
    fontPath: "STXINWEI.TTF"
    contents:
      #配置为列表，每个列表项目以- 开始
        #text是文字内容
        #特殊格式 ${Date}表示日期
      - text: "文档解释权归技术支持部门所有   ${Date}"
        #字号
        fontSize: 13
        #是否启用字体
        useFont: true
        #水平起始位置
        posX0: 280
        #竖直起始位置
        posY0: 30
        #颜色
        rgb: [63,68,76]
        #日期格式 2016 01 02 15 04 05 分别表示年月日 时分秒，不可用错
        dateFormat: "2006年01月02日15时"
security:
  userpass:
  #是否添加打开密码
    enable: false
  #打开密码设定，不开启需要保留空
    password2add: ""
  ownerpass:
    #是否添加修改密码
    enable: true
    #修改密码设定，不开启需要保留空
    password2add: "Lt12345"
    printing: false
    fullprintquality: false
    modify: false
    annotate: false
    fillforms: false
    rotateinsert: false
    extractgraphics: false
    disabilityextract: false
enabled: true
path: ""
path2: ""

```



