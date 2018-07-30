package greet

import  "strings"

func Greet(name string) string {
	if (len(name) == 0) {
		return "HELLO WORLD!"
	} else {
		name = strings.ToUpper(name)
		return "HELLO " + name + "!"
	}
}
