#!/usr/local/bin/python
# This cannot be in go, since you cannot execute a go command from within go
import os

# get data from go-file
with open('default_parameter.go', encoding='utf-8') as fIn:
    text = fIn.read()
    text = text.split('return Parameter{')[1]
    lines = text.split('}')[0].split('\n')
    lines = [i.strip() for i in lines if i.strip()]
    data  = {i.split(":")[0]:i.split(":")[1].replace('"','').replace(',','').strip() for i in lines}

osDict = {
    #label   os          arch     extension
    "linux":["linux",   "amd64", "out"] ,
    "winxp":["windows", "386",   "exe"],
    "win10":["windows", "amd64", "exe"]
}
osList = osDict[data['os'].lower()]

os.system(f"env GOOS={osList[0]} GOARCH={osList[1]} go build -o shuttle.{osList[2]}")

