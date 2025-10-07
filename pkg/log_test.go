package pkg

import (
	"reflect"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestCreateLogService(t *testing.T) {
	type args struct {
		configEnv DatabaseConfig
	}
	tests := []struct {
		name string
		args args
		want LogService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLogService(tt.args.configEnv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLogService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildZapConfig(t *testing.T) {
	tests := []struct {
		name string
		want zap.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildZapConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildZapConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLogLevel(t *testing.T) {
	tests := []struct {
		name string
		want zapcore.Level
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_determineOutput(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineOutput(); got != tt.want {
				t.Errorf("determineOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toString(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.args.v); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZapWrapper_LogInfo(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		args []interface{}
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogInfo(tt.args.args...)
		})
	}
}

func TestZapWrapper_LogInfoFormat(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		template string
		args     []interface{}
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogInfoFormat(tt.args.template, tt.args.args...)
		})
	}
}

func TestZapWrapper_LogErrorFormat(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		template string
		args     []interface{}
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogErrorFormat(tt.args.template, tt.args.args...)
		})
	}
}

func TestZapWrapper_LogError(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		msg    string
		err    error
		fields []zap.Field
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogError(tt.args.msg, tt.args.err, tt.args.fields...)
		})
	}
}

func TestZapWrapper_LogFatalFormat(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		template string
		args     []interface{}
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogFatalFormat(tt.args.template, tt.args.args...)
		})
	}
}

func TestZapWrapper_LogPanicFormat(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		template string
		args     []interface{}
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
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			z.LogPanicFormat(tt.args.template, tt.args.args...)
		})
	}
}

func TestZapWrapper_WithContext(t *testing.T) {
	type fields struct {
		core *zap.Logger
	}
	type args struct {
		attributes map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LogService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZapWrapper{
				core: tt.fields.core,
			}
			if got := z.WithContext(tt.args.attributes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZapWrapper.WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatArgs(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatArgs(tt.args.args); got != tt.want {
				t.Errorf("formatArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatTemplate(t *testing.T) {
	type args struct {
		template string
		args     []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTemplate(tt.args.template, tt.args.args); got != tt.want {
				t.Errorf("formatTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToZapFields(t *testing.T) {
	type args struct {
		attributes map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []zap.Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToZapFields(tt.args.attributes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToZapFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
