package command

import (
	"fmt"
	"os"

	"my-help/src/common"

	commandCli "github.com/urfave/cli"
)

// define const keys for fast build main
const (
	FlagURL    = "url"
	FlagFile   = "file"
	FlagDir    = "dir"
	FlagString = "string"

	CommandDownload = "download"
	CommandMd5sum   = "md5sum"
	CommandUnzip    = "unzip"
	CommandTime     = "time"
)

// Name : return client name
func Name() string {
	return "my-help"
}

// UsageDesc : return usage describe
func UsageDesc() string {
	return "a help tool"
}

// VersionDesc : return version describe
func VersionDesc() string {
	return fmt.Sprintf("Version:   %s\n\t Tag:       %s\n\t BuildTime: %s\n\t GitHash:   %s",
		common.Version, common.Tag, common.BuildTime, common.GitHash)
}

// Run : start main thread
func Run() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	client := commandCli.NewApp()
	client.Name = Name()
	client.Action = MainActionFuncEntry
	client.Usage = UsageDesc()
	client.Version = VersionDesc()

	client.Commands = []commandCli.Command{
		{
			Name:    CommandDownload,
			Aliases: []string{"dl"},
			Usage:   "download file by http url",
			Action:  MainActionFuncEntry,
			Flags: []commandCli.Flag{
				commandCli.StringFlag{
					Name:  "url",
					Usage: "http url ready to download",
				},
				commandCli.StringFlag{
					Name:  "file, f",
					Usage: "specified local file to save",
				},
			},
		},
		{
			Name:    CommandMd5sum,
			Aliases: []string{"md5sum"},
			Usage:   "md5sum specified file",
			Action:  MainActionFuncEntry,
			Flags: []commandCli.Flag{
				commandCli.StringFlag{
					Name:  "file, f",
					Usage: "file to md5sum",
				},
			},
		},
		{
			Name:    CommandUnzip,
			Aliases: []string{"unzip"},
			Usage:   "unzip file to specified dir",
			Action:  MainActionFuncEntry,
			Flags: []commandCli.Flag{
				commandCli.StringFlag{
					Name:  "file, f",
					Usage: "local zip file",
				},
				commandCli.StringFlag{
					Name:  "dir",
					Usage: "directory to save unziped files",
				},
			},
		},
		{
			Name:    CommandTime,
			Aliases: []string{"time"},
			Usage:   "convert time format",
			Action:  MainActionFuncEntry,
			Flags: []commandCli.Flag{
				commandCli.StringFlag{
					Name:  "string, s",
					Usage: "string of time to convert",
				},
			},
		},
	}

	// override the version printer
	commandCli.VersionPrinter = func(c *commandCli.Context) {
		fmt.Printf("Version:   %s\nTag:       %s\nBuildTime: %s\nGitHash:   %s\n",
			common.Version, common.Tag, common.BuildTime, common.GitHash)
	}

	return client.Run(os.Args)
}
