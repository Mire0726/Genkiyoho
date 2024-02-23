package auth

import (
	"context"
)

type contextKey string

var userIDKey = contextKey("userID")

// SetUserID はContextにユーザID（整数型）を保存します。
func SetUserID(ctx context.Context, userID int) context.Context {
    return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext ContextからユーザIDを取得する
func GetUserIDFromContext(ctx context.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}
