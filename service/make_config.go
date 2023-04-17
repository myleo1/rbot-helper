package service

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/mizuki1412/go-core-kit/class/exception"
	"github.com/mizuki1412/go-core-kit/library/filekit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/spf13/cast"
)

type Config struct {
	OracleName string
	OracleVal  string
	AzureName  string
	AzureVal   string
	Other      string
}

func NewConfig(oracleName, azureName string) *Config {
	configOracle := configkit.Get("rbot.oracle", "")
	configAzure := configkit.Get("rbot.azure", "")
	if configOracle == "" && configAzure == "" {
		logkit.Fatal("read rbot config error")
	}
	configListOracle := cast.ToSlice(configOracle)
	configListAzure := cast.ToSlice(configAzure)
	if len(configListOracle) == 0 && len(configListAzure) == 0 {
		logkit.Fatal("rbot oracle and azure config length 0")
	}
	other := configkit.GetString("rbot.other", "")
	if other == "" {
		logkit.Fatal("rbot other information null")
	}
	c := &Config{}
	for _, v := range configListOracle {
		m := cast.ToStringMapString(v)
		if m["name"] == oracleName {
			c.OracleVal = m["value"]
			break
		}
	}
	for _, v := range configListAzure {
		m := cast.ToStringMapString(v)
		if m["name"] == azureName {
			c.AzureVal = m["value"]
			break
		}
	}
	c.OracleName = oracleName
	c.AzureName = azureName
	c.Other = other
	return c
}

func (th *Config) MakeConfig() {
	//oci
	c := "oci=begin\n\n[" + th.OracleName + "]\n" + th.OracleVal + "\n\noci=end\n\n"
	//other
	c = c + th.Other + "\n\n"
	//azure
	c = c + "azure=begin\n\n[" + th.AzureName + "]\n" + th.AzureVal + "\n\nazure=end\n\n"
	err := filekit.WriteFile("/root/client_config", []byte(c))
	if err != nil {
		panic(exception.New("write client_config file error"))
	}
}

func (th *Config) RestartRBot() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	// 获取需要重启的容器 ID
	containerName := configkit.GetString("rbot.container", "")
	if containerName == "" {
		panic(exception.New("rbot container name get error"))
	}
	containerListOptions := types.ContainerListOptions{All: true, Filters: filters.NewArgs(filters.Arg("name", containerName))}
	containers, err := cli.ContainerList(context.Background(), containerListOptions)
	if err != nil {
		panic(exception.New(err.Error()))
	}
	if len(containers) == 0 {
		panic(exception.New(fmt.Sprintf("no such container: %s", containerName)))
	}
	containerID := containers[0].ID
	// 重启容器
	if err = cli.ContainerRestart(context.Background(), containerID, container.StopOptions{}); err != nil {
		panic(exception.New(err.Error()))
	}
}
