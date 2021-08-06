package schedule

import (
	`strconv`
)

var _ option = (*optionId)(nil)

type optionId struct {
	id string
}

// Id 配置编号
func Id(id int64) *optionId {
	return StringId(strconv.FormatInt(id, 10))
}

// StringId 配置字符串形式的编号
func StringId(id string) *optionId {
	return &optionId{id: id}
}

func (i *optionId) apply(options *options) {
	options.id = i.id
}
