package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Bucket is the Schema for the Buckets API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}
