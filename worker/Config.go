package worker

import (
	"encoding/json"
	"io/ioutil"
)

//程序配置
type Config struct {
	EtcdEndpoints         []string `json:"etcdEndpoints"`
	EtcdDialTimeout       int      `json:"etcdDialTimeout"`
	MongodbUri            string   `json:"mongodbUri"`
	MongodbConnectTimeout int      `json:"mongodbConnectTimeout"`
	JobLogBatchSize       int      `json:"jobLogBatchSize"`
	JobLogCommitTimeout   int      `json:"jobLogCommitTimeout"`
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
