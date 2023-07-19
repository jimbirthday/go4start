// A small SSH daemon providing bash sessions
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {

	// In the latest version of crypto/ssh (after Go 1.3), the SSH server type has been removed
	// in favour of an SSH connection type. A ssh.ServerConn is created by passing an existing
	// net.Conn and a ssh.ServerConfig to ssh.NewServerConn, in effect, upgrading the net.Conn
	// into an ssh.ServerConn

	config := &ssh.ServerConfig{
		PublicKeyCallback: func(c ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			// Should use constant-time compare (or better, salt+hash) in a production setting.
			//if c.User() == "foo" && string(key) == "bar" {
			//	return nil, nil
			//}
			//return nil, fmt.Errorf("password rejected for %q", c.User())
			return nil, nil
		},
		// You may also explicitly allow anonymous client authentication, though anon bash
		// sessions may not be a wise idea
		// NoClientAuth: true,
	}

	// You can generate a keypair with 'ssh-keygen -t rsa'
	privateBytes, err := ioutil.ReadFile("/Users/jim/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("Failed to load private key (~/.ssh/id_rsa):%+v", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	config.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:2222")
	if err != nil {
		log.Fatalf("Failed to listen on 2200 (%s)", err)
	}

	// Accept all connections
	log.Print("Listening on 2222...")
	for {
		tcpConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection (%s)", err)
			continue
		}
		// Before use, a handshake must be performed on the incoming net.Conn.
		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, config)
		if err != nil {
			log.Printf("Failed to handshake (%s)", err)
			continue
		}

		log.Printf("New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())
		// Discard all global out-of-band Requests
		go ssh.DiscardRequests(reqs)
		// Accept all channels
		go handleChannels(chans)
	}
}

func handleChannels(chans <-chan ssh.NewChannel) {
	// Service the incoming Channel channel in go routine
	for newChannel := range chans {
		go handleChannel(newChannel)
	}
}

func handleChannel(newChannel ssh.NewChannel) {
	// Since we're handling a shell, we expect a
	// channel type of "session". The also describes
	// "x11", "direct-tcpip" and "forwarded-tcpip"
	// channel types.
	//log.Printf("channel info type:%s\n", newChannel.ChannelType())
	if t := newChannel.ChannelType(); t != "session" {
		newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
		return
	}

	// At this point, we have the opportunity to reject the client's
	// request for another logical connection
	channel, requests, err := newChannel.Accept()
	if err != nil {
		log.Printf("Could not accept channel (%s)", err)
		return
	}
	//defer channel.Close()

	// Sessions have out-of-band requests such as "shell", "pty-req" and "env"
	go func(in <-chan *ssh.Request) {
		for req := range in {
			//ok := false
			log.Printf("do")
			log.Printf("type:%s payload:%s", req.Type, string(req.Payload))
			switch req.Type {
			case "exec":
				//ok = true
				log.Printf("exec")
				if code, err := channel.Write([]byte("/home/op\\n")); err != nil {
					log.Printf("channel write err code:%d, err:%+v\n", code, err)
				}
				sendExitStatus(0, channel)
			case "shell":
				log.Printf("shell")
				if len(req.Payload) == 0 {
					log.Printf("shell payload:%s\n", req.Payload)
					//ok = true
				}
			case "subsystem":
				log.Printf("subsystem")
				if string(req.Payload[4:]) == "sftp" {
					//ok = true
				}
			}
			//channel.SendRequest("exit-status", false, []byte{0, 0, 0, byte(0)})
			//log.Printf("done")
		}
	}(requests)
	//serverOptions := []sftp.ServerOption{
	//	sftp.WithDebug(debugStream),
	//}
	//
	//server, err := sftp.NewServer(
	//	channel,
	//	serverOptions...,
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//go func() {
	//	if err := server.Serve(); err == io.EOF {
	//		server.Close()
	//		log.Print("sftp client exited session.")
	//	} else if err != nil {
	//		log.Fatal("sftp server completed with error:", err)
	//	}
	//}()
}

func sendExitStatus(status uint32, channel ssh.Channel) error {
	defer channel.Close()
	exit := struct{ Status uint32 }{uint32(0)}
	_, err := channel.SendRequest("exit-status", false, ssh.Marshal(exit))
	return err
}
