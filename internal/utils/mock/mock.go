package mock

import "github.com/stretchr/testify/mock"

func ReturnNilOrError(args mock.Arguments, index int) error {
	if len(args) >= index {
		if err, ok := args.Get(index).(error); ok {
			return err
		}
	}

	return nil
}
