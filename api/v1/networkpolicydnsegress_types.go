/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DomainNamePorts is a set of domain names + ports
type DomainNamePorts struct {
	DomainName string `json:"domainName,omitempty"`
	Ports      []int  `json:"ports,omitempty"`
}

// ResolutionSet is a set of networkPolicy + domain names to resolve to IPs
type ResolutionSet struct {
	PolicyName          string            `json:"policyName,omitempty"`
	DomainNamePortsSets []DomainNamePorts `json:"domainNamePortsSets,omitempty"`
}

// NetworkPolicyDNSEgressSpec defines the desired state of NetworkPolicyDNSEgress
type NetworkPolicyDNSEgressSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	ResolutionSets []ResolutionSet `json:"resolutionSets,omitempty"`
}

// PolicyStatus defines that update status of a networkPolicy
type PolicyStatus struct {
	PolicyName string `json:"policyName,omitempty"`
	Updated    bool   `json:"updated,omitempty"`
}

// NetworkPolicyDNSEgressStatus defines the observed state of NetworkPolicyDNSEgress
type NetworkPolicyDNSEgressStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []PolicyStatus `json:"policyStatus,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPolicyDNSEgress is the Schema for the networkpolicydnsegresses API
type NetworkPolicyDNSEgress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkPolicyDNSEgressSpec   `json:"spec,omitempty"`
	Status NetworkPolicyDNSEgressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkPolicyDNSEgressList contains a list of NetworkPolicyDNSEgress
type NetworkPolicyDNSEgressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkPolicyDNSEgress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkPolicyDNSEgress{}, &NetworkPolicyDNSEgressList{})
}
