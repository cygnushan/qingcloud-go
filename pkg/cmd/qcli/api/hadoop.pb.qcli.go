// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: hadoop.proto

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
	AllCommands = append(AllCommands, CmdHadoopService)
}

var CmdHadoopService = cli.Command{
	Name:    "hadoop",
	Aliases: []string{},
	Usage:   "manage HadoopService",
	Subcommands: []cli.Command{
		{
			Name:    "AddHadoopNodes",
			Aliases: []string{},
			Usage:   "AddHadoopNodes",
			Action:  cmdAddHadoopNodes,
		},
		{
			Name:    "DeleteHadoopNodes",
			Aliases: []string{},
			Usage:   "DeleteHadoopNodes",
			Action:  cmdDeleteHadoopNodes,
		},
		{
			Name:    "StartHadoops",
			Aliases: []string{},
			Usage:   "StartHadoops",
			Action:  cmdStartHadoops,
		},
		{
			Name:    "StopHadoops",
			Aliases: []string{},
			Usage:   "StopHadoops",
			Action:  cmdStopHadoops,
		},
	},
}

func cmdAddHadoopNodes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewHadoopService(conf, zone)

	in := new(pb.AddHadoopNodesInput)

	// TODO: fill field from flags

	out, err := qc.AddHadoopNodes(in)
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

func cmdDeleteHadoopNodes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewHadoopService(conf, zone)

	in := new(pb.DeleteHadoopNodesInput)

	// TODO: fill field from flags

	out, err := qc.DeleteHadoopNodes(in)
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

func cmdStartHadoops(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewHadoopService(conf, zone)

	in := new(pb.StartHadoopsInput)

	// TODO: fill field from flags

	out, err := qc.StartHadoops(in)
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

func cmdStopHadoops(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewHadoopService(conf, zone)

	in := new(pb.StopHadoopsInput)

	// TODO: fill field from flags

	out, err := qc.StopHadoops(in)
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
