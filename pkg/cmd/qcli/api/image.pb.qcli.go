// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: image.proto

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
	AllCommands = append(AllCommands, CmdImageService)
}

var CmdImageService = cli.Command{
	Name:    "image",
	Aliases: []string{},
	Usage:   "manage ImageService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeImages",
			Aliases: []string{},
			Usage:   "DescribeImages",
			Action:  cmdDescribeImages,
		},
		{
			Name:    "CaptureInstance",
			Aliases: []string{},
			Usage:   "CaptureInstance",
			Action:  cmdCaptureInstance,
		},
		{
			Name:    "DeleteImages",
			Aliases: []string{},
			Usage:   "DeleteImages",
			Action:  cmdDeleteImages,
		},
		{
			Name:    "ModifyImageAttributes",
			Aliases: []string{},
			Usage:   "ModifyImageAttributes",
			Action:  cmdModifyImageAttributes,
		},
		{
			Name:    "GrantImageToUsers",
			Aliases: []string{},
			Usage:   "GrantImageToUsers",
			Action:  cmdGrantImageToUsers,
		},
		{
			Name:    "RevokeImageFromUsers",
			Aliases: []string{},
			Usage:   "RevokeImageFromUsers",
			Action:  cmdRevokeImageFromUsers,
		},
		{
			Name:    "DescribeImageUsers",
			Aliases: []string{},
			Usage:   "DescribeImageUsers",
			Action:  cmdDescribeImageUsers,
		},
		{
			Name:    "CloneImages",
			Aliases: []string{},
			Usage:   "CloneImages",
			Action:  cmdCloneImages,
		},
	},
}

func cmdDescribeImages(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImagesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeImages(in)
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

func cmdCaptureInstance(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CaptureInstanceInput)

	// TODO: fill field from flags

	out, err := qc.CaptureInstance(in)
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

func cmdDeleteImages(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DeleteImagesInput)

	// TODO: fill field from flags

	out, err := qc.DeleteImages(in)
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

func cmdModifyImageAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.ModifyImageAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyImageAttributes(in)
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

func cmdGrantImageToUsers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.GrantImageToUsersInput)

	// TODO: fill field from flags

	out, err := qc.GrantImageToUsers(in)
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

func cmdRevokeImageFromUsers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.RevokeImageFromUsersInput)

	// TODO: fill field from flags

	out, err := qc.RevokeImageFromUsers(in)
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

func cmdDescribeImageUsers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.DescribeImageUsersInput)

	// TODO: fill field from flags

	out, err := qc.DescribeImageUsers(in)
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

func cmdCloneImages(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewImageService(conf, zone)

	in := new(pb.CloneImagesInput)

	// TODO: fill field from flags

	out, err := qc.CloneImages(in)
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
