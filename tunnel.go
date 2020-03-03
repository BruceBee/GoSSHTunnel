// @Version: 1.0.0
// @Author: Bruce
// @TIME  : 2020/3/2 20:26
// @E-mail: mzpy_1119@126.com

package main

import (
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"fmt"
	"strconv"
)



func forward(localConn net.Conn, cliconfig *ssh.ClientConfig, config tomlConf) {

	sshAddr := fmt.Sprintf("%s:%s", config.SSH.Host, strconv.Itoa(config.SSH.Port))
	bindAddr := fmt.Sprintf("%s:%s", config.SSH_Bind.RemotAddr, strconv.Itoa(config.SSH_Bind.RemotPort))

	// Setup sshClientConn (type *ssh.ClientConn)
	sshClientConn, err := ssh.Dial("tcp", sshAddr, cliconfig)
	if err != nil {
		log.Fatalf("ssh.Dial failed: %s", err)
	}


	// Setup sshConn (type net.Conn)
	sshConn, err := sshClientConn.Dial("tcp", bindAddr)

	// Copy localConn.Reader to sshConn.Writer
	go func() {
		_, err = io.Copy(sshConn, localConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()

	// Copy sshConn.Reader to localConn.Writer
	go func() {
		_, err = io.Copy(localConn, sshConn)
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}()
}
