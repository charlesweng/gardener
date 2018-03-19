// +build !ignore_autogenerated

/*
Copyright 2018 The Gardener Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProfileControllerConfiguration) DeepCopyInto(out *CloudProfileControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProfileControllerConfiguration.
func (in *CloudProfileControllerConfiguration) DeepCopy() *CloudProfileControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudProfileControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	if in.GardenerClientConnection != nil {
		in, out := &in.GardenerClientConnection, &out.GardenerClientConnection
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClientConnectionConfiguration)
			**out = **in
		}
	}
	in.Controllers.DeepCopyInto(&out.Controllers)
	out.LeaderElection = in.LeaderElection
	out.Metrics = in.Metrics
	out.Server = in.Server
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerControllerConfiguration) DeepCopyInto(out *ControllerManagerControllerConfiguration) {
	*out = *in
	if in.CloudProfile != nil {
		in, out := &in.CloudProfile, &out.CloudProfile
		if *in == nil {
			*out = nil
		} else {
			*out = new(CloudProfileControllerConfiguration)
			**out = **in
		}
	}
	if in.SecretBinding != nil {
		in, out := &in.SecretBinding, &out.SecretBinding
		if *in == nil {
			*out = nil
		} else {
			*out = new(SecretBindingControllerConfiguration)
			**out = **in
		}
	}
	if in.CrossSecretBinding != nil {
		in, out := &in.CrossSecretBinding, &out.CrossSecretBinding
		if *in == nil {
			*out = nil
		} else {
			*out = new(CrossSecretBindingControllerConfiguration)
			**out = **in
		}
	}
	if in.PrivateSecretBinding != nil {
		in, out := &in.PrivateSecretBinding, &out.PrivateSecretBinding
		if *in == nil {
			*out = nil
		} else {
			*out = new(PrivateSecretBindingControllerConfiguration)
			**out = **in
		}
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		if *in == nil {
			*out = nil
		} else {
			*out = new(QuotaControllerConfiguration)
			**out = **in
		}
	}
	if in.Seed != nil {
		in, out := &in.Seed, &out.Seed
		if *in == nil {
			*out = nil
		} else {
			*out = new(SeedControllerConfiguration)
			**out = **in
		}
	}
	in.Shoot.DeepCopyInto(&out.Shoot)
	out.ShootCare = in.ShootCare
	out.ShootMaintenance = in.ShootMaintenance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerControllerConfiguration.
func (in *ControllerManagerControllerConfiguration) DeepCopy() *ControllerManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossSecretBindingControllerConfiguration) DeepCopyInto(out *CrossSecretBindingControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossSecretBindingControllerConfiguration.
func (in *CrossSecretBindingControllerConfiguration) DeepCopy() *CrossSecretBindingControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CrossSecretBindingControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfiguration) DeepCopyInto(out *MetricsConfiguration) {
	*out = *in
	out.Interval = in.Interval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfiguration.
func (in *MetricsConfiguration) DeepCopy() *MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateSecretBindingControllerConfiguration) DeepCopyInto(out *PrivateSecretBindingControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateSecretBindingControllerConfiguration.
func (in *PrivateSecretBindingControllerConfiguration) DeepCopy() *PrivateSecretBindingControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PrivateSecretBindingControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaControllerConfiguration) DeepCopyInto(out *QuotaControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaControllerConfiguration.
func (in *QuotaControllerConfiguration) DeepCopy() *QuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(QuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBindingControllerConfiguration) DeepCopyInto(out *SecretBindingControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBindingControllerConfiguration.
func (in *SecretBindingControllerConfiguration) DeepCopy() *SecretBindingControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretBindingControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedControllerConfiguration) DeepCopyInto(out *SeedControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedControllerConfiguration.
func (in *SeedControllerConfiguration) DeepCopy() *SeedControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SeedControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootCareControllerConfiguration) DeepCopyInto(out *ShootCareControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootCareControllerConfiguration.
func (in *ShootCareControllerConfiguration) DeepCopy() *ShootCareControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootCareControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootControllerConfiguration) DeepCopyInto(out *ShootControllerConfiguration) {
	*out = *in
	out.RetryDuration = in.RetryDuration
	out.SyncPeriod = in.SyncPeriod
	if in.WatchNamespace != nil {
		in, out := &in.WatchNamespace, &out.WatchNamespace
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootControllerConfiguration.
func (in *ShootControllerConfiguration) DeepCopy() *ShootControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMaintenanceControllerConfiguration) DeepCopyInto(out *ShootMaintenanceControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMaintenanceControllerConfiguration.
func (in *ShootMaintenanceControllerConfiguration) DeepCopy() *ShootMaintenanceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootMaintenanceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}
