// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha2

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSpec) DeepCopyInto(out *CommonSpec) {
	*out = *in
	in.Affinity.DeepCopyInto(&out.Affinity)
	in.Resources.DeepCopyInto(&out.Resources)
	in.ServiceTemplate.DeepCopyInto(&out.ServiceTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSpec.
func (in *CommonSpec) DeepCopy() *CommonSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonSplunkSpec) DeepCopyInto(out *CommonSplunkSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.LicenseMasterRef = in.LicenseMasterRef
	out.IndexerClusterRef = in.IndexerClusterRef
	out.SecretsRef = in.SecretsRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonSplunkSpec.
func (in *CommonSplunkSpec) DeepCopy() *CommonSplunkSpec {
	if in == nil {
		return nil
	}
	out := new(CommonSplunkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerCluster) DeepCopyInto(out *IndexerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerCluster.
func (in *IndexerCluster) DeepCopy() *IndexerCluster {
	if in == nil {
		return nil
	}
	out := new(IndexerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterList) DeepCopyInto(out *IndexerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IndexerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterList.
func (in *IndexerClusterList) DeepCopy() *IndexerClusterList {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IndexerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterMemberStatus) DeepCopyInto(out *IndexerClusterMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterMemberStatus.
func (in *IndexerClusterMemberStatus) DeepCopy() *IndexerClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterSpec) DeepCopyInto(out *IndexerClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterSpec.
func (in *IndexerClusterSpec) DeepCopy() *IndexerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IndexerClusterStatus) DeepCopyInto(out *IndexerClusterStatus) {
	*out = *in
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]IndexerClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IndexerClusterStatus.
func (in *IndexerClusterStatus) DeepCopy() *IndexerClusterStatus {
	if in == nil {
		return nil
	}
	out := new(IndexerClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMaster) DeepCopyInto(out *LicenseMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMaster.
func (in *LicenseMaster) DeepCopy() *LicenseMaster {
	if in == nil {
		return nil
	}
	out := new(LicenseMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterList) DeepCopyInto(out *LicenseMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LicenseMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterList.
func (in *LicenseMasterList) DeepCopy() *LicenseMasterList {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LicenseMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterSpec) DeepCopyInto(out *LicenseMasterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterSpec.
func (in *LicenseMasterSpec) DeepCopy() *LicenseMasterSpec {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseMasterStatus) DeepCopyInto(out *LicenseMasterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseMasterStatus.
func (in *LicenseMasterStatus) DeepCopy() *LicenseMasterStatus {
	if in == nil {
		return nil
	}
	out := new(LicenseMasterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadCluster) DeepCopyInto(out *SearchHeadCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadCluster.
func (in *SearchHeadCluster) DeepCopy() *SearchHeadCluster {
	if in == nil {
		return nil
	}
	out := new(SearchHeadCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterList) DeepCopyInto(out *SearchHeadClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SearchHeadCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterList.
func (in *SearchHeadClusterList) DeepCopy() *SearchHeadClusterList {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SearchHeadClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterMemberStatus) DeepCopyInto(out *SearchHeadClusterMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterMemberStatus.
func (in *SearchHeadClusterMemberStatus) DeepCopy() *SearchHeadClusterMemberStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterSpec) DeepCopyInto(out *SearchHeadClusterSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	out.SparkRef = in.SparkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterSpec.
func (in *SearchHeadClusterSpec) DeepCopy() *SearchHeadClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchHeadClusterStatus) DeepCopyInto(out *SearchHeadClusterStatus) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]SearchHeadClusterMemberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchHeadClusterStatus.
func (in *SearchHeadClusterStatus) DeepCopy() *SearchHeadClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SearchHeadClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spark) DeepCopyInto(out *Spark) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spark.
func (in *Spark) DeepCopy() *Spark {
	if in == nil {
		return nil
	}
	out := new(Spark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spark) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkList) DeepCopyInto(out *SparkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Spark, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkList.
func (in *SparkList) DeepCopy() *SparkList {
	if in == nil {
		return nil
	}
	out := new(SparkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SparkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkSpec) DeepCopyInto(out *SparkSpec) {
	*out = *in
	in.CommonSpec.DeepCopyInto(&out.CommonSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkSpec.
func (in *SparkSpec) DeepCopy() *SparkSpec {
	if in == nil {
		return nil
	}
	out := new(SparkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SparkStatus) DeepCopyInto(out *SparkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SparkStatus.
func (in *SparkStatus) DeepCopy() *SparkStatus {
	if in == nil {
		return nil
	}
	out := new(SparkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Standalone) DeepCopyInto(out *Standalone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Standalone.
func (in *Standalone) DeepCopy() *Standalone {
	if in == nil {
		return nil
	}
	out := new(Standalone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Standalone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneList) DeepCopyInto(out *StandaloneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Standalone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneList.
func (in *StandaloneList) DeepCopy() *StandaloneList {
	if in == nil {
		return nil
	}
	out := new(StandaloneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StandaloneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSpec) DeepCopyInto(out *StandaloneSpec) {
	*out = *in
	in.CommonSplunkSpec.DeepCopyInto(&out.CommonSplunkSpec)
	out.SparkRef = in.SparkRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSpec.
func (in *StandaloneSpec) DeepCopy() *StandaloneSpec {
	if in == nil {
		return nil
	}
	out := new(StandaloneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneStatus) DeepCopyInto(out *StandaloneStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneStatus.
func (in *StandaloneStatus) DeepCopy() *StandaloneStatus {
	if in == nil {
		return nil
	}
	out := new(StandaloneStatus)
	in.DeepCopyInto(out)
	return out
}
