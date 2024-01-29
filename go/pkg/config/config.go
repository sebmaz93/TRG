package cfg

import (
	"fmt"
	"os"
	"path"
)

type Operation = int

const (
	Print Operation = iota
	Add
	Remove
)

type Config struct {
	Args      []string
	Operation Operation
	Pwd       string
	Config    string
}

func getPwd(opts Opts) (string, error) {
	if opts.Pwd != "" {
		return opts.Pwd, nil
	}

	return os.Getwd()
	// we can opt this out since we are returning the same type from the func (string, error)
	// if err != nil {
	// 	return "", err
	// }
}

func getConfig(opts Opts) (string, error) {
	if opts.Config != "" {
		return opts.Config, nil
	}

	xdg, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(config, "projector", "projector.json"), nil
}

func getArgs(opts Opts) ([]string, error) {
	if len(opts.Args) == 0 {
		return []string{}, nil
	}

	if opts.Args[0] == "add" {
		if len(opts.Args) != 3 {
			return nil, fmt.Errorf("add reqires 2 args, but received %v", len(opts.args))
		}
		return opts.Args[1:], nil
	}

	if opts.Args[0] == "rm" {
	}
}

func getOperation(opts Opts) Operation {
	if len(opts.Args) == 0 {
		return Print
	}

	if opts.Args[0] == "add" {
		return Add
	}

	if opts.Args[0] == "rm" {
		return Remove
	}

	return Print
}

func NewConfig() (*Config, error) {
}
