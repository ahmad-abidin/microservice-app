package model

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

// Log Warning Error Log and Info
func Log(notifcode string, errcode string, err error) error {
	switch notifcode {
	case "w":
		log.Printf("[%vWarning%v] (code%v %v%v) : %v", Yellow, Reset, Yellow, errcode, Reset, err)
		break
	case "e":
		log.Printf("[%vError%v] (code%v %v%v) : %v", Red, Reset, Red, errcode, Reset, err)
		break
	case "i":
		log.Printf("[%vInfo%v] (code%v %v%v) : %v", Blue, Reset, Blue, errcode, Reset, err)
		break
	case "s":
		log.Printf("[%vSuccess%v] (code%v %v%v) : %v", Green, Reset, Green, errcode, Reset, err)
		break
	}

	return errors.New(errcode)
}
