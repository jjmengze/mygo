package app

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"mygo/cmd/client/app/config"
	"sigs.k8s.io/yaml"
)

type Options struct {
	// ComponentConfig is the  server's configuration object.
	ComponentConfig config.Config
	// ConfigFile is the location of the client's configuration file.
	ConfigFile string

	// errCh is the channel that errors will be sent
	errCh chan error
}

// NewOptions returns initialized Options
func NewOptions() *Options {

	return &Options{
		//healthzPort: ports.ProxyHealthzPort,
		//metricsPort: ports.ProxyStatusPort,
		errCh: make(chan error),
	}
}

// Flags returns flags for a specific scheduler by section name
func (o *Options) Flags(fs *pflag.FlagSet) {
	miscFlagSet := pflag.NewFlagSet("misc", pflag.ExitOnError)
	miscFlagSet.SetNormalizeFunc(pflag.CommandLine.GetNormalizeFunc())
	fs.StringVar(&o.ConfigFile, "config", o.ConfigFile, `The path to the configuration file. The flags can overwrite fields in this file:`)
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	// Load the config file here in Complete, so that Validate validates the fully-resolved config.
	if len(o.ConfigFile) > 0 {
		c, err := o.loadConfigFromFile(o.ConfigFile)
		if err != nil {
			return err
		}
		o.ComponentConfig = *c
	}
	return nil
}

// loadConfigFromFile loads the contents of file and decodes it as a
// Options object.
func (o *Options) loadConfigFromFile(file string) (*config.Config, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return o.loadConfig(data)
}

// loadConfig decodes a serialized Options to the internal type.
func (o *Options) loadConfig(data []byte) (*config.Config, error) {
	c := &config.Config{}

	if err := yaml.UnmarshalStrict(data, c); err != nil {
		return nil, fmt.Errorf("couldn't decode as server config, got %s: ", err)
	}
	return c, nil

}
