// @Version: 1.0.0
// @Author: Bruce
// @TIME  : 2020/3/2 20:26
// @E-mail: mzpy_1119@126.com

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"strconv"
)

type tomlConf struct {
	Title string `toml:"title"`
	SSH sshConf
	SSH_Bind sshBindConf
}

type sshConf struct {
	Host string `toml:"ssh_host"`
	Port int `toml:"ssh_port"`
	User string `toml:"ssh_user"`
	Pkey string `toml:"ssh_pkey"`
	Pwd string `toml:"ssh_pass"`
}

type sshBindConf struct {
	RemoteAddr string `toml:"remote_bind_addr"`
	LocalAddr string `toml:"local_bind_addr"`
	BindPort [][]int `toml:"bind_port"`

}

func main()  {
	var filePath = "./conf/app.toml"
	var config tomlConf

	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		fmt.Println(err)
		return
	}

	key, err := ioutil.ReadFile(config.SSH.Pkey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(config.SSH.Pwd))
	if err != nil {
		signer, err = ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
		}
	}

	conf := &ssh.ClientConfig{
		User: "work",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Loop local port array
	for _, p := range config.SSH_Bind.BindPort {
		go func(port []int) {
			localAddr := fmt.Sprintf("%s:%s", config.SSH_Bind.LocalAddr, strconv.Itoa(port[0]))
			fmt.Print("Listen success: ", localAddr, "\n")

			localListener, err := net.Listen("tcp", localAddr)

			if err != nil {
				log.Fatalf("net.Listen failed: %v", err)
			}

			for {
				// Setup localConn (type net.Conn)
				localConn, err := localListener.Accept()
				if err != nil {
					log.Fatalf("listen.Accept failed: %v", err)
				}
				go forward(localConn, conf, config, port[1])
			}
		}(p)
	}
	select{}
}
