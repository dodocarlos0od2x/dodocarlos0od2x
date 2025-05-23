//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceAction) DeepCopyInto(out *AvoidanceAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceAction.
func (in *AvoidanceAction) DeepCopy() *AvoidanceAction {
	if in == nil {
		return nil
	}
	out := new(AvoidanceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoidanceAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionList) DeepCopyInto(out *AvoidanceActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AvoidanceAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionList.
func (in *AvoidanceActionList) DeepCopy() *AvoidanceActionList {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AvoidanceActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionSpec) DeepCopyInto(out *AvoidanceActionSpec) {
	*out = *in
	if in.Throttle != nil {
		in, out := &in.Throttle, &out.Throttle
		*out = new(ThrottleAction)
		**out = **in
	}
	if in.Eviction != nil {
		in, out := &in.Eviction, &out.Eviction
		*out = new(EvictionAction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionSpec.
func (in *AvoidanceActionSpec) DeepCopy() *AvoidanceActionSpec {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvoidanceActionStatus) DeepCopyInto(out *AvoidanceActionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvoidanceActionStatus.
func (in *AvoidanceActionStatus) DeepCopy() *AvoidanceActionStatus {
	if in == nil {
		return nil
	}
	out := new(AvoidanceActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUBurst) DeepCopyInto(out *CPUBurst) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUBurst.
func (in *CPUBurst) DeepCopy() *CPUBurst {
	if in == nil {
		return nil
	}
	out := new(CPUBurst)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUQOS) DeepCopyInto(out *CPUQOS) {
	*out = *in
	if in.CPUPriority != nil {
		in, out := &in.CPUPriority, &out.CPUPriority
		*out = new(int32)
		**out = **in
	}
	if in.ContainerPriority != nil {
		in, out := &in.ContainerPriority, &out.ContainerPriority
		*out = make(map[string]*int32, len(*in))
		for key, val := range *in {
			var outVal *int32
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(int32)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	out.CPUBurst = in.CPUBurst
	out.HtIsolation = in.HtIsolation
	out.CPUSet = in.CPUSet
	in.RDT.DeepCopyInto(&out.RDT)
	if in.ContainerRDT != nil {
		in, out := &in.ContainerRDT, &out.ContainerRDT
		*out = make(map[string]RDT, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUQOS.
func (in *CPUQOS) DeepCopy() *CPUQOS {
	if in == nil {
		return nil
	}
	out := new(CPUQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUSet) DeepCopyInto(out *CPUSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUSet.
func (in *CPUSet) DeepCopy() *CPUSet {
	if in == nil {
		return nil
	}
	out := new(CPUSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUThrottle) DeepCopyInto(out *CPUThrottle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUThrottle.
func (in *CPUThrottle) DeepCopy() *CPUThrottle {
	if in == nil {
		return nil
	}
	out := new(CPUThrottle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DevNetIOLimits) DeepCopyInto(out *DevNetIOLimits) {
	{
		in := &in
		*out = make(DevNetIOLimits, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevNetIOLimits.
func (in DevNetIOLimits) DeepCopy() DevNetIOLimits {
	if in == nil {
		return nil
	}
	out := new(DevNetIOLimits)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOLimit) DeepCopyInto(out *DiskIOLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOLimit.
func (in *DiskIOLimit) DeepCopy() *DiskIOLimit {
	if in == nil {
		return nil
	}
	out := new(DiskIOLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOQOS) DeepCopyInto(out *DiskIOQOS) {
	*out = *in
	out.DiskIOWeight = in.DiskIOWeight
	out.DiskIOLimit = in.DiskIOLimit
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOQOS.
func (in *DiskIOQOS) DeepCopy() *DiskIOQOS {
	if in == nil {
		return nil
	}
	out := new(DiskIOQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskIOWeight) DeepCopyInto(out *DiskIOWeight) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskIOWeight.
func (in *DiskIOWeight) DeepCopy() *DiskIOWeight {
	if in == nil {
		return nil
	}
	out := new(DiskIOWeight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticCoreCpuLimit) DeepCopyInto(out *ElasticCoreCpuLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticCoreCpuLimit.
func (in *ElasticCoreCpuLimit) DeepCopy() *ElasticCoreCpuLimit {
	if in == nil {
		return nil
	}
	out := new(ElasticCoreCpuLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticCoreCpuLimitPeriod) DeepCopyInto(out *ElasticCoreCpuLimitPeriod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticCoreCpuLimitPeriod.
func (in *ElasticCoreCpuLimitPeriod) DeepCopy() *ElasticCoreCpuLimitPeriod {
	if in == nil {
		return nil
	}
	out := new(ElasticCoreCpuLimitPeriod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticCpuAvoidance) DeepCopyInto(out *ElasticCpuAvoidance) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticCpuAvoidance.
func (in *ElasticCpuAvoidance) DeepCopy() *ElasticCpuAvoidance {
	if in == nil {
		return nil
	}
	out := new(ElasticCpuAvoidance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticCpuLimit) DeepCopyInto(out *ElasticCpuLimit) {
	*out = *in
	out.ElasticNodeCpuLimit = in.ElasticNodeCpuLimit
	if in.ElasticCoreCpuLimit != nil {
		in, out := &in.ElasticCoreCpuLimit, &out.ElasticCoreCpuLimit
		*out = make([]ElasticCoreCpuLimit, len(*in))
		copy(*out, *in)
	}
	if in.ElasticCoreCpuLimitPeriod != nil {
		in, out := &in.ElasticCoreCpuLimitPeriod, &out.ElasticCoreCpuLimitPeriod
		*out = make([]ElasticCoreCpuLimitPeriod, len(*in))
		copy(*out, *in)
	}
	out.ElasticCpuAvoidance = in.ElasticCpuAvoidance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticCpuLimit.
func (in *ElasticCpuLimit) DeepCopy() *ElasticCpuLimit {
	if in == nil {
		return nil
	}
	out := new(ElasticCpuLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticNodeCpuLimit) DeepCopyInto(out *ElasticNodeCpuLimit) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticNodeCpuLimit.
func (in *ElasticNodeCpuLimit) DeepCopy() *ElasticNodeCpuLimit {
	if in == nil {
		return nil
	}
	out := new(ElasticNodeCpuLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvictionAction) DeepCopyInto(out *EvictionAction) {
	*out = *in
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvictionAction.
func (in *EvictionAction) DeepCopy() *EvictionAction {
	if in == nil {
		return nil
	}
	out := new(EvictionAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HtIsolation) DeepCopyInto(out *HtIsolation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HtIsolation.
func (in *HtIsolation) DeepCopy() *HtIsolation {
	if in == nil {
		return nil
	}
	out := new(HtIsolation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemAsyncReclaim) DeepCopyInto(out *MemAsyncReclaim) {
	*out = *in
	if in.AsyncRatio != nil {
		in, out := &in.AsyncRatio, &out.AsyncRatio
		*out = new(int64)
		**out = **in
	}
	if in.AsyncDistanceFactor != nil {
		in, out := &in.AsyncDistanceFactor, &out.AsyncDistanceFactor
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemAsyncReclaim.
func (in *MemAsyncReclaim) DeepCopy() *MemAsyncReclaim {
	if in == nil {
		return nil
	}
	out := new(MemAsyncReclaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemLimit) DeepCopyInto(out *MemLimit) {
	*out = *in
	if in.PageCacheLimitGlobal != nil {
		in, out := &in.PageCacheLimitGlobal, &out.PageCacheLimitGlobal
		*out = new(bool)
		**out = **in
	}
	if in.PageCacheLimitRetryTimes != nil {
		in, out := &in.PageCacheLimitRetryTimes, &out.PageCacheLimitRetryTimes
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemLimit.
func (in *MemLimit) DeepCopy() *MemLimit {
	if in == nil {
		return nil
	}
	out := new(MemLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemPageCacheLimit) DeepCopyInto(out *MemPageCacheLimit) {
	*out = *in
	if in.PageCacheMaxRatio != nil {
		in, out := &in.PageCacheMaxRatio, &out.PageCacheMaxRatio
		*out = new(int64)
		**out = **in
	}
	if in.PageCacheReclaimRatio != nil {
		in, out := &in.PageCacheReclaimRatio, &out.PageCacheReclaimRatio
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemPageCacheLimit.
func (in *MemPageCacheLimit) DeepCopy() *MemPageCacheLimit {
	if in == nil {
		return nil
	}
	out := new(MemPageCacheLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemWatermark) DeepCopyInto(out *MemWatermark) {
	*out = *in
	if in.WatermarkRatio != nil {
		in, out := &in.WatermarkRatio, &out.WatermarkRatio
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemWatermark.
func (in *MemWatermark) DeepCopy() *MemWatermark {
	if in == nil {
		return nil
	}
	out := new(MemWatermark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryQOS) DeepCopyInto(out *MemoryQOS) {
	*out = *in
	in.MemAsyncReclaim.DeepCopyInto(&out.MemAsyncReclaim)
	in.MemWatermark.DeepCopyInto(&out.MemWatermark)
	in.MemPageCacheLimit.DeepCopyInto(&out.MemPageCacheLimit)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryQOS.
func (in *MemoryQOS) DeepCopy() *MemoryQOS {
	if in == nil {
		return nil
	}
	out := new(MemoryQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryThrottle) DeepCopyInto(out *MemoryThrottle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryThrottle.
func (in *MemoryThrottle) DeepCopy() *MemoryThrottle {
	if in == nil {
		return nil
	}
	out := new(MemoryThrottle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricRule) DeepCopyInto(out *MetricRule) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Value = in.Value.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricRule.
func (in *MetricRule) DeepCopy() *MetricRule {
	if in == nil {
		return nil
	}
	out := new(MetricRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetIOLimits) DeepCopyInto(out *NetIOLimits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetIOLimits.
func (in *NetIOLimits) DeepCopy() *NetIOLimits {
	if in == nil {
		return nil
	}
	out := new(NetIOLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetIOQOS) DeepCopyInto(out *NetIOQOS) {
	*out = *in
	if in.NetIOPriority != nil {
		in, out := &in.NetIOPriority, &out.NetIOPriority
		*out = new(int64)
		**out = **in
	}
	if in.ContainersPriority != nil {
		in, out := &in.ContainersPriority, &out.ContainersPriority
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.NetIOLimits = in.NetIOLimits
	if in.DevNetIOLimits != nil {
		in, out := &in.DevNetIOLimits, &out.DevNetIOLimits
		*out = make(DevNetIOLimits, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.WhitelistPorts = in.WhitelistPorts
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetIOQOS.
func (in *NetIOQOS) DeepCopy() *NetIOQOS {
	if in == nil {
		return nil
	}
	out := new(NetIOQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetLimits) DeepCopyInto(out *NetLimits) {
	*out = *in
	if in.RXBpsMin != nil {
		in, out := &in.RXBpsMin, &out.RXBpsMin
		*out = new(int64)
		**out = **in
	}
	if in.RXBpsMax != nil {
		in, out := &in.RXBpsMax, &out.RXBpsMax
		*out = new(int64)
		**out = **in
	}
	if in.TXBpsMin != nil {
		in, out := &in.TXBpsMin, &out.TXBpsMin
		*out = new(int64)
		**out = **in
	}
	if in.TXBpsMax != nil {
		in, out := &in.TXBpsMax, &out.TXBpsMax
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetLimits.
func (in *NetLimits) DeepCopy() *NetLimits {
	if in == nil {
		return nil
	}
	out := new(NetLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLocalGet) DeepCopyInto(out *NodeLocalGet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLocalGet.
func (in *NodeLocalGet) DeepCopy() *NodeLocalGet {
	if in == nil {
		return nil
	}
	out := new(NodeLocalGet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOS) DeepCopyInto(out *NodeQOS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOS.
func (in *NodeQOS) DeepCopy() *NodeQOS {
	if in == nil {
		return nil
	}
	out := new(NodeQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeQOS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSList) DeepCopyInto(out *NodeQOSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeQOS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSList.
func (in *NodeQOSList) DeepCopy() *NodeQOSList {
	if in == nil {
		return nil
	}
	out := new(NodeQOSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeQOSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSSpec) DeepCopyInto(out *NodeQOSSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.NodeQualityProbe.DeepCopyInto(&out.NodeQualityProbe)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ElasticCpuLimit.DeepCopyInto(&out.ElasticCpuLimit)
	in.MemoryLimit.DeepCopyInto(&out.MemoryLimit)
	in.NetLimits.DeepCopyInto(&out.NetLimits)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSSpec.
func (in *NodeQOSSpec) DeepCopy() *NodeQOSSpec {
	if in == nil {
		return nil
	}
	out := new(NodeQOSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQOSStatus) DeepCopyInto(out *NodeQOSStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQOSStatus.
func (in *NodeQOSStatus) DeepCopy() *NodeQOSStatus {
	if in == nil {
		return nil
	}
	out := new(NodeQOSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeQualityProbe) DeepCopyInto(out *NodeQualityProbe) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(corev1.HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeLocalGet != nil {
		in, out := &in.NodeLocalGet, &out.NodeLocalGet
		*out = new(NodeLocalGet)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeQualityProbe.
func (in *NodeQualityProbe) DeepCopy() *NodeQualityProbe {
	if in == nil {
		return nil
	}
	out := new(NodeQualityProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOS) DeepCopyInto(out *PodQOS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOS.
func (in *PodQOS) DeepCopy() *PodQOS {
	if in == nil {
		return nil
	}
	out := new(PodQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodQOS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSEnsurancePolicyStatus) DeepCopyInto(out *PodQOSEnsurancePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSEnsurancePolicyStatus.
func (in *PodQOSEnsurancePolicyStatus) DeepCopy() *PodQOSEnsurancePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(PodQOSEnsurancePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSList) DeepCopyInto(out *PodQOSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodQOS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSList.
func (in *PodQOSList) DeepCopy() *PodQOSList {
	if in == nil {
		return nil
	}
	out := new(PodQOSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodQOSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSSpec) DeepCopyInto(out *PodQOSSpec) {
	*out = *in
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.ScopeSelector != nil {
		in, out := &in.ScopeSelector, &out.ScopeSelector
		*out = new(ScopeSelector)
		(*in).DeepCopyInto(*out)
	}
	in.ResourceQOS.DeepCopyInto(&out.ResourceQOS)
	in.PodQualityProbe.DeepCopyInto(&out.PodQualityProbe)
	if in.AllowedActions != nil {
		in, out := &in.AllowedActions, &out.AllowedActions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSSpec.
func (in *PodQOSSpec) DeepCopy() *PodQOSSpec {
	if in == nil {
		return nil
	}
	out := new(PodQOSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQOSStatus) DeepCopyInto(out *PodQOSStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQOSStatus.
func (in *PodQOSStatus) DeepCopy() *PodQOSStatus {
	if in == nil {
		return nil
	}
	out := new(PodQOSStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodQualityProbe) DeepCopyInto(out *PodQualityProbe) {
	*out = *in
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(corev1.HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodQualityProbe.
func (in *PodQualityProbe) DeepCopy() *PodQualityProbe {
	if in == nil {
		return nil
	}
	out := new(PodQualityProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QOSEnsurance) DeepCopyInto(out *QOSEnsurance) {
	*out = *in
	in.Rule.DeepCopyInto(&out.Rule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QOSEnsurance.
func (in *QOSEnsurance) DeepCopy() *QOSEnsurance {
	if in == nil {
		return nil
	}
	out := new(QOSEnsurance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RDT) DeepCopyInto(out *RDT) {
	*out = *in
	if in.L3 != nil {
		in, out := &in.L3, &out.L3
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MB != nil {
		in, out := &in.MB, &out.MB
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDT.
func (in *RDT) DeepCopy() *RDT {
	if in == nil {
		return nil
	}
	out := new(RDT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RDTValue) DeepCopyInto(out *RDTValue) {
	{
		in := &in
		*out = make(RDTValue, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RDTValue.
func (in RDTValue) DeepCopy() RDTValue {
	if in == nil {
		return nil
	}
	out := new(RDTValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQOS) DeepCopyInto(out *ResourceQOS) {
	*out = *in
	if in.CPUQOS != nil {
		in, out := &in.CPUQOS, &out.CPUQOS
		*out = new(CPUQOS)
		(*in).DeepCopyInto(*out)
	}
	if in.MemoryQOS != nil {
		in, out := &in.MemoryQOS, &out.MemoryQOS
		*out = new(MemoryQOS)
		(*in).DeepCopyInto(*out)
	}
	if in.NetIOQOS != nil {
		in, out := &in.NetIOQOS, &out.NetIOQOS
		*out = new(NetIOQOS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskIOQOS != nil {
		in, out := &in.DiskIOQOS, &out.DiskIOQOS
		*out = new(DiskIOQOS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQOS.
func (in *ResourceQOS) DeepCopy() *ResourceQOS {
	if in == nil {
		return nil
	}
	out := new(ResourceQOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.MetricRule != nil {
		in, out := &in.MetricRule, &out.MetricRule
		*out = new(MetricRule)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeSelector) DeepCopyInto(out *ScopeSelector) {
	*out = *in
	if in.MatchExpressions != nil {
		in, out := &in.MatchExpressions, &out.MatchExpressions
		*out = make([]ScopedResourceSelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeSelector.
func (in *ScopeSelector) DeepCopy() *ScopeSelector {
	if in == nil {
		return nil
	}
	out := new(ScopeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopedResourceSelectorRequirement) DeepCopyInto(out *ScopedResourceSelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopedResourceSelectorRequirement.
func (in *ScopedResourceSelectorRequirement) DeepCopy() *ScopedResourceSelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(ScopedResourceSelectorRequirement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThrottleAction) DeepCopyInto(out *ThrottleAction) {
	*out = *in
	out.CPUThrottle = in.CPUThrottle
	out.MemoryThrottle = in.MemoryThrottle
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThrottleAction.
func (in *ThrottleAction) DeepCopy() *ThrottleAction {
	if in == nil {
		return nil
	}
	out := new(ThrottleAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhitelistPorts) DeepCopyInto(out *WhitelistPorts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhitelistPorts.
func (in *WhitelistPorts) DeepCopy() *WhitelistPorts {
	if in == nil {
		return nil
	}
	out := new(WhitelistPorts)
	in.DeepCopyInto(out)
	return out
}
