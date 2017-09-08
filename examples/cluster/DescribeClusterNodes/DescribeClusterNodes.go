// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"

	qc "github.com/chai2010/qingcloud-go"
	"github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/spec.pb"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()

	conf := config.MustLoadUserConfig()
	conf.JSONAllowUnknownFields = true
	conf.LogLevel = "debug" // debug/warn

	qc.SetLogLevel(conf.LogLevel)

	clusterService, err := pb.NewClusterService(conf, "pek3a")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := clusterService.DescribeClusterNodes(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(encodeJSON(reply))
}

func encodeJSON(m proto.Message) string {
	jsonMarshaler := &jsonpb.Marshaler{
		OrigName:     true,
		Indent:       "  ",
		EnumsAsInts:  true,
		EmitDefaults: true,
	}

	s, err := jsonMarshaler.MarshalToString(m)
	if err != nil {
		log.Fatal(err)
	}
	return s
}