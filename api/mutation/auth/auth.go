package auth

import "github.com/steebchen/graphql/api/super"

type Auth struct {
	*super.Resolver
}

func New(super *super.Resolver) *Auth {
	return &Auth{
		Resolver: super,
	}
}
