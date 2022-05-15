package user

import "context"

type User struct {
	ID    string
	Name  string
	Email string
}

const userKey = "user"

func FromContext(ctx context.Context) *User {
	usr, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil
	}

	return usr
}
