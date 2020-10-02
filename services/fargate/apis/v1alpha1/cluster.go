package v1alpha1

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BucketSpec defines the desired state of Bucket
type ClusterSpec struct {
	ACL                        *string                    `json:"acl,omitempty"`
	//CreateBucketConfiguration  *CreateBucketConfiguration `json:"createBucketConfiguration,omitempty"`
	GrantFullControl           *string                    `json:"grantFullControl,omitempty"`
	GrantRead                  *string                    `json:"grantRead,omitempty"`
	GrantReadACP               *string                    `json:"grantReadACP,omitempty"`
	GrantWrite                 *string                    `json:"grantWrite,omitempty"`
	GrantWriteACP              *string                    `json:"grantWriteACP,omitempty"`
	Name                       *string                    `json:"name,omitempty"`
	ObjectLockEnabledForBucket *bool                      `json:"objectLockEnabledForBucket,omitempty"`
}

// BucketStatus defines the observed state of Bucket
type ClusterStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	Location   *string                  `json:"location,omitempty"`
}

// Cluster is the Schema for the Clusters API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// ClusterList contains a list of Bucket
// +kubebuilder:object:root=true
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
