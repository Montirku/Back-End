package bengkel

import (
	bc "github.com/fazaalexander/montirku-be/modules/usecase/user/bengkel"
)

type BengkelHandler struct {
	bengkelUseCase bc.BengkelUseCase
}

func New(bengkelUseCase bc.BengkelUseCase) *BengkelHandler {
	return &BengkelHandler{
		bengkelUseCase,
	}
}
