// Code generated by mockery v2.15.0. DO NOT EDIT.

package repositories

import (
	context "context"

	models "github.com/nronas/invasion_sim/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockAliensRepository is an autogenerated mock type for the AliensRepository type
type MockAliensRepository struct {
	mock.Mock
}

type MockAliensRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAliensRepository) EXPECT() *MockAliensRepository_Expecter {
	return &MockAliensRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, alien
func (_m *MockAliensRepository) Create(ctx context.Context, alien *models.Alien) (*models.Alien, error) {
	ret := _m.Called(ctx, alien)

	var r0 *models.Alien
	if rf, ok := ret.Get(0).(func(context.Context, *models.Alien) *models.Alien); ok {
		r0 = rf(ctx, alien)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Alien)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Alien) error); ok {
		r1 = rf(ctx, alien)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAliensRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAliensRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - alien *models.Alien
func (_e *MockAliensRepository_Expecter) Create(ctx interface{}, alien interface{}) *MockAliensRepository_Create_Call {
	return &MockAliensRepository_Create_Call{Call: _e.mock.On("Create", ctx, alien)}
}

func (_c *MockAliensRepository_Create_Call) Run(run func(ctx context.Context, alien *models.Alien)) *MockAliensRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Alien))
	})
	return _c
}

func (_c *MockAliensRepository_Create_Call) Return(_a0 *models.Alien, _a1 error) *MockAliensRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockAliensRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAliensRepository creates a new instance of MockAliensRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAliensRepository(t mockConstructorTestingTNewMockAliensRepository) *MockAliensRepository {
	mock := &MockAliensRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
