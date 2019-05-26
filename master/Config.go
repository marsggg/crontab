package master

import (
	"encoding/json"
	"io/ioutil"
)

//程序配置
type Config struct {
	ApiPort               int      `json:"apiPort"`
	ApiReadTimeout        int      `json:"apiReadTimeout"`
	ApiWriteTimeout       int      `json:"apiWriteTimeout"`
	EtcdEndpoints         []string `json:"etcdEndpoints"`
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`
	Webroot               string   `json:"webroot"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"`
}

var (
	G_config *Config
)

//加载配置文件
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	//把文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return err
	}
	//做Json反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return err
	}

	//赋值单例
	G_config = &conf

	//fmt.Println(conf)

	return nil
}
