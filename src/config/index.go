package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var configFilePath = ".env"

var configMap = make(map[string]string)

func Load() error {

	configFilePath, error := filepath.Abs(configFilePath)

	if error != nil {
		return error
	}

	configFile, error := os.Open(configFilePath)
	if error != nil {
		return error
	}

	defer configFile.Close()

	fd, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	fileString := string(fd)
	fileContent := strings.Split(fileString, "\n")

	for _, fileLine := range fileContent {

		// 过滤掉空白行
		if 0 == len(fileLine) {
			continue
		}

		// 过滤掉注释行
		if strings.HasPrefix(fileLine, "#") {
			continue
		}

		configItem := strings.Split(fileLine, "=")
		// 过滤掉不符合配置的配置项
		if 2 != len(configItem) {
			continue
		}

		// 过滤空格
		tmpKey := strings.ToLower(strings.TrimSpace(configItem[0]))
		tmpVal := strings.TrimSpace(configItem[1])

		configMap[tmpKey] = tmpVal
	}

	return nil
}

func Get(name string) string {
	if val, ok := configMap[name]; ok {
		return val
	}

	return ""
}
