#!/bin/bash

go install github.com/cheekybits/genny

# Generate code for single-type functions.
for file in ./*function.1.go; do
  echo "$file"
	genny -out=../../xstring/"$file" -pkg="xstring" gen "TA=string" <"$file"
done

# Generate code for double-type functions.
for file in ./*function.2.go; do
  echo "$file"
	genny -out=../../xstring/"$file" -pkg="xstring" gen "TA=string TB=BUILTINS" <"$file"
done

# Generate function containers.
for file in ./*container*; do
  echo "$file"
	genny -out=../../xstring/"$file" -pkg="xstring" gen "TA=string" <"$file"
done