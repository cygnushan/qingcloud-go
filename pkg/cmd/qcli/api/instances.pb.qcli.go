// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: instances.proto

package qcli_pb

import (
	"fmt"

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

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdInstanceService)
}

var CmdInstanceService = cli.Command{
	Name:    "instance",
	Aliases: []string{},
	Usage:   "manage InstanceService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeInstances",
			Aliases: []string{},
			Usage:   "DescribeInstances",
			Action:  cmdDescribeInstances,
		},
		{
			Name:    "RunInstances",
			Aliases: []string{},
			Usage:   "RunInstances",
			Action:  cmdRunInstances,
		},
		{
			Name:    "TerminateInstances",
			Aliases: []string{},
			Usage:   "TerminateInstances",
			Action:  cmdTerminateInstances,
		},
		{
			Name:    "StartInstances",
			Aliases: []string{},
			Usage:   "StartInstances",
			Action:  cmdStartInstances,
		},
		{
			Name:    "StopInstances",
			Aliases: []string{},
			Usage:   "StopInstances",
			Action:  cmdStopInstances,
		},
		{
			Name:    "RestartInstances",
			Aliases: []string{},
			Usage:   "RestartInstances",
			Action:  cmdRestartInstances,
		},
		{
			Name:    "ResetInstances",
			Aliases: []string{},
			Usage:   "ResetInstances",
			Action:  cmdResetInstances,
		},
		{
			Name:    "ResizeInstances",
			Aliases: []string{},
			Usage:   "ResizeInstances",
			Action:  cmdResizeInstances,
		},
		{
			Name:    "ModifyInstanceAttributes",
			Aliases: []string{},
			Usage:   "ModifyInstanceAttributes",
			Action:  cmdModifyInstanceAttributes,
		},
		{
			Name:    "DescribeInstanceTypes",
			Aliases: []string{},
			Usage:   "DescribeInstanceTypes",
			Action:  cmdDescribeInstanceTypes,
		},
		{
			Name:    "CreateBrokers",
			Aliases: []string{},
			Usage:   "CreateBrokers",
			Action:  cmdCreateBrokers,
		},
		{
			Name:    "DeleteBrokers",
			Aliases: []string{},
			Usage:   "DeleteBrokers",
			Action:  cmdDeleteBrokers,
		},
	},
}

func cmdDescribeInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.DescribeInstancesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeInstances(in)
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

func cmdRunInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.RunInstancesInput)

	// TODO: fill field from flags

	out, err := qc.RunInstances(in)
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

func cmdTerminateInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.TerminateInstancesInput)

	// TODO: fill field from flags

	out, err := qc.TerminateInstances(in)
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

func cmdStartInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.StartInstancesInput)

	// TODO: fill field from flags

	out, err := qc.StartInstances(in)
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

func cmdStopInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.StopInstancesInput)

	// TODO: fill field from flags

	out, err := qc.StopInstances(in)
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

func cmdRestartInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.RestartInstancesInput)

	// TODO: fill field from flags

	out, err := qc.RestartInstances(in)
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

func cmdResetInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.ResetInstancesInput)

	// TODO: fill field from flags

	out, err := qc.ResetInstances(in)
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

func cmdResizeInstances(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.ResizeInstancesInput)

	// TODO: fill field from flags

	out, err := qc.ResizeInstances(in)
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

func cmdModifyInstanceAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.ModifyInstanceAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyInstanceAttributes(in)
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

func cmdDescribeInstanceTypes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.DescribeInstanceTypesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeInstanceTypes(in)
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

func cmdCreateBrokers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.CreateBrokersInput)

	// TODO: fill field from flags

	out, err := qc.CreateBrokers(in)
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

func cmdDeleteBrokers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewInstanceService(conf, zone)

	in := new(pb.DeleteBrokersInput)

	// TODO: fill field from flags

	out, err := qc.DeleteBrokers(in)
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
