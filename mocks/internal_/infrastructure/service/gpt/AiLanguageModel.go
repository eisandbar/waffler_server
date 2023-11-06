// Code generated by mockery v2.36.0. DO NOT EDIT.

package gpt

import (
	openai "github.com/sashabaranov/go-openai"
	mock "github.com/stretchr/testify/mock"
)

// AiLanguageModel is an autogenerated mock type for the AiLanguageModel type
type AiLanguageModel struct {
	mock.Mock
}

type AiLanguageModel_Expecter struct {
	mock *mock.Mock
}

func (_m *AiLanguageModel) EXPECT() *AiLanguageModel_Expecter {
	return &AiLanguageModel_Expecter{mock: &_m.Mock}
}

// QuestionForGPT provides a mock function with given fields: answer
func (_m *AiLanguageModel) QuestionForGPT(answer string) (*openai.ChatCompletionResponse, error) {
	ret := _m.Called(answer)

	var r0 *openai.ChatCompletionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*openai.ChatCompletionResponse, error)); ok {
		return rf(answer)
	}
	if rf, ok := ret.Get(0).(func(string) *openai.ChatCompletionResponse); ok {
		r0 = rf(answer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*openai.ChatCompletionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(answer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AiLanguageModel_QuestionForGPT_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QuestionForGPT'
type AiLanguageModel_QuestionForGPT_Call struct {
	*mock.Call
}

// QuestionForGPT is a helper method to define mock.On call
//   - answer string
func (_e *AiLanguageModel_Expecter) QuestionForGPT(answer interface{}) *AiLanguageModel_QuestionForGPT_Call {
	return &AiLanguageModel_QuestionForGPT_Call{Call: _e.mock.On("QuestionForGPT", answer)}
}

func (_c *AiLanguageModel_QuestionForGPT_Call) Run(run func(answer string)) *AiLanguageModel_QuestionForGPT_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AiLanguageModel_QuestionForGPT_Call) Return(_a0 *openai.ChatCompletionResponse, _a1 error) *AiLanguageModel_QuestionForGPT_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AiLanguageModel_QuestionForGPT_Call) RunAndReturn(run func(string) (*openai.ChatCompletionResponse, error)) *AiLanguageModel_QuestionForGPT_Call {
	_c.Call.Return(run)
	return _c
}

// NewAiLanguageModel creates a new instance of AiLanguageModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAiLanguageModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *AiLanguageModel {
	mock := &AiLanguageModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}