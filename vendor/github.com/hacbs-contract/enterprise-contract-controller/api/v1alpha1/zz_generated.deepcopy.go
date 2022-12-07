//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]AuthorizedComponent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizedComponent) DeepCopyInto(out *AuthorizedComponent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizedComponent.
func (in *AuthorizedComponent) DeepCopy() *AuthorizedComponent {
	if in == nil {
		return nil
	}
	out := new(AuthorizedComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicy) DeepCopyInto(out *EnterpriseContractPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicy.
func (in *EnterpriseContractPolicy) DeepCopy() *EnterpriseContractPolicy {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseContractPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicyConfiguration) DeepCopyInto(out *EnterpriseContractPolicyConfiguration) {
	*out = *in
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicyConfiguration.
func (in *EnterpriseContractPolicyConfiguration) DeepCopy() *EnterpriseContractPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicyExceptions) DeepCopyInto(out *EnterpriseContractPolicyExceptions) {
	*out = *in
	if in.NonBlocking != nil {
		in, out := &in.NonBlocking, &out.NonBlocking
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicyExceptions.
func (in *EnterpriseContractPolicyExceptions) DeepCopy() *EnterpriseContractPolicyExceptions {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicyExceptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicyList) DeepCopyInto(out *EnterpriseContractPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnterpriseContractPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicyList.
func (in *EnterpriseContractPolicyList) DeepCopy() *EnterpriseContractPolicyList {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnterpriseContractPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicySpec) DeepCopyInto(out *EnterpriseContractPolicySpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(Authorization)
		(*in).DeepCopyInto(*out)
	}
	if in.Exceptions != nil {
		in, out := &in.Exceptions, &out.Exceptions
		*out = new(EnterpriseContractPolicyExceptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(EnterpriseContractPolicyConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicySpec.
func (in *EnterpriseContractPolicySpec) DeepCopy() *EnterpriseContractPolicySpec {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnterpriseContractPolicyStatus) DeepCopyInto(out *EnterpriseContractPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnterpriseContractPolicyStatus.
func (in *EnterpriseContractPolicyStatus) DeepCopy() *EnterpriseContractPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(EnterpriseContractPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
