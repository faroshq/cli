package base

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kcp-dev/kcp/pkg/cliplugins/base"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	utilprint "github.com/faroshq/cli/pkg/util/print"
)

// Options contains options common to most CLI plugins, including settings for connecting to faros
type Options struct {
	*base.Options
	// Output specifies output format
	Output string

	// APIEndpoint is the URL of the faros API
	APIEndpoint string

	// APIEndpointFaros is the URL for a given faros deployment
	APIEndpointFaros string
	// APIEndpointFaros is the URL for a given ingress deployment
	APIEndpointIngress string
}

// NewOptions provides an instance of Options with default values.
func NewOptions(streams genericclioptions.IOStreams) *Options {
	return &Options{
		Options: base.NewOptions(streams),
	}
}

// BindFlags binds options fields to cmd's flagset.
func (o *Options) BindFlags(cmd *cobra.Command) {
	o.Options.BindFlags(cmd)
	cmd.Flags().StringVarP(&o.Output, "output", "o", o.Output, "output format [table,json,yaml]")

	cmd.Flags().StringVarP(&o.APIEndpoint, "endpoint", "e", "faros.sh", "Faros API's endpoints. Different services url will be generated based on this value. For example, if the value is 'faros.sh', the url for the faros service will be 'faros.faros.sh' and the url for the ingress service will be 'ingress.faros.sh'")

}

// Complete initializes ClientConfig based on Kubeconfig and KubectlOverrides.
func (o *Options) Complete() error {
	if err := o.Options.Complete(); err != nil {
		return err
	}
	if o.Output == "" {
		o.Output = utilprint.FormatTable
	}

	o.APIEndpointFaros = fmt.Sprintf("https://faros.%s", o.APIEndpoint)
	o.APIEndpointIngress = fmt.Sprintf("https://ingress.%s", o.APIEndpoint)

	switch o.Output {
	case utilprint.FormatJSON, utilprint.FormatYAML, utilprint.FormatTable, utilprint.FormatJSONStream:
		return nil
	default:
		return fmt.Errorf("invalid output format: %s", o.Output)
	}
}

// Validate validates the configured options.
func (o *Options) Validate() error {
	return nil
}
