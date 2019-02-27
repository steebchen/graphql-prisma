package mutation

import (
	"github.com/steebchen/graphql/api/mutation/auth"
	"github.com/steebchen/graphql/api/super"
	"github.com/steebchen/graphql/gqlgen"
)

type Mutation struct {
	*super.Resolver
	*auth.Auth
}

func New(super *super.Resolver) gqlgen.MutationResolver {
	return &Mutation{
		Resolver: super,
		Auth:     auth.New(super),
	}
}
