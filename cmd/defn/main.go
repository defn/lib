package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/util/errors"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/term"
	ctrlmanageropts "k8s.io/controller-manager/options"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"

	"github.com/authzed/controller-idioms/manager"
)

func main() {
	root := &cobra.Command{
		Use:     os.Args[0],
		Short:   "an operator for managing SpiceDB clusters",
		Version: "0.0.1",
	}

	root.AddCommand(NewCmdRun(RecommendedOptions()))

	var includeDeps bool
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "display operator version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("nope")
		},
	}
	versionCmd.Flags().BoolVar(&includeDeps, "include-deps", false, "include dependencies' versions")
	root.AddCommand(versionCmd)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

// Options contains the input to the run command.
type Options struct {
	ConfigFlags  *genericclioptions.ConfigFlags
	DebugFlags   *ctrlmanageropts.DebuggingOptions
	DebugAddress string

	BootstrapCRDs         bool
	BootstrapSpicedbsPath string
	OperatorConfigPath    string
}

// RecommendedOptions builds a new options config with default values
func RecommendedOptions() *Options {
	return &Options{
		ConfigFlags:  genericclioptions.NewConfigFlags(true),
		DebugFlags:   ctrlmanageropts.RecommendedDebuggingOptions(),
		DebugAddress: ":8080",
	}
}

// NewCmdRun creates a command object for "run"
func NewCmdRun(o *Options) *cobra.Command {
	f := cmdutil.NewFactory(o.ConfigFlags)

	cmd := &cobra.Command{
		Use:                   "run [flags]",
		DisableFlagsInUseLine: true,
		Short:                 "run SpiceDB operator",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := genericapiserver.SetupSignalContext()
			cmdutil.CheckErr(o.Validate())
			cmdutil.CheckErr(o.Run(ctx, f))
		},
	}

	namedFlagSets := &cliflag.NamedFlagSets{}
	bootstrapFlags := namedFlagSets.FlagSet("bootstrap")
	bootstrapFlags.BoolVar(&o.BootstrapCRDs, "crd", true, "if set, the operator will attempt to install/update the CRDs before starting up.")
	bootstrapFlags.StringVar(&o.BootstrapSpicedbsPath, "bootstrap-spicedbs", "", "set a path to a config file for spicedbs to load on start up.")
	debugFlags := namedFlagSets.FlagSet("debug")
	debugFlags.StringVar(&o.DebugAddress, "debug-address", o.DebugAddress, "address where debug information is served (/healthz, /metrics/, /debug/pprof, etc)")
	o.ConfigFlags.AddFlags(namedFlagSets.FlagSet("kubernetes"))
	o.DebugFlags.AddFlags(debugFlags)
	globalFlags := namedFlagSets.FlagSet("global")
	globalflag.AddGlobalFlags(globalFlags, cmd.Name())
	globalFlags.StringVar(&o.OperatorConfigPath, "config", "", "set a path to the operator's config file (configure registries, image tags, etc)")

	for _, f := range namedFlagSets.FlagSets {
		cmd.Flags().AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, *namedFlagSets, cols)

	return cmd
}

// Validate checks the set of flags provided by the user.
func (o *Options) Validate() error {
	return errors.NewAggregate(o.DebugFlags.Validate())
}

// Run performs the apply operation.
func (o *Options) Run(ctx context.Context, f cmdutil.Factory) error {
	restConfig, err := f.ToRESTConfig()
	if err != nil {
		return err
	}
	DisableClientRateLimits(restConfig)

	kclient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return err
	}

	eventSink := &typedcorev1.EventSinkImpl{Interface: kclient.CoreV1().Events("")}
	broadcaster := record.NewBroadcaster()

	controllers := make([]manager.Controller, 0)

	if ctx.Err() != nil {
		return ctx.Err()
	}

	mgr := manager.NewManager(o.DebugFlags.DebuggingConfiguration, o.DebugAddress, broadcaster, eventSink)

	return mgr.Start(ctx, controllers...)
}

// DisableClientRateLimits removes rate limiting against the apiserver; we
// respect priority and fairness and will back off if the server tells us to
func DisableClientRateLimits(restConfig *rest.Config) {
	restConfig.Burst = 2000
	restConfig.QPS = -1
}
