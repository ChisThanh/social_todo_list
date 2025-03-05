package biz

import (
	"context"
	"social_todo_list/modules/item/model"
)

type UpdateItemStorage interface {
	FindById(context.Context, int) (*model.TodoItems, error)
	UpdateByCondition(context.Context, *model.TodoItemUpdate, map[string]interface{}) error
}

type updateItemBiz struct {
	storage UpdateItemStorage
}

func NewUpdateItemBiz(storage UpdateItemStorage) *updateItemBiz {
	return &updateItemBiz{storage: storage}
}

func (b *updateItemBiz) UpdateItem(ctx context.Context, data *model.TodoItemUpdate, id int) error {
	item, err := b.storage.FindById(ctx, id)
	if err != nil {
		return model.ErrItemNotFound
	}

	if item.Status != nil && *item.Status == model.StatusDelete {
		return model.ErrItemDeleted
	}

	cond := map[string]interface{}{"id": id}
	if err := b.storage.UpdateByCondition(ctx, data, cond); err != nil {
		return err
	}

	return nil
}
