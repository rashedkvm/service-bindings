/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned"
	duckv1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/duck/v1alpha2"
	fakeduckv1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/duck/v1alpha2/fake"
	bindingsv1alpha1 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/labs/v1alpha1"
	fakebindingsv1alpha1 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/labs/v1alpha1/fake"
	internalv1alpha1 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/labsinternal/v1alpha1"
	fakeinternalv1alpha1 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/labsinternal/v1alpha1/fake"
	servicev1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/servicebinding/v1alpha2"
	fakeservicev1alpha2 "github.com/vmware-labs/service-bindings/pkg/client/clientset/versioned/typed/servicebinding/v1alpha2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// DuckV1alpha2 retrieves the DuckV1alpha2Client
func (c *Clientset) DuckV1alpha2() duckv1alpha2.DuckV1alpha2Interface {
	return &fakeduckv1alpha2.FakeDuckV1alpha2{Fake: &c.Fake}
}

// BindingsV1alpha1 retrieves the BindingsV1alpha1Client
func (c *Clientset) BindingsV1alpha1() bindingsv1alpha1.BindingsV1alpha1Interface {
	return &fakebindingsv1alpha1.FakeBindingsV1alpha1{Fake: &c.Fake}
}

// InternalV1alpha1 retrieves the InternalV1alpha1Client
func (c *Clientset) InternalV1alpha1() internalv1alpha1.InternalV1alpha1Interface {
	return &fakeinternalv1alpha1.FakeInternalV1alpha1{Fake: &c.Fake}
}

// ServiceV1alpha2 retrieves the ServiceV1alpha2Client
func (c *Clientset) ServiceV1alpha2() servicev1alpha2.ServiceV1alpha2Interface {
	return &fakeservicev1alpha2.FakeServiceV1alpha2{Fake: &c.Fake}
}
