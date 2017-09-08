// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	qc "github.com/chai2010/qingcloud-go"
	"github.com/chai2010/qingcloud-go/config"
	pb "github.com/chai2010/qingcloud-go/spec.pb"
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

	reply, err := clusterService.DescribeClusters(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(encodeJSON(reply)))
}

func encodeJSON(m interface{}) []byte {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return nil
	}
	data = bytes.Replace(data, []byte("\n"), []byte("\r\n"), -1)
	return data
}
