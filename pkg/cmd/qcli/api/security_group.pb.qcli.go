// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: security_group.proto

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
	AllCommands = append(AllCommands, CmdSecurityGroupService)
}

var CmdSecurityGroupService = cli.Command{
	Name:    "securitygroup",
	Aliases: []string{},
	Usage:   "manage SecurityGroupService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeSecurityGroups",
			Aliases: []string{},
			Usage:   "DescribeSecurityGroups",
			Flags:   _flag_SecurityGroupService_DescribeSecurityGroups,
			Action:  _cmd_SecurityGroupService_DescribeSecurityGroups,
		},
		{
			Name:    "CreateSecurityGroup",
			Aliases: []string{},
			Usage:   "CreateSecurityGroup",
			Flags:   _flag_SecurityGroupService_CreateSecurityGroup,
			Action:  _cmd_SecurityGroupService_CreateSecurityGroup,
		},
		{
			Name:    "DeleteSecurityGroups",
			Aliases: []string{},
			Usage:   "DeleteSecurityGroups",
			Flags:   _flag_SecurityGroupService_DeleteSecurityGroups,
			Action:  _cmd_SecurityGroupService_DeleteSecurityGroups,
		},
		{
			Name:    "ApplySecurityGroup",
			Aliases: []string{},
			Usage:   "ApplySecurityGroup",
			Flags:   _flag_SecurityGroupService_ApplySecurityGroup,
			Action:  _cmd_SecurityGroupService_ApplySecurityGroup,
		},
		{
			Name:    "ModifySecurityGroupAttributes",
			Aliases: []string{},
			Usage:   "ModifySecurityGroupAttributes",
			Flags:   _flag_SecurityGroupService_ModifySecurityGroupAttributes,
			Action:  _cmd_SecurityGroupService_ModifySecurityGroupAttributes,
		},
		{
			Name:    "DescribeSecurityGroupRules",
			Aliases: []string{},
			Usage:   "DescribeSecurityGroupRules",
			Flags:   _flag_SecurityGroupService_DescribeSecurityGroupRules,
			Action:  _cmd_SecurityGroupService_DescribeSecurityGroupRules,
		},
		{
			Name:    "AddSecurityGroupRules",
			Aliases: []string{},
			Usage:   "AddSecurityGroupRules",
			Flags:   _flag_SecurityGroupService_AddSecurityGroupRules,
			Action:  _cmd_SecurityGroupService_AddSecurityGroupRules,
		},
		{
			Name:    "DeleteSecurityGroupRules",
			Aliases: []string{},
			Usage:   "DeleteSecurityGroupRules",
			Flags:   _flag_SecurityGroupService_DeleteSecurityGroupRules,
			Action:  _cmd_SecurityGroupService_DeleteSecurityGroupRules,
		},
		{
			Name:    "ModifySecurityGroupRuleAttributes",
			Aliases: []string{},
			Usage:   "ModifySecurityGroupRuleAttributes",
			Flags:   _flag_SecurityGroupService_ModifySecurityGroupRuleAttributes,
			Action:  _cmd_SecurityGroupService_ModifySecurityGroupRuleAttributes,
		},
		{
			Name:    "CreateSecurityGroupSnapshot",
			Aliases: []string{},
			Usage:   "CreateSecurityGroupSnapshot",
			Flags:   _flag_SecurityGroupService_CreateSecurityGroupSnapshot,
			Action:  _cmd_SecurityGroupService_CreateSecurityGroupSnapshot,
		},
		{
			Name:    "DescribeSecurityGroupSnapshots",
			Aliases: []string{},
			Usage:   "DescribeSecurityGroupSnapshots",
			Flags:   _flag_SecurityGroupService_DescribeSecurityGroupSnapshots,
			Action:  _cmd_SecurityGroupService_DescribeSecurityGroupSnapshots,
		},
		{
			Name:    "DeleteSecurityGroupSnapshots",
			Aliases: []string{},
			Usage:   "DeleteSecurityGroupSnapshots",
			Flags:   _flag_SecurityGroupService_DeleteSecurityGroupSnapshots,
			Action:  _cmd_SecurityGroupService_DeleteSecurityGroupSnapshots,
		},
		{
			Name:    "RollbackSecurityGroup",
			Aliases: []string{},
			Usage:   "RollbackSecurityGroup",
			Flags:   _flag_SecurityGroupService_RollbackSecurityGroup,
			Action:  _cmd_SecurityGroupService_RollbackSecurityGroup,
		},
		{
			Name:    "DescribeSecurityGroupIPSets",
			Aliases: []string{},
			Usage:   "DescribeSecurityGroupIPSets",
			Flags:   _flag_SecurityGroupService_DescribeSecurityGroupIPSets,
			Action:  _cmd_SecurityGroupService_DescribeSecurityGroupIPSets,
		},
		{
			Name:    "CreateSecurityGroupIPSet",
			Aliases: []string{},
			Usage:   "CreateSecurityGroupIPSet",
			Flags:   _flag_SecurityGroupService_CreateSecurityGroupIPSet,
			Action:  _cmd_SecurityGroupService_CreateSecurityGroupIPSet,
		},
		{
			Name:    "DeleteSecurityGroupIPSets",
			Aliases: []string{},
			Usage:   "DeleteSecurityGroupIPSets",
			Flags:   _flag_SecurityGroupService_DeleteSecurityGroupIPSets,
			Action:  _cmd_SecurityGroupService_DeleteSecurityGroupIPSets,
		},
		{
			Name:    "ModifySecurityGroupIPSetAttributes",
			Aliases: []string{},
			Usage:   "ModifySecurityGroupIPSetAttributes",
			Flags:   _flag_SecurityGroupService_ModifySecurityGroupIPSetAttributes,
			Action:  _cmd_SecurityGroupService_ModifySecurityGroupIPSetAttributes,
		},
		{
			Name:    "CopySecurityGroupIPSets",
			Aliases: []string{},
			Usage:   "CopySecurityGroupIPSets",
			Flags:   _flag_SecurityGroupService_CopySecurityGroupIPSets,
			Action:  _cmd_SecurityGroupService_CopySecurityGroupIPSets,
		},
	},
}

