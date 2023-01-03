package repository

import (
	"context"

	"github.com/deall-users/internal/model"
	orm "github.com/deall-users/pkg/gorm"
	"gorm.io/gorm"
)

func NewUserRepo(db *orm.DbPgSql) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	db *orm.DbPgSql
}

func (rp userRepository) FindAllUsers(ctx context.Context, pagination model.Pagination, criteria map[string]interface{}) ([]*model.User, error) {
	db := rp.db.Db
	users := make([]*model.User, 0)

	err := db.Model(&model.User{}).Select("users.*, r.role_name").
		Joins("LEFT JOIN roles r ON users.role_id = r.role_id").Where(criteria).
		Order("user_id DESC").Limit(pagination.Limit).Offset(pagination.Offset()).Find(&users).Error

	if err != nil {
		return nil, rp.db.WrapErrorGorm(ctx, err, "FindAllUsers")
	}

	return users, nil
}

func (rp userRepository) FindSingleUser(ctx context.Context, criteria map[string]interface{}) (*model.User, error) {
	db := rp.db.Db
	user := model.User{}

	err := db.Model(&model.User{}).Select("users.*, r.role_name").
		Joins("LEFT JOIN roles r ON users.role_id = r.role_id").Where(criteria).
		Take(&user).Error

	if err != nil {
		return nil, rp.db.WrapErrorGorm(ctx, err, "FindSingleUser")
	}

	return &user, nil
}

func (rp userRepository) FindUserByID(ctx context.Context, userID uint64) (*model.User, error) {
	return rp.FindSingleUser(ctx, map[string]interface{}{"user_id": userID})
}

func (rp userRepository) CountUsers(ctx context.Context, criteria map[string]interface{}) (int, error) {
	db := rp.db.Db
	var total int64

	err := db.Model(&model.User{}).Where(criteria).Count(&total).Error

	if err != nil {
		return 0, rp.db.WrapErrorGorm(ctx, err, "CountUsers")
	}

	return int(total), nil
}

func (rp userRepository) CreateUser(ctx context.Context, user *model.User) error {
	db := rp.db.Db

	err := db.Create(user).Error
	if err != nil {
		return rp.db.WrapErrorGorm(ctx, err, "CreateUser")
	}

	return nil
}

func (rp userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	db := rp.db.Db

	db = db.Updates(user)
	if db.Error != nil {
		return rp.db.WrapErrorGorm(ctx, db.Error, "UpdateUser")
	}

	if db.RowsAffected == 0 {
		return rp.db.WrapErrorGorm(ctx, gorm.ErrRecordNotFound, "UpdateUser")
	}

	return nil
}

func (rp userRepository) DeleteUser(ctx context.Context, user *model.User) error {
	db := rp.db.Db

	db = db.Delete(user)
	if db.Error != nil {
		return rp.db.WrapErrorGorm(ctx, db.Error, "DeleteUser")
	}

	if db.RowsAffected == 0 {
		return rp.db.WrapErrorGorm(ctx, gorm.ErrRecordNotFound, "DeleteUser")
	}

	return nil
}

func (rp userRepository) FindSingleRole(ctx context.Context, criteria map[string]interface{}) (*model.Role, error) {
	db := rp.db.Db
	role := model.Role{}

	err := db.Model(&model.Role{}).Where(criteria).Take(&role).Error

	if err != nil {
		return nil, rp.db.WrapErrorGorm(ctx, err, "FindSingleRole")
	}

	return &role, nil
}

func (rp userRepository) FindRoleByID(ctx context.Context, roleID uint64) (*model.Role, error) {
	return rp.FindSingleRole(ctx, map[string]interface{}{"role_id": roleID})
}
