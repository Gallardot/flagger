/*
Copyright 2022 The Flux authors

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

package router

import (
	"context"
	"fmt"
	"testing"

	a6v2 "github.com/fluxcd/flagger/pkg/apis/apisix/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apimachinery/pkg/util/intstr"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	ar := &a6v2.ApisixRoute{
		TypeMeta: metav1.TypeMeta{APIVersion: a6v2.SchemeGroupVersion.String()},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "podinfo",
		},
		Spec: a6v2.ApisixRouteSpec{HTTP: []a6v2.ApisixRouteHTTP{
			{
				Name: "method",
				Match: a6v2.ApisixRouteHTTPMatch{
					Hosts:   []string{"foobar.com"},
					Methods: []string{"GET"},
					Paths:   []string{"/*"},
				},
				Plugins: []a6v2.ApisixRoutePlugin{
					{
						Name:   "prometheus",
						Enable: true,
						Config: a6v2.ApisixRoutePluginConfig{
							"disable":     "false",
							"prefer_name": "true",
						},
					},
				},
				Backends: []a6v2.ApisixRouteHTTPBackend{
					{ServiceName: "podinfo",
						ServicePort: intstr.IntOrString{
							Type:   intstr.Int,
							IntVal: 80,
						}},
				},
			},
		},
		},
	}
	br := &a6v2.ApisixRoute{
		TypeMeta: metav1.TypeMeta{APIVersion: a6v2.SchemeGroupVersion.String()},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "podinfo",
		},
		Spec: a6v2.ApisixRouteSpec{HTTP: []a6v2.ApisixRouteHTTP{
			{
				Name: "method",
				Match: a6v2.ApisixRouteHTTPMatch{
					Hosts:   []string{"foobar.com"},
					Methods: []string{"GET"},
					Paths:   []string{"/*"},
				},
				Plugins: []a6v2.ApisixRoutePlugin{
					{
						Name:   "prometheus",
						Enable: true,
						Config: a6v2.ApisixRoutePluginConfig{
							"disable":     "false",
							"prefer_name": "true",
						},
					},
				},
				Backends: []a6v2.ApisixRouteHTTPBackend{
					{ServiceName: "podinfo",
						ServicePort: intstr.IntOrString{
							Type:   intstr.Int,
							IntVal: 80,
						}},
				},
			},
		},
		},
	}

	assert.Equal(t, "", cmp.Diff(ar.Spec.HTTP[0], br.Spec.HTTP[0], cmpopts.IgnoreFields(a6v2.ApisixRouteHTTP{}, "Name")))
}

func TestApisixRouter_Reconcile(t *testing.T) {
	mocks := newFixture(nil)
	mocks.canary.Spec.RouteRef = &v1beta1.LocalObjectReference{
		Name:       "podinfo",
		Kind:       "ApisixRoute",
		APIVersion: "apisix.apache.org/v2",
	}
	router := &ApisixRouter{
		apisixClient: mocks.flaggerClient,
		logger:       mocks.logger,
	}
	err := router.Reconcile(mocks.canary)
	require.NoError(t, err)
	canaryName := fmt.Sprintf("%s-canary", mocks.canary.Spec.RouteRef.Name)
	arCanary, err := router.apisixClient.ApisixV2().ApisixRoutes("default").Get(context.TODO(), canaryName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, 2, len(arCanary.Spec.HTTP[0].Backends))
}

func TestApisixRouter_GetSetRoutes(t *testing.T) {
	mocks := newFixture(nil)
	mocks.canary.Spec.RouteRef = &v1beta1.LocalObjectReference{
		Name:       "podinfo",
		Kind:       "ApisixRoute",
		APIVersion: "apisix.apache.org/v2",
	}
	router := &ApisixRouter{
		apisixClient: mocks.flaggerClient,
		logger:       mocks.logger,
	}
	err := router.Reconcile(mocks.canary)
	require.NoError(t, err)
	p, c, m, err := router.GetRoutes(mocks.canary)
	require.NoError(t, err)
	assert.Equal(t, 100, p)
	assert.Equal(t, 0, c)
	assert.False(t, m)

	p = 50
	c = 50
	m = false
	err = router.SetRoutes(mocks.canary, p, c, m)
	require.NoError(t, err)

	p, c, m, err = router.GetRoutes(mocks.canary)
	require.NoError(t, err)
	assert.Equal(t, 50, p)
	assert.Equal(t, 50, c)
	assert.False(t, m)

	canaryName := fmt.Sprintf("%s-canary", mocks.canary.Spec.RouteRef.Name)
	arRouter, err := router.apisixClient.ApisixV2().ApisixRoutes("default").Get(context.TODO(), canaryName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, 2, len(arRouter.Spec.HTTP[0].Backends))
	assert.Equal(t, 50, *arRouter.Spec.HTTP[0].Backends[0].Weight)
	assert.Equal(t, 50, *arRouter.Spec.HTTP[0].Backends[1].Weight)
}
