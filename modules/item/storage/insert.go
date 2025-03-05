package storage

import (
	"context"
	"social_todo_list/common"
	"social_todo_list/modules/item/model"
)

func (s *sqlStorage) Insert(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
