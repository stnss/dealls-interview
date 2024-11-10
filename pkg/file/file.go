package file

type ReadFileFunc func(string) ([]byte, error)
type YAMLUnmarshalFunc func([]byte, any) error

func ReadFromYAML(path string, target any, readFile ReadFileFunc, yamlUnmarshal YAMLUnmarshalFunc) error {
	yf, err := readFile(path)
	if err != nil {
		return err
	}
	return yamlUnmarshal(yf, target)
}
