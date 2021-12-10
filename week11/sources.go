package news

import (
	"context"
)

type Source interface {
	ConnectSource(context.Context) (<-chan Article)
	Stop() // error?
}
