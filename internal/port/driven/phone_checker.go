package driven

import "context"

type PhoneChecker interface {
	IsPhoneValid(ctx context.Context, id, phone string) bool
}
