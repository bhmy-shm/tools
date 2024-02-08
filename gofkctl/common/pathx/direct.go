package pathx

import (
	"fmt"
	"os"
)

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// CreateDirectoryIfNotExist 生成目录
func CreateDirectoryIfNotExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateFileIfNotExist 生成文件
func CreateFileIfNotExist(path string) (*os.File, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Create(path)
	}
	return nil, fmt.Errorf("%s already exist", path)
}

// NewFileBuild 基于template模板生成文件
func NewFileBuild(tp *TemplateData, filePath string) error {

	tmpl, err := tp.Parse()
	if err != nil {
		return err
	}

	file, err := CreateFileIfNotExist(filePath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filePath, err)
		return err
	}

	err = tmpl.Execute(file, tp)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
		return err
	}

	return file.Close()
}

// NewServiceWithContent 基于content []byte 内容生成文件
//func NewServiceWithContent(filePath string, content []byte) error {
//	tmpl, err := template.New("fileTemplate").Parse(string(content))
//	if err != nil {
//		fmt.Printf("Error parsing template: %v\n", err)
//		return err
//	}
//
//	file, err := CreateFileIfNotExist(filePath)
//	if err != nil {
//		fmt.Printf("Error creating file %s: %v\n", filePath, err)
//		return err
//	}
//
//	err = tmpl.Execute(file, content)
//	if err != nil {
//		fmt.Printf("Error executing template: %v\n", err)
//		return err
//	}
//
//	return file.Close()
//}
