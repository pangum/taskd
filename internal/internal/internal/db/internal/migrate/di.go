package migrate

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newM1InitSchedule,
	).Group("migrations").Build().Apply()
}
