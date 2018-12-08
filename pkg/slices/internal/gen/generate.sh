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
	big_Int
	big_Float
	big_Rat
)

for type_b in ${types[@]}; do 
  type_b_list="$type_b_list,$type_b"
done
type_b_list=${type_b_list#,}

for type_a in ${types[@]}; do
	package=x${type_a/_/}
	echo "  generating transforms for package $package..."

	# Generate code for single-type_a functions.
	for file_in in ./*function.1.go; do
		echo "    $file_in..."
		file_out=../../$package/$file_in
		genny -out=$file_out -pkg=$package gen "TA=$type_a" <"$file_in"
		sed -i "s|$type_a|${type_a/_/.}|g" $file_out
		goimports -w $file_out
	done

	# Generate code for double-type_a functions.
	for file_in in ./*function.2.go; do
		echo "    $file_in..."
		file_out=../../$package/$file_in
		genny -out=$file_out -pkg=$package gen "TA=$type_a TB=$type_b_list" <"$file_in"
		sed -i "s|$type_a|${type_a/_/.}|g" $file_out
    for b in ${types[@]}; do 
		  sed -i "s|$b|${b/_/.}|g" $file_out
    done
		goimports -w $file_out
	done

	# Generate function containers.
	for file in ./*container*; do
		echo "    $file..."
		genny -out=../../$package/$file -pkg=$package gen "TA=$type_a" <"$file"
	done

done

echo "Done."
