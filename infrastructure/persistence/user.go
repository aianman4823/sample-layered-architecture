package persistence

import (
	"local.packages/domain"
	"local.packages/domain/repository"
	"database/sql"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// ユーザ登録処理
func (up userPersistence) Insert(DB *sql.DB, userID, name, email string) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
	if err != nil {
		return nil
	}

	_, err = stmt.Exec(userID, name, email)
	return err
}

// userIDによってユーザー情報を取得する処理
func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM user where user_id = ?", userID)

	// row型をgolangで利用できる形にキャストする
	return convertToUser(row)
}
//         ↑
// row型をuser型に紐付ける
func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	// domain.User{}にマッチする構造になっているかチェック
	err := row.Scan(&user.UserID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// 何も返さないのは型チェックに反するからnilを返す(構造体でもnilは持てる)
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
