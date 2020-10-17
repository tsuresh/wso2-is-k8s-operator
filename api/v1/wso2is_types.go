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
	Size           int32          `json:"replicas"`
	Configurations Configurations `json:"configurations"`
}
type Configurations struct {
	Server         Server         `json:"server" toml:"server"`
	SuperAdmin     SuperAdmin     `json:"superAdmin" toml:"super_admin"`
	UserStore      UserStore      `json:"userStore" toml:"user_store"`
	Database       Database       `json:"database" toml:"database"`
	Transport      Transport      `json:"transport" toml:"transport"`
	Datasource     []Datasource   `json:"datasource" toml:"datasource"`
	Authentication Authentication `json:"authentication" toml:"authentication"`
	Keystore       Keystore       `json:"keystore" toml:"keystore"`
	Clustering     Clustering     `json:"clustering" toml:"clustering"`
	Monitoring     Monitoring     `json:"monitoring" toml:"monitoring"`
}
type Server struct {
	Hostname string `json:"hostname" toml:"hostname"`
	NodeIP   string `json:"nodeIp" toml:"node_ip"`
}
type SuperAdmin struct {
	Username           string `json:"username" toml:"username"`
	Password           string `json:"password" toml:"password"`
	CreateAdminAccount bool   `json:"createAdminAccount" toml:"create_admin_account"`
}
type UserStore struct {
	Type string `json:"type" toml:"type"`
}
type PoolOptions struct {
	ValidationQuery string `json:"validationQuery" toml:"validationQuery"`
}
type User struct {
	URL         string      `json:"url" toml:"url"`
	Username    string      `json:"username" toml:"username"`
	Password    string      `json:"password" toml:"password"`
	Driver      string      `json:"driver" toml:"driver"`
	PoolOptions PoolOptions `json:"pool_options" toml:"pool_options"`
}
type IdentityDb struct {
	URL         string      `json:"url" toml:"url"`
	Username    string      `json:"username" toml:"username"`
	Password    string      `json:"password" toml:"password"`
	Driver      string      `json:"driver" toml:"driver"`
	PoolOptions PoolOptions `json:"pool_options"`
}
type SharedDb struct {
	URL         string      `json:"url" toml:"url"`
	Username    string      `json:"username" toml:"username"`
	Password    string      `json:"password" toml:"password"`
	Driver      string      `json:"driver" toml:"driver"`
	PoolOptions PoolOptions `json:"pool_options" toml:"pool_options"`
}
type BpsDatabase struct {
	URL         string      `json:"url" toml:"url"`
	Username    string      `json:"username" toml:"username"`
	Password    string      `json:"password" toml:"password"`
	Driver      string      `json:"driver" toml:"driver"`
	PoolOptions PoolOptions `json:"pool_options" toml:"pool_options"`
}
type Database struct {
	User        User        `json:"user" toml:"user"`
	IdentityDb  IdentityDb  `json:"identityDb" toml:"identity_db"`
	SharedDb    SharedDb    `json:"sharedDb" toml:"shared_db"`
	BpsDatabase BpsDatabase `json:"bpsDb" toml:"bps_database"`
}
type Properties struct {
	ProxyPort int `json:"proxyPort" toml:"proxyPort"`
}
type HTTPS struct {
	Properties Properties `json:"properties" toml:"properties"`
}
type Transport struct {
	HTTPS HTTPS `json:"https" toml:"https"`
}
type Datasource struct {
	ID                         string `json:"id" toml:"id"`
	URL                        string `json:"url" toml:"url"`
	Username                   string `json:"username" toml:"username"`
	Password                   string `json:"password" toml:"password"`
	Driver                     string `json:"driver" toml:"driver"`
	PoolOptionsValidationQuery string `json:"pool_options_validationQuery" toml:"pool_options.validationQuery"`
	PoolOptionsMaxActive       int    `json:"pool_options_maxActive" toml:"pool_options.maxActive"`
	PoolOptionsMaxWait         int    `json:"pool_options_maxWait" toml:"pool_options.maxWait"`
	PoolOptionsTestOnBorrow    bool   `json:"pool_options_testOnBorrow" toml:"pool_options.testOnBorrow"`
	PoolOptionsJmxEnabled      bool   `json:"pool_options_jmxEnabled" toml:"pool_options.jmxEnabled"`
}
type Consent struct {
	DataSource string `json:"data_source" toml:"data_source"`
}
type Authentication struct {
	Consent Consent `json:"consent" toml:"consent"`
}
type Primary struct {
	Name     string `json:"name" toml:"name"`
	Password string `json:"password" toml:"password"`
}
type Keystore struct {
	Primary Primary `json:"primary" toml:"primary"`
}
type Clustering struct {
	MembershipScheme                              string `json:"membership_scheme" toml:"membership_scheme"`
	Domain                                        string `json:"domain" toml:"domain"`
	PropertiesMembershipSchemeClassName           string `json:"properties_membershipSchemeClassName" toml:"properties.membershipSchemeClassName"`
	PropertiesKUBERNETESNAMESPACE                 string `json:"properties_KUBERNETES_NAMESPACE" toml:"properties.KUBERNETES_NAMESPACE"`
	PropertiesKUBERNETESSERVICES                  string `json:"properties_KUBERNETES_SERVICES" toml:"properties.KUBERNETES_SERVICES"`
	PropertiesKUBERNETESMASTERSKIPSSLVERIFICATION bool   `json:"properties_KUBERNETES_MASTER_SKIP_SSL_VERIFICATION" toml:"properties.KUBERNETES_MASTER_SKIP_SSL_VERIFICATION"`
	PropertiesUSEDNS                              bool   `json:"properties_USE_DNS" toml:"properties.USE_DNS"`
	PropertiesKUBERNETES_API_SERVER               string `json:"properties_KUBERNETES_API_SERVER" toml:"properties.KUBERNETES_API_SERVER"`
}
type Jmx struct {
	RmiServerStart bool `toml:"rmi_server_start" json:"rmi_server_start"`
}
type Monitoring struct {
	Jmx Jmx `toml:"jmx" json:"jmx"`
}

// Wso2IsStatus defines the observed state of Wso2Is
type Wso2IsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes" toml:"nodes"`
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
