// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: router.proto

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
	AllCommands = append(AllCommands, CmdRouterService)
}

var CmdRouterService = cli.Command{
	Name:    "router",
	Aliases: []string{},
	Usage:   "manage RouterService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeRouters",
			Aliases: []string{},
			Usage:   "DescribeRouters",
			Flags:   _flag_RouterService_DescribeRouters,
			Action:  _cmd_RouterService_DescribeRouters,
		},
		{
			Name:    "CreateRouters",
			Aliases: []string{},
			Usage:   "CreateRouters",
			Flags:   _flag_RouterService_CreateRouters,
			Action:  _cmd_RouterService_CreateRouters,
		},
		{
			Name:    "DeleteRouters",
			Aliases: []string{},
			Usage:   "DeleteRouters",
			Flags:   _flag_RouterService_DeleteRouters,
			Action:  _cmd_RouterService_DeleteRouters,
		},
		{
			Name:    "UpdateRouters",
			Aliases: []string{},
			Usage:   "UpdateRouters",
			Flags:   _flag_RouterService_UpdateRouters,
			Action:  _cmd_RouterService_UpdateRouters,
		},
		{
			Name:    "PowerOffRouters",
			Aliases: []string{},
			Usage:   "PowerOffRouters",
			Flags:   _flag_RouterService_PowerOffRouters,
			Action:  _cmd_RouterService_PowerOffRouters,
		},
		{
			Name:    "PowerOnRouters",
			Aliases: []string{},
			Usage:   "PowerOnRouters",
			Flags:   _flag_RouterService_PowerOnRouters,
			Action:  _cmd_RouterService_PowerOnRouters,
		},
		{
			Name:    "JoinRouter",
			Aliases: []string{},
			Usage:   "JoinRouter",
			Flags:   _flag_RouterService_JoinRouter,
			Action:  _cmd_RouterService_JoinRouter,
		},
		{
			Name:    "LeaveRouter",
			Aliases: []string{},
			Usage:   "LeaveRouter",
			Flags:   _flag_RouterService_LeaveRouter,
			Action:  _cmd_RouterService_LeaveRouter,
		},
		{
			Name:    "ModifyRouterAttributes",
			Aliases: []string{},
			Usage:   "ModifyRouterAttributes",
			Flags:   _flag_RouterService_ModifyRouterAttributes,
			Action:  _cmd_RouterService_ModifyRouterAttributes,
		},
		{
			Name:    "DescribeRouterStatics",
			Aliases: []string{},
			Usage:   "DescribeRouterStatics",
			Flags:   _flag_RouterService_DescribeRouterStatics,
			Action:  _cmd_RouterService_DescribeRouterStatics,
		},
		{
			Name:    "AddRouterStatics",
			Aliases: []string{},
			Usage:   "AddRouterStatics",
			Flags:   _flag_RouterService_AddRouterStatics,
			Action:  _cmd_RouterService_AddRouterStatics,
		},
		{
			Name:    "ModifyRouterStaticAttributes",
			Aliases: []string{},
			Usage:   "ModifyRouterStaticAttributes",
			Flags:   _flag_RouterService_ModifyRouterStaticAttributes,
			Action:  _cmd_RouterService_ModifyRouterStaticAttributes,
		},
		{
			Name:    "DeleteRouterStatics",
			Aliases: []string{},
			Usage:   "DeleteRouterStatics",
			Flags:   _flag_RouterService_DeleteRouterStatics,
			Action:  _cmd_RouterService_DeleteRouterStatics,
		},
		{
			Name:    "CopyRouterStatics",
			Aliases: []string{},
			Usage:   "CopyRouterStatics",
			Flags:   _flag_RouterService_CopyRouterStatics,
			Action:  _cmd_RouterService_CopyRouterStatics,
		},
		{
			Name:    "DescribeRouterVxnets",
			Aliases: []string{},
			Usage:   "DescribeRouterVxnets",
			Flags:   _flag_RouterService_DescribeRouterVxnets,
			Action:  _cmd_RouterService_DescribeRouterVxnets,
		},
		{
			Name:    "AddRouterStaticEntries",
			Aliases: []string{},
			Usage:   "AddRouterStaticEntries",
			Flags:   _flag_RouterService_AddRouterStaticEntries,
			Action:  _cmd_RouterService_AddRouterStaticEntries,
		},
		{
			Name:    "DeleteRouterStaticEntries",
			Aliases: []string{},
			Usage:   "DeleteRouterStaticEntries",
			Flags:   _flag_RouterService_DeleteRouterStaticEntries,
			Action:  _cmd_RouterService_DeleteRouterStaticEntries,
		},
		{
			Name:    "ModifyRouterStaticEntryAttributes",
			Aliases: []string{},
			Usage:   "ModifyRouterStaticEntryAttributes",
			Flags:   _flag_RouterService_ModifyRouterStaticEntryAttributes,
			Action:  _cmd_RouterService_ModifyRouterStaticEntryAttributes,
		},
		{
			Name:    "DescribeRouterStaticEntries",
			Aliases: []string{},
			Usage:   "DescribeRouterStaticEntries",
			Flags:   _flag_RouterService_DescribeRouterStaticEntries,
			Action:  _cmd_RouterService_DescribeRouterStaticEntries,
		},
	},
}

