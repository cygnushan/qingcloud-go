// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package qcli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/cmd/qcli/api"
	_ "github.com/chai2010/qingcloud-go/pkg/internal/glogger"
	verpkg "github.com/chai2010/qingcloud-go/pkg/version"
)

func Main() {
	app := cli.NewApp()
	app.Name = "qcli"
	app.Usage = "QingCloud Command Line Interface"
	app.Version = verpkg.Version

	app.Authors = []cli.Author{
		{Name: "ChaiShushan", Email: "chaishushan@gmail.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "config file",
			Value:  "~/.qingcloud/config.yaml",
			EnvVar: "QCLI_CONFIG_FILE",
		},
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "ERR: command %q not found!\n", command)
	}

	app.Commands = []cli.Command{
		cmdInstance,
		cmdCluster,
		cmdVolume,
		cmdNic,
		cmdVxnet,
		cmdRouter,
		cmdEip,
		cmdSecurityGroup,
		cmdKeyPair,
		cmdImage,
		cmdLoadBalancer,
		cmdMonitor,
		cmdSnapshot,
		cmdJob,
	}

	app.Commands = pb.AllCommands

	app.Run(os.Args)
}