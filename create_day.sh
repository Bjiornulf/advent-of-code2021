dir_name="day$1"
if [ -z "$1" ]
then
	echo "Please provide argument: $0 dayNum"
	exit
fi
mkdir "$dir_name"
cd $dir_name

printf "package main

import (
	\"fmt\"
	\"aoc2021/utils\"
)

func main() {
	fmt.Println(\"$dir_name\")
	input := utils.ReadLines(\"input\")

	/* ---------- Puzzle 1 ---------- */

	/* ---------- Puzzle 2 ---------- */
}" > $dir_name.go

if [ -f "../.session" ]
then
	curl --cookie "session=$(cat ../.session)" "https://adventofcode.com/2021/day/$1/input" > input
else
	touch input
fi

cd ..
