@rem go build -ldflags "-w -s" ./src
py build.py
upx src.exe
del doc2pdf.exe
copy src.exe doc2pdf.exe