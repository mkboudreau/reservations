package controller

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/mkboudreau/loggo"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var logger *loggo.LevelLogger = loggo.DefaultLevelLogger()
var defaultTimeout time.Time = time.Now().Add(time.Second * 10)
var portToServerMap map[int]*CommandServer = make(map[int]*CommandServer)

type CommandServer struct {
	Port        int
	Timeout     time.Time
	stopChannel chan bool
	waitgroup   *sync.WaitGroup
	running     bool
}

func (server *CommandServer) String() string {
	return fmt.Sprintf("Port: %v; Timeout: %v; Running [%v], Pointer, %p", server.Port, server.Timeout, server.running, server)
}

func newCommandServerOnPort(port int) *CommandServer {
	s := &CommandServer{
		Timeout:     defaultTimeout,
		Port:        port,
		stopChannel: make(chan bool),
		waitgroup:   new(sync.WaitGroup),
		running:     false,
	}
	s.waitgroup.Add(1)

	return s
}

func getCommandServerOnPort(port int) *CommandServer {
	return portToServerMap[port]
}
func registerCommandServer(server *CommandServer) {
	portToServerMap[server.Port] = server
	logger.Debug("Registering new command server:", server)
	logger.Trace("Currently Registered Server List:", portToServerMap)
}
func unregisterCommandServer(server *CommandServer) {
	delete(portToServerMap, server.Port)
	logger.Debug("UN-Registering new command server:", server)
	logger.Trace("Currently Registered Server List:", portToServerMap)
}

func ShutdownListenerOnPort(port int) error {
	server := getCommandServerOnPort(port)
	if server != nil {
		server.ShutdownListener()
		unregisterCommandServer(server)
		return nil
	} else {
		// log: not running
		return errors.New(fmt.Sprintf("Nothing listening on port: %v", port))
	}
}
func SetupNewListenerOnPort(port int) (*CommandServer, error) {
	server := getCommandServerOnPort(port)
	if server != nil {
		// log: already running
		return nil, errors.New(fmt.Sprintf("Already listening on port: %v", port))
	}
	server = newCommandServerOnPort(port)
	registerCommandServer(server)
	return server, nil
}
func (server *CommandServer) ShutdownListener() {
	if server.running {
		close(server.stopChannel) // should trigger all loops to finish
		server.waitgroup.Wait()   // waits until each level is able to finish
		server.running = false
	}
}
func (server *CommandServer) StartListenerInBackground() (err error) {
	go server.StartListener()
	return nil
}
func (server *CommandServer) StartListener() (err error) {
	defer server.waitgroup.Done()
	var tcpAddr *net.TCPAddr
	var tcpListener *net.TCPListener
	var tcpConn *net.TCPConn

	server.running = true

	tcpAddr, err = getTCPAddrHostStringForPort(server.Port)
	if err != nil {
		return err
	}
	tcpListener, err = net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}
	for {
		select {
		case <-server.stopChannel:
			//handle stop... log ?
			tcpListener.Close()
			return
		default:
			// did not receive a stop... keep going
		}
		tcpListener.SetDeadline(server.Timeout)
		tcpConn, err = tcpListener.AcceptTCP()
		if err != nil {
			if oppErr, ok := err.(*net.OpError); ok && oppErr.Timeout() {
				continue
			}
			// handle non-timeout error...
			// log?
			// return err
		} else {
			server.waitgroup.Add(1)
			go server.handleTcpConnection(tcpConn)
		}
	}
}

func (server *CommandServer) handleTcpConnection(conn *net.TCPConn) {
	defer conn.Close()
	defer server.waitgroup.Done()

	for {
		select {
		case <-server.stopChannel:
			// disconnecting
			return
		default:
		}

		server.waitgroup.Add(1)
		readBytes := server.readBytesFromTcpConnection(conn)
		if len(readBytes) == 0 {
			return
		}
		command := decodeCommand(readBytes)
		go server.executeCommand(command)

	}
}

func (server *CommandServer) readBytesFromTcpConnection(conn *net.TCPConn) []byte {
	defer server.waitgroup.Done()

	var err error
	allBytes := make([]byte, 0)
	reader := bufio.NewReader(conn)

	for err != io.EOF {
		select {
		case <-server.stopChannel:
			// disconnecting
			return make([]byte, 0)
		default:
		}

		conn.SetDeadline(server.Timeout)
		buf := make([]byte, 4096)
		readCount := 0
		if readCount, err = reader.Read(buf); err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				return make([]byte, 0)
			}
		}

		allBytes = append(allBytes, buf[:readCount]...)
	}

	return allBytes
}
func (server *CommandServer) executeCommand(command *ControlCommand) {
	fmt.Printf("Control Command: %v", command)
}

/////
///// private utility functions
/////
func getHostStringForPort(port int) string {
	return net.JoinHostPort("localhost", strconv.Itoa(port))
}
func getTCPAddrHostStringForPort(port int) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", getHostStringForPort(port))
}
func decodeCommand(gobBytes []byte) *ControlCommand {
	gobReader := bytes.NewReader(gobBytes)
	gobDecoder := gob.NewDecoder(gobReader)
	// Decode (receive) and print the values.
	var command *ControlCommand
	err := gobDecoder.Decode(&command)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	return command
}
