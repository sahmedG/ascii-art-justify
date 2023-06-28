#!/bin/bash

#* test file for multiple diff cases

go run . "\"test\""

echo

go run . "HELLO"

echo

go run . "HeLlo HuMaN"

echo

go run . "1Hello 2There"

echo


go run . "Hello\nThere"

echo

go run . --color=red "{Hello & There #}"

echo

go run . --color=yellow 'hello' 'hello There 1 to 2!' 

echo

go run . --output=output.txt "MaD3IrA&LiSboN"

echo

go run . "1a\"#FdwHywR&/()="

echo


go run . "{|}~"

echo 

go run . --color=green --output=output2.txt "dwH" "1a\"#FdwHywR&/()="

echo


go run . ":;<=>?@"

echo

go run . "[\]^_ 'a"

echo

go run . '\!" #$%&'"'"'()*+,-./'

echo


go run . "RGB"

echo

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

echo

go run . "abcdefghijklmnopqrstuvwxyz"