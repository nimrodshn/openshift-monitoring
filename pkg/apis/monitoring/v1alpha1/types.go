/*
Copyright (c) 2018 Red Hat, Inc.

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

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Alert is the description of an alert.
//
type Alert struct {
	meta.TypeMeta `json:",inline"`

	// Standard object metadata.
	// +optional
	meta.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// The status of the alert.
	// +optional
	Status AlertStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// AlertStatus is the status for an alert.
//
type AlertStatus struct {
	// Prometheus labels.
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,1,rep,name=labels"`

	// Prometheus annotations.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,2,rep,name=annotations"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertList is a list of alerts.
//
type AlertList struct {
	meta.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	meta.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of alerts.
	Items []Alert `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertRule is the description of an alerting rule.
//
type AlertingRule struct {
	meta.TypeMeta `json:",inline"`

	// Standard object metadata.
	// +optional
	meta.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the alerting rule.
	// +optional
	Spec AlertingRuleSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec`

	// Status represents the current information about an alenrting rule.
	// +optional
	Status AlertingRuleStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=spec`
}

// AlertingRuleSpec is the specification for an alerting rule.
//
type AlertingRuleSpec struct {
	// Group is the name of the group that the alerting rule should belong to.
	// +optional
	Group string `json:"group,omitempty" protobuf:"bytes,1,opt,name=group`

	// Alert is the name of the rule that will be used inside Prometheus, and inside the alert
	// manager. If not specified will use the name of the object.
	// +optional
	Alert string `json:"alert,omitempty" protobuf:"bytes,2,opt,name=alert`

	// Expr is the expresion that defines the rule. The syntax is exactly the same used
	// to define alerting rules in Prometheus.
	Expr string `json:"expr,omitempty" protobuf:"bytes,3,opt,name=expr`

	// For indicates how long Prometheus will wait till the expression evaluates to true
	// till the alert is actually fired.
	// +optional
	For time.Duration `json:"for,omitempty" protobuf:"bytes,4,opt,name=spec`

	// Prometheus labels.
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,5,rep,name=labels"`

	// Prometheus annotations.
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,6,rep,name=annotations"`
}

// AlertingRuleStatus is the status for an alerting rule.
//
type AlertingRuleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertingRuleList is a list of alerting rules.
//
type AlertingRuleList struct {
	meta.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	meta.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of alerting rules.
	Items []AlertingRule `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HealingRule is the description of an healing rule.
//
type HealingRule struct {
	meta.TypeMeta `json:",inline"`

	// Standard object metadata.
	// +optional
	meta.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the healing rule.
	// +optional
	Spec HealingRuleSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec`

	// Status represents the current information about an healing rule.
	// +optional
	Status HealingRuleStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=spec`
}

// HealingRuleSpec is the specification for a healing rule.
//
type HealingRuleSpec struct {
	// Conditions is the set of conditions that should be true for the rule to be activated.
	// +optional
	Conditions []HealingCondition `json:"conditions" protobuf:"bytes,1,rep,name=conditions`

	// Actions is the set of healing actions that will be executed when the rule is activated.
	// +optional
	Actions []HealingAction `json:"actions" protobuf:"bytes,2,rep,name=conditions`
}

// HealingCondition represents a condition for the activation of a healing rule.
//
type HealingCondition struct {
	// Alert is the name of an alert that has to exist for this condition to be true.
	// +optional
	Alert string `json:"alert,omitempty" protobuf:"bytes,1,opt,name=alert`
}

// HealingAction represents an action that will be performed when a healing rule is activated.
//
type HealingAction struct {
	// AWX is used when the healing action is implemented by an Ansible AWX job.
	// +optional
	AWX *AWXAction `json:"awx,omitempty" protobuf:"bytes,1,opt,name=awx`
}

// AWXAction describes how to run an Ansible AWX job.
//
type AWXAction struct {
	// The name of the project that contains the job template.
	// +optional
	Project string `json:"project,omitempty" protobuf:"bytes,1,opt,name=project`

	// The name of the job template.
	// +optional
	JobTemplate string `json:"jobTemplate,omitempty" protobuf:"bytes,2,opt,name=jobTemplate`
}

// HealingRuleStatus is the status for an alerting rule.
//
type HealingRuleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HealingRuleList is a list of alerting rules.
//
type HealingRuleList struct {
	meta.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	meta.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of alerting rules.
	Items []HealingRule `json:"items" protobuf:"bytes,2,rep,name=items"`
}
