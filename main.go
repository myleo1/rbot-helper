package main

import (
	"fmt"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"robot-helper/cmd"
)

var (
	version   string
	date      string
	goVersion string
)

func main() {
	fmt.Println(` ____  _           _        _   _      _
|  _ \| |__   ___ | |_     | | | | ___| |_ __   ___ _ __
| |_) | '_ \ / _ \| __|____| |_| |/ _ \ | '_ \ / _ \ '__|
|  _ <| |_) | (_) | ||_____|  _  |  __/ | |_) |  __/ |
|_| \_\_.__/ \___/ \__|    |_| |_|\___|_| .__/ \___|_|
                                        |_|`)
	info := fmt.Sprintf("***Version %s***\n***BuildDate %s***\n***%s***\n", version, date, goVersion)
	fmt.Print(info)
	initkit.LoadConfig()
	cmd.Execute()
}
