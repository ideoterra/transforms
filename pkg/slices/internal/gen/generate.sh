#!/bin/bash

go install github.com/cheekybits/genny

files=$(ls ./*.go | grep _ -v | grep type -v)
for file in $files; do
  echo "$file"
	genny -out=../../xstring/"$file" -pkg="xstring" gen "TA=string TB=BUILTINS" <"$file"
done

files=$(ls *type* | grep types -v)
for file in $files; do
  echo "$file"
	genny -out=../../xstring/"$file" -pkg="xstring" gen "TA=string" <"$file"
done