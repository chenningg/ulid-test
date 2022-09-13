//go:generate go run github.com/dmarkham/enumer -type=AuthRole -json -text -sql

package auth

import (
	"fmt"
)

type AuthRole int

const (
	Demo AuthRole = iota + 1
	Free
	Plus
	Pro
	Enterprise
	Support
	Admin
	SuperAdmin
)

func (authRole *AuthRole) Validate() error {
	if authRole.IsAAuthRole() {
		return nil
	}

	return fmt.Errorf("invalid value for AuthRole, expected %v, got %v", AuthRoleValues(), authRole.String())
}

func (authRole AuthRole) Values() []string {
	return AuthRoleStrings()
}
