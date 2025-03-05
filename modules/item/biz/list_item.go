package biz

import (
	"context"
	"social_todo_list/common"
	"social_todo_list/modules/item/model"
)

type ListItemStorage interface {
	List(c context.Context,
		paging *common.Paging,
		filter *model.Filter,
		moreKeys ...string,
	) ([]model.TodoItems, error)
}

type listItemBiz struct {
	storage ListItemStorage
}

func NewListItemBiz(storage ListItemStorage) *listItemBiz {
	return &listItemBiz{storage: storage}
}

func (b *listItemBiz) ListItems(
	c context.Context,
	paging *common.Paging,
	filter *model.Filter,
) ([]model.TodoItems, error) {
	items, err := b.storage.List(c, paging, filter)
	if err != nil {
		return nil, err
	}
	return items, nil
}
