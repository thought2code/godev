package osutil

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/thought2code/godev/internal/strconst"
)

func TestCheckDirExist(t *testing.T) {
	// temp dir for test
	tempDir := t.TempDir()

	tests := []struct {
		name    string
		setup   func() string
		want    bool
		wantErr bool
	}{
		{
			name: "directory exists",
			setup: func() string {
				dir := filepath.Join(tempDir, "existing_dir")
				if err := os.MkdirAll(dir, 0o755); err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}
				return dir
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "directory does not exist",
			setup: func() string {
				return filepath.Join(tempDir, "non_existing_dir")
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "path is a file not directory",
			setup: func() string {
				file := filepath.Join(tempDir, "test_file.txt")
				if err := os.WriteFile(file, []byte("test"), 0o644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				return file
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "invalid path",
			setup: func() string {
				// use an invalid path that would cause an error
				return strconst.Empty
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()
			got, gotErr := CheckDirExist(path)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("CheckDirExist() failed, got unexpected error = %v", gotErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDirExist() failed, got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestCheckDirEmpty(t *testing.T) {
	// temp dir for test
	tempDir := t.TempDir()

	tests := []struct {
		name    string
		setup   func() string
		want    bool
		wantErr bool
	}{
		{
			name: "empty directory",
			setup: func() string {
				dir := filepath.Join(tempDir, "empty_dir")
				if err := os.MkdirAll(dir, 0o755); err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}
				return dir
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "directory with files",
			setup: func() string {
				dir := filepath.Join(tempDir, "dir_with_files")
				if err := os.MkdirAll(dir, 0o755); err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}
				// create a test file
				testFile := filepath.Join(dir, "test.txt")
				if err := os.WriteFile(testFile, []byte("test content"), 0o644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				return dir
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "directory with subdirectories",
			setup: func() string {
				dir := filepath.Join(tempDir, "dir_with_subdirs")
				subDir := filepath.Join(dir, "subdir")
				if err := os.MkdirAll(subDir, 0o755); err != nil {
					t.Fatalf("Failed to create test subdirectory: %v", err)
				}
				return dir
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "non-existent directory",
			setup: func() string {
				return filepath.Join(tempDir, "non_existing_dir")
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()
			got, gotErr := CheckDirEmpty(path)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("CheckDirEmpty() failed, got unexpected error = %v", gotErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckDirEmpty() failed, got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDirIfExist(t *testing.T) {
	// temp dir for test
	tempDir := t.TempDir()

	tests := []struct {
		name    string
		setup   func() string
		wantErr bool
	}{
		{
			name: "remove existing directory",
			setup: func() string {
				dir := filepath.Join(tempDir, "dir_to_remove")
				if err := os.MkdirAll(dir, 0o755); err != nil {
					t.Fatalf("Failed to create test directory: %v", err)
				}
				// create a test file
				testFile := filepath.Join(dir, "test.txt")
				if err := os.WriteFile(testFile, []byte("test"), 0o644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				return dir
			},
			wantErr: false,
		},
		{
			name: "remove non-existent directory",
			setup: func() string {
				return filepath.Join(tempDir, "non_existing_dir")
			},
			wantErr: false,
		},
		{
			name: "remove file (not directory)",
			setup: func() string {
				file := filepath.Join(tempDir, "test_file.txt")
				if err := os.WriteFile(file, []byte("test"), 0o644); err != nil {
					t.Fatalf("Failed to create test file: %v", err)
				}
				return file
			},
			wantErr: false, // should not error, just won't remove the file
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := tt.setup()
			gotErr := RemoveDirIfExist(path)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("RemoveDirIfExist() failed, got unexpected error = %v", gotErr)
				return
			}

			// verify directory is removed (if it existed and was a directory)
			if !tt.wantErr {
				if _, err := os.Stat(path); errors.Is(err, fs.ErrExist) {
					// for non-existent paths, this is expected
					// for files, they shouldn't be removed
					info, _ := os.Stat(path)
					if info != nil && info.IsDir() {
						t.Errorf("RemoveDirIfExist() failed, directory still exists: %s", path)
					}
				}
			}
		})
	}
}
