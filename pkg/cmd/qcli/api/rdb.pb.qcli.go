// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: rdb.proto

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
	AllCommands = append(AllCommands, CmdRDBService)
}

var CmdRDBService = cli.Command{
	Name:    "rdb",
	Aliases: []string{},
	Usage:   "manage RDBService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateRDB",
			Aliases: []string{},
			Usage:   "CreateRDB",
			Action:  cmdCreateRDB,
		},
		{
			Name:    "DescribeRDBs",
			Aliases: []string{},
			Usage:   "DescribeRDBs",
			Action:  cmdDescribeRDBs,
		},
		{
			Name:    "DeleteRDBs",
			Aliases: []string{},
			Usage:   "DeleteRDBs",
			Action:  cmdDeleteRDBs,
		},
		{
			Name:    "StartRDBs",
			Aliases: []string{},
			Usage:   "StartRDBs",
			Action:  cmdStartRDBs,
		},
		{
			Name:    "StopRDBs",
			Aliases: []string{},
			Usage:   "StopRDBs",
			Action:  cmdStopRDBs,
		},
		{
			Name:    "ResizeRDBs",
			Aliases: []string{},
			Usage:   "ResizeRDBs",
			Action:  cmdResizeRDBs,
		},
		{
			Name:    "RDBsLeaveVxnet",
			Aliases: []string{},
			Usage:   "RDBsLeaveVxnet",
			Action:  cmdRDBsLeaveVxnet,
		},
		{
			Name:    "RDBsJoinVxnet",
			Aliases: []string{},
			Usage:   "RDBsJoinVxnet",
			Action:  cmdRDBsJoinVxnet,
		},
		{
			Name:    "CreateRDBFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateRDBFromSnapshot",
			Action:  cmdCreateRDBFromSnapshot,
		},
		{
			Name:    "CreateTempRDBInstanceFromSnapshot",
			Aliases: []string{},
			Usage:   "CreateTempRDBInstanceFromSnapshot",
			Action:  cmdCreateTempRDBInstanceFromSnapshot,
		},
		{
			Name:    "GetRDBInstanceFiles",
			Aliases: []string{},
			Usage:   "GetRDBInstanceFiles",
			Action:  cmdGetRDBInstanceFiles,
		},
		{
			Name:    "CopyRDBInstanceFilesToFTP",
			Aliases: []string{},
			Usage:   "CopyRDBInstanceFilesToFTP",
			Action:  cmdCopyRDBInstanceFilesToFTP,
		},
		{
			Name:    "PurgeRDBLogs",
			Aliases: []string{},
			Usage:   "PurgeRDBLogs",
			Action:  cmdPurgeRDBLogs,
		},
		{
			Name:    "CeaseRDBInstance",
			Aliases: []string{},
			Usage:   "CeaseRDBInstance",
			Action:  cmdCeaseRDBInstance,
		},
		{
			Name:    "ModifyRDBParameters",
			Aliases: []string{},
			Usage:   "ModifyRDBParameters",
			Action:  cmdModifyRDBParameters,
		},
		{
			Name:    "ApplyRDBParameterGroup",
			Aliases: []string{},
			Usage:   "ApplyRDBParameterGroup",
			Action:  cmdApplyRDBParameterGroup,
		},
		{
			Name:    "DescribeRDBParameters",
			Aliases: []string{},
			Usage:   "DescribeRDBParameters",
			Action:  cmdDescribeRDBParameters,
		},
	},
}

func cmdCreateRDB(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.CreateRDBInput)

	// TODO: fill field from flags

	out, err := qc.CreateRDB(in)
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

func cmdDescribeRDBs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.DescribeRDBsInput)

	// TODO: fill field from flags

	out, err := qc.DescribeRDBs(in)
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

func cmdDeleteRDBs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.DeleteRDBsInput)

	// TODO: fill field from flags

	out, err := qc.DeleteRDBs(in)
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

func cmdStartRDBs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.StartRDBsInput)

	// TODO: fill field from flags

	out, err := qc.StartRDBs(in)
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

func cmdStopRDBs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.StopRDBsInput)

	// TODO: fill field from flags

	out, err := qc.StopRDBs(in)
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

func cmdResizeRDBs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.ResizeRDBsInput)

	// TODO: fill field from flags

	out, err := qc.ResizeRDBs(in)
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

func cmdRDBsLeaveVxnet(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.RDBsLeaveVxnetInput)

	// TODO: fill field from flags

	out, err := qc.RDBsLeaveVxnet(in)
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

func cmdRDBsJoinVxnet(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.RDBsJoinVxnetInput)

	// TODO: fill field from flags

	out, err := qc.RDBsJoinVxnet(in)
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

func cmdCreateRDBFromSnapshot(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.CreateRDBFromSnapshotInput)

	// TODO: fill field from flags

	out, err := qc.CreateRDBFromSnapshot(in)
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

func cmdCreateTempRDBInstanceFromSnapshot(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.CreateTempRDBInstanceFromSnapshotInput)

	// TODO: fill field from flags

	out, err := qc.CreateTempRDBInstanceFromSnapshot(in)
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

func cmdGetRDBInstanceFiles(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.GetRDBInstanceFilesInput)

	// TODO: fill field from flags

	out, err := qc.GetRDBInstanceFiles(in)
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

func cmdCopyRDBInstanceFilesToFTP(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.CopyRDBInstanceFilesToFTPInput)

	// TODO: fill field from flags

	out, err := qc.CopyRDBInstanceFilesToFTP(in)
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

func cmdPurgeRDBLogs(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.PurgeRDBLogsInput)

	// TODO: fill field from flags

	out, err := qc.PurgeRDBLogs(in)
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

func cmdCeaseRDBInstance(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.CeaseRDBInstanceInput)

	// TODO: fill field from flags

	out, err := qc.CeaseRDBInstance(in)
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

func cmdModifyRDBParameters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.ModifyRDBParametersInput)

	// TODO: fill field from flags

	out, err := qc.ModifyRDBParameters(in)
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

func cmdApplyRDBParameterGroup(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.ApplyRDBParameterGroupInput)

	// TODO: fill field from flags

	out, err := qc.ApplyRDBParameterGroup(in)
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

func cmdDescribeRDBParameters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewRDBService(conf, zone)

	in := new(pb.DescribeRDBParametersInput)

	// TODO: fill field from flags

	out, err := qc.DescribeRDBParameters(in)
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
