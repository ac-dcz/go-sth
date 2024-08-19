package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type MysqlConf struct {
	Host     string `yaml:"Host"`
	DataBase string `yaml:"DataBase"`
	User     string `yaml:"User"`
	PassWord string `yaml:"PassWord"`
}

// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (m *MysqlConf) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", m.User, m.PassWord, m.Host, m.DataBase)
}

func Load(filePath string) (*MysqlConf, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Printf("Open file error: %v \n", err)
		return nil, err
	}
	defer file.Close()
	c := &MysqlConf{}
	if err := yaml.NewDecoder(file).Decode(c); err != nil {
		fmt.Printf("Decode config file error: %v\n", err)
		return nil, err
	}
	return c, nil
}
