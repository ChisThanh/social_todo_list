package biz

import (
	"context"
	"social_todo_list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	Insert(context.Context, *model.TodoItemCreation) error
}

type createItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *createItemBiz {
	return &createItemBiz{storage: storage}
}

func (b *createItemBiz) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleEmpty
	}

	if err := b.storage.Insert(ctx, data); err != nil {
		return err
	}

	return nil
}
