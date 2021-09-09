// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2beta2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runner) DeepCopyInto(out *Runner) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.WorkVolumeClaimTemplate != nil {
		in, out := &in.WorkVolumeClaimTemplate, &out.WorkVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(map[v1.ResourceName]resource.Quantity)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(map[v1.ResourceName]resource.Quantity)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MountDockerSock != nil {
		in, out := &in.MountDockerSock, &out.MountDockerSock
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runner.
func (in *Runner) DeepCopy() *Runner {
	if in == nil {
		return nil
	}
	out := new(Runner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunner) DeepCopyInto(out *ScaledActionRunner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunner.
func (in *ScaledActionRunner) DeepCopy() *ScaledActionRunner {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledActionRunner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerCore) DeepCopyInto(out *ScaledActionRunnerCore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerCore.
func (in *ScaledActionRunnerCore) DeepCopy() *ScaledActionRunnerCore {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledActionRunnerCore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerCoreList) DeepCopyInto(out *ScaledActionRunnerCoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledActionRunnerCore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerCoreList.
func (in *ScaledActionRunnerCoreList) DeepCopy() *ScaledActionRunnerCoreList {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerCoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledActionRunnerCoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerCoreSpec) DeepCopyInto(out *ScaledActionRunnerCoreSpec) {
	*out = *in
	if in.ApiServerExtraArgs != nil {
		in, out := &in.ApiServerExtraArgs, &out.ApiServerExtraArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CreateApiServer != nil {
		in, out := &in.CreateApiServer, &out.CreateApiServer
		*out = new(bool)
		**out = **in
	}
	if in.CreateMemcached != nil {
		in, out := &in.CreateMemcached, &out.CreateMemcached
		*out = new(bool)
		**out = **in
	}
	if in.CreateAuthentication != nil {
		in, out := &in.CreateAuthentication, &out.CreateAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.MemcachedUser != nil {
		in, out := &in.MemcachedUser, &out.MemcachedUser
		*out = new(string)
		**out = **in
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerCoreSpec.
func (in *ScaledActionRunnerCoreSpec) DeepCopy() *ScaledActionRunnerCoreSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerCoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerCoreStatus) DeepCopyInto(out *ScaledActionRunnerCoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerCoreStatus.
func (in *ScaledActionRunnerCoreStatus) DeepCopy() *ScaledActionRunnerCoreStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerCoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerList) DeepCopyInto(out *ScaledActionRunnerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledActionRunner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerList.
func (in *ScaledActionRunnerList) DeepCopy() *ScaledActionRunnerList {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledActionRunnerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerSpec) DeepCopyInto(out *ScaledActionRunnerSpec) {
	*out = *in
	if in.RunnerSecrets != nil {
		in, out := &in.RunnerSecrets, &out.RunnerSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(Scaling)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleFactor != nil {
		in, out := &in.ScaleFactor, &out.ScaleFactor
		*out = new(string)
		**out = **in
	}
	if in.MetricsSelector != nil {
		in, out := &in.MetricsSelector, &out.MetricsSelector
		*out = new(string)
		**out = **in
	}
	if in.Runner != nil {
		in, out := &in.Runner, &out.Runner
		*out = new(Runner)
		(*in).DeepCopyInto(*out)
	}
	if in.ForceScaleUpWindowSecs != nil {
		in, out := &in.ForceScaleUpWindowSecs, &out.ForceScaleUpWindowSecs
		*out = new(int32)
		**out = **in
	}
	if in.ForceScaleUpFrequencyDays != nil {
		in, out := &in.ForceScaleUpFrequencyDays, &out.ForceScaleUpFrequencyDays
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerSpec.
func (in *ScaledActionRunnerSpec) DeepCopy() *ScaledActionRunnerSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledActionRunnerStatus) DeepCopyInto(out *ScaledActionRunnerStatus) {
	*out = *in
	if in.ReferencedSecrets != nil {
		in, out := &in.ReferencedSecrets, &out.ReferencedSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledActionRunnerStatus.
func (in *ScaledActionRunnerStatus) DeepCopy() *ScaledActionRunnerStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledActionRunnerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Scaling) DeepCopyInto(out *Scaling) {
	*out = *in
	if in.Behavior != nil {
		in, out := &in.Behavior, &out.Behavior
		*out = new(v2beta2.HorizontalPodAutoscalerBehavior)
		(*in).DeepCopyInto(*out)
	}
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Scaling.
func (in *Scaling) DeepCopy() *Scaling {
	if in == nil {
		return nil
	}
	out := new(Scaling)
	in.DeepCopyInto(out)
	return out
}
