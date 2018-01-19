#!/usr/bin/env python3

import time, subprocess

def runCmd(cmd):
    p = subprocess.Popen(cmd, shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
    stdout = p.communicate()[0].decode('utf-8').strip()
    return stdout

# Assemble build command.
def buildCmd():
    buildFlag = []

    # current time
    buildFlag.append("-X main.compileDate='{}'".format(time.strftime("%Y-%m-%d")))

    return 'go build -ldflags "{}" ./src'.format(" ".join(buildFlag))

if subprocess.call(buildCmd(), shell = True) == 0:
    print("build finished.")