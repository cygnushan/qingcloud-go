// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: vxnet.proto

package qcli_pb

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/urfave/cli"

	pb "github.com/chai2010/qingcloud-go/pkg/api"
	"github.com/chai2010/qingcloud-go/pkg/config"
	"github.com/chai2010/qingcloud-go/pkg/logger"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdVxnetService)
}

var CmdVxnetService = cli.Command{
	Name:    "vxnet",
	Aliases: []string{},
	Usage:   "manage VxnetService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeVxnets",
			Aliases: []string{},
			Usage:   "DescribeVxnets",
			Flags:   _flag_VxnetService_DescribeVxnets,
			Action:  _cmd_VxnetService_DescribeVxnets,
		},
		{
			Name:    "CreateVxnets",
			Aliases: []string{},
			Usage:   "CreateVxnets",
			Flags:   _flag_VxnetService_CreateVxnets,
			Action:  _cmd_VxnetService_CreateVxnets,
		},
		{
			Name:    "DeleteVxnets",
			Aliases: []string{},
			Usage:   "DeleteVxnets",
			Flags:   _flag_VxnetService_DeleteVxnets,
			Action:  _cmd_VxnetService_DeleteVxnets,
		},
		{
			Name:    "JoinVxnet",
			Aliases: []string{},
			Usage:   "JoinVxnet",
			Flags:   _flag_VxnetService_JoinVxnet,
			Action:  _cmd_VxnetService_JoinVxnet,
		},
		{
			Name:    "LeaveVxnet",
			Aliases: []string{},
			Usage:   "LeaveVxnet",
			Flags:   _flag_VxnetService_LeaveVxnet,
			Action:  _cmd_VxnetService_LeaveVxnet,
		},
		{
			Name:    "ModifyVxnetAttributes",
			Aliases: []string{},
			Usage:   "ModifyVxnetAttributes",
			Flags:   _flag_VxnetService_ModifyVxnetAttributes,
			Action:  _cmd_VxnetService_ModifyVxnetAttributes,
		},
		{
			Name:    "DescribeVxnetInstances",
			Aliases: []string{},
			Usage:   "DescribeVxnetInstances",
			Flags:   _flag_VxnetService_DescribeVxnetInstances,
			Action:  _cmd_VxnetService_DescribeVxnetInstances,
		},
	},
}

var _flag_VxnetService_DescribeVxnets = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_DescribeVxnets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.DescribeVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeVxnets(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_CreateVxnets = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_CreateVxnets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.CreateVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateVxnets(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_DeleteVxnets = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_DeleteVxnets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.DeleteVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteVxnets(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_JoinVxnet = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_JoinVxnet(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.JoinVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.JoinVxnet(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_LeaveVxnet = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_LeaveVxnet(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.LeaveVxnetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.LeaveVxnet(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_ModifyVxnetAttributes = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_ModifyVxnetAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.ModifyVxnetAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyVxnetAttributes(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}

var _flag_VxnetService_DescribeVxnetInstances = []cli.Flag{ /* fields */ }

func _cmd_VxnetService_DescribeVxnetInstances(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewVxnetService(conf, zone)

	in := new(pb.DescribeVxnetInstancesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeVxnetInstances(in)
	if err != nil {
		logger.Fatal(err)
	}

	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
		Indent:       "  ",
	}
	s, err := jsonMarshaler.MarshalToString(out)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println(s)
	return nil
}
