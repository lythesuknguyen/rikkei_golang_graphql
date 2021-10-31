package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lythesuknguyen/rikkei/graph/model"
)

type MemberRepo struct {
	DB *pg.DB
}

func (m *MemberRepo) getMembers() ([]*model.Member, error) {
	var members []*model.Member
	err := m.DB.Model(&members).Select()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (m *MemberRepo) getMembersByID(id string) (*model.Member, error) {
	var member model.Member
	err := m.DB.Model(&member).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *MemberRepo) getMembersByIDs(ids []string) ([]*model.Member, error) {
	var members []*model.Member
	err := m.DB.Model(&members).Where("id IN ?", ids).Distinct()
	return members, nil
}
