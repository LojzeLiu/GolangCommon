package Common

import (
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	var debug Debugtrace
	if err := debug.Init("./log/", "test", 5); err != nil {
		t.Error(err)
	}
}

func TestTRACE(t *testing.T) {
	SetLogger("./log/", "testapp", DEBUG_LEVE)
	DEBUG("Test logger file")
	WARN("Test logger file 2")
	ERROR("Test logger file 3")
	//FATAL("Test logger file 4")
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 10)
		DEBUG(i)
	}
	DEBUG("End.")
	time.Sleep(time.Millisecond * 1000 * 10)
}
