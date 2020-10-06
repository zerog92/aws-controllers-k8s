package cluster

import (
	"context"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ctrlrt "sigs.k8s.io/controller-runtime"
)

var (
	sdkLog           = ctrlrt.Log.WithName("sdk")
)


// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	sdkLog.Info("sdkFind")

	ko := r.ko.DeepCopy()

	sdkLog.Info("sdkFind", "ko", ko)

	return &resource{ko}, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	sdkLog.Info("sdkCreate")

	return nil, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	sdkLog.Info("sdkUpdate")

	return nil, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	sdkLog.Info("sdkDelete")
	return nil
}
