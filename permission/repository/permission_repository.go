package repository

import (
	"jargonjester/domain"
	"jargonjester/entity"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewConversationRepository(db *gorm.DB) domain.PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) AddChannelPermission(channelID string) error {
	channelPermission := entity.ChannelPermission{ChannelID: channelID}
	result := r.db.First(&channelPermission)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			channelPermission.Active = true
			err := r.db.Create(&channelPermission).Error
			if err != nil {
				return err
			}
			return nil
		}
		return result.Error
	}
	channelPermission.Active = true

	err := r.db.Create(&channelPermission).Error

	if err != nil {
		return err
	}

	return nil
}
func (r *permissionRepository) RemoveChannelPermission(string) error { return nil }
func (r *permissionRepository) CheckChannelPermission(string) bool   { return false }
func (r *permissionRepository) AdduserPermission(string) error       { return nil }
func (r *permissionRepository) RemoveUserPermission(string) error    { return nil }
func (r *permissionRepository) CheckUserPermission(string) bool      { return false }
