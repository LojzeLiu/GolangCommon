package Common

import (
	"fmt"
	"io"
	"testing"
)

func TestGetConf(t *testing.T) {
	var conf Configer
	fileName := "./test.conf"
	confGroup := "TEST_CONF"
	if err := conf.Init(fileName); err != nil {
		t.Error(err)
	}
	defer conf.Destroy()
	confs, err := conf.GetConf(confGroup)
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	fmt.Println(confs)
}
