package localization

import (
	"gopkg.in/yaml.v3"
	"localizer/pkg/file"
	"localizer/pkg/yamlx"
	"path"
)

func ReadFileMap(filename string) (yamlx.Map, error) {
	data, err := file.Read(filename)
	if err != nil {
		return nil, err
	}
	var d any
	err = yaml.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	ym := yamlx.Map{}
	ym.Set(d)
	return ym, nil
}

func writeFileMap(filename string, data yamlx.Map) error {
	obj, err := data.Marshal()
	if err != nil {
		return err
	}
	d, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	return file.Write(filename, d)
}

func yamlPath(p, filename string) string {
	return path.Join(p, filename+".yaml")
}
