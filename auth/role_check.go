package auth

import "fmt"

func RoleCheck(path string, role *int) bool {
	fmt.Println(path, role)
	return true
}
