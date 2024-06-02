package configs

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name:    "Test load config done",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
