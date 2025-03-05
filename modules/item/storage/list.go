package storage

import (
	"context"
	"social_todo_list/common"
	"social_todo_list/modules/item/model"
)

func (s *sqlStorage) List(c context.Context,
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]model.TodoItems, error) {
	var results []model.TodoItems

	query := s.db.WithContext(c)
	query = query.Where("status <> ?", "Deleted")

	if filter != nil {
		if v := filter.Status; v != "" {
			query = query.Where("status = ?", v)
		}
	}

	if len(moreKeys) > 0 {
		query = query.Preload(moreKeys[0])
	}

	if err := query.Table(model.TodoItems{}.TableName()).
		Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := query.Order("id desc").
		Offset(paging.Offset()).
		Limit(paging.Limit).
		Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
