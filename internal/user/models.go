package user

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type AccessRole string

const (
	UserRole  AccessRole = "USER"
	AdminRole AccessRole = "ADMIN"
	GuestRole AccessRole = "GUEST"
)

func (r AccessRole) IsValid() bool {
	switch r {
	case UserRole, AdminRole, GuestRole:
		return true
	}
	return false
}

type AccessRoles []AccessRole

func (roles AccessRoles) Value() (driver.Value, error) {
	if len(roles) == 0 {
		return "{}", nil
	}

	var roleStrings []string
	for _, role := range roles {
		if !role.IsValid() {
			return nil, fmt.Errorf("Invalid access role: %s", role)
		}
		roleStrings = append(roleStrings, string(role))
	}
	return "{" + strings.Join(roleStrings, ",") + "}", nil
}

func (roles *AccessRoles) Scan(src interface{}) error {
	if src == nil {
		*roles = AccessRoles{}
		return nil
	}

	var data string
	switch v := src.(type) {
	case string:
		data = v
	case []byte:
		data = string(v)
	default:
		return fmt.Errorf("Cannot scan %T into AccessRoles", src)
	}

	data = strings.Trim(data, "{}")
	if data == "" {
		*roles = AccessRoles{}
		return nil
	}

	parts := strings.Split(data, ",")
	result := make(AccessRoles, len(parts))
	for i, part := range parts {
		role := AccessRole(strings.TrimSpace(part))
		if !role.IsValid() {
			return fmt.Errorf("Invalid role in database: %s", role)
		}
		result[i] = role
	}

	*roles = result
	return nil
}

func (roles AccessRoles) HasRole(role AccessRole) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func (roles *AccessRoles) AddRole(role AccessRole) error {
	if !role.IsValid() {
		return fmt.Errorf("Invalid role: %s", role)
	}
	if !roles.HasRole(role) {
		*roles = append(*roles, role)
	}
	return nil
}

func (roles *AccessRoles) RemoveRole(role AccessRole) {
	for i, r := range *roles {
		if r == role {
			*roles = append((*roles)[:i], (*roles)[i+1:]...)
			break
		}
	}
}

type User struct {
	ID          uint
	Username    string `gorm:"uniqueIndex"`
	Email       string
	Birthday    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Password    string
	AccessRoles AccessRoles `gorm:"type:text[]"`
}
