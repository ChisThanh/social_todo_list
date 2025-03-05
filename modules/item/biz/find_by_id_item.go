package biz

import (
	"context"
	"social_todo_list/modules/item/model"
)

type GetItemStorage interface {
	FindById(context.Context, int) (*model.TodoItems, error)
	FindByCondition(context.Context, map[string]interface{}) (*model.TodoItems, error)
}

type findIdItemBiz struct {
	storage GetItemStorage
}

func NewFindIdItemBiz(storage GetItemStorage) *findIdItemBiz {
	return &findIdItemBiz{storage: storage}
}

func (b *findIdItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItems, error) {
	cond := map[string]interface{}{"id": id}
	item, err := b.storage.FindByCondition(ctx, cond)
	if err != nil {
		return nil, err
	}
	return item, nil
}
