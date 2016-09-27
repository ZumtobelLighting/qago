
//# Change /etc/ssh/sshd_config: on the target server to enable password authentication
//#PasswordAuthentication yes

package main

import (
	"fmt"


	"os"
	"bytes"
)
import "golang.org/x/crypto/ssh"


func main() {
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: "dlmaint",
		Auth: []ssh.AuthMethod{
			ssh.Password(os.Getenv("PASSWORD")),
		},
	}
	client, err := ssh.Dial("tcp", "10.1.30.126:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("cd ../..;ls; whoami; "); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
