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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type StorageType string

const (
	DefaultReaperImage = "thelastpickle/cassandra-reaper:2.0.5"

	StorageTypeMemory    = StorageType("memory")
	StorageTypeCassandra = StorageType("cassandra")

	DefaultKeyspace    = "reaper_db"
	DefaultStorageType = StorageTypeMemory
)

type ServerConfig struct {
	StorageType StorageType `json:"storageType,omitempty"`

	CassandraBackend *CassandraBackend `json:"cassandraBackend,omitempty" yaml:"cassandra,omitempty"`
}

// Specifies the replication strategy for a keyspace
type ReplicationConfig struct {
	// Specifies the replication_factor when SimpleStrategy is used
	SimpleStrategy *int32 `json:"simpleStrategy,omitempty"`

	// Specifies the replication_factor when NetworkTopologyStrategy is used. The mapping is data center name to RF.
	NetworkTopologyStrategy *map[string]int32 `json:"networkTopologyStrategy,omitempty"`
}

type AuthProvider struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

type CassandraBackend struct {
	ClusterName string `json:"clusterName" yaml:"clusterName"`

	// The headless service that provides endpoints for the StorageTypeCassandra pods
	CassandraService string `json:"cassandraService" yaml:"contactPoints"`

	// Defaults to reaper
	Keyspace string `json:"keyspace,omitempty" yaml:"keyspace,omitempty"`

	Replication ReplicationConfig `json:"replication" yaml:"-"`

	AuthProvider AuthProvider `json:"authProvider,omitempty" yaml:"authProvider,omitempty"`
}

// ReaperSpec defines the desired state of Reaper
type ReaperSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Image string `json:"image,omitempty"`

	ServerConfig ServerConfig `json:"serverConfig,omitempty" yaml:"serverConfig,omitempty"`
}

// ReaperStatus defines the observed state of Reaper
type ReaperStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Reaper is the Schema for the reapers API
type Reaper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReaperSpec   `json:"spec,omitempty"`
	Status ReaperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReaperList contains a list of Reaper
type ReaperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reaper `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Reaper{}, &ReaperList{})
}