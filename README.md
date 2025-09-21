# TColor

## Color pkg for golang

### How to Install 
```bash
go get github.com/munnaMia/TColor
```
### How to USE
```go
package main

import tcolor "github.com/munnaMia/TColor"

func main(){
    tcolor.PrintMsg(tcolor.BgBlue, "bg color is blue and foreground is white")
}

```

### Color option
```go
const (
	Reset = "\033[0m"

	// Foreground Colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"

	// Blod Foreground Colors
	BlodRed     = "\033[31;1m"
	BlodGreen   = "\033[32;1m"
	BlodYellow  = "\033[33;1m"
	BlodBlue    = "\033[34;1m"
	BlodMagenta = "\033[35;1m"
	BlodCyan    = "\033[36;1m"
	BlodGray    = "\033[37;1m"
	BlodWhite   = "\033[97;1m"

	// Background Colors
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgGray    = "\033[47m"
	BgWhite   = "\033[107m"
)
```

## Always reset the ternal is you use color for you output heiglighting like this
```go
package main
import "fmt"
func main(){
    fmt.Println(tcolor.BgRed + "Your text" + tcolor.Reset) // reset to go back your default terminal
}
```