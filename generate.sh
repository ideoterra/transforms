#!/bin/bash

genny -in=gen.go -out=../../xint32/api.go -pkg="xint32" gen "AA=[]int32 A=int32 BB=[]string B=string"