var _flag_SecurityGroupService_DescribeSecurityGroups = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DescribeSecurityGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DescribeSecurityGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSecurityGroups(in)
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

var _flag_SecurityGroupService_CreateSecurityGroup = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_CreateSecurityGroup(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.CreateSecurityGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateSecurityGroup(in)
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

var _flag_SecurityGroupService_DeleteSecurityGroups = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DeleteSecurityGroups(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DeleteSecurityGroupsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteSecurityGroups(in)
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

var _flag_SecurityGroupService_ApplySecurityGroup = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_ApplySecurityGroup(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.ApplySecurityGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ApplySecurityGroup(in)
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

var _flag_SecurityGroupService_ModifySecurityGroupAttributes = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_ModifySecurityGroupAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.ModifySecurityGroupAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifySecurityGroupAttributes(in)
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

var _flag_SecurityGroupService_DescribeSecurityGroupRules = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DescribeSecurityGroupRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DescribeSecurityGroupRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSecurityGroupRules(in)
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

var _flag_SecurityGroupService_AddSecurityGroupRules = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_AddSecurityGroupRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.AddSecurityGroupRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.AddSecurityGroupRules(in)
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

var _flag_SecurityGroupService_DeleteSecurityGroupRules = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DeleteSecurityGroupRules(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DeleteSecurityGroupRulesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteSecurityGroupRules(in)
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

var _flag_SecurityGroupService_ModifySecurityGroupRuleAttributes = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_ModifySecurityGroupRuleAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.ModifySecurityGroupRuleAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifySecurityGroupRuleAttributes(in)
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

var _flag_SecurityGroupService_CreateSecurityGroupSnapshot = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_CreateSecurityGroupSnapshot(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.CreateSecurityGroupSnapshotInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateSecurityGroupSnapshot(in)
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

var _flag_SecurityGroupService_DescribeSecurityGroupSnapshots = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DescribeSecurityGroupSnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DescribeSecurityGroupSnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSecurityGroupSnapshots(in)
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

var _flag_SecurityGroupService_DeleteSecurityGroupSnapshots = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DeleteSecurityGroupSnapshots(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DeleteSecurityGroupSnapshotsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteSecurityGroupSnapshots(in)
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

var _flag_SecurityGroupService_RollbackSecurityGroup = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_RollbackSecurityGroup(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.RollbackSecurityGroupInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.RollbackSecurityGroup(in)
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

var _flag_SecurityGroupService_DescribeSecurityGroupIPSets = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DescribeSecurityGroupIPSets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DescribeSecurityGroupIPSetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DescribeSecurityGroupIPSets(in)
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

var _flag_SecurityGroupService_CreateSecurityGroupIPSet = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_CreateSecurityGroupIPSet(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.CreateSecurityGroupIPSetInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CreateSecurityGroupIPSet(in)
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

var _flag_SecurityGroupService_DeleteSecurityGroupIPSets = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_DeleteSecurityGroupIPSets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.DeleteSecurityGroupIPSetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.DeleteSecurityGroupIPSets(in)
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

var _flag_SecurityGroupService_ModifySecurityGroupIPSetAttributes = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_ModifySecurityGroupIPSetAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.ModifySecurityGroupIPSetAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.ModifySecurityGroupIPSetAttributes(in)
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

var _flag_SecurityGroupService_CopySecurityGroupIPSets = []cli.Flag{ /* fields */ }

func _cmd_SecurityGroupService_CopySecurityGroupIPSets(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewSecurityGroupService(conf, zone)

	in := new(pb.CopySecurityGroupIPSetsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
	}

	out, err := qc.CopySecurityGroupIPSets(in)
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
