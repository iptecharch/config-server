/*
Copyright 2024 Nokia.

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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Encoding string

const (
	Encoding_UNKNOWN   Encoding = "UNKNOWN"
	Encoding_JSON      Encoding = "JSON"
	Encoding_JSON_IETF Encoding = "JSON_IETF"
	Encoding_BYTES     Encoding = "BYTES"
	Encoding_PROTO     Encoding = "PROTO"
	Encoding_ASCII     Encoding = "ASCII"
	Encoding_CONFIG    Encoding = "CONFIG"
)

func (r Encoding) String() string {
	switch r {
	case Encoding_UNKNOWN:
		return "UNKNOWN"
	case Encoding_JSON:
		return "JSON"
	case Encoding_JSON_IETF:
		return "JSON_IETF"
	case Encoding_BYTES:
		return "BYTES"
	case Encoding_PROTO:
		return "PROTO"
	case Encoding_ASCII:
		return "ASCII"
	case Encoding_CONFIG:
		return "CONFIG"
	default:
		return "UNKNOWN"
	}
}

type Protocol string

const (
	Protocol_UNKNOWN Protocol = "unknown"
	Protocol_GNMI    Protocol = "gnmi"
	Protocol_NETCONF Protocol = "netconf"
	Protocol_NOOP    Protocol = "noop"
	Protocol_NONE    Protocol = "none"
)

type CommitCandidate string

const (
	CommitCandidate_Candidate CommitCandidate = "candidate"
	CommitCandidate_Running   CommitCandidate = "running"
)

// TargetConnectionProfileSpec defines the desired state of TargetConnectionProfile
type TargetConnectionProfileSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="connectRetry is immutable"
	// +kubebuilder:default:="10s"
	ConnectRetry metav1.Duration `json:"connectRetry,omitempty" yaml:"connectRetry,omitempty" protobuf:"bytes,1,opt,name=connectRetry"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="timeout is immutable"
	// +kubebuilder:default:="10s"
	Timeout metav1.Duration `json:"timeout,omitempty" yaml:"timeout,omitempty" protobuf:"bytes,2,opt,name=timeout"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="protocol is immutable"
	// +kubebuilder:validation:Enum=unknown;gnmi;netconf;noop;
	// +kubebuilder:default:="gnmi"
	Protocol Protocol `json:"protocol" yaml:"protocol" protobuf:"bytes,3,opt,name=protocol,casttype=Protocol"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="port is immutable"
	// +kubebuilder:default:=57400
	// Port defines the port on which the scan runs
	Port uint32 `json:"port" yaml:"port" protobuf:"varint,4,opt,name=port"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="encoding is immutable"
	// +kubebuilder:validation:Enum=UNKNOWN;JSON;JSON_IETF;PROTO;
	Encoding *Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty" protobuf:"bytes,5,opt,name=encoding,casttype=Encoding"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="preferredNetconfVersion is immutable"
	// +kubebuilder:validation:Enum="1.0";"1.1";
	// +kubebuilder:default:="1.0"
	PreferredNetconfVersion *string `json:"preferredNetconfVersion,omitempty" yaml:"preferredNetconfVersion,omitempty" protobuf:"bytes,6,opt,name=preferredNetconfVersion"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="insecure is immutable"
	// +kubebuilder:default:=false
	Insecure *bool `json:"insecure,omitempty" yaml:"insecure,omitempty" protobuf:"varint,7,opt,name=insecure"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="skipVerify is immutable"
	// +kubebuilder:default:=true
	SkipVerify *bool `json:"skipVerify,omitempty" yaml:"skipVerify,omitempty" protobuf:"varint,8,opt,name=skipVerify"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="includeNS is immutable"
	// +kubebuilder:default:=false
	IncludeNS *bool `json:"includeNS,omitempty" yaml:"includeNS,omitempty" protobuf:"varint,9,opt,name=includeNS"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="operationWithNS is immutable"
	// +kubebuilder:default:=false
	OperationWithNS *bool `json:"operationWithNS,omitempty" yaml:"operationWithNS,omitempty" protobuf:"varint,10,opt,name=operationWithNS"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UseOperationRemove is immutable"
	// +kubebuilder:default:=false
	UseOperationRemove *bool `json:"useOperationRemove,omitempty" yaml:"useOperationRemove,omitempty" protobuf:"varint,11,opt,name=useOperationRemove"`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UseOperationRemove is immutable"
	// +kubebuilder:validation:Enum=candidate;running;
	// +kubebuilder:default:="candidate"
	CommitCandidate *CommitCandidate `json:"commitCandidate,omitempty" yaml:"commitCandidate,omitempty" protobuf:"bytes,12,opt,name=commitCandidate,casttype=CommitCandidate"`
}

// +kubebuilder:object:root=true
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:name="PROTOCOL",type="string",JSONPath=".spec.protocol"
// +kubebuilder:printcolumn:name="PORT",type="string",JSONPath=".spec.port"
// +kubebuilder:printcolumn:name="ENCODING",type="string",JSONPath=".spec.encoding"
// +kubebuilder:printcolumn:name="INSECURE",type="string",JSONPath=".spec.insecure"
// +kubebuilder:printcolumn:name="SKIPVERIFY",type="string",JSONPath=".spec.skipVerify"
// +kubebuilder:resource:categories={sdc,inv}
// TargetConnectionProfile is the Schema for the TargetConnectionProfile API
// +k8s:openapi-gen=true
type TargetConnectionProfile struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec TargetConnectionProfileSpec `json:"spec,omitempty" yaml:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +kubebuilder:object:root=true
// TargetConnectionProfileList contains a list of TargetConnectionProfile
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TargetConnectionProfileList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []TargetConnectionProfile `json:"items" yaml:"items" protobuf:"bytes,2,rep,name=items"`
}

func init() {
	localSchemeBuilder.Register(&TargetConnectionProfile{}, &TargetConnectionProfileList{})
}

var (
	TargetConnectionProfileKind             = reflect.TypeOf(TargetConnectionProfile{}).Name()
	TargetConnectionProfileGroupKind        = schema.GroupKind{Group: SchemeGroupVersion.Group, Kind: TargetConnectionProfileKind}.String()
	TargetConnectionProfileKindAPIVersion   = TargetKind + "." + SchemeGroupVersion.String()
	TargetConnectionProfileGroupVersionKind = SchemeGroupVersion.WithKind(TargetConnectionProfileKind)
)
