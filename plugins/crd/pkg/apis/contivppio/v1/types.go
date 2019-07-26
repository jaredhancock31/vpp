// Copyright (c) 2019 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomNetwork define custom network for contiv/vpp
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CustomNetwork struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the custom resource spec
	Spec CustomNetworkSpec `json:"spec"`
}

// CustomNetworkSpec is the spec for custom network configuration resource
type CustomNetworkSpec struct {
	Type                   string `json:"type"`
	SubnetCIDR             string `json:"subnetCIDR"`
	SubnetOneNodePrefixLen uint32 `json:"subnetOneNodePrefixLen"`
}

// CustomNetworkList is a list of CustomNetwork resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CustomNetworkList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []CustomNetwork `json:"items"`
}

// ExternalInterface is used to store definition of an external interface defined via CRD.
// It is a logical entity that may mean different physical interfaces on different nodes.
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExternalInterface struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the custom resource spec
	Spec ExternalInterfaceSpec `json:"spec"`
}

// ExternalInterfaceSpec is the spec for external interface configuration resource
type ExternalInterfaceSpec struct {
	Name  string          `json:"name"`
	Type  string          `json:"type"`
	Nodes []NodeInterface `json:"nodes"`
}

// NodeInterface describe config for an interface referenced by logical name on a node
type NodeInterface struct {
	Node             string `json:"node"`
	VppInterfaceName string `json:"vppInterfaceName"`
	IP               string `json:"IP"`
	VLAN             uint32 `json:"VLAN"`
}

// ExternalInterfaceList is a list of ExternalInterface resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExternalInterfaceList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []ExternalInterface `json:"items"`
}

// ServiceFunctionChain define service function chain crd for contiv/vpp
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ServiceFunctionChain struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	meta_v1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	// Spec is the custom resource spec
	Spec ServiceFunctionChainSpec `json:"spec"`
}

// ServiceFunctionChainSpec describe service function chain
type ServiceFunctionChainSpec struct {
	Network string            `json:"network"`
	Chain   []ServiceFunction `json:"chain"`
}

// ServiceFunction describes single segment of the chain
type ServiceFunction struct {
	Name            string            `json:"name"`
	Type            string            `json:"type"`
	PodSelector     map[string]string `json:"podSelector"`
	Interface       string            `json:"interface"`
	InputInterface  string            `json:"inputInterface"`
	OutputInterface string            `json:"outputInterface"`
}

// ServiceFunctionChainList is a list of ServiceFunctionChain resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ServiceFunctionChainList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []ServiceFunctionChain `json:"items"`
}
