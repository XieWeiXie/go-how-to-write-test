package part_six

import "fmt"

func One() string {
	var s string
	s += "hello python + \n"
	s += "hello golang + \n"
	return s
}

func Two() string {
	return fmt.Sprintf("%s+\n%s +\n", "hello python", "hello golang")
}
