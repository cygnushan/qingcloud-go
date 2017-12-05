// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: cluster.proto

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
	AllCommands = append(AllCommands, CmdClusterService)
}

var CmdClusterService = cli.Command{
	Name:    "cluster",
	Aliases: []string{},
	Usage:   "manage ClusterService",
	Subcommands: []cli.Command{
		{
			Name:    "CreateCluster",
			Aliases: []string{},
			Usage:   "CreateCluster",
			Action:  cmdCreateCluster,
		},
		{
			Name:    "DescribeClusters",
			Aliases: []string{},
			Usage:   "DescribeClusters",
			Action:  cmdDescribeClusters,
		},
		{
			Name:    "DescribeClusterNodes",
			Aliases: []string{},
			Usage:   "DescribeClusterNodes",
			Action:  cmdDescribeClusterNodes,
		},
		{
			Name:    "StopClusters",
			Aliases: []string{},
			Usage:   "StopClusters",
			Action:  cmdStopClusters,
		},
		{
			Name:    "StartClusters",
			Aliases: []string{},
			Usage:   "StartClusters",
			Action:  cmdStartClusters,
		},
		{
			Name:    "DeleteClusters",
			Aliases: []string{},
			Usage:   "DeleteClusters",
			Action:  cmdDeleteClusters,
		},
		{
			Name:    "Lease",
			Aliases: []string{},
			Usage:   "Lease",
			Action:  cmdLease,
		},
		{
			Name:    "AddClusterNodes",
			Aliases: []string{},
			Usage:   "AddClusterNodes",
			Action:  cmdAddClusterNodes,
		},
		{
			Name:    "DeleteClusterNodes",
			Aliases: []string{},
			Usage:   "DeleteClusterNodes",
			Action:  cmdDeleteClusterNodes,
		},
		{
			Name:    "ResizeCluster",
			Aliases: []string{},
			Usage:   "ResizeCluster",
			Action:  cmdResizeCluster,
		},
		{
			Name:    "ChangeClusterVxnet",
			Aliases: []string{},
			Usage:   "ChangeClusterVxnet",
			Action:  cmdChangeClusterVxnet,
		},
		{
			Name:    "SuspendClusters",
			Aliases: []string{},
			Usage:   "SuspendClusters",
			Action:  cmdSuspendClusters,
		},
		{
			Name:    "UpdateClusterEnvironment",
			Aliases: []string{},
			Usage:   "UpdateClusterEnvironment",
			Action:  cmdUpdateClusterEnvironment,
		},
		{
			Name:    "ModifyClusterAttributes",
			Aliases: []string{},
			Usage:   "ModifyClusterAttributes",
			Action:  cmdModifyClusterAttributes,
		},
		{
			Name:    "ModifyClusterNodeAttributes",
			Aliases: []string{},
			Usage:   "ModifyClusterNodeAttributes",
			Action:  cmdModifyClusterNodeAttributes,
		},
		{
			Name:    "GetClustersStats",
			Aliases: []string{},
			Usage:   "GetClustersStats",
			Action:  cmdGetClustersStats,
		},
		{
			Name:    "DescribeClusterUsers",
			Aliases: []string{},
			Usage:   "DescribeClusterUsers",
			Action:  cmdDescribeClusterUsers,
		},
		{
			Name:    "RestartClusterService",
			Aliases: []string{},
			Usage:   "RestartClusterService",
			Action:  cmdRestartClusterService,
		},
		{
			Name:    "UpgradeClusters",
			Aliases: []string{},
			Usage:   "UpgradeClusters",
			Action:  cmdUpgradeClusters,
		},
		{
			Name:    "AuthorizeClustersBrokerToDeveloper",
			Aliases: []string{},
			Usage:   "AuthorizeClustersBrokerToDeveloper",
			Action:  cmdAuthorizeClustersBrokerToDeveloper,
		},
		{
			Name:    "RevokeClustersBrokerFromDeveloper",
			Aliases: []string{},
			Usage:   "RevokeClustersBrokerFromDeveloper",
			Action:  cmdRevokeClustersBrokerFromDeveloper,
		},
	},
}

func cmdCreateCluster(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.CreateClusterInput)

	// TODO: fill field from flags

	out, err := qc.CreateCluster(in)
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

func cmdDescribeClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.DescribeClustersInput)

	// TODO: fill field from flags

	out, err := qc.DescribeClusters(in)
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

func cmdDescribeClusterNodes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.DescribeClusterNodesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeClusterNodes(in)
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

func cmdStopClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.StopClustersInput)

	// TODO: fill field from flags

	out, err := qc.StopClusters(in)
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

func cmdStartClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.StartClustersInput)

	// TODO: fill field from flags

	out, err := qc.StartClusters(in)
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

func cmdDeleteClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.DeleteClustersInput)

	// TODO: fill field from flags

	out, err := qc.DeleteClusters(in)
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

func cmdLease(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.LeaseInput)

	// TODO: fill field from flags

	out, err := qc.Lease(in)
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

func cmdAddClusterNodes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.AddClusterNodesInput)

	// TODO: fill field from flags

	out, err := qc.AddClusterNodes(in)
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

func cmdDeleteClusterNodes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.DeleteClusterNodesInput)

	// TODO: fill field from flags

	out, err := qc.DeleteClusterNodes(in)
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

func cmdResizeCluster(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.ResizeClusterInput)

	// TODO: fill field from flags

	out, err := qc.ResizeCluster(in)
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

func cmdChangeClusterVxnet(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.ChangeClusterVxnetInput)

	// TODO: fill field from flags

	out, err := qc.ChangeClusterVxnet(in)
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

func cmdSuspendClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.SuspendClustersInput)

	// TODO: fill field from flags

	out, err := qc.SuspendClusters(in)
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

func cmdUpdateClusterEnvironment(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.UpdateClusterEnvironmentInput)

	// TODO: fill field from flags

	out, err := qc.UpdateClusterEnvironment(in)
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

func cmdModifyClusterAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.ModifyClusterAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyClusterAttributes(in)
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

func cmdModifyClusterNodeAttributes(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.ModifyClusterNodeAttributesInput)

	// TODO: fill field from flags

	out, err := qc.ModifyClusterNodeAttributes(in)
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

func cmdGetClustersStats(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.GetClustersStatsInput)

	// TODO: fill field from flags

	out, err := qc.GetClustersStats(in)
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

func cmdDescribeClusterUsers(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.DescribeClusterUsersInput)

	// TODO: fill field from flags

	out, err := qc.DescribeClusterUsers(in)
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

func cmdRestartClusterService(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.RestartClusterServiceInput)

	// TODO: fill field from flags

	out, err := qc.RestartClusterService(in)
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

func cmdUpgradeClusters(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.UpgradeClustersInput)

	// TODO: fill field from flags

	out, err := qc.UpgradeClusters(in)
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

func cmdAuthorizeClustersBrokerToDeveloper(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.AuthorizeClustersBrokerToDeveloperInput)

	// TODO: fill field from flags

	out, err := qc.AuthorizeClustersBrokerToDeveloper(in)
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

func cmdRevokeClustersBrokerFromDeveloper(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewClusterService(conf, zone)

	in := new(pb.RevokeClustersBrokerFromDeveloperInput)

	// TODO: fill field from flags

	out, err := qc.RevokeClustersBrokerFromDeveloper(in)
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
