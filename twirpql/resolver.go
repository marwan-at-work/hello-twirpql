// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package twirpql

import (
	"context"

	"marwan.io/hello-twirpql"
)

type Resolver struct {
	hellotwirpql.Greeter
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Hello(ctx context.Context, req *hellotwirpql.HelloReq) (*hellotwirpql.HelloResp, error) {
	return r.Greeter.Hello(ctx, req)
}
