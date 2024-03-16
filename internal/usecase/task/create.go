package usecase_task

import (
	"context"
	"fmt"
	"fun-with-gin/domain/entity"
)

func (u taskUseCaseInit) CreateOne(ctx context.Context, payload *entity.Task) error {
	fmt.Println(payload)
	//TODO implement me
	panic("implement me")
}
