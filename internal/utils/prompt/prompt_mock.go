package prompt

import (
	"github.com/stretchr/testify/mock"
	utilsMock "github.com/wilian746/go-generator/internal/utils/mock"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Ask(label, defaultValue string) (string, error) {
	args := m.MethodCalled("Ask")
	return args.Get(0).(string), utilsMock.ReturnNilOrError(args, 1)
}
