# TColor

## Simple teminal color pkg for golang

### How to Install 
```bash
go get github.com/munnaMia/TColor
```


### Options
```go
const (
	Reset Color = "\033[0m"
	None  Color = ""

	// Foreground Colors
	Black   Color = "\033[30m"
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Magenta Color = "\033[35m"
	Cyan    Color = "\033[36m"
	Gray    Color = "\033[37m"
	White   Color = "\033[97m"

	// Blod Foreground Colors
	BlodRed     Color = "\033[31;1m"
	BlodGreen   Color = "\033[32;1m"
	BlodYellow  Color = "\033[33;1m"
	BlodBlue    Color = "\033[34;1m"
	BlodMagenta Color = "\033[35;1m"
	BlodCyan    Color = "\033[36;1m"
	BlodGray    Color = "\033[37;1m"
	BlodWhite   Color = "\033[97;1m"

	// Background Colors
	BgRed     Color = "\033[41m"
	BgGreen   Color = "\033[42m"
	BgYellow  Color = "\033[43m"
	BgBlue    Color = "\033[44m"
	BgMagenta Color = "\033[45m"
	BgCyan    Color = "\033[46m"
	BgGray    Color = "\033[47m"
	BgWhite   Color = "\033[107m"
)
```

## Examples
```go
package main

import (
	"fmt"

	tcolor "github.com/munnaMia/TColor"
)

func main() {
	// Print string
	tcolor.Println(tcolor.Black, tcolor.BgCyan, "test code epc")

	// Create new String
	txt := tcolor.Sprintf(tcolor.BlodGreen, tcolor.None,  "test code. %d, %v",3,true)
	fmt.Println(txt)

	// Customize text
	text := tcolor.Colorize(tcolor.Yellow, tcolor.BgWhite, "test code epc")
	fmt.Println(text)

	// Print error msg and exit the program
	tcolor.ErrorMsg(fmt.Errorf("this is error"))

}

```
