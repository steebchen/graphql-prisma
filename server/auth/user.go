package auth

import (
	"context"
	"github.com/steebchen/graphql/prisma"
)

func User(ctx context.Context) (*prisma.User, error) {
	user, ok := ctx.Value(UserContextKey).(*prisma.User)

	if !ok {
		// TODO: return graphql error
	}

	return user, nil
}
