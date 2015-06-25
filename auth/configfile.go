package auth

import (
	"fmt"
	"path"

	"github.com/codegangsta/cli"
	"github.com/jrperritt/rack/util"
	"gopkg.in/ini.v1"
)

func configfile(c *cli.Context, have map[string]string, need map[string]string) error {
	dir, err := util.RackDir()
	if err != nil {
		// return fmt.Errorf("Error retrieving rack directory: %s\n", err)
		return nil
	}
	f := path.Join(dir, "config")
	cfg, err := ini.Load(f)
	if err != nil {
		// return fmt.Errorf("Error loading config file: %s\n", err)
		return nil
	}
	cfg.BlockMode = false
	var profile string
	if c.GlobalIsSet("profile") {
		profile = c.GlobalString("profile")
	} else if c.IsSet("profile") {
		profile = c.String("profile")
	section, err := cfg.GetSection(profile)
	if err != nil && profile != "" {
		return fmt.Errorf("Invalid config file profile: %s\n", profile)
	}

	for opt := range need {
		if val := section.Key(opt).String(); val != "" {
			have[opt] = val
			delete(need, opt)
		}
	}
	return nil
}
