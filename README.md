# doc2pdf

# 功能 [Function]
convert word to pdf with watermark and owner password protect.

用于将word、excel、ppt文件转换为pdf并增加水印和修改限制等功能



# 使用 [Usage] 

0. 需要本机安装office

1. 将需要转换的文件放到doc2pdf目录中来。

2. 将需要转换的文件拖到doc2pdf.exe之上

3. 假设拖入的文件名为XXX.doc，则生成的文件中，不加密不加水印的为XXX.doc.pdf，加密加水印的为Done_XXX.doc.pdf



# 配置

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



