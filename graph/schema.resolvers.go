package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.72

import (
	"context"
	"user-graphql-api/graph/model"
)

// UserInfo is the resolver for the userInfo field.
func (r *queryResolver) UserInfo(ctx context.Context, userID string) (*model.User, error) {
	// 从用户映射中获取用户
	user, exists := r.users[userID]
	if !exists {
		return nil, nil // 如果用户不存在，返回nil
	}
	return user, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
