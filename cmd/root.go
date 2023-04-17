package cmd

import (
	"github.com/mizuki1412/go-core-kit/init/configkey"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/mizuki1412/go-core-kit/service/restkit"
	"github.com/spf13/cobra"
	"robot-helper/constant"
	"robot-helper/controller"
)

func init() {
	defFlags(rootCmd)
}

var rootCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		restkit.AddActions(controller.Init)
		_ = restkit.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logkit.Fatal(err.Error())
	}
}

func defFlags(cmd *cobra.Command) {
	cmd.Flags().String(configkey.RestServerPort, "9000", "Api listen port")
	cmd.Flags().String(configkey.RestServerBase, "rbot-helper", "Api base path")
	cmd.Flags().String(constant.ConfigKeyToken, "", "Api token")
}
