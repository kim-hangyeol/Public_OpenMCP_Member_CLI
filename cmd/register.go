/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net"
	cobrautil "openmcp/omctl/util"
	"os"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 && args[0] == "openmcp" {
			registerASOpenMCP()
		} else if len(args) != 0 && args[0] == "member" {
			if args[1] == "" {
				fmt.Println("You Must Provide Cluster IP")
			} else {
				registerMemberToOpenMCP(args[1])
			}
		}
	},
}

func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func registerASOpenMCP() {
	c := cobrautil.GetOmcpctlConf("/var/lib/omctl/config.yaml")

	cobrautil.CmdExec("umount -l /mnt")
	defer cobrautil.CmdExec("umount -l /mnt")

	cobrautil.CmdExec("mount -t nfs " + c.NfsServer + ":/home/nfs/ /mnt")

	openmcpIP := GetOutboundIP()

	if fileExists("/mnt/openmcp/"+openmcpIP) {
		fmt.Println("Failed Register OpenMCP Master")
		fmt.Println("=> Already Registered OpenMCP :"+openmcpIP)
		return
	}

	cobrautil.CmdExec("mkdir /mnt/openmcp")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP)
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/master")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/master/config")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/master/pki")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members/join")
	cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members/unjoin")

	cobrautil.CmdExec("cp ~/.kube/config /mnt/openmcp/" + openmcpIP + "/master/config/config")
	cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/ca.crt /mnt/openmcp/" + openmcpIP + "/master/pki/ca.crt")
	cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/server.crt /mnt/openmcp/" + openmcpIP + "/master/pki/server.crt")
	cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/server.key /mnt/openmcp/" + openmcpIP + "/master/pki/server.key")

	//SSH Public Key Copy
	cobrautil.CmdExec("cat /mnt/ssh/id_rsa.pub >> /root/.ssh/authorized_keys")

	fmt.Println("Success OpenMCP Master Register '" + openmcpIP + "'")
	return

}

func registerMemberToOpenMCP(openmcpIP string) {
	c := cobrautil.GetOmcpctlConf("/var/lib/omctl/config.yaml")

	cobrautil.CmdExec("umount -l /mnt")
	defer cobrautil.CmdExec("umount -l /mnt")

	cobrautil.CmdExec("mount -t nfs " + c.NfsServer + ":/home/nfs/ /mnt")

	memberIP := GetOutboundIP()
	//openmcpIP := cobrautil.Option_ip

	if !fileExists("/mnt/openmcp/" + openmcpIP + "/master") {
		fmt.Println("Failed Register '"+memberIP +"' in OpenMCP Master: "+openmcpIP)
		fmt.Println("=> Not Yet Register OpenMCP.")
		fmt.Println("=> First You Must be Input the Next Command in 'OpenMCP Master Server(" + openmcpIP + ")' : omctl register openmcp")
		return
	}

	if memberIP == openmcpIP {
		fmt.Println("Failed Register '"+memberIP +"' in OpenMCP Master: "+openmcpIP)
		fmt.Println("=> Can Not Self regist. [My_IP '"+memberIP+"', OpenMCP_IP '"+openmcpIP+"']")
		return
	}

	// Already Regist
	if fileExists("/mnt/openmcp/"+openmcpIP+"/members/unjoin/"+memberIP) {
		fmt.Println("Failed Register '"+memberIP+"' in OpenMCP Master: "+openmcpIP)
		fmt.Println("=> Already Regist")
		return

	} else if fileExists("/mnt/openmcp/"+openmcpIP+"/members/join/"+memberIP) {
		fmt.Println("Failed Register '"+memberIP +"' in OpenMCP Master: "+openmcpIP)
		fmt.Println("=> Already Joined by OpenMCP '"+openmcpIP+"'")
		return

	} else {
		cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP)
		cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/config")
		cobrautil.CmdExec("mkdir /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/pki")

		cobrautil.CmdExec("cp ~/.kube/config /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/config/config")
		cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/ca.crt /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/pki/ca.crt")
		cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/server.crt /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/pki/server.crt")
		cobrautil.CmdExec("cp /etc/kubernetes/pki/etcd/server.key /mnt/openmcp/" + openmcpIP + "/members/unjoin/" + memberIP + "/pki/server.key")

		// SSH Public Key Copy
		cobrautil.CmdExec("cat /mnt/ssh/id_rsa.pub >> /root/.ssh/authorized_keys")

		fmt.Println("Success Register '"+memberIP +"' in OpenMCP Master: "+openmcpIP)
		return
	}
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
