// Code generated by protoc-gen-qingcloud. DO NOT EDIT.
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud
// plugin: https://github.com/chai2010/qingcloud-go/tree/master/pkg/cmd/protoc-gen-qingcloud/generator/qcli
// source: key_pair.proto

package qcli_pb

import (
	"encoding/json"
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
	_ = json.Marshal
	_ = os.Stdin

	_ = cli.Command{}
	_ = jsonpb.Unmarshal
	_ = proto.Marshal

	_ = config.Config{}
	_ = logger.Info
	_ = pb.AlarmService{}
)

func init() {
	AllCommands = append(AllCommands, CmdKeyPairService)
}

var CmdKeyPairService = cli.Command{
	Name:    "keypair",
	Aliases: []string{},
	Usage:   "manage KeyPairService",
	Subcommands: []cli.Command{
		{
			Name:    "DescribeKeyPairs",
			Aliases: []string{},
			Usage:   "DescribeKeyPairs",
			Flags:   _flag_KeyPairService_DescribeKeyPairs,
			Action:  _func_KeyPairService_DescribeKeyPairs,
		},
		{
			Name:    "CreateKeyPair",
			Aliases: []string{},
			Usage:   "CreateKeyPair",
			Flags:   _flag_KeyPairService_CreateKeyPair,
			Action:  _func_KeyPairService_CreateKeyPair,
		},
		{
			Name:    "DeleteKeyPairs",
			Aliases: []string{},
			Usage:   "DeleteKeyPairs",
			Flags:   _flag_KeyPairService_DeleteKeyPairs,
			Action:  _func_KeyPairService_DeleteKeyPairs,
		},
		{
			Name:    "AttachKeyPairs",
			Aliases: []string{},
			Usage:   "AttachKeyPairs",
			Flags:   _flag_KeyPairService_AttachKeyPairs,
			Action:  _func_KeyPairService_AttachKeyPairs,
		},
		{
			Name:    "DetachKeyPairs",
			Aliases: []string{},
			Usage:   "DetachKeyPairs",
			Flags:   _flag_KeyPairService_DetachKeyPairs,
			Action:  _func_KeyPairService_DetachKeyPairs,
		},
		{
			Name:    "ModifyKeyPairAttributes",
			Aliases: []string{},
			Usage:   "ModifyKeyPairAttributes",
			Flags:   _flag_KeyPairService_ModifyKeyPairAttributes,
			Action:  _func_KeyPairService_ModifyKeyPairAttributes,
		},
	},
}

var _flag_KeyPairService_DescribeKeyPairs = []cli.Flag{
	cli.StringFlag{
		Name:  "keypairs",
		Usage: "keypairs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instance_id",
		Usage: "instance id",
		Value: "",
	},
	cli.StringFlag{
		Name:  "owner",
		Usage: "owner",
		Value: "",
	},
	cli.StringFlag{
		Name:  "encrypt_method",
		Usage: "encrypt method",
		Value: "",
	},
	cli.StringFlag{
		Name:  "search_word",
		Usage: "search word",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tags",
		Usage: "tags",
		Value: "", // json: slice/message/map/time
	},
	cli.IntFlag{
		Name:  "verbose",
		Usage: "verbose",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "offset",
		Usage: "offset",
		Value: 0,
	},
	cli.IntFlag{
		Name:  "limit",
		Usage: "limit",
		Value: 0,
	},
}

func _func_KeyPairService_DescribeKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DescribeKeyPairsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypairs") {
			if err := json.Unmarshal([]byte(c.String("keypairs")), &in.Keypairs); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("instance_id") {
			in.InstanceId = proto.String(c.String("instance_id"))
		}
		if c.IsSet("owner") {
			in.Owner = proto.String(c.String("owner"))
		}
		if c.IsSet("encrypt_method") {
			in.EncryptMethod = proto.String(c.String("encrypt_method"))
		}
		if c.IsSet("search_word") {
			in.SearchWord = proto.String(c.String("search_word"))
		}
		if c.IsSet("tags") {
			if err := json.Unmarshal([]byte(c.String("tags")), &in.Tags); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("verbose") {
			in.Verbose = proto.Int32(int32(c.Int("verbose")))
		}
		if c.IsSet("offset") {
			in.Offset = proto.Int32(int32(c.Int("offset")))
		}
		if c.IsSet("limit") {
			in.Limit = proto.Int32(int32(c.Int("limit")))
		}
	}

	out, err := qc.DescribeKeyPairs(in)
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

