/*
Copyright © 2021 Yugabyte Support

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package xcluster

import (
	"bytes"
	fmt "fmt"

	"github.com/go-logr/logr"
	. "github.com/icza/gox/gox"
	"github.com/spf13/cobra"
	"github.com/yugabyte/yb-tools/yugatool/api/yb/common"
	"github.com/yugabyte/yb-tools/yugatool/api/yb/master"
	"github.com/yugabyte/yb-tools/yugatool/pkg/client"
	"github.com/yugabyte/yb-tools/yugatool/pkg/cmdutil"
	"github.com/yugabyte/yb-tools/yugatool/pkg/util"
)

func InitProducerCmd(ctx *cmdutil.YugatoolContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init_producer",
		Short: "Generate commands to init xCluster replication",
		Long:  `Generate commands to init xCluster replication`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := ctx.WithCmd(cmd).Setup()
			if err != nil {
				return err
			}
			defer ctx.Client.Close()

			return runInitProducer(ctx.Log, ctx.Client)
		},
	}

	return cmd
}

func runInitProducer(log logr.Logger, c *client.YBClient) error {
	tables, err := c.Master.MasterService.ListTables(&master.ListTablesRequestPB{
		ExcludeSystemTables: NewBool(true),
		RelationTypeFilter:  []master.RelationType{master.RelationType_USER_TABLE_RELATION},
	})
	if err != nil {
		return err
	}

	//masters, err := c.Master.MasterService.ListMasters(&master.ListMastersRequestPB{})
	//if err != nil {
	//	return err
	//}

	var initCDCCommand bytes.Buffer

	initCDCCommand.WriteString("yb-admin -master_addresses ")
	initCDCCommand.WriteString("$PRODUCER_MASTERS")
	//for i, master := range masters.GetMasters() {
	//	hostports := master.Registration.GetPrivateRpcAddresses()
	//	initCDCCommand.WriteString(hostports[0].GetHost())
	//	initCDCCommand.WriteRune(':')
	//	initCDCCommand.WriteString(strconv.Itoa(int(hostports[0].GetPort())))
	//	if i+1 < len(masters.GetMasters()) {
	//		initCDCCommand.WriteRune(',')
	//	}
	//}

	initCDCCommand.WriteRune(' ')

	if util.HasTLS(c.Config.GetTlsOpts()) {
		initCDCCommand.WriteString("-certs_dir_name $CERTS_DIR ")
	}

	initCDCCommand.WriteString("bootstrap_cdc_producer ")
	for i, table := range tables.GetTables() {
		if table.TableType.Number() == common.TableType_YQL_TABLE_TYPE.Number() {
			initCDCCommand.Write(table.GetId())
			if i+1 < len(tables.GetTables()) {
				initCDCCommand.WriteRune(',')
			}
		}
	}

	fmt.Println(initCDCCommand.String())
	return nil
}
