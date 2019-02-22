package mutation

import (
	"context"
	"github.com/steebchen/graphql/api/super"
	"github.com/steebchen/graphql/prisma"
)

type Mutation struct {
	*super.Resolver
}

func (*Mutation) Signup(ctx context.Context, email string, name string) (prisma.User, error) {
	panic("implement me")
}
