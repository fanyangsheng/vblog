package comment

import "context"

type impl struct {
}

func (*impl) CreateComment(ctx context.Context, comment *Comment) error {
	return nil
}
