package model

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type BaseJob struct {
	FileUrl string
	Type    string
}

func (b *BaseJob) SyncRun() (res []byte, err error) {
	var (
		cmder      *exec.Cmd
		scanner    = bufio.NewScanner(os.Stdout)
		errScanner = bufio.NewScanner(os.Stderr)
	)
	switch b.Type {

	default:
		cmder = exec.Command("bin/bash", b.FileUrl)
	}
	cmder.Stdout = os.Stdout
	err = cmder.Start()
	if err != nil {
		goto ERR
	}
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	go func() {
		for errScanner.Scan() {
			fmt.Println(errScanner.Text())
		}
	}()
	err = cmder.Wait()
	if err != nil {
		goto ERR
	}
ERR:
	{
		log.Printf("任务执行失败")
		return nil, err
	}
}
