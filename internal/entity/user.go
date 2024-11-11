package entity

import "time"

type (
	User struct {
		ID        string     `db:"id,omitempty" json:"id,omitempty"`
		Name      string     `db:"name,omitempty" json:"name,omitempty"`
		Email     string     `db:"email,omitempty" json:"email,omitempty"`
		Password  string     `db:"password,omitempty" json:"password,omitempty"`
		CreatedAt time.Time  `db:"created_at,omitempty" json:"created_at,omitempty"`
		UpdatedAt time.Time  `db:"updated_at,omitempty" json:"updated_at,omitempty"`
		DeletedAt *time.Time `db:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	}
)
