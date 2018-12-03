package Common

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var gDebugtrace *Debugtrace

type LOG_LEVE int32

const LogBuffLen int = 1024

const (
	FATAL_LEVE LOG_LEVE = iota
	ERROR_LEVE
	WARN_LEVE
	DEBUG_LEVE
)

func SetLogger(path, appName string, OutLevel LOG_LEVE) error {
	if gDebugtrace != nil {
		return errors.New("Error: Repeated initialization the logger.")
	}

	gDebugtrace = &Debugtrace{}
	return gDebugtrace.Init(path, appName, OutLevel)
}

type Debugtrace struct {
	log.Logger
	mLogBuf       *bufio.Writer
	mLogFD        *os.File
	mLastBuildDay int
	mLogLevel     LOG_LEVE
	mLogPath      string
	mLogAppName   string
}

func (d *Debugtrace) Init(path, appName string, level LOG_LEVE) error {
	d.mLogLevel = level
	d.mLogPath = path
	d.mLogAppName = appName
	if err := d.buildLogFile(); err != nil {
		return nil
	}

	go d.ProceLog()
	return nil
}

func (d *Debugtrace) ProceLog() {
	for {
		d.mLogBuf.Flush()
		d.rebuildLogFile()
		time.Sleep(1000 * time.Millisecond * 10)
	}
}

func (d *Debugtrace) rebuildLogFile() {
	Today := time.Now().Day()
	if d.mLastBuildDay != Today {
		fmt.Println("Rebuild log fileing. Todat:", Today, "; last day:", d.mLastBuildDay)
		if d.mLastBuildDay > 0 {
			if err := d.buildLogFile(); err != nil {
				d.Println("Rebuild log file failed, Reason: ", err)
				return
			}
		}
		d.mLastBuildDay = Today
	}
}

func (d *Debugtrace) buildLogFile() error {
	if d.mLogFD != nil {
		d.mLogFD.Close()
		d.mLogFD = nil
	}
	logName := d.createLogFile(d.mLogPath, d.mLogAppName)
	err := os.MkdirAll(d.mLogPath, os.ModePerm)
	if err != nil {
		return err
	}
	d.mLogFD, err = os.Create(logName)
	if err != nil {
		return err
	}

	d.mLogBuf = bufio.NewWriterSize(d.mLogFD, LogBuffLen)
	d.SetOutput(d.mLogBuf)
	d.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	d.SetPrefix("logger:")
	return nil
}

func (d *Debugtrace) createLogFile(path, appName string) string {
	currTime := time.Now()
	return fmt.Sprintf("%s/%s-%d-%d%02d%02d%02d%02d%02d.log", path, appName, os.Getegid(), currTime.Year(), currTime.Month(),
		currTime.Day(), currTime.Hour(), currTime.Minute(), currTime.Second())
}
