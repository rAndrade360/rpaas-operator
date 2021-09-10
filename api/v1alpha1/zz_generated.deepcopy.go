// +build !ignore_autogenerated

// Copyright 2020 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "github.com/tsuru/nginx-operator/api/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocatedPort) DeepCopyInto(out *AllocatedPort) {
	*out = *in
	out.Owner = in.Owner
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocatedPort.
func (in *AllocatedPort) DeepCopy() *AllocatedPort {
	if in == nil {
		return nil
	}
	out := new(AllocatedPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedUpstream) DeepCopyInto(out *AllowedUpstream) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedUpstream.
func (in *AllowedUpstream) DeepCopy() *AllowedUpstream {
	if in == nil {
		return nil
	}
	out := new(AllowedUpstream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bind) DeepCopyInto(out *Bind) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bind.
func (in *Bind) DeepCopy() *Bind {
	if in == nil {
		return nil
	}
	out := new(Bind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSnapshotStorage) DeepCopyInto(out *CacheSnapshotStorage) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.StorageSize != nil {
		in, out := &in.StorageSize, &out.StorageSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeLabels != nil {
		in, out := &in.VolumeLabels, &out.VolumeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSnapshotStorage.
func (in *CacheSnapshotStorage) DeepCopy() *CacheSnapshotStorage {
	if in == nil {
		return nil
	}
	out := new(CacheSnapshotStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheSnapshotSyncSpec) DeepCopyInto(out *CacheSnapshotSyncSpec) {
	*out = *in
	if in.CmdPodToPVC != nil {
		in, out := &in.CmdPodToPVC, &out.CmdPodToPVC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CmdPVCToPod != nil {
		in, out := &in.CmdPVCToPod, &out.CmdPVCToPod
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheSnapshotSyncSpec.
func (in *CacheSnapshotSyncSpec) DeepCopy() *CacheSnapshotSyncSpec {
	if in == nil {
		return nil
	}
	out := new(CacheSnapshotSyncSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManager) DeepCopyInto(out *CertManager) {
	*out = *in
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManager.
func (in *CertManager) DeepCopy() *CertManager {
	if in == nil {
		return nil
	}
	out := new(CertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfig) DeepCopyInto(out *DNSConfig) {
	*out = *in
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfig.
func (in *DNSConfig) DeepCopy() *DNSConfig {
	if in == nil {
		return nil
	}
	out := new(DNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicCertificates) DeepCopyInto(out *DynamicCertificates) {
	*out = *in
	if in.CertManager != nil {
		in, out := &in.CertManager, &out.CertManager
		*out = new(CertManager)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicCertificates.
func (in *DynamicCertificates) DeepCopy() *DynamicCertificates {
	if in == nil {
		return nil
	}
	out := new(DynamicCertificates)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Location) DeepCopyInto(out *Location) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(Value)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Location.
func (in *Location) DeepCopy() *Location {
	if in == nil {
		return nil
	}
	out := new(Location)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedOwner) DeepCopyInto(out *NamespacedOwner) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedOwner.
func (in *NamespacedOwner) DeepCopy() *NamespacedOwner {
	if in == nil {
		return nil
	}
	out := new(NamespacedOwner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxConfig) DeepCopyInto(out *NginxConfig) {
	*out = *in
	if in.CacheEnabled != nil {
		in, out := &in.CacheEnabled, &out.CacheEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CacheSize != nil {
		in, out := &in.CacheSize, &out.CacheSize
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.CacheZoneSize != nil {
		in, out := &in.CacheZoneSize, &out.CacheZoneSize
		x := (*in).DeepCopy()
		*out = &x
	}
	in.CacheSnapshotStorage.DeepCopyInto(&out.CacheSnapshotStorage)
	in.CacheSnapshotSync.DeepCopyInto(&out.CacheSnapshotSync)
	if in.VTSEnabled != nil {
		in, out := &in.VTSEnabled, &out.VTSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyslogEnabled != nil {
		in, out := &in.SyslogEnabled, &out.SyslogEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxConfig.
func (in *NginxConfig) DeepCopy() *NginxConfig {
	if in == nil {
		return nil
	}
	out := new(NginxConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasFlavor) DeepCopyInto(out *RpaasFlavor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasFlavor.
func (in *RpaasFlavor) DeepCopy() *RpaasFlavor {
	if in == nil {
		return nil
	}
	out := new(RpaasFlavor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasFlavor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasFlavorList) DeepCopyInto(out *RpaasFlavorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RpaasFlavor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasFlavorList.
func (in *RpaasFlavorList) DeepCopy() *RpaasFlavorList {
	if in == nil {
		return nil
	}
	out := new(RpaasFlavorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasFlavorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasFlavorSpec) DeepCopyInto(out *RpaasFlavorSpec) {
	*out = *in
	if in.InstanceTemplate != nil {
		in, out := &in.InstanceTemplate, &out.InstanceTemplate
		*out = new(RpaasInstanceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasFlavorSpec.
func (in *RpaasFlavorSpec) DeepCopy() *RpaasFlavorSpec {
	if in == nil {
		return nil
	}
	out := new(RpaasFlavorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasInstance) DeepCopyInto(out *RpaasInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasInstance.
func (in *RpaasInstance) DeepCopy() *RpaasInstance {
	if in == nil {
		return nil
	}
	out := new(RpaasInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasInstanceAutoscaleSpec) DeepCopyInto(out *RpaasInstanceAutoscaleSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.TargetCPUUtilizationPercentage != nil {
		in, out := &in.TargetCPUUtilizationPercentage, &out.TargetCPUUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
	if in.TargetMemoryUtilizationPercentage != nil {
		in, out := &in.TargetMemoryUtilizationPercentage, &out.TargetMemoryUtilizationPercentage
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasInstanceAutoscaleSpec.
func (in *RpaasInstanceAutoscaleSpec) DeepCopy() *RpaasInstanceAutoscaleSpec {
	if in == nil {
		return nil
	}
	out := new(RpaasInstanceAutoscaleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasInstanceList) DeepCopyInto(out *RpaasInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RpaasInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasInstanceList.
func (in *RpaasInstanceList) DeepCopy() *RpaasInstanceList {
	if in == nil {
		return nil
	}
	out := new(RpaasInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasInstanceSpec) DeepCopyInto(out *RpaasInstanceSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Flavors != nil {
		in, out := &in.Flavors, &out.Flavors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PlanTemplate != nil {
		in, out := &in.PlanTemplate, &out.PlanTemplate
		*out = new(RpaasPlanSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Binds != nil {
		in, out := &in.Binds, &out.Binds
		*out = make([]Bind, len(*in))
		copy(*out, *in)
	}
	if in.Blocks != nil {
		in, out := &in.Blocks, &out.Blocks
		*out = make(map[BlockType]Value, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]Location, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = new(DNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = make([]apiv1alpha1.NginxTLS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(apiv1alpha1.NginxService)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraFiles != nil {
		in, out := &in.ExtraFiles, &out.ExtraFiles
		*out = new(apiv1alpha1.FilesRef)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigHistoryLimit != nil {
		in, out := &in.ConfigHistoryLimit, &out.ConfigHistoryLimit
		*out = new(int)
		**out = **in
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
	if in.AllocateContainerPorts != nil {
		in, out := &in.AllocateContainerPorts, &out.AllocateContainerPorts
		*out = new(bool)
		**out = **in
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(RpaasInstanceAutoscaleSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Lifecycle != nil {
		in, out := &in.Lifecycle, &out.Lifecycle
		*out = new(apiv1alpha1.NginxLifecycle)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSSessionResumption != nil {
		in, out := &in.TLSSessionResumption, &out.TLSSessionResumption
		*out = new(TLSSessionResumption)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowedUpstreams != nil {
		in, out := &in.AllowedUpstreams, &out.AllowedUpstreams
		*out = make([]AllowedUpstream, len(*in))
		copy(*out, *in)
	}
	if in.DynamicCertificates != nil {
		in, out := &in.DynamicCertificates, &out.DynamicCertificates
		*out = new(DynamicCertificates)
		(*in).DeepCopyInto(*out)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(apiv1alpha1.NginxIngress)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasInstanceSpec.
func (in *RpaasInstanceSpec) DeepCopy() *RpaasInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(RpaasInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasInstanceStatus) DeepCopyInto(out *RpaasInstanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasInstanceStatus.
func (in *RpaasInstanceStatus) DeepCopy() *RpaasInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(RpaasInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPlan) DeepCopyInto(out *RpaasPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPlan.
func (in *RpaasPlan) DeepCopy() *RpaasPlan {
	if in == nil {
		return nil
	}
	out := new(RpaasPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPlanList) DeepCopyInto(out *RpaasPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RpaasPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPlanList.
func (in *RpaasPlanList) DeepCopy() *RpaasPlanList {
	if in == nil {
		return nil
	}
	out := new(RpaasPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPlanSpec) DeepCopyInto(out *RpaasPlanSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(Value)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPlanSpec.
func (in *RpaasPlanSpec) DeepCopy() *RpaasPlanSpec {
	if in == nil {
		return nil
	}
	out := new(RpaasPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPortAllocation) DeepCopyInto(out *RpaasPortAllocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPortAllocation.
func (in *RpaasPortAllocation) DeepCopy() *RpaasPortAllocation {
	if in == nil {
		return nil
	}
	out := new(RpaasPortAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasPortAllocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPortAllocationList) DeepCopyInto(out *RpaasPortAllocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RpaasPortAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPortAllocationList.
func (in *RpaasPortAllocationList) DeepCopy() *RpaasPortAllocationList {
	if in == nil {
		return nil
	}
	out := new(RpaasPortAllocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RpaasPortAllocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RpaasPortAllocationSpec) DeepCopyInto(out *RpaasPortAllocationSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]AllocatedPort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RpaasPortAllocationSpec.
func (in *RpaasPortAllocationSpec) DeepCopy() *RpaasPortAllocationSpec {
	if in == nil {
		return nil
	}
	out := new(RpaasPortAllocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSessionResumption) DeepCopyInto(out *TLSSessionResumption) {
	*out = *in
	if in.SessionTicket != nil {
		in, out := &in.SessionTicket, &out.SessionTicket
		*out = new(TLSSessionTicket)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSessionResumption.
func (in *TLSSessionResumption) DeepCopy() *TLSSessionResumption {
	if in == nil {
		return nil
	}
	out := new(TLSSessionResumption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSessionTicket) DeepCopyInto(out *TLSSessionTicket) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSessionTicket.
func (in *TLSSessionTicket) DeepCopy() *TLSSessionTicket {
	if in == nil {
		return nil
	}
	out := new(TLSSessionTicket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Value) DeepCopyInto(out *Value) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ValueSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Value.
func (in *Value) DeepCopy() *Value {
	if in == nil {
		return nil
	}
	out := new(Value)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueSource) DeepCopyInto(out *ValueSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueSource.
func (in *ValueSource) DeepCopy() *ValueSource {
	if in == nil {
		return nil
	}
	out := new(ValueSource)
	in.DeepCopyInto(out)
	return out
}
