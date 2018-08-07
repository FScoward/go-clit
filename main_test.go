package main_test

import (
	"testing"
	"github.com/urfave/cli"
	)

func Test_タスク作成(t *testing.T) {
	command := []string{"n"}
	app := cli.NewApp()
	actual := app.Run(command)

	if actual != nil {
		t.Errorf("actual %v\n want %v", actual, nil)
	}
}