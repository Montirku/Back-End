package common

import (
	ah "github.com/fazaalexander/montirku-be/modules/handler/api/auth"
	bh "github.com/fazaalexander/montirku-be/modules/handler/api/user/bengkel"
)

type Handler struct {
	AuthHandler    *ah.AuthHandler
	BengkelHandler *bh.BengkelHandler
}
