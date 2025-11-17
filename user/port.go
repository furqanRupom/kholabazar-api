package user
import (
	"kholabazar/domain"
	userHandler "kholabazar/rest/handlers/user"
)
type Service interface {
	userHandler.Service
}
type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, password string) (*domain.User, error)
}