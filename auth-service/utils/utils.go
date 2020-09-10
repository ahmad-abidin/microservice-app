package utils

import (
	"errors"
	"log"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// WELI Warning Error Log and Info
func WELI(notifcode string, errcode string, err error) error {
	switch notifcode {
	case "w":
		log.Printf("[%vWarning%v] (%vcode %v%v) : %v", Yellow, Reset, Yellow, Reset, errcode, err)
		break
	case "e":
		log.Printf("[%vError%v] (%vcode %v%v) : %v", Red, Reset, Red, Reset, errcode, err)
		break
	case "i":
		log.Printf("[%vInfo%v] (%vcode %v%v) : %v", Blue, Reset, Blue, Reset, errcode, err)
		break
	}
	return errors.New(errcode)
}
