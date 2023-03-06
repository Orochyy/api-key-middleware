package repositories

import (
	"api-key-middleware/internal/core/domain"
	"api-key-middleware/internal/core/ports"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
)

var (
	usersTable  = "user"
	authTable   = "auth"
	userProfile = "user_profile"
)

type UserRepository struct {
	conn *sql.DB
	sql  squirrel.StatementBuilderType
}

var _ ports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
		sql:  builderPSQL,
	}
}

func (repo *UserRepository) FindAllProfile(ctx context.Context) ([]*domain.UserProfile, error) {
	rows, err := repo.conn.QueryContext(ctx, `SELECT u.id, u.username, up.first_name, up.last_name,
       up.phone, up.address, up.city, ud.school
		 FROM user u
         JOIN user_profile up ON u.id = up.user_id
         JOIN user_data ud ON u.id = ud.user_id;`)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("can't query db in user repo: %v", err)
	}
	fmt.Println(rows)
	var profiles []*domain.UserProfile
	for rows.Next() {
		profile := &domain.UserProfile{}
		if err := rows.Scan(
			&profile.ID,
			&profile.UserName,
			&profile.FirstName,
			&profile.LastName,
			&profile.Phone,
			&profile.Address,
			&profile.City,
			&profile.School,
		); err != nil {
			return nil, fmt.Errorf("can't scan row from user repo: %v", err)
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (repo *UserRepository) FindByName(ctx context.Context, name string) (*domain.UserProfile, error) {
	rows, err := repo.conn.QueryContext(ctx, ` SELECT
     		u.id, u.username, up.first_name, up.last_name, up.city, ud.school, up.phone, up.address
            FROM user u
            JOIN user_profile up ON u.id = up.user_id
            JOIN user_data ud ON u.id = ud.user_id
            WHERE u.username = ?`, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("can't query db in user repo: %v", err)
	}

	var profile *domain.UserProfile
	for rows.Next() {
		profile = &domain.UserProfile{}
		if err := rows.Scan(
			&profile.ID,
			&profile.UserName,
			&profile.FirstName,
			&profile.LastName,
			&profile.Phone,
			&profile.Address,
			&profile.City,
			&profile.School,
		); err != nil {
			return nil, fmt.Errorf("can't scan row from user repo: %v", err)
		}
	}

	return profile, nil
}