var _flag_KeyPairService_CreateKeyPair = []cli.Flag{
	cli.StringFlag{
		Name:  "keypair_name",
		Usage: "keypair name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "mode",
		Usage: "mode",
		Value: "",
	},
	cli.StringFlag{
		Name:  "encrypt_method",
		Usage: "encrypt method",
		Value: "",
	},
	cli.StringFlag{
		Name:  "public_key",
		Usage: "public key",
		Value: "",
	},
}

func _func_KeyPairService_CreateKeyPair(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.CreateKeyPairInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypair_name") {
			in.KeypairName = proto.String(c.String("keypair_name"))
		}
		if c.IsSet("mode") {
			in.Mode = proto.String(c.String("mode"))
		}
		if c.IsSet("encrypt_method") {
			in.EncryptMethod = proto.String(c.String("encrypt_method"))
		}
		if c.IsSet("public_key") {
			in.PublicKey = proto.String(c.String("public_key"))
		}
	}

	out, err := qc.CreateKeyPair(in)
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

var _flag_KeyPairService_DeleteKeyPairs = []cli.Flag{
	cli.StringFlag{
		Name:  "keypairs",
		Usage: "keypairs",
		Value: "", // json: slice/message/map/time
	},
}

func _func_KeyPairService_DeleteKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DeleteKeyPairsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypairs") {
			if err := json.Unmarshal([]byte(c.String("keypairs")), &in.Keypairs); err != nil {
				logger.Fatal(err)
			}
		}
	}

	out, err := qc.DeleteKeyPairs(in)
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

var _flag_KeyPairService_AttachKeyPairs = []cli.Flag{
	cli.StringFlag{
		Name:  "keypairs",
		Usage: "keypairs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
}

func _func_KeyPairService_AttachKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.AttachKeyPairsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypairs") {
			if err := json.Unmarshal([]byte(c.String("keypairs")), &in.Keypairs); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				logger.Fatal(err)
			}
		}
	}

	out, err := qc.AttachKeyPairs(in)
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

var _flag_KeyPairService_DetachKeyPairs = []cli.Flag{
	cli.StringFlag{
		Name:  "keypairs",
		Usage: "keypairs",
		Value: "", // json: slice/message/map/time
	},
	cli.StringFlag{
		Name:  "instances",
		Usage: "instances",
		Value: "", // json: slice/message/map/time
	},
}

func _func_KeyPairService_DetachKeyPairs(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.DetachKeyPairsInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypairs") {
			if err := json.Unmarshal([]byte(c.String("keypairs")), &in.Keypairs); err != nil {
				logger.Fatal(err)
			}
		}
		if c.IsSet("instances") {
			if err := json.Unmarshal([]byte(c.String("instances")), &in.Instances); err != nil {
				logger.Fatal(err)
			}
		}
	}

	out, err := qc.DetachKeyPairs(in)
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

var _flag_KeyPairService_ModifyKeyPairAttributes = []cli.Flag{
	cli.StringFlag{
		Name:  "keypair",
		Usage: "keypair",
		Value: "",
	},
	cli.StringFlag{
		Name:  "keypair_name",
		Usage: "keypair name",
		Value: "",
	},
	cli.StringFlag{
		Name:  "description",
		Usage: "description",
		Value: "",
	},
}

func _func_KeyPairService_ModifyKeyPairAttributes(c *cli.Context) error {
	conf := config.MustLoadConfigFromFilepath(c.GlobalString("config"))
	zone := c.GlobalString("zone")
	qc := pb.NewKeyPairService(conf, zone)

	in := new(pb.ModifyKeyPairAttributesInput)

	if c.NArg() == 1 && c.Args().Get(0) == "-" {
		// read from stdin json
		err := jsonpb.Unmarshal(os.Stdin, in)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		// read from flags
		if c.IsSet("keypair") {
			in.Keypair = proto.String(c.String("keypair"))
		}
		if c.IsSet("keypair_name") {
			in.KeypairName = proto.String(c.String("keypair_name"))
		}
		if c.IsSet("description") {
			in.Description = proto.String(c.String("description"))
		}
	}

	out, err := qc.ModifyKeyPairAttributes(in)
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
