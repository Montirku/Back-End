package bengkel

import (
	"gorm.io/gorm"
)

type BengkelCategory struct {
	*gorm.Model `json:"-"`
	Category    string `json:"Category" form:"Category"`
	Bengkels    []Bengkel
}
