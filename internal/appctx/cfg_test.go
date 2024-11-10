package appctx

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock implementation of the file reader and YAML unmarshaller.
type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

type MockYAMLUnmarshaler struct {
	mock.Mock
}

func (m *MockYAMLUnmarshaler) Unmarshal(data []byte, v any) error {
	args := m.Called(data, v)
	return args.Error(0)
}

func TestReadConfig(t *testing.T) {
	testCases := []struct {
		name           string
		fileContent    string
		fileErr        error
		unmarshalErr   error
		expectedCfg    *Config
		expectedErrStr string
	}{
		{
			name: "Valid configuration",
			fileContent: `
app:
  name: "TestApp"
  port: 8080
  debug: true
  timezone: "UTC"
  env: "development"
  read_timeout: 30s
  write_timeout: 30s
log:
  level: "info"
`,
			fileErr: nil,
			expectedCfg: &Config{
				App: App{
					Name:         "TestApp",
					Port:         8080,
					Debug:        true,
					Timezone:     "UTC",
					Env:          "development",
					ReadTimeout:  30 * time.Second,
					WriteTimeout: 30 * time.Second,
				},
				Logger: Logger{
					Level: "info",
				},
			},
		},
		{
			name:           "Invalid YAML format",
			fileContent:    `invalid_yaml`,
			fileErr:        nil,
			unmarshalErr:   errors.New("yaml: unmarshal error"),
			expectedCfg:    nil,
			expectedErrStr: "file config parse error",
		},
		{
			name:           "File not found",
			fileContent:    "",
			fileErr:        errors.New("file not found"),
			expectedCfg:    nil,
			expectedErrStr: "file config parse error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mocking file reader
			mockFileReader := new(MockFileReader)
			mockFileReader.On("ReadFile", mock.Anything).Return([]byte(tc.fileContent), tc.fileErr)

			// Mocking YAML unmarshaller
			mockYAMLUnmarshaler := new(MockYAMLUnmarshaler)
			if tc.fileErr == nil {
				mockYAMLUnmarshaler.On("Unmarshal", mock.Anything, mock.Anything).Return(tc.unmarshalErr).Run(func(args mock.Arguments) {
					arg := args.Get(1).(**Config)
					if tc.expectedCfg != nil {
						*arg = tc.expectedCfg // Assign expected configuration
					}
				})
			}

			// Call the readConfig function
			config, err := readConfig("app.yaml", mockFileReader.ReadFile, mockYAMLUnmarshaler.Unmarshal, "config/")
			fmt.Println(config)
			// Validate results
			if tc.expectedErrStr != "" {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tc.expectedErrStr)
				assert.Nil(t, config)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, config)
				assert.Equal(t, tc.expectedCfg, config)
			}

			mockFileReader.AssertExpectations(t)
			if tc.fileErr == nil {
				mockYAMLUnmarshaler.AssertExpectations(t)
			}
		})
	}
}
