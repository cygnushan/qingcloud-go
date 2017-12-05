// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: span.proto

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
	AllCommands = append(AllCommands, CmdSpanService)
}

var CmdSpanService = cli.Command{
	Name:    "span",
	Aliases: []string{},
	Usage:   "manage SpanService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateSpan",
			Aliases: []string{},
			Usage:   "CreateSpan",
			Action:  cmdCreateSpan,
		},
		{
			Name:    "DescribeSpans",
			Aliases: []string{},
			Usage:   "DescribeSpans",
			Action:  cmdDescribeSpans,
		},
		{
			Name:    "DeleteSpans",
			Aliases: []string{},
			Usage:   "DeleteSpans",
			Action:  cmdDeleteSpans,
		},
		{
			Name:    "AddSpanMembers",
			Aliases: []string{},
			Usage:   "AddSpanMembers",
			Action:  cmdAddSpanMembers,
		},
		{
			Name:    "RemoveSpanMembers",
			Aliases: []string{},
			Usage:   "RemoveSpanMembers",
			Action:  cmdRemoveSpanMembers,
		},
		{
			Name:    "ModifySpanAttributes",
			Aliases: []string{},
			Usage:   "ModifySpanAttributes",
			Action:  cmdModifySpanAttributes,
		},
		{
			Name:    "UpdateSpan",
			Aliases: []string{},
			Usage:   "UpdateSpan",
			Action:  cmdUpdateSpan,
		},
	},
}

func cmdCreateSpan(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.CreateSpanInput)

	// TODO: fill field from flags

	out, err := qc.CreateSpan(in)
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

func cmdDescribeSpans(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.DescribeSpansInput)

	// TODO: fill field from flags

	out, err := qc.DescribeSpans(in)
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

func cmdDeleteSpans(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.DeleteSpansInput)

	// TODO: fill field from flags

	out, err := qc.DeleteSpans(in)
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

func cmdAddSpanMembers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.AddSpanMembersInput)

	// TODO: fill field from flags

	out, err := qc.AddSpanMembers(in)
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

func cmdRemoveSpanMembers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.RemoveSpanMembersInput)

	// TODO: fill field from flags

	out, err := qc.RemoveSpanMembers(in)
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

func cmdModifySpanAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.ModifySpanAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifySpanAttributes(in)
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

func cmdUpdateSpan(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewSpanService(conf, zone)

	in := new(pb.UpdateSpanInput)

	// TODO: fill field from flags

	out, err := qc.UpdateSpan(in)
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