package biz

import (
	"context"
	"social_todo_list/modules/item/model"
)

type DeleteItemStorage interface {
	FindById(context.Context, int) (*model.TodoItems, error)
	DeleteByCondition(context.Context, map[string]interface{}) error
}

type deleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *deleteItemBiz {
	return &deleteItemBiz{storage: storage}
}

func (b *deleteItemBiz) DeleteItem(ctx context.Context, id int) error {
	item, err := b.storage.FindById(ctx, id)
	if err != nil {
		return model.ErrItemNotFound
	}

	if item.Status != nil && *item.Status == model.StatusDelete {
		return model.ErrItemDeleted
	}

	cond := map[string]interface{}{"id": id}
	if err := b.storage.DeleteByCondition(ctx, cond); err != nil {
		return err
	}

	return nil
}
