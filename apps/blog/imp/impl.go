package blog

import "context"

type impl struct {
}

func (*impl) CreateBlog(context.Context, *Blog) error {
	return nil
}

//func (*impl) UpdateBlog(context.Context, *Blog) error {
//	return nil
//}
//
//func (*impl) DeleteBlog(context.Context, *Blog) error {
//	return nil
//}
//
//func (*impl) QueryBlog(ctx context.Context, blog *Blog) error {
//	return nil
//}
//
//func (*impl) DescribeBlog(context.Context, *Blog) error {
//	return nil
//}
