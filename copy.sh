#!/bin/bash
# Copy files in template/ with target directory name replaced
set -e

mkdir -p "$1"

target_dir="$1"
filename="${1##*/}"
basename="${filename%.*}"

placeholder='PROBLEM'

for filename in ./template/*; do
    if [ ! -f "${filename}" ]; then
        continue
    fi
    new_filename=$(echo ${filename##*/} | sed "s/$placeholder/$basename/")
    new_filename="$target_dir/$new_filename"
    cp -n "$filename" "$new_filename"
done