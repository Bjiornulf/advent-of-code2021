#!/bin/bash

dir_name="day$1"
if [ -z "$1" ]
then
	echo "Please provide argument: $0 dayNum"
	exit
fi
mkdir "$dir_name"
cd $dir_name
go mod init
printf "package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\tfmt.Println(\"$dir_name\")\n}" > $dir_name.go
cd ..
