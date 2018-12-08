#!/bin/bash

go install github.com/cheekybits/genny

types=(bool rune byte string int int8 int16 int32 int64 float32 float64 uint uint8 uint16 uint32 uint64 complex64 complex128 uintptr error)
for type in ${types[@]}; do

  # Generate code for single-type functions.
  for file in ./*function.1.go; do
    echo "$file"
    genny -out=../../x$type/$file -pkg=x$type gen "TA=$type" <"$file"
  done

  # Generate code for double-type functions.
  for file in ./*function.2.go; do
    echo "$file"
    genny -out=../../x$type/$file -pkg=x$type gen "TA=$type TB=BUILTINS" <"$file"
  done

  # Generate function containers.
  for file in ./*container*; do
    echo "$file"
    genny -out=../../x$type/$file -pkg=x$type gen "TA=$type" <"$file"
  done

done