package generate

import _ "github.com/golang/mock/gomock"

//go:generate mockgen -destination=mock_foo.go -package=generate . Doer
type Doer interface {
	DoSomething(int, string) error
}
