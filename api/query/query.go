package query

import (
	"context"
	"github.com/steebchen/graphql/api/super"
	"github.com/steebchen/graphql/prisma"
	"github.com/steebchen/graphql/server/session_context"
)

type Resolver struct {
	*super.Resolver
}

func (r *Resolver) User(ctx context.Context) (prisma.User, error) {
	user, err := session_context.User(ctx)

	// needed because the return type cannot be nil
	if err != nil {
		return prisma.User{}, err
	}

	return *user, err
}
