package storage

import (
	"context"
	"social_todo_list/modules/item/model"
)

func (s *sqlStorage) FindById(c context.Context, id int) (*model.TodoItems, error) {
	var item model.TodoItems
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *sqlStorage) FindByCondition(c context.Context, conditions map[string]interface{}) (*model.TodoItems, error) {
	var item model.TodoItems
	if err := s.db.Where(conditions).First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
