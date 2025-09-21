package tcolor

import (
	"fmt"
)

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

// take error msg and print that
func ErrorMsg(errMsg error) {
	fmt.Printf("%s%s%s \n", BlodRed, errMsg, Reset)
}

// return error msg as string
func SerrorMsg(errMsg error) string {
	if errMsg != nil {
		return fmt.Sprintf("%s%s%s", BlodRed, errMsg, Reset)
	}
	return ""
}

// Print massage. takes color or background and string
func PrintMsg(code, str string) {
	fmt.Printf("%s%s%s \n", code, str, Reset)
}

// return massage as string. takes color or background and string
func SprintMsg(code, str string) string {
	if code == "" || str == "" {
		return str
	}
	return fmt.Sprintf("%s%s%s", code, str, Reset)
}
