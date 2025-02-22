// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.Installation":       schema_pkg_apis_integreatly_v1alpha1_Installation(ref),
		"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationSpec":   schema_pkg_apis_integreatly_v1alpha1_InstallationSpec(ref),
		"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationStatus": schema_pkg_apis_integreatly_v1alpha1_InstallationStatus(ref),
	}
}

func schema_pkg_apis_integreatly_v1alpha1_Installation(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Installation is the Schema for the installations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationSpec", "github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_InstallationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationSpec defines the desired state of Installation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"routingSubdomain": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"masterURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespacePrefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"selfSignedCerts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"pullSecret": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.PullSecretSpec"),
						},
					},
				},
				Required: []string{"type", "namespacePrefix", "selfSignedCerts"},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.PullSecretSpec"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_InstallationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "InstallationStatus defines the observed state of Installation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"stages": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationStageStatus"),
									},
								},
							},
						},
					},
					"preflightStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"preflightMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"stages"},
			},
		},
		Dependencies: []string{
			"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1.InstallationStageStatus"},
	}
}
