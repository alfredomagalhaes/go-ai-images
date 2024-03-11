package view

import (
	"context"

	"github.com/alfredomagalhaes/go-ai-images/types"
)

func AuthenticadeUser(ctx context.Context) types.AuthenticatedUser {
	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{}
	}
	return user
}
