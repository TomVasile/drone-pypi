package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	_ "github.com/joho/godotenv/autoload"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "pypi plugin"
	app.Usage = "pypi publish plugin"
	app.Action = run
	app.Version = fmt.Sprintf("0.0.%s", build)
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "repository",
			Usage:  "pypi repository URL",
			Value:  "https://upload.pypi.org/legacy/",
			EnvVar: "PLUGIN_REPOSITORY",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "pypi username",
			Value:  "guido",
			EnvVar: "PLUGIN_USERNAME,PYPI_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "pypi password",
			Value:  "secret",
			EnvVar: "PLUGIN_PASSWORD,PYPI_PASSWORD",
		},
		cli.StringFlag{
			Name:   "setupfile",
			Usage:  "relative location of setup.py file",
			Value:  "setup.py",
			EnvVar: "PLUGIN_SETUPFILE",
		},
		cli.StringSliceFlag{
			Name:   "distributions",
			Usage:  "distribution types to deploy",
			EnvVar: "PLUGIN_DISTRIBUTIONS",
		},
		cli.BoolFlag{
			Name:   "skip_build",
			Usage:  "skip build and only upload pre-build packages",
			EnvVar: "PLUGIN_SKIP_BUILD",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Repository:    c.String("repository"),
		Username:      c.String("username"),
		Password:      c.String("password"),
		SetupFile:     c.String("setupfile"),
		Distributions: c.StringSlice("distributions"),
		SkipBuild:     c.Bool("skip_build"),
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
