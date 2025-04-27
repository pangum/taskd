package migrate

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newInitializer,
	).Build().Apply()
}
