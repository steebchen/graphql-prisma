package query

import (
	"context"
	"github.com/steebchen/graphql/api/super"
	"github.com/steebchen/graphql/lib/session_context"
	"github.com/steebchen/graphql/prisma"
)

type Resolver struct {
	*super.Resolver
}

func New(super *super.Resolver) *Resolver {
	return &Resolver{
		Resolver: super,
	}
}

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	user, err := session_context.User(ctx)

	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
