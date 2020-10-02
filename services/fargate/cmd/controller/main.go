package main

import (
	ackrt "github.com/aws/aws-controllers-k8s/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime"
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

	setupLog.Info(
		"initializing service controller",
		"aws.service", awsServiceAlias,
	)


	setupLog.Info(
		"starting manager",
		"aws.service", awsServiceAlias,
	)

}
