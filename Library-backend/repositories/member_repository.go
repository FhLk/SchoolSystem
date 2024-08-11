package repositories

import (
	"Library-backend/models"
	"context"
	"gorm.io/gorm"
)

type MemberRepository interface {
	CreateMember(ctx context.Context, member *models.Member) error
	GetMemberByID(ctx context.Context, id string) (*models.Member, error)
	GetAllMember(ctx context.Context) ([]*models.Member, error)
	DeleteMember(ctx context.Context, id string) error
	UpdateMember(ctx context.Context, member *models.Member) error
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) CreateMember(ctx context.Context, member *models.Member) error {
	return r.db.Create(member).Error
}

func (r *memberRepository) GetMemberByID(ctx context.Context, id string) (*models.Member, error) {
	var member models.Member
	if err := r.db.Find(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) GetAllMember(ctx context.Context) ([]*models.Member, error) {
	var members []*models.Member
	if err := r.db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) DeleteMember(ctx context.Context, id string) error {
	var member models.Member
	if err := r.db.Delete(&member, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *memberRepository) UpdateMember(ctx context.Context, member *models.Member) error {
	return r.db.Save(member).Error
}
