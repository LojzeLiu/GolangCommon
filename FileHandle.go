package Common

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

type Configer struct {
	mFD *os.File
}

func (c *Configer) Init(fname string) error {
	if c.mFD != nil {
		c.mFD.Close()
		c.mFD = nil
	}
	var err error
	c.mFD, err = os.Open(fname)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configer) Destroy() error {
	if c.mFD != nil {
		return c.mFD.Close()
	}
	return nil
}

func (c *Configer) GetConf(confGroup string) (confs map[string]string, err error) {
	confs = make(map[string]string)
	if c.mFD == nil {
		return nil, errors.New("Error, Not initializtion.")
	}
	rd := bufio.NewReader(c.mFD)
	regexpCG := regexp.MustCompile(`^\[[a-zA-Z0-9_]*\]$`)
	regexpNote := regexp.MustCompile(`^#.*`)

	TargetGroup := "[" + confGroup + "]"
	for {

		CurrData, prefix, err := rd.ReadLine()
		if err != nil {
			break
		}
		CurrString := string(CurrData)
		if strings.Compare(CurrString, TargetGroup) == 0 {
			for {
				//找到对应组
				CurrData, prefix, err = rd.ReadLine()
				if err != nil {
					break
				}
				CurrString = string(CurrData)
				if prefix {
				}
				//判断是否注释了该项
				if regexpNote.MatchString(CurrString) {
					continue
				}

				//判断是否到了另一组
				if regexpCG.MatchString(CurrString) {
					break
				}
				conf := strings.Split(CurrString, "=")
				if len(conf) != 2 {
					//配置错误
					continue
				}
				confs[conf[0]] = conf[1]
			}
		}
	}
	return
}
