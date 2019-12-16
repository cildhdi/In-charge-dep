package auth

import "github.com/cildhdi/In-charge/models"

var privileges = map[string][]int{
	"/api/auth/reachable":   []int{models.SuperUser, models.AdminUser, models.CustomerUser, models.MerChantUser},
	"/api/auth/unreachable": []int{},
	"/api/admin-register":   []int{models.SuperUser},
}

func RoleCheck(path string, role int) bool {
	if v, ok := privileges[path]; ok {
		for _, roleValue := range v {
			if roleValue == role {
				return true
			}
		}
	}
	return false
}
