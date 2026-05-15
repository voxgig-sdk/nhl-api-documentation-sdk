package core

type NhlApiDocumentationError struct {
	IsNhlApiDocumentationError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewNhlApiDocumentationError(code string, msg string, ctx *Context) *NhlApiDocumentationError {
	return &NhlApiDocumentationError{
		IsNhlApiDocumentationError: true,
		Sdk:              "NhlApiDocumentation",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *NhlApiDocumentationError) Error() string {
	return e.Msg
}
