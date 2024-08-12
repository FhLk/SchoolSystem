package services

import (
	"Library-backend/models"
	"Library-backend/repositories"
	"context"
)

type MemberService interface {
	CreateMember(ctx context.Context, member *models.Member) error
	GetMemberByID(ctx context.Context, id string) (*models.Member, error)
	GetAllMembers(ctx context.Context) ([]*models.Member, error)
	DeleteMember(ctx context.Context, id string) error
	UpdateMember(ctx context.Context, book *models.Member) error
}

type MemberServiceImpl struct {
	memberRepository repositories.MemberRepository
}

func NewMemberService(memberRepository repositories.MemberRepository) *MemberServiceImpl {
	return &MemberServiceImpl{memberRepository: memberRepository}
}

func (s *MemberServiceImpl) CreateMember(ctx context.Context, member *models.Member) error {
	return s.memberRepository.CreateMember(ctx, member)
}

func (s *MemberServiceImpl) GetAllMembers(ctx context.Context) ([]*models.Member, error) {
	return s.memberRepository.GetAllMember(ctx)
}

func (s *MemberServiceImpl) DeleteMember(ctx context.Context, id string) error {
	return s.memberRepository.DeleteMember(ctx, id)
}

func (s *MemberServiceImpl) UpdateMember(ctx context.Context, book *models.Member) error {
	return s.memberRepository.UpdateMember(ctx, book)
}
