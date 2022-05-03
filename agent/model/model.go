package model

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"goOrigin/pkg/utils"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

var Pool TaskPool

type TaskPool struct {
	Tasks   []Task
	Lock    sync.RWMutex
	TaskMap sync.Map
}

func (tp *TaskPool) Reg(j Task) {
	tp.Tasks = append(tp.Tasks, j)
}

type Task interface {
	Exec() error
}

type ShellTask struct {
	Content string
	Path    string
	Timeout string
}

type SyncTask struct {
	Url     string
	Content string
	Timeout int
	Ctx     context.Context
}

type AsyncTask struct {
	ID      string
	Url     string
	Content string
	Timeout int
	Ctx     context.Context
	Status  string
	Writer *bufio.Writer
	FilePath string
	Stdout io.ReadCloser
}

type TaskResult struct {
	Err    error
	Result []byte
}

func (t *SyncTask) ExecSingle(ctx context.Context) (res []byte, err error) {
	var (
		ch = make(chan *TaskResult)
	)

	go func() {
		var r = &TaskResult{}
		r.Result, r.Err = exec.Command("/bin/bash", utils.GetFilePath("template/test.sh")).CombinedOutput()
		ch <- r
	}()
	select {
	case <-ctx.Done():
		return nil, errors.New("canceled")
	case <-time.After(time.Duration(t.Timeout) * time.Second):
		return nil, errors.New("timeout")
	case result, ok := <-ch:
		if ok {
			return result.Result, result.Err
		}
		return nil, errors.New("unknown error")
	}

}

func (t *AsyncTask) Exec(ctx context.Context) (res []byte, err error) {
	var (
		//ch       = make(chan *TaskResult)
		waitPool = time.NewTicker(1 * time.Second)
	)
wait:
	for {
		select {
		case <-waitPool.C:
			if t.Reg() {
				break wait
			}
		}
	}
	var (
		fileObj, fileErr = os.OpenFile("test", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		outWriter        = bufio.NewWriterSize(fileObj, 1)
		errStderr        error
	)

	if fileErr != nil {
		return nil, err
	}
	defer func() {
		_ = outWriter.Flush()
		_ = fileObj.Close()
	}()
	cmd := exec.Command("/bin/bash", utils.GetFilePath("template/test.sh"))
	stdoutIn, _ := cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
		logrus.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStderr = io.Copy(outWriter, stdoutIn)

	}()
	err = cmd.Wait()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStderr != nil {
		logrus.Fatal("failed to capture stdout or stderr\n")
	}

	return nil, nil
}

func (t *AsyncTask) Exec2(ctx context.Context)  (res []byte, err error) {

	var (
		fileObj, fileErr = os.OpenFile("test", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		outWriter        = bufio.NewWriterSize(fileObj, 1)
		errStderr        error
	)
	t.Reg()

	if fileErr != nil {
		return nil, err
	}
	defer func() {
		_ = outWriter.Flush()
		_ = fileObj.Close()
	}()
	cmd := exec.Command("/bin/bash", utils.GetFilePath("template/test.sh"))
	t.Stdout, _ = cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
		logrus.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStderr = io.Copy(outWriter, t.Stdout)
	}()
	err = cmd.Wait()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStderr != nil {
		logrus.Fatal("failed to capture stdout or stderr\n")
	}

	return nil, nil
}

func Query(id string)  {
	res, ok := Pool.TaskMap.Load(id)
	if ! ok{
		fmt.Println("任务已完成")
	}
	buf := make([]byte, 62)
	t,ok := res.(*AsyncTask)
	t.FilePath = utils.GetFilePath("agent/model/test")
	fileObj,err := os.Open(t.FilePath)
	if err != nil {
		fmt.Printf("read file %s", err)
	}
	defer func() {_ = fileObj.Close()}()
	stat, err := os.Stat(t.FilePath)
	if err != nil {
		return
	}
	start := stat.Size() - 62
	_ ,err = fileObj.ReadAt(buf, start)
	if err != nil {
		fmt.Println()
	}
	fmt.Println(string(buf))


}
func (t *AsyncTask) Reg() bool {
	Pool.TaskMap.Store(t.ID, t)
	return true
}

func Regis(t Task) (ok bool, err error) {
	if Pool.GetLen() > 1000 {
		return false, nil
	}
	Pool.Lock.RLock()
	defer Pool.Lock.RUnlock()
	Pool.Tasks = append(Pool.Tasks, t)
	return true, nil
}

func (tp *TaskPool) GetLen() int {
	return len(tp.Tasks)
}

func init() {
	Pool = TaskPool{
		Tasks: nil,
		Lock:  sync.RWMutex{},
	}
}
