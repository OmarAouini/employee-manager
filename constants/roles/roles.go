package roles

import "strings"

type Roles int

const (
	ADMIN Roles = iota
	MANAGER
	EMPLOYEE
	rolesLimit
)

func (r Roles) String() string {
	strings := [...]string{"ADMIN", "MANAGER", "EMPLOYEE"}

	// prevent panicking in case of status is out-of-range
	if r < ADMIN || r > EMPLOYEE {
		return "Unknown"
	}
	return strings[r]
}

func (r Roles) EnumIndex() int {
	return int(r)
}

func IsPresent(value string) bool {
	for role := Roles(0); role < rolesLimit; role++ {
		if strings.ToUpper(value) == role.String() {
			return true
		}
	}
	return false
}
