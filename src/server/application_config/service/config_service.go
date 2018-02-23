package service

import (
	"fmt"
	"os"
)

type ApplicationConfigService struct {
}

func (s *ApplicationConfigService) WriteToFile(name, path, content string) error {
	if err := checkPath(path); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fileName := path + "/" + name
	exits, errFile := checkFileIsExist(fileName)
	if errFile != nil {
		fmt.Println(errFile.Error())
		return errFile
	}
	if exits {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

func checkPath(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func checkFileIsExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
