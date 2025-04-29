package migrate

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newInitializer,
	).Build().Apply()
}