var _flag_RouterService_DescribeRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DescribeRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DescribeRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeRouters(in)
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

var _flag_RouterService_CreateRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_CreateRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.CreateRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateRouters(in)
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

var _flag_RouterService_DeleteRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DeleteRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DeleteRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteRouters(in)
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

var _flag_RouterService_UpdateRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_UpdateRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.UpdateRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.UpdateRouters(in)
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

var _flag_RouterService_PowerOffRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_PowerOffRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.PowerOffRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.PowerOffRouters(in)
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

var _flag_RouterService_PowerOnRouters = []cli.Flag{ /* fields */ }

func _cmd_RouterService_PowerOnRouters(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.PowerOnRoutersInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.PowerOnRouters(in)
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

var _flag_RouterService_JoinRouter = []cli.Flag{ /* fields */ }

func _cmd_RouterService_JoinRouter(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.JoinRouterInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.JoinRouter(in)
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

var _flag_RouterService_LeaveRouter = []cli.Flag{ /* fields */ }

func _cmd_RouterService_LeaveRouter(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.LeaveRouterInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.LeaveRouter(in)
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

var _flag_RouterService_ModifyRouterAttributes = []cli.Flag{ /* fields */ }

func _cmd_RouterService_ModifyRouterAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.ModifyRouterAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyRouterAttributes(in)
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

var _flag_RouterService_DescribeRouterStatics = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DescribeRouterStatics(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DescribeRouterStaticsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeRouterStatics(in)
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

var _flag_RouterService_AddRouterStatics = []cli.Flag{ /* fields */ }

func _cmd_RouterService_AddRouterStatics(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.AddRouterStaticsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddRouterStatics(in)
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

var _flag_RouterService_ModifyRouterStaticAttributes = []cli.Flag{ /* fields */ }

func _cmd_RouterService_ModifyRouterStaticAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.ModifyRouterStaticAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyRouterStaticAttributes(in)
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

var _flag_RouterService_DeleteRouterStatics = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DeleteRouterStatics(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DeleteRouterStaticsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteRouterStatics(in)
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

var _flag_RouterService_CopyRouterStatics = []cli.Flag{ /* fields */ }

func _cmd_RouterService_CopyRouterStatics(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.CopyRouterStaticsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CopyRouterStatics(in)
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

var _flag_RouterService_DescribeRouterVxnets = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DescribeRouterVxnets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DescribeRouterVxnetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeRouterVxnets(in)
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

var _flag_RouterService_AddRouterStaticEntries = []cli.Flag{ /* fields */ }

func _cmd_RouterService_AddRouterStaticEntries(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.AddRouterStaticEntriesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddRouterStaticEntries(in)
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

var _flag_RouterService_DeleteRouterStaticEntries = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DeleteRouterStaticEntries(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DeleteRouterStaticEntriesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteRouterStaticEntries(in)
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

var _flag_RouterService_ModifyRouterStaticEntryAttributes = []cli.Flag{ /* fields */ }

func _cmd_RouterService_ModifyRouterStaticEntryAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.ModifyRouterStaticEntryAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifyRouterStaticEntryAttributes(in)
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

var _flag_RouterService_DescribeRouterStaticEntries = []cli.Flag{ /* fields */ }

func _cmd_RouterService_DescribeRouterStaticEntries(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewRouterService(conf, zone)

	in := new(pb.DescribeRouterStaticEntriesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeRouterStaticEntries(in)
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
