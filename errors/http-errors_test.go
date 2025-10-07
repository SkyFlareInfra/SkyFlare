package errors

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDispatchInternalError(t *testing.T) {
	type args struct {
		c   *gin.Context
		err interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DispatchInternalError(tt.args.c, tt.args.err)
		})
	}
}

func TestDispatchBadRequest(t *testing.T) {
	type args struct {
		c         *gin.Context
		msg       interface{}
		err       error
		errorType []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DispatchBadRequest(tt.args.c, tt.args.msg, tt.args.err, tt.args.errorType...)
		})
	}
}

func TestDispatchNotFound(t *testing.T) {
	type args struct {
		c   *gin.Context
		msg string
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DispatchNotFound(tt.args.c, tt.args.msg, tt.args.err)
		})
	}
}

func TestDispatchForbidden(t *testing.T) {
	type args struct {
		c   *gin.Context
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DispatchForbidden(tt.args.c, tt.args.msg)
		})
	}
}

func TestDispatchMethodNotAllowed(t *testing.T) {
	type args struct {
		c   *gin.Context
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DispatchMethodNotAllowed(tt.args.c, tt.args.msg)
		})
	}
}

func TestNewErrorBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *ErrorBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_New(t *testing.T) {
	type args struct {
		statusCode int
		message    string
		success    bool
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.New(tt.args.statusCode, tt.args.message, tt.args.success); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_BadRequest(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.BadRequest(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.BadRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_Unauthorized(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.Unauthorized(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.Unauthorized() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_NotFound(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.NotFound(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.NotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_InternalError(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.InternalError(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.InternalError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_Forbidden(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.Forbidden(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.Forbidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_TooManyRequests(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.TooManyRequests(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.TooManyRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorBuilder_Accepted(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		b    *ErrorBuilder
		args args
		want *APIError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &ErrorBuilder{}
			if got := b.Accepted(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorBuilder.Accepted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIError_Send(t *testing.T) {
	type fields struct {
		Message    string
		Success    bool
		StatusCode int
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &APIError{
				Message:    tt.fields.Message,
				Success:    tt.fields.Success,
				StatusCode: tt.fields.StatusCode,
			}
			e.Send(tt.args.c)
		})
	}
}
