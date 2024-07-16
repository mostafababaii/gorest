package interfaces

import "github.com/mostafababaii/gorest/internal/models"

type Tokenizer interface {
	Generate(user *models.User) (string, error)
	Validate(token string) (*models.User, error)
}
