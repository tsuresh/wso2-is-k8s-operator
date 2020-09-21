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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Wso2IsSpec defines the desired state of Wso2Is
type Wso2IsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Wso2Is. Edit Wso2Is_types.go to remove/update
	Size                int32          `json:"replicas"`
	Configurations      Configurations `json:"configurations"`
	ContainerVersion    string         `json:"containerVersion"`
	InitialDelaySeconds int32          `json:"initialDelaySeconds"`
	PeriodSeconds       int32          `json:"periodSeconds"`
}

type Configurations struct {
	Server     Server     `json:"server"`
	SuperAdmin SuperAdmin `json:"super_admin"`
	UserStore  UserStore  `json:"user_store"`
	Database   Database   `json:"database"`
	Keystore   Keystore   `json:"keystore"`
}
type Server struct {
	Hostname string `json:"hostname"`
	NodeIP   string `json:"node_ip"`
}
type SuperAdmin struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	CreateAdminAccount bool   `json:"create_admin_account"`
}
type UserStore struct {
	Type               string `json:"type"`
	ConnectionURL      string `json:"connection_url"`
	ConnectionName     string `json:"connection_name"`
	ConnectionPassword string `json:"connection_password"`
	BaseDn             string `json:"base_dn"`
}
type IdentityDb struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type SharedDb struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Database struct {
	IdentityDb IdentityDb `json:"identity_db"`
	SharedDb   SharedDb   `json:"shared_db"`
}
type Primary struct {
	FileName string `json:"file_name"`
	Password string `json:"password"`
}
type Keystore struct {
	Primary Primary `json:"primary"`
}

// Wso2IsStatus defines the observed state of Wso2Is
type Wso2IsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Wso2Is is the Schema for the wso2is API
type Wso2Is struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wso2IsSpec   `json:"spec,omitempty"`
	Status Wso2IsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wso2IsList contains a list of Wso2Is
type Wso2IsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wso2Is `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Wso2Is{}, &Wso2IsList{})
}
