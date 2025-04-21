package graph

import (
	"user-graphql-api/graph/model"
)

type Resolver struct {
	users map[string]*model.User
}

// NewResolver 创建一个新的resolver并初始化用户数据
func NewResolver() *Resolver {
	users := make(map[string]*model.User)
	users["aaa"] = &model.User{
		ID:   "aaa",
		Name: "张三",
		Age:  28,
		Sex:  "男",
	}

	return &Resolver{
		users: users,
	}
}
