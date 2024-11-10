package file

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock ReadFileFunc for testing
func MockReadFile(content []byte, err error) ReadFileFunc {
	return func(path string) ([]byte, error) {
		return content, err
	}
}

// Mock YAMLUnmarshalFunc for testing
func MocYAMLUnmarshal(content []byte, err error) YAMLUnmarshalFunc {
	return func(data []byte, target any) error {
		return err
	}
}

func TestReadFromYAML(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		fileContent  []byte
		fileErr      error
		unmarshalErr error
		expectedErr  error
		target       interface{}
	}{
		{
			name:         "Successful read and unmarshal",
			path:         "valid.yaml",
			fileContent:  []byte("key: value"),
			fileErr:      nil,
			unmarshalErr: nil,
			expectedErr:  nil,
			target:       &map[string]string{},
		},
		{
			name:         "File not found",
			path:         "invalid.yaml",
			fileContent:  nil,
			fileErr:      errors.New("file not found"),
			unmarshalErr: nil,
			expectedErr:  errors.New("file not found"),
			target:       &map[string]string{},
		},
		{
			name:         "Invalid YAML content",
			path:         "invalid.yaml",
			fileContent:  []byte("key: :value"),
			fileErr:      nil,
			unmarshalErr: errors.New("yaml: line 1: did not find expected key"),
			expectedErr:  errors.New("yaml: line 1: did not find expected key"),
			target:       &map[string]string{},
		},
		{
			name:         "Empty file",
			path:         "empty.yaml",
			fileContent:  []byte(""),
			fileErr:      nil,
			unmarshalErr: nil,
			expectedErr:  nil,
			target:       &map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test with mock ReadFileFunc and mock YAMLUnmarshalFunc
			err := ReadFromYAML(tt.path, tt.target, MockReadFile(tt.fileContent, tt.fileErr), MocYAMLUnmarshal(tt.fileContent, tt.unmarshalErr))

			// Assertions
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
