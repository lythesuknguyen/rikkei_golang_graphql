package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lythesuknguyen/rikkei/graph/model"
)

type SkillRepo struct {
	DB *pg.DB
}

func (s *SkillRepo) getSkills() ([]*model.Skill, error) {
	var skills []*model.Skill
	err := s.DB.Model(&skills).Select()
	if err != nil {
		return nil, err
	}
	return skills, nil
}
