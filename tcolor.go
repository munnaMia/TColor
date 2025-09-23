package tcolor

import (
	"fmt"
	"os"
)

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

type Color string

func (c Color) String() string {
	return string(c)
}

func (c Color) With(str string) string {
	return fmt.Sprintf("%s%s%s", c, str, Reset)
}

// takes a string with foreground and a background color and transfrom it.
func Colorize(frColor, bgColor Color, str string) string {
	return fmt.Sprintf("%s%s%s%s", frColor, bgColor, str, Reset)
}

// Print an error massage with Blod red color and exit with status 1
func ErrorMsg(err error) {
	fmt.Printf("%s%s%s \n", BlodRed, err.Error(), Reset)
	os.Exit(1)
}

// Take a massage & foreground with background color for the massage and print that
func Println(frColor, bgColor Color, str string) {
	fmt.Printf("%s%s%s%s \n", frColor, bgColor, str, Reset)
}

// Take a massage & foreground with background color for the massage. and return a customized string.
func Sprintf(frColor, bgColor Color, format string, a ...any) string {
	return fmt.Sprintf("%s%s%s%s", frColor, bgColor, fmt.Sprintf(format, a...), Reset)
}
