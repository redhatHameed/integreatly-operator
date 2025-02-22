package fuseonopenshift

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"

	"github.com/integr8ly/integreatly-operator/pkg/apis"
	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/marketplace"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/products/config"

	imagev1 "github.com/openshift/api/image/v1"
	samplesv1 "github.com/openshift/cluster-samples-operator/pkg/apis/samples/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	pkgclient "sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	OPERATOR_NAMESPACE = "integreatly-operator"
)

type FuseOnOpenShiftScenario struct {
	Name           string
	ExpectError    bool
	ExpectedError  string
	ExpectedStatus v1alpha1.StatusPhase
	FakeConfig     *config.ConfigReadWriterMock
	FakeClient     pkgclient.Client
	FakeMPM        *marketplace.MarketplaceInterfaceMock
	Installation   *v1alpha1.Installation
	Product        *v1alpha1.InstallationProductStatus
}

func getFakeConfig() *config.ConfigReadWriterMock {
	return &config.ConfigReadWriterMock{
		GetOperatorNamespaceFunc: func() string {
			return OPERATOR_NAMESPACE
		},
		ReadFuseOnOpenshiftFunc: func() (ready *config.FuseOnOpenshift, e error) {
			return config.NewFuseOnOpenshift(config.ProductConfig{}), nil
		},
	}
}

func TestFuseOnOpenShift(t *testing.T) {
	// Initialize scheme so that types required by the scenarios are available
	scheme := scheme.Scheme
	if err := apis.AddToScheme(scheme); err != nil {
		t.Fatalf("failed to initialize scheme: %s", err)
	}

	sampleClusterConfig := &samplesv1.Config{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
		},
		Spec: samplesv1.ConfigSpec{
			SkippedImagestreams: []string{"fis-java-openshift"},
			SkippedTemplates:    []string{"fuse74-console-cluster"},
		},
	}

	// Sample imagestream that's managed by the sample cluster operator
	sampleClusterImgStream := &imagev1.ImageStream{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ImageStream",
			APIVersion: "image.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fis-java-openshift",
			Namespace: fuseOnOpenshiftNs,
			Labels: map[string]string{
				"samples.operator.openshift.io/managed": "true",
			},
		},
	}

	// Sample imagestream that's created by integreatly
	integreatlyImgStream := &imagev1.ImageStream{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ImageStream",
			APIVersion: "image.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fis-java-openshift",
			Namespace: fuseOnOpenshiftNs,
			Labels: map[string]string{
				"integreatly": "true",
			},
		},
	}

	cases := []FuseOnOpenShiftScenario{
		{
			Name:           "test error on failed config read",
			ExpectError:    true,
			ExpectedError:  fmt.Sprintf("could not retrieve %[1]s config: could not read %[1]s config", v1alpha1.ProductFuseOnOpenshift),
			ExpectedStatus: v1alpha1.PhaseFailed,
			Installation:   &v1alpha1.Installation{},
			FakeClient:     fakeclient.NewFakeClient(),
			FakeConfig: &config.ConfigReadWriterMock{
				ReadFuseOnOpenshiftFunc: func() (ready *config.FuseOnOpenshift, e error) {
					return nil, errors.Errorf("could not read %s config", v1alpha1.ProductFuseOnOpenshift)
				},
			},
			Product: &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test error on invalid image stream file content",
			ExpectError:    true,
			ExpectedStatus: v1alpha1.PhaseFailed,
			Installation:   &v1alpha1.Installation{},
			FakeClient: fakeclient.NewFakeClient(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      templatesConfigMapName,
					Namespace: OPERATOR_NAMESPACE,
				},
				Data: map[string]string{},
			}),
			FakeConfig: getFakeConfig(),
			Product:    &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test error on invalid image stream",
			ExpectError:    true,
			ExpectedStatus: v1alpha1.PhaseFailed,
			Installation:   &v1alpha1.Installation{},
			FakeClient: fakeclient.NewFakeClient(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      templatesConfigMapName,
					Namespace: OPERATOR_NAMESPACE,
				},
				Data: map[string]string{
					"fis-image-streams.json": `{ "items": [{
						"name": "invalid-image-stream"
					}] }`,
				},
			}),
			FakeConfig: getFakeConfig(),
			Product:    &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test error on invalid template file content",
			ExpectError:    true,
			ExpectedStatus: v1alpha1.PhaseFailed,
			Installation:   &v1alpha1.Installation{},
			FakeClient: fakeclient.NewFakeClient(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      templatesConfigMapName,
					Namespace: OPERATOR_NAMESPACE,
				},
				Data: map[string]string{
					"fis-image-streams.json":        `{ "items": [] }`,
					"fuse-console-cluster-os4.json": "invalid-file-content",
				},
			}),
			FakeConfig: getFakeConfig(),
			Product:    &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test error on invalid template object",
			ExpectError:    true,
			ExpectedStatus: v1alpha1.PhaseFailed,
			Installation:   &v1alpha1.Installation{},
			FakeClient: fakeclient.NewFakeClient(&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      templatesConfigMapName,
					Namespace: OPERATOR_NAMESPACE,
				},
				Data: map[string]string{
					"fis-image-streams.json": `{ "items": [] }`,
					"fuse-console-cluster-os4.json": `{
						"name": "invalid-template"
					}`,
				},
			}),
			FakeConfig: getFakeConfig(),
			Product:    &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test successful reconcile when resource already exists",
			ExpectedStatus: v1alpha1.PhaseCompleted,
			Installation:   &v1alpha1.Installation{},
			FakeClient:     fakeclient.NewFakeClient(integreatlyImgStream),
			FakeConfig:     getFakeConfig(),
			Product:        &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test successful reconcile without sample cluster operator installed",
			ExpectError:    false,
			ExpectedStatus: v1alpha1.PhaseCompleted,
			Installation:   &v1alpha1.Installation{},
			FakeClient:     fakeclient.NewFakeClient(),
			FakeConfig:     getFakeConfig(),
			Product:        &v1alpha1.InstallationProductStatus{},
		},
		{
			Name:           "test successful reconcile with sample cluster operator installed",
			ExpectError:    false,
			ExpectedStatus: v1alpha1.PhaseCompleted,
			Installation:   &v1alpha1.Installation{},
			FakeClient:     fakeclient.NewFakeClient(sampleClusterConfig, sampleClusterImgStream),
			FakeConfig:     getFakeConfig(),
			Product:        &v1alpha1.InstallationProductStatus{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			testReconciler, err := NewReconciler(
				tc.FakeConfig,
				tc.Installation,
				tc.FakeMPM,
			)

			if err != nil && err.Error() != tc.ExpectedError {
				t.Fatalf("unexpected error : '%v', expected: '%v'", err, tc.ExpectedError)
			}

			if err == nil && tc.ExpectedError != "" {
				t.Fatalf("expected error '%v' and got nil", tc.ExpectedError)
			}

			// if we expect errors creating the reconciler, don't try to use it
			if tc.ExpectedError != "" {
				return
			}

			status, err := testReconciler.Reconcile(context.TODO(), tc.Installation, tc.Product, tc.FakeClient)
			if err != nil && !tc.ExpectError {
				t.Fatalf("expected error but got one: %v", err)
			}

			if err == nil && tc.ExpectError {
				t.Fatal("expected error but got none")
			}

			if status != tc.ExpectedStatus {
				t.Fatalf("Expected status: '%v', got: '%v'", tc.ExpectedStatus, status)
			}
		})
	}
}
