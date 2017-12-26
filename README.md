# doc2pdf

# 功能 [Function]
Use OLE to convert a word/excel/powerpoint file to pdf ,and add  watermark and owner password protect to the pdf.

---

通过OLE使用OFFICE等将word、excel、ppt文件转换为pdf
并对PDF增加水印和修改限制等功能
常用于发布一些不希望别人修改的内容



# 使用 [Usage]
0. install Office

1. Put the file to convert in the .exe directory

2. drag the file to doc2pdf.exe

3. if the file drag in is XXX.doc，then it will produce an unlock file name XXX.doc.pdf and a locked file Done_XXX.doc.pdf

4. if a pdf is dragged in, A locked pdf will generate then. 

---

0. 需要本机安装office 

1. 将需要转换的文件放到doc2pdf目录中来。 

2. 将需要转换的文件拖到doc2pdf.exe之上

3. 假设拖入的文件名为XXX.doc，则生成的文件中，不加密不加水印的为XXX.doc.pdf，加密加水印的为Done_XXX.doc.pdf

4. 拖入未加密的pdf文件，会生成加密加水印的pdf文件

# 配置
Use the conf.yaml to stroe the configuration info.The Default one is below.

---

使用conf.yaml保存配置，修改配置需要遵循YAML语法，默认配置如下

```
loglevel: INFO
compress: false
convert:
#是否开启转换
  enable: true
  suffixallow: [doc, docx, xls, xlsx, ppt, pptx]
watermark:
#是否开启水印
  enable: true
  path: "aaa.png"
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



