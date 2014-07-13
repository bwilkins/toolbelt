package config

import (
  "os"
  "io/ioutil"
  "log"
  yaml "gopkg.in/yaml.v1"
)

type ConfigBlock struct {
  OpsWorks OpsWorksBlock `yaml:"ops_works"`
  EngineYard interface{} `yaml:"engine_yard,omitempty"`
}

type OpsWorksBlock struct {
  AccessId string `yaml:"aws_access_key_id"`
  SecretKey string `yaml:"aws_secret_access_key"`
  Ssh SshBlock `yaml:"ssh"`
}

type SshBlock struct {
  DefaultUser string `yaml:"default_username"`
  DefaultKeys []string `yaml:"default_keys"`
}

var Config ConfigBlock

func SetConfig(configFile string) {
  file, err := os.Open(configFile)
  defer file.Close()

  if err != nil {
    log.Fatal(err.Error())
  }

  contents, err := ioutil.ReadAll(file)
  if err != nil {
    log.Fatal(err.Error())
  }

  yaml.Unmarshal(contents, &Config)

  //Env overrides
  access_id := os.Getenv("AWS_ACCESS_KEY_ID")
  if access_id != "" {
    Config.OpsWorks.AccessId = access_id
  }

  secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
  if secret_key != "" {
    Config.OpsWorks.SecretKey = secret_key
  }

  return
}
