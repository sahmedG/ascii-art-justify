# ascii-art-color

This program will color the output of ascii-art.
The program will not accept any invalid words (that are above the decimal ascii value 126)

## Avialable colors
If you want to use `string` color the available are:
`black, red, green, yellow, blue, magenta, cyan, white` and `orange`

## Usage
`Usage: go run . [OPTION] [STRING]`


## Example
`EX: go run . --color=<color> <letters to be colored> "something"`

## Notes
- Letters to be colored will color the letters of the input string in sequence. If the sequence does does not match it will not color it.
- If color is not available it will output the art with default color.