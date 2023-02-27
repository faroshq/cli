package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"

	logincmd "github.com/faroshq/cli/pkg/login/cmd"
	organizationcmd "github.com/faroshq/faros/pkg/cliplugins/organization/cmd"
	workspacecmd "github.com/faroshq/faros/pkg/cliplugins/workspace/cmd"

	connectioncmd "github.com/faroshq/faros-ingress/pkg/cliplugins/connection/cmd"
	exposecmd "github.com/faroshq/faros-ingress/pkg/cliplugins/expose/cmd"
)

// New returns a cobra.Command for faros actions.
func New(streams genericclioptions.IOStreams) (*cobra.Command, error) {
	// generic login command
	loginCmd, err := logincmd.New(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// faros specific commands
	organizationCmd, err := organizationcmd.New(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	workspaceCmd, err := workspacecmd.New(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// faros ingress specific commands
	connectionCmd, err := connectioncmd.New(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	exposeCmd, err := exposecmd.New(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	cmd := &cobra.Command{
		Use:   "faros",
		Short: "Manage faros",
	}

	cmd.AddCommand(loginCmd)

	cmd.AddCommand(organizationCmd)
	cmd.AddCommand(workspaceCmd)

	cmd.AddCommand(connectionCmd)
	cmd.AddCommand(exposeCmd)

	return cmd, nil
}
