package schedule

import (
	"github.com/pangum/pangu"
	"github.com/pangum/schedule/internal/plugin"
)

func init() {
	creator := new(plugin.Constructor)
	pangu.New().Get().Dependency().Put(
		creator.New,
	).Build().Build().Apply()
}
