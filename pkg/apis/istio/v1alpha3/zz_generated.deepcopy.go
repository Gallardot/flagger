//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Flux authors

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

package v1alpha3

import (
	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/istio/common/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolSettings) DeepCopyInto(out *ConnectionPoolSettings) {
	*out = *in
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPSettings)
		**out = **in
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(HTTPSettings)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings.
func (in *ConnectionPoolSettings) DeepCopy() *ConnectionPoolSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsistentHashLB) DeepCopyInto(out *ConsistentHashLB) {
	*out = *in
	if in.HTTPCookie != nil {
		in, out := &in.HTTPCookie, &out.HTTPCookie
		*out = new(HTTPCookie)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsistentHashLB.
func (in *ConsistentHashLB) DeepCopy() *ConsistentHashLB {
	if in == nil {
		return nil
	}
	out := new(ConsistentHashLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CorsPolicy) DeepCopyInto(out *CorsPolicy) {
	*out = *in
	if in.AllowOrigin != nil {
		in, out := &in.AllowOrigin, &out.AllowOrigin
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowOrigins != nil {
		in, out := &in.AllowOrigins, &out.AllowOrigins
		*out = make([]*v1alpha1.StringMatch, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(v1alpha1.StringMatch)
				**out = **in
			}
		}
	}
	if in.AllowMethods != nil {
		in, out := &in.AllowMethods, &out.AllowMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AllowHeaders != nil {
		in, out := &in.AllowHeaders, &out.AllowHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExposeHeaders != nil {
		in, out := &in.ExposeHeaders, &out.ExposeHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CorsPolicy.
func (in *CorsPolicy) DeepCopy() *CorsPolicy {
	if in == nil {
		return nil
	}
	out := new(CorsPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(PortSelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationRule) DeepCopyInto(out *DestinationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRule.
func (in *DestinationRule) DeepCopy() *DestinationRule {
	if in == nil {
		return nil
	}
	out := new(DestinationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DestinationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationRuleList) DeepCopyInto(out *DestinationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DestinationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRuleList.
func (in *DestinationRuleList) DeepCopy() *DestinationRuleList {
	if in == nil {
		return nil
	}
	out := new(DestinationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DestinationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationRuleSpec) DeepCopyInto(out *DestinationRuleSpec) {
	*out = *in
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(TrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Subsets != nil {
		in, out := &in.Subsets, &out.Subsets
		*out = make([]Subset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationRuleSpec.
func (in *DestinationRuleSpec) DeepCopy() *DestinationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DestinationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Distribute) DeepCopyInto(out *Distribute) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make(map[string]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Distribute.
func (in *Distribute) DeepCopy() *Distribute {
	if in == nil {
		return nil
	}
	out := new(Distribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Failover) DeepCopyInto(out *Failover) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Failover.
func (in *Failover) DeepCopy() *Failover {
	if in == nil {
		return nil
	}
	out := new(Failover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPCookie) DeepCopyInto(out *HTTPCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPCookie.
func (in *HTTPCookie) DeepCopy() *HTTPCookie {
	if in == nil {
		return nil
	}
	out := new(HTTPCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPFaultInjection) DeepCopyInto(out *HTTPFaultInjection) {
	*out = *in
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(InjectDelay)
		**out = **in
	}
	if in.Abort != nil {
		in, out := &in.Abort, &out.Abort
		*out = new(InjectAbort)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPFaultInjection.
func (in *HTTPFaultInjection) DeepCopy() *HTTPFaultInjection {
	if in == nil {
		return nil
	}
	out := new(HTTPFaultInjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPMatchRequest) DeepCopyInto(out *HTTPMatchRequest) {
	*out = *in
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(v1alpha1.StringMatch)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(v1alpha1.StringMatch)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(v1alpha1.StringMatch)
		**out = **in
	}
	if in.Authority != nil {
		in, out := &in.Authority, &out.Authority
		*out = new(v1alpha1.StringMatch)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]v1alpha1.StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.QueryParams != nil {
		in, out := &in.QueryParams, &out.QueryParams
		*out = make(map[string]v1alpha1.StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.WithoutHeaders != nil {
		in, out := &in.WithoutHeaders, &out.WithoutHeaders
		*out = make(map[string]v1alpha1.StringMatch, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMatchRequest.
func (in *HTTPMatchRequest) DeepCopy() *HTTPMatchRequest {
	if in == nil {
		return nil
	}
	out := new(HTTPMatchRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRedirect) DeepCopyInto(out *HTTPRedirect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRedirect.
func (in *HTTPRedirect) DeepCopy() *HTTPRedirect {
	if in == nil {
		return nil
	}
	out := new(HTTPRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRetry) DeepCopyInto(out *HTTPRetry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRetry.
func (in *HTTPRetry) DeepCopy() *HTTPRetry {
	if in == nil {
		return nil
	}
	out := new(HTTPRetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRewrite) DeepCopyInto(out *HTTPRewrite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRewrite.
func (in *HTTPRewrite) DeepCopy() *HTTPRewrite {
	if in == nil {
		return nil
	}
	out := new(HTTPRewrite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRoute) DeepCopyInto(out *HTTPRoute) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = make([]HTTPRouteDestination, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(HTTPRedirect)
		**out = **in
	}
	if in.Rewrite != nil {
		in, out := &in.Rewrite, &out.Rewrite
		*out = new(HTTPRewrite)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(HTTPRetry)
		**out = **in
	}
	if in.Fault != nil {
		in, out := &in.Fault, &out.Fault
		*out = new(HTTPFaultInjection)
		(*in).DeepCopyInto(*out)
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.MirrorPercentage != nil {
		in, out := &in.MirrorPercentage, &out.MirrorPercentage
		*out = new(Percent)
		**out = **in
	}
	if in.CorsPolicy != nil {
		in, out := &in.CorsPolicy, &out.CorsPolicy
		*out = new(CorsPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(Headers)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRoute.
func (in *HTTPRoute) DeepCopy() *HTTPRoute {
	if in == nil {
		return nil
	}
	out := new(HTTPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteDestination) DeepCopyInto(out *HTTPRouteDestination) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(Headers)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteDestination.
func (in *HTTPRouteDestination) DeepCopy() *HTTPRouteDestination {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSettings) DeepCopyInto(out *HTTPSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSettings.
func (in *HTTPSettings) DeepCopy() *HTTPSettings {
	if in == nil {
		return nil
	}
	out := new(HTTPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderOperations) DeepCopyInto(out *HeaderOperations) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderOperations.
func (in *HeaderOperations) DeepCopy() *HeaderOperations {
	if in == nil {
		return nil
	}
	out := new(HeaderOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Headers) DeepCopyInto(out *Headers) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(HeaderOperations)
		(*in).DeepCopyInto(*out)
	}
	if in.Response != nil {
		in, out := &in.Response, &out.Response
		*out = new(HeaderOperations)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Headers.
func (in *Headers) DeepCopy() *Headers {
	if in == nil {
		return nil
	}
	out := new(Headers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InjectAbort) DeepCopyInto(out *InjectAbort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InjectAbort.
func (in *InjectAbort) DeepCopy() *InjectAbort {
	if in == nil {
		return nil
	}
	out := new(InjectAbort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InjectDelay) DeepCopyInto(out *InjectDelay) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InjectDelay.
func (in *InjectDelay) DeepCopy() *InjectDelay {
	if in == nil {
		return nil
	}
	out := new(InjectDelay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4MatchAttributes) DeepCopyInto(out *L4MatchAttributes) {
	*out = *in
	if in.SourceLabel != nil {
		in, out := &in.SourceLabel, &out.SourceLabel
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4MatchAttributes.
func (in *L4MatchAttributes) DeepCopy() *L4MatchAttributes {
	if in == nil {
		return nil
	}
	out := new(L4MatchAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSettings) DeepCopyInto(out *LoadBalancerSettings) {
	*out = *in
	if in.ConsistentHash != nil {
		in, out := &in.ConsistentHash, &out.ConsistentHash
		*out = new(ConsistentHashLB)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalityLbSetting != nil {
		in, out := &in.LocalityLbSetting, &out.LocalityLbSetting
		*out = new(LocalityLbSetting)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSettings.
func (in *LoadBalancerSettings) DeepCopy() *LoadBalancerSettings {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalityLbSetting) DeepCopyInto(out *LocalityLbSetting) {
	*out = *in
	if in.Distribute != nil {
		in, out := &in.Distribute, &out.Distribute
		*out = make([]Distribute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Failover != nil {
		in, out := &in.Failover, &out.Failover
		*out = make([]Failover, len(*in))
		copy(*out, *in)
	}
	if in.FailoverPriority != nil {
		in, out := &in.FailoverPriority, &out.FailoverPriority
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalityLbSetting.
func (in *LocalityLbSetting) DeepCopy() *LocalityLbSetting {
	if in == nil {
		return nil
	}
	out := new(LocalityLbSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	*out = *in
	if in.ConsecutiveGatewayErrors != nil {
		in, out := &in.ConsecutiveGatewayErrors, &out.ConsecutiveGatewayErrors
		*out = new(uint32)
		**out = **in
	}
	if in.Consecutive5xxErrors != nil {
		in, out := &in.Consecutive5xxErrors, &out.Consecutive5xxErrors
		*out = new(uint32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Percent) DeepCopyInto(out *Percent) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percent.
func (in *Percent) DeepCopy() *Percent {
	if in == nil {
		return nil
	}
	out := new(Percent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortTrafficPolicy) DeepCopyInto(out *PortTrafficPolicy) {
	*out = *in
	out.Port = in.Port
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectionPool != nil {
		in, out := &in.ConnectionPool, &out.ConnectionPool
		*out = new(ConnectionPoolSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.OutlierDetection != nil {
		in, out := &in.OutlierDetection, &out.OutlierDetection
		*out = new(OutlierDetection)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSSettings)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortTrafficPolicy.
func (in *PortTrafficPolicy) DeepCopy() *PortTrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(PortTrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subset) DeepCopyInto(out *Subset) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(TrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subset.
func (in *Subset) DeepCopy() *Subset {
	if in == nil {
		return nil
	}
	out := new(Subset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPRoute) DeepCopyInto(out *TCPRoute) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]L4MatchAttributes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Route.DeepCopyInto(&out.Route)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPRoute.
func (in *TCPRoute) DeepCopy() *TCPRoute {
	if in == nil {
		return nil
	}
	out := new(TCPRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPSettings) DeepCopyInto(out *TCPSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPSettings.
func (in *TCPSettings) DeepCopy() *TCPSettings {
	if in == nil {
		return nil
	}
	out := new(TCPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSettings) DeepCopyInto(out *TLSSettings) {
	*out = *in
	if in.SubjectAltNames != nil {
		in, out := &in.SubjectAltNames, &out.SubjectAltNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSettings.
func (in *TLSSettings) DeepCopy() *TLSSettings {
	if in == nil {
		return nil
	}
	out := new(TLSSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficPolicy) DeepCopyInto(out *TrafficPolicy) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.ConnectionPool != nil {
		in, out := &in.ConnectionPool, &out.ConnectionPool
		*out = new(ConnectionPoolSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.OutlierDetection != nil {
		in, out := &in.OutlierDetection, &out.OutlierDetection
		*out = new(OutlierDetection)
		(*in).DeepCopyInto(*out)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.PortLevelSettings != nil {
		in, out := &in.PortLevelSettings, &out.PortLevelSettings
		*out = make([]PortTrafficPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficPolicy.
func (in *TrafficPolicy) DeepCopy() *TrafficPolicy {
	if in == nil {
		return nil
	}
	out := new(TrafficPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualService) DeepCopyInto(out *VirtualService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualService.
func (in *VirtualService) DeepCopy() *VirtualService {
	if in == nil {
		return nil
	}
	out := new(VirtualService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceList) DeepCopyInto(out *VirtualServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceList.
func (in *VirtualServiceList) DeepCopy() *VirtualServiceList {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServiceSpec) DeepCopyInto(out *VirtualServiceSpec) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = make([]HTTPRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tcp != nil {
		in, out := &in.Tcp, &out.Tcp
		*out = make([]TCPRoute, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServiceSpec.
func (in *VirtualServiceSpec) DeepCopy() *VirtualServiceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServiceSpec)
	in.DeepCopyInto(out)
	return out
}
