package mail

import "github.com/stretchr/testify/mock"

type MailMock struct {
	Mock mock.Mock
}

func (mailMock *MailMock) SendOTP(name, email, code string) error {
	args := mailMock.Mock.Called(name, email, code)
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}
