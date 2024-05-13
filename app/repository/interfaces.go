package repository

import (
	"github.com/amirrstm/go-clinique/app/dto"
)

type UserRepository interface {
	Create(b *dto.CreateUser) error
}
