package contextext

import (
	"context"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/typesext"
)

func GetValuesForKeys(ctx context.Context, keys []typesext.ContextKey) (values []any) {
	for _, key := range keys {
		values = append(values, ctx.Value(key))
	}
	return values
}
