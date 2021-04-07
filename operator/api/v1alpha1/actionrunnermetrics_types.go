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

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ActionRunnerMetricsSpec defines the desired state of ActionRunnerMetrics
type ActionRunnerMetricsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ActionRunnerMetrics. Edit ActionRunnerMetrics_types.go to remove/update
	Namespace                   string        `json:"namespace"`
	Name                        string        `json:"name"`
	Image                       string        `json:"image,omitempty"`
	Replicas                    int32         `json:"replicas,omitempty"`
	CreateApiServer             *bool         `json:"createApiServer,omitempty"`
	CreateMemcached             *bool         `json:"createMemcached,omitempty"`
	CreateAuthentication        *bool         `json:"createAuthentication,omitempty"`
	PrometheusNamespace         string        `json:"prometheusNamespace,omitempty"`
	MemcachedReplicas           int32         `json:"memcachedReplicas,omitempty"`
	MemcachedImage              string        `json:"memcachedImage,omitempty"`
	ExistingSslCertSecret       string        `json:"existingSslCertSecret,omitempty"`
	KedaNamespace               string        `json:"kedaNamespace,omitempty"`
	ExistingMemcacheCredsSecret string        `json:"existingMemcacheCredsSecret,omitempty"`
	MemcachedUser               *string       `json:"memcacheUser,omitempty"`
	ExistingMemcacheServers     string        `json:"existingMemcacheServers,omitempty"`
	CacheWindow                 time.Duration `json:"cacheWindow,omitempty"`
	CacheWindowWhenEmpty        time.Duration `json:"cacheWindowWhenEmpty,omitempty"`
	ResyncInterval              time.Duration `json:"resyncInterval,omitempty"`
	Namespaces                  []string      `json:"namespaces,omitempty"`
}

// ActionRunnerMetricsStatus defines the observed state of ActionRunnerMetrics
type ActionRunnerMetricsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=actionrunnermetrics,scope=Cluster
// ActionRunnerMetrics is the Schema for the actionrunnermetrics API
type ActionRunnerMetrics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionRunnerMetricsSpec   `json:"spec,omitempty"`
	Status ActionRunnerMetricsStatus `json:"status,omitempty"`
}

func (a *ActionRunnerMetrics) Setup() {
	if a.Spec.CacheWindow.Milliseconds() == 0 {
		a.Spec.CacheWindow, _ = time.ParseDuration("1m")
	}
	if a.Spec.CacheWindowWhenEmpty.Milliseconds() == 0 {
		a.Spec.CacheWindowWhenEmpty, _ = time.ParseDuration("30s")
	}
	if a.Spec.ResyncInterval.Milliseconds() == 0 {
		a.Spec.ResyncInterval, _ = time.ParseDuration("5m")
	}
	if a.Spec.MemcachedReplicas == 0 {
		a.Spec.MemcachedReplicas = 2
	}
	if a.Spec.Replicas == 0 {
		a.Spec.Replicas = 2
	}
	boolTrue := true
	if a.Spec.CreateApiServer == nil {
		a.Spec.CreateApiServer = &boolTrue
	}
	if a.Spec.CreateAuthentication == nil {
		a.Spec.CreateAuthentication = &boolTrue
	}
	if a.Spec.CreateMemcached == nil {
		a.Spec.CreateMemcached = &boolTrue
	}
	if a.Spec.Image == "" {
		a.Spec.Image = "joeshearn/github-runner-autoscaler-apiserver:latest"
	}
	if a.Spec.KedaNamespace == "" {
		a.Spec.KedaNamespace = "keda"
	}
	if a.Spec.MemcachedImage == "" {
		a.Spec.MemcachedImage = "docker.io/bitnami/memcached:1.6.9-debian-10-r86"
	}
	if a.Spec.MemcachedUser == nil {
		user := "user"
		a.Spec.MemcachedUser = &user
	}
}

// +kubebuilder:object:root=true

// ActionRunnerMetricsList contains a list of ActionRunnerMetrics
type ActionRunnerMetricsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionRunnerMetrics `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ActionRunnerMetrics{}, &ActionRunnerMetricsList{})
}
