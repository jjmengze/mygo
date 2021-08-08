package app

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"github.com/jjmengze/mygo/cmd/server/app/config"
	"github.com/jjmengze/mygo/pkg/server"
	"sigs.k8s.io/yaml"
)

type Options struct {
	// ComponentConfig is the  server's configuration object.
	ComponentConfig config.Config

	// ConfigFile is the location of the proxy server's configuration file.
	ConfigFile string
	// WriteConfigTo is the path where the default configuration will be written.
	WriteConfigTo string
	// CleanupAndExit, when true, makes the proxy server clean up iptables and ipvs rules, then exit.
	CleanupAndExit bool
	// WindowsService should be set to true if kube-proxy is running as a service on Windows.
	// Its corresponding flag only gets registered in Windows builds
	WindowsService bool
	// config is the proxy server's configuration object.
	InsecureServing *server.Config
	// errCh is the channel that errors will be sent
	errCh chan error

	// The fields below here are placeholders for flags that can't be directly mapped into
	// config.KubeProxyConfiguration.
	//
	// TODO remove these fields once the deprecated flags are removed.

	// master is used to override the kubeconfig's URL to the apiserver.
	master string
	// healthzPort is the port to be used by the healthz server.
	healthzPort int32
	// metricsPort is the port to be used by the metrics server.
	metricsPort int32

	// hostnameOverride, if set from the command line flag, takes precedence over the `HostnameOverride` value from the config file
	hostnameOverride string
}

// NewOptions returns initialized Options
func NewOptions() *Options {

	return &Options{
		InsecureServing: new(server.Config),
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

	o.InsecureServing.AddFlags(miscFlagSet)
	fs.AddFlagSet(miscFlagSet)

	flagSet := pflag.NewFlagSet("insecure serving", pflag.ExitOnError)
	flagSet.SetNormalizeFunc(pflag.CommandLine.GetNormalizeFunc())

	o.InsecureServing.AddFlags(flagSet)
	fs.AddFlagSet(flagSet)
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
