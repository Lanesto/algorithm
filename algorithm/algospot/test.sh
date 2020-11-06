#!/bin/bash
# A simple test automation shell script

# Arguments:
# $1: filename like: {problem}_{lang}.{extension}
#     e.g) BOGGLE_go.go
program=$1
problem=$(echo $1 | sed -E 's/(.*)_.*\..*/\1/')

# Automatically find datafile; name should be like: {problem}.dat
# MUST datafile have block [in], [out] for work becuz no exception is handled
# [in]
# a
# b
# [out]
# 1
# e.g) BOGGLE.dat
data_file="$problem.dat"
if [ ! -f $data_file ]; then
	echo "Requires data file but does not exists: $data_file"
	exit 1
fi

echo "Testing $program with $data_file"

# Separate in/out to temporary file
input_tmpfile=$(mktemp)
sed -n '/[in]/,/[out]/ p' $data_file | sed -e '2,$!d' -e '$d' > $input_tmpfile
output_tmpfile=$(mktemp)
sed -n '/[out]/,$ p' $data_file | sed -n '2,$ p' > $output_tmpfile

# Run program and get output
# For testing the program should write answers to stdout
result_tmpfile=$(mktemp)
extension=$(echo $1 | sed 's/.*\.//')
if [ "$extension" == "go" ]; then
	go run $program < $input_tmpfile > $result_tmpfile
	diff -q $result_tmpfile $output_tmpfile &> /dev/null
	if [ "$?" == "0" ]; then
		echo "Test is done."
	else
		echo "Test failed."
		echo
		echo | diff $result_tmpfile $output_tmpfile
		echo
		exit 1
	fi
# elif ...
else
	echo "No runner for program(.$extension) is not defined."
fi

