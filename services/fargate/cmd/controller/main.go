package main

import (
	ackrt "github.com/aws/aws-controllers-k8s/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	ctrlrt "sigs.k8s.io/controller-runtime"
	flag "github.com/spf13/pflag"
)

var (
	awsServiceAPIGroup = "fargate.services.k8s.aws"
	awsServiceAlias    = "fargate"
	scheme             = runtime.NewScheme()
	setupLog           = ctrlrt.Log.WithName("setup")
)

func main() {

	var ackCfg ackrt.Config
	ackCfg.BindFlags()
	flag.Parse()
	ackCfg.SetupLogger()

	if err := ackCfg.Validate(); err != nil {
		setupLog.Error(
			err, "Unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	mgr, err := ctrlrt.NewManager(ctrlrt.GetConfigOrDie(), ctrlrt.Options{
		Scheme:             scheme,
		Port:               ackCfg.BindPort,
		MetricsBindAddress: ackCfg.MetricsAddr,
		LeaderElection:     ackCfg.EnableLeaderElection,
		LeaderElectionID:   awsServiceAPIGroup,
	})
	if err != nil {
		setupLog.Error(
			err, "unable to create controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

	stopChan := ctrlrt.SetupSignalHandler()

	setupLog.Info(
		"initializing service controller",
		"aws.service", awsServiceAlias,
		"mgr", mgr,
	)


	setupLog.Info(
		"starting manager",
		"aws.service", awsServiceAlias,
	)
	if err := mgr.Start(stopChan); err != nil {
		setupLog.Error(
			err, "unable to start controller manager",
			"aws.service", awsServiceAlias,
		)
		os.Exit(1)
	}

}
