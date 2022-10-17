package mocks

import (
	"github.com/amir/raidman"
	"github.com/stretchr/testify/mock"
)

type MockRiemannClient struct {
	mock.Mock
}

func (m *MockRiemannClient) Close() {
	m.Called()
}

func (m *MockRiemannClient) Send(event *raidman.Event) error {
	return m.Called(event).Error(0)
}
