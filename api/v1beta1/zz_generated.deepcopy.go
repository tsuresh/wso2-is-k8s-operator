// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.Consent = in.Consent
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpsDatabase) DeepCopyInto(out *BpsDatabase) {
	*out = *in
	out.PoolOptions = in.PoolOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpsDatabase.
func (in *BpsDatabase) DeepCopy() *BpsDatabase {
	if in == nil {
		return nil
	}
	out := new(BpsDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clustering) DeepCopyInto(out *Clustering) {
	*out = *in
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clustering.
func (in *Clustering) DeepCopy() *Clustering {
	if in == nil {
		return nil
	}
	out := new(Clustering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusteringProperties) DeepCopyInto(out *ClusteringProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusteringProperties.
func (in *ClusteringProperties) DeepCopy() *ClusteringProperties {
	if in == nil {
		return nil
	}
	out := new(ClusteringProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configurations) DeepCopyInto(out *Configurations) {
	*out = *in
	out.Server = in.Server
	out.SuperAdmin = in.SuperAdmin
	out.UserStore = in.UserStore
	out.Database = in.Database
	out.Transport = in.Transport
	out.Keystore = in.Keystore
	out.Clustering = in.Clustering
	out.Monitoring = in.Monitoring
	out.Hazelcast = in.Hazelcast
	if in.SecondaryUserstore != nil {
		in, out := &in.SecondaryUserstore, &out.SecondaryUserstore
		*out = make([]SecondaryUserstore, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configurations.
func (in *Configurations) DeepCopy() *Configurations {
	if in == nil {
		return nil
	}
	out := new(Configurations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consent) DeepCopyInto(out *Consent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consent.
func (in *Consent) DeepCopy() *Consent {
	if in == nil {
		return nil
	}
	out := new(Consent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.IdentityDb = in.IdentityDb
	out.SharedDb = in.SharedDb
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datasource) DeepCopyInto(out *Datasource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datasource.
func (in *Datasource) DeepCopy() *Datasource {
	if in == nil {
		return nil
	}
	out := new(Datasource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPS) DeepCopyInto(out *HTTPS) {
	*out = *in
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPS.
func (in *HTTPS) DeepCopy() *HTTPS {
	if in == nil {
		return nil
	}
	out := new(HTTPS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hazelcast) DeepCopyInto(out *Hazelcast) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hazelcast.
func (in *Hazelcast) DeepCopy() *Hazelcast {
	if in == nil {
		return nil
	}
	out := new(Hazelcast)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityDb) DeepCopyInto(out *IdentityDb) {
	*out = *in
	out.PoolOptions = in.PoolOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityDb.
func (in *IdentityDb) DeepCopy() *IdentityDb {
	if in == nil {
		return nil
	}
	out := new(IdentityDb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jmx) DeepCopyInto(out *Jmx) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jmx.
func (in *Jmx) DeepCopy() *Jmx {
	if in == nil {
		return nil
	}
	out := new(Jmx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Keystore) DeepCopyInto(out *Keystore) {
	*out = *in
	out.Primary = in.Primary
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Keystore.
func (in *Keystore) DeepCopy() *Keystore {
	if in == nil {
		return nil
	}
	out := new(Keystore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	out.Jmx = in.Jmx
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolOptions) DeepCopyInto(out *PoolOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolOptions.
func (in *PoolOptions) DeepCopy() *PoolOptions {
	if in == nil {
		return nil
	}
	out := new(PoolOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Primary) DeepCopyInto(out *Primary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Primary.
func (in *Primary) DeepCopy() *Primary {
	if in == nil {
		return nil
	}
	out := new(Primary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Properties) DeepCopyInto(out *Properties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Properties.
func (in *Properties) DeepCopy() *Properties {
	if in == nil {
		return nil
	}
	out := new(Properties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecondaryUserstore) DeepCopyInto(out *SecondaryUserstore) {
	*out = *in
	out.Properties = in.Properties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecondaryUserstore.
func (in *SecondaryUserstore) DeepCopy() *SecondaryUserstore {
	if in == nil {
		return nil
	}
	out := new(SecondaryUserstore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedDb) DeepCopyInto(out *SharedDb) {
	*out = *in
	out.PoolOptions = in.PoolOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedDb.
func (in *SharedDb) DeepCopy() *SharedDb {
	if in == nil {
		return nil
	}
	out := new(SharedDb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SuperAdmin) DeepCopyInto(out *SuperAdmin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SuperAdmin.
func (in *SuperAdmin) DeepCopy() *SuperAdmin {
	if in == nil {
		return nil
	}
	out := new(SuperAdmin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Transport) DeepCopyInto(out *Transport) {
	*out = *in
	out.HTTPS = in.HTTPS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Transport.
func (in *Transport) DeepCopy() *Transport {
	if in == nil {
		return nil
	}
	out := new(Transport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.PoolOptions = in.PoolOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStore) DeepCopyInto(out *UserStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStore.
func (in *UserStore) DeepCopy() *UserStore {
	if in == nil {
		return nil
	}
	out := new(UserStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Userstore) DeepCopyInto(out *Userstore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Userstore.
func (in *Userstore) DeepCopy() *Userstore {
	if in == nil {
		return nil
	}
	out := new(Userstore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Userstore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserstoreList) DeepCopyInto(out *UserstoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Userstore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserstoreList.
func (in *UserstoreList) DeepCopy() *UserstoreList {
	if in == nil {
		return nil
	}
	out := new(UserstoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserstoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserstoreProperties) DeepCopyInto(out *UserstoreProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserstoreProperties.
func (in *UserstoreProperties) DeepCopy() *UserstoreProperties {
	if in == nil {
		return nil
	}
	out := new(UserstoreProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserstoreSpec) DeepCopyInto(out *UserstoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserstoreSpec.
func (in *UserstoreSpec) DeepCopy() *UserstoreSpec {
	if in == nil {
		return nil
	}
	out := new(UserstoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserstoreStatus) DeepCopyInto(out *UserstoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserstoreStatus.
func (in *UserstoreStatus) DeepCopy() *UserstoreStatus {
	if in == nil {
		return nil
	}
	out := new(UserstoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wso2Is) DeepCopyInto(out *Wso2Is) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wso2Is.
func (in *Wso2Is) DeepCopy() *Wso2Is {
	if in == nil {
		return nil
	}
	out := new(Wso2Is)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Wso2Is) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wso2IsList) DeepCopyInto(out *Wso2IsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Wso2Is, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wso2IsList.
func (in *Wso2IsList) DeepCopy() *Wso2IsList {
	if in == nil {
		return nil
	}
	out := new(Wso2IsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Wso2IsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wso2IsSpec) DeepCopyInto(out *Wso2IsSpec) {
	*out = *in
	in.Configurations.DeepCopyInto(&out.Configurations)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wso2IsSpec.
func (in *Wso2IsSpec) DeepCopy() *Wso2IsSpec {
	if in == nil {
		return nil
	}
	out := new(Wso2IsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wso2IsStatus) DeepCopyInto(out *Wso2IsStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wso2IsStatus.
func (in *Wso2IsStatus) DeepCopy() *Wso2IsStatus {
	if in == nil {
		return nil
	}
	out := new(Wso2IsStatus)
	in.DeepCopyInto(out)
	return out
}
