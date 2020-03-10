package usecase

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"local.packages/domain"
	"local.packages/domain/repository"
)

// UserにおけるUseCaseのインターフェース
type UserUseCase interface {
	GetByUserID(DB *sql.DB, userID string) (domain.User, error)
	Insert(DB *sql.DB, userID, name, email string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// Userデータに対するusecaseを生成
// NewとUserUseCaseインターフェースがinfrastructure層にusecaseを依存させない鍵になるもの
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu UserUseCase) GetByUserID(DB *sql.DB, userID string) (domain.User, error) {
	user, err := uu.userRepository.GetByUserID(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu UserUseCase) Insert(DB *sql.DB, userID, name, email string) error {
	// TODO: emailの形式が正しいかバリデーションを入れる

	// userIDが一意な値を持つようにする
	userID, err := uuid.NewRandom() // 返り血はuuid型(uuidからNewしてるから)
	if err != nil {
		return err
	}

	// domainを介してinfrastructureで実装した関数を呼び出す
	err := uu.userRepository.Insert(DB *sql.DB, userID.String(), name, email)
	if err != nil {
		return err
	}
	return nil
}
