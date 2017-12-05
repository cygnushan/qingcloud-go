// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/golang
// source: dns.proto

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
	AllCommands = append(AllCommands, CmdDNSAliasService)
}

var CmdDNSAliasService = cli.Command{
	Name:    "dnsalias",
	Aliases: []string{},
	Usage:   "manage DNSAliasService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeDNSAliases",
			Aliases: []string{},
			Usage:   "DescribeDNSAliases",
			Action:  cmdDescribeDNSAliases,
		},
		{
			Name:    "AssociateDNSAlias",
			Aliases: []string{},
			Usage:   "AssociateDNSAlias",
			Action:  cmdAssociateDNSAlias,
		},
		{
			Name:    "DissociateDNSAliases",
			Aliases: []string{},
			Usage:   "DissociateDNSAliases",
			Action:  cmdDissociateDNSAliases,
		},
		{
			Name:    "GetDNSLabel",
			Aliases: []string{},
			Usage:   "GetDNSLabel",
			Action:  cmdGetDNSLabel,
		},
	},
}

func cmdDescribeDNSAliases(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewDNSAliasService(conf, zone)

	in := new(pb.DescribeDNSAliasesInput)

	// TODO: fill field from flags

	out, err := qc.DescribeDNSAliases(in)
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

func cmdAssociateDNSAlias(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewDNSAliasService(conf, zone)

	in := new(pb.AssociateDNSAliasInput)

	// TODO: fill field from flags

	out, err := qc.AssociateDNSAlias(in)
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

func cmdDissociateDNSAliases(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewDNSAliasService(conf, zone)

	in := new(pb.DissociateDNSAliasesInput)

	// TODO: fill field from flags

	out, err := qc.DissociateDNSAliases(in)
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

func cmdGetDNSLabel(c *cli.Context) error {
	var conf *config.Config
	var zone string
	qc := pb.NewDNSAliasService(conf, zone)

	in := new(pb.GetDNSLabelInput)

	// TODO: fill field from flags

	out, err := qc.GetDNSLabel(in)
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
