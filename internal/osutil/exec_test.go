package osutil

import (
	"testing"
)

func TestRunCommand(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		args    []string
		wantErr bool
	}{
		{
			name:    "run command success",
			cmd:     "go",
			args:    []string{"version"},
			wantErr: false,
		},
		{
			name:    "run command failed",
			cmd:     "this is a invalid command",
			args:    []string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := RunCommand(tt.cmd, tt.args...)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("RunCommand() failed, got unexpected error: %v", gotErr)
				return
			}
		})
	}
}
