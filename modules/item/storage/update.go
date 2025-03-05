package storage

import (
	"context"
	"social_todo_list/modules/item/model"
)

func (s *sqlStorage) UpdateByCondition(c context.Context, data *model.TodoItemUpdate, conditions map[string]interface{}) error {
	if err := s.db.Where(conditions).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
