package buildercli

import (
	"github.com/nikhilsbhat/renderer/version"
	"github.com/urfave/cli/v2"
)

const (
	flagEnvVars      = "environment-variable"
	flagTemplatePath = "go-template-path"
	flagTargetPath   = "target-config-path"
)

func CliApp() *cli.App {
	return &cli.App{
		Name:                 "renderer",
		Usage:                "Utility which helps in wrapping the build variables with go releaser",
		UsageText:            "renderer [flags]",
		EnableBashCompletion: true,
		HideHelp:             false,
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "version of the renderer",
				Action:  version.AppVersion,
			},
			{
				Name:    "render",
				Aliases: []string{"r"},
				Usage:   "render template with environment variables passed",
				Action:  Render,
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:     flagEnvVars,
						Usage:    "comma separated environment variables to be evaluated against the template",
						Aliases:  []string{"e"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    flagTemplatePath,
						Usage:   "path to go-releaser config template",
						EnvVars: []string{"GO_TEMPLATE_PATH"},
						Aliases: []string{"path"},
					},
					&cli.StringFlag{
						Name:    flagTargetPath,
						Usage:   "target path where go-releaser config has to be rendered",
						EnvVars: []string{"TARGET_FILE_PATH"},
						Aliases: []string{"target"},
					},
				},
			},
		},
	}
}
