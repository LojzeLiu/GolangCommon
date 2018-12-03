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
	TRACE(DEBUG_LEVE, "This is test log,Level: ", DEBUG_LEVE)
	TRACE(WARN_LEVE, "This is test log,Level: ", WARN_LEVE)
	TRACE(ERROR_LEVE, "This is test log,Level: ", ERROR_LEVE)
	time.Sleep(240)
}
