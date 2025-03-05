package storage

import (
	"context"
	"social_todo_list/modules/item/model"
)

func (s *sqlStorage) DeleteByCondition(c context.Context, conditions map[string]interface{}) error {
	// Delete record
	// if err := db.Delete(&item).Error; err != nil {
	// 	return
	// }

	if err := s.db.Where(conditions).Delete(&model.TodoItems{}).Error; err != nil {
		return err
	}
	return nil
}
