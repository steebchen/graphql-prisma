package api

import (
	"github.com/steebchen/graphql/api/mutation"
	"github.com/steebchen/graphql/api/query"
	"github.com/steebchen/graphql/api/super"
	"github.com/steebchen/graphql/gqlgen"
	"github.com/steebchen/graphql/prisma"
)

type Resolver struct {
	*super.Resolver
}

func New(client *prisma.Client) *Resolver {
	return &Resolver{
		Resolver: &super.Resolver{
			Prisma: client,
		},
	}
}

func (r *Resolver) Mutation() gqlgen.MutationResolver {
	return mutation.New(r.Resolver)
}

func (r *Resolver) Query() gqlgen.QueryResolver {
	return query.New(r.Resolver)
}
