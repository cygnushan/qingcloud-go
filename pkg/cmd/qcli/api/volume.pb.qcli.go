// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: volume.proto

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
	AllCommands = append(AllCommands, CmdVolumesService)
}

var CmdVolumesService = cli.Command{
	Name:    "volumes",
	Aliases: []string{},
	Usage:   "manage VolumesService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeVolumes",
			Aliases: []string{},
			Usage:   "DescribeVolumes",
			Action:  cmdDescribeVolumes,
		},
		{
			Name:    "CreateVolumes",
			Aliases: []string{},
			Usage:   "CreateVolumes",
			Action:  cmdCreateVolumes,
		},
		{
			Name:    "DeleteVolumes",
			Aliases: []string{},
			Usage:   "DeleteVolumes",
			Action:  cmdDeleteVolumes,
		},
		{
			Name:    "AttachVolumes",
			Aliases: []string{},
			Usage:   "AttachVolumes",
			Action:  cmdAttachVolumes,
		},
		{
			Name:    "DetachVolumes",
			Aliases: []string{},
			Usage:   "DetachVolumes",
			Action:  cmdDetachVolumes,
		},
		{
			Name:    "ResizeVolumes",
			Aliases: []string{},
			Usage:   "ResizeVolumes",
			Action:  cmdResizeVolumes,
		},
		{
			Name:    "ModifyVolumeAttributes",
			Aliases: []string{},
			Usage:   "ModifyVolumeAttributes",
			Action:  cmdModifyVolumeAttributes,
		},
	},
}

func cmdDescribeVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.DescribeVolumesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeVolumes(in)
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

func cmdCreateVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.CreateVolumesInput)

	// TODO: fill field from flags

	out, err := qc.CreateVolumes(in)
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

func cmdDeleteVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.DeleteVolumesInput)

	// TODO: fill field from flags

	out, err := qc.DeleteVolumes(in)
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

func cmdAttachVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.AttachVolumesInput)

	// TODO: fill field from flags

	out, err := qc.AttachVolumes(in)
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

func cmdDetachVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.DetachVolumesInput)

	// TODO: fill field from flags

	out, err := qc.DetachVolumes(in)
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

func cmdResizeVolumes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.ResizeVolumesInput)

	// TODO: fill field from flags

	out, err := qc.ResizeVolumes(in)
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

func cmdModifyVolumeAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewVolumesService(conf, zone)

	in := new(pb.ModifyVolumeAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyVolumeAttributes(in)
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
