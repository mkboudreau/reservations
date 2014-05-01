package controller

import (
	"testing"
	//"log"
	"github.com/mkboudreau/loggo"
)

var testLogger *loggo.LevelLogger = loggo.DefaultLevelLogger()

func TestServerStart(t *testing.T) {
	_, err := SetupNewListenerOnPort(56999)
	if err != nil {
		t.Error(err)
	}
}
func TestServerStartTwice(t *testing.T) {
	_, err := SetupNewListenerOnPort(7777)
	if err != nil {
		t.Error(err)
	}
	_, err = SetupNewListenerOnPort(7777)
	if err == nil {
		t.Error("Should not be able to start another instance")
	}
}

func TestServerStopWithoutAStart(t *testing.T) {
	err := ShutdownListenerOnPort(8888)
	if err == nil {
		t.Error(err)
	}
}
func TestServerStopAfterStart(t *testing.T) {
	_, err := SetupNewListenerOnPort(9999)
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(9999)
	if err != nil {
		t.Error(err)
	}
}
func TestServerStartAfterRestart(t *testing.T) {
	_, err := SetupNewListenerOnPort(2222)
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(2222)
	if err != nil {
		t.Error(err)
	}
	_, err = SetupNewListenerOnPort(2222)
	if err != nil {
		t.Error(err)
	}
}

func TestTwoStartsAndTwoStopsWithErrorChecks(t *testing.T) {
	_, err := SetupNewListenerOnPort(4444)
	if err != nil {
		t.Error(err)
	}
	_, err = SetupNewListenerOnPort(4445)
	if err != nil {
		t.Error(err)
	}
	_, err = SetupNewListenerOnPort(4445)
	if err == nil {
		t.Error("Should not be able to start another instance")
	}

	err = ShutdownListenerOnPort(4444)
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(4445)
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(4445)
	if err == nil {
		t.Error("Should not be able to stop another instance")
	}
}

func TestListenServerStopAfterStart(t *testing.T) {
	server, err := SetupNewListenerOnPort(10001)
	if err != nil {
		t.Error(err)
	}
	err = server.StartListenerInBackground()
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(10001)
	if err != nil {
		t.Error(err)
	}
}
func TestListenTwoServerStopAfterStart(t *testing.T) {
	server, err := SetupNewListenerOnPort(20001)
	if err != nil {
		t.Error(err)
	}
	err = server.StartListenerInBackground()
	if err != nil {
		t.Error(err)
	}
	server, err = SetupNewListenerOnPort(20002)
	if err != nil {
		t.Error(err)
	}
	err = server.StartListenerInBackground()
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(20001)
	if err != nil {
		t.Error(err)
	}
	err = ShutdownListenerOnPort(20002)
	if err != nil {
		t.Error(err)
	}
}
