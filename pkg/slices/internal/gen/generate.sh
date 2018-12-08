#!/bin/bash

echo "Ensuring genny is installed..."
go install github.com/cheekybits/genny

echo "Removing old generated transforms..."
for x in ../../x*/; do
  rm -drf "$x"
done

echo "Generating new transforms..."
types=(
	bool
	rune
	byte
	string
	int
	int8
	int16
	int32
	int64
	float32
	float64
	uint
	uint8
	uint16
	uint32
	uint64
	complex64
	complex128
	uintptr
	error
)

for type_a in ${types[@]}; do
	for type_b in ${types[@]}; do
    echo "  generating transforms from $type_a to $type_b..."
		# Generate code for single-type_a functions.
		for file in ./*function.1.go; do
			echo "    $file..."
			genny -out=../../x$type_a/$file -pkg=x$type_a gen "TA=$type_a" <"$file"
		done

		# Generate code for double-type_a functions.
		for file in ./*function.2.go; do
			echo "    $file..."
			genny -out=../../x$type_a/$file -pkg=x$type_a gen "TA=$type_a TB=$type_b" <"$file"
		done

		# Generate function containers.
		for file in ./*container*; do
			echo "    $file..."
			genny -out=../../x$type_a/$file -pkg=x$type_a gen "TA=$type_a" <"$file"
		done

	done
done

echo "Done."
