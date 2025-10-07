package acme

import (
	"reflect"
	"testing"

	"github.com/go-acme/lego/v4/challenge"
)

func TestHTTP01Provider_Present(t *testing.T) {
	type args struct {
		domain  string
		token   string
		keyAuth string
	}
	tests := []struct {
		name    string
		p       *HTTP01Provider
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HTTP01Provider{}
			if err := p.Present(tt.args.domain, tt.args.token, tt.args.keyAuth); (err != nil) != tt.wantErr {
				t.Errorf("HTTP01Provider.Present() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHTTP01Provider_CleanUp(t *testing.T) {
	type args struct {
		domain  string
		token   string
		keyAuth string
	}
	tests := []struct {
		name    string
		p       *HTTP01Provider
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HTTP01Provider{}
			if err := p.CleanUp(tt.args.domain, tt.args.token, tt.args.keyAuth); (err != nil) != tt.wantErr {
				t.Errorf("HTTP01Provider.CleanUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProvider(t *testing.T) {
	tests := []struct {
		name string
		want challenge.Provider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Provider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChallengeResponse(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetChallengeResponse(tt.args.token)
			if got != tt.want {
				t.Errorf("GetChallengeResponse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetChallengeResponse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
