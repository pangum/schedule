package schedule

import (
	"github.com/pangum/pangu"
	"github.com/pangum/schedule/internal/plugin"
)

func init() {
	pangu.New().Get().Dependency().Build().Provide(new(plugin.Creator).New)
}
