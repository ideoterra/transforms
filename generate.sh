#!/bin/bash

genny -in=all.go -out=../../xint32/all.go -pkg="xint32"  gen "TA=int32 TB=string"
genny -in=map.go -out=../../xint32/map.go -pkg="xint32"  gen "TA=int32 TB=string"