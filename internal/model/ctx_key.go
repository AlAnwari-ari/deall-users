package model

type contextKey int

const (
	UserIDCtxKey contextKey = iota
	RoleIDCtxKey contextKey = iota
	TokenCtxKey
)
