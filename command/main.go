/*
  Copyright 2018 Davide Madrisan <davide.madrisan@gmail.com>

  Licensed under the Mozilla Public License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      https://www.mozilla.org/en-US/MPL/2.0/

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package command

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/madrisan/hashicorp-vault-monitor/version"
	"github.com/mitchellh/cli"
)

// Exit (Nagios compatible) return codes constants
const (
	StateOk int = iota
	_
	StateCritical
	StateError
)

const (
	addressDefault = "https://127.0.0.1:8200"
	addressDescr   = "The address of the Vault server. " +
		"Overrides the " + api.EnvVaultAddress + " environment variable if set"

	tokenDefault = ""
	tokenDescr   = "The token to access Vault. " +
		"Overrides the " + api.EnvVaultToken + " environment variable if set"
)

// Run initializes a CLI instance and its command state engine.
func Run(args []string) int {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	verInfo := version.GetVersion()
	version := verInfo.FullVersionNumber(true)

	c := cli.NewCLI("hashicorp-vault-monitor", version)
	c.Args = args

	c.Commands = map[string]cli.CommandFactory{
		"get": func() (cli.Command, error) {
			return &GetCommand{
				BaseCommand: &BaseCommand{
					Ui: &cli.ColoredUi{
						Ui:          ui,
						ErrorColor:  cli.UiColorRed,
						OutputColor: cli.UiColorGreen,
						WarnColor:   cli.UiColorYellow,
					},
				},
			}, nil
		},
		"policies": func() (cli.Command, error) {
			return &PoliciesCommand{
				BaseCommand: &BaseCommand{
					Ui: &cli.ColoredUi{
						Ui:          ui,
						ErrorColor:  cli.UiColorRed,
						OutputColor: cli.UiColorGreen,
						WarnColor:   cli.UiColorYellow,
					},
				},
			}, nil
		},
		"status": func() (cli.Command, error) {
			return &StatusCommand{
				BaseCommand: &BaseCommand{
					Ui: &cli.ColoredUi{
						Ui:          ui,
						ErrorColor:  cli.UiColorRed,
						OutputColor: cli.UiColorGreen,
						WarnColor:   cli.UiColorYellow,
					},
				},
			}, nil
		},
	}

	exitStatus, err := c.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return StateError
	}

	return exitStatus
}