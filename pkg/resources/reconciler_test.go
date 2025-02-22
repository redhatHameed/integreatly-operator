package resources

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/marketplace"
	"github.com/integr8ly/integreatly-operator/pkg/controller/installation/products/config"

	oauthv1 "github.com/openshift/api/oauth/v1"

	alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/ownerutil"
	marketplacev1 "github.com/operator-framework/operator-marketplace/pkg/apis/operators/v1"
	marketplacev2 "github.com/operator-framework/operator-marketplace/pkg/apis/operators/v2"

	corev1 "k8s.io/api/core/v1"
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func basicConfigMock() *config.ConfigReadWriterMock {
	return &config.ConfigReadWriterMock{
		ReadFuseFunc: func() (ready *config.Fuse, e error) {
			return config.NewFuse(config.ProductConfig{}), nil
		},
		ReadRHSSOFunc: func() (*config.RHSSO, error) {
			return config.NewRHSSO(config.ProductConfig{
				"NAMESPACE": "fuse",
				"URL":       "fuse.openshift-cluster.com",
			}), nil
		},
		WriteConfigFunc: func(config config.ConfigReadable) error {
			return nil
		},
	}
}

func buildScheme() (*runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	err := alpha1.AddToScheme(scheme)
	err = oauthv1.AddToScheme(scheme)
	err = marketplacev1.SchemeBuilder.AddToScheme(scheme)
	err = marketplacev2.SchemeBuilder.AddToScheme(scheme)
	err = corev1.SchemeBuilder.AddToScheme(scheme)
	return scheme, err
}

func TestNewReconciler_ReconcileSubscription(t *testing.T) {
	scheme, err := buildScheme()
	if err != nil {
		t.Fatalf("error creating scheme: %s", err.Error())
	}
	ownerInstall := &v1alpha1.Installation{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1alpha1.SchemeGroupVersion.String(),
			Kind:       "Installation",
		},
	}
	catalogSourceConfig := &marketplacev2.CatalogSourceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "installed-integreatly-test-ns",
			Namespace: "openshift-marketplace",
		},
	}
	ownerutil.AddOwner(catalogSourceConfig, ownerInstall, true, true)
	cases := []struct {
		Name             string
		FakeMPM          marketplace.MarketplaceInterface
		client           client.Client
		SubscriptionName string
		ExpectErr        bool
		ExpectedStatus   v1alpha1.StatusPhase
		Installation     *v1alpha1.Installation
		Target           marketplace.Target
		Validate         func(t *testing.T, mock *marketplace.MarketplaceInterfaceMock)
	}{
		{
			Name: "test reconcile subscription creates a new subscription  completes successfully ",
			FakeMPM: &marketplace.MarketplaceInterfaceMock{
				InstallOperatorFunc: func(ctx context.Context, serverClient client.Client, owner ownerutil.Owner, t marketplace.Target, operatorGroupNamespaces []string, approvalStrategy alpha1.Approval) error {

					return nil
				},
				GetSubscriptionInstallPlansFunc: func(ctx context.Context, serverClient client.Client, subName string, ns string) (plans *alpha1.InstallPlanList, subscription *alpha1.Subscription, e error) {
					return &alpha1.InstallPlanList{Items: []alpha1.InstallPlan{alpha1.InstallPlan{Status: alpha1.InstallPlanStatus{Phase: alpha1.InstallPlanPhaseComplete}}}}, &alpha1.Subscription{}, nil
				},
			},
			SubscriptionName: "something",
			ExpectedStatus:   v1alpha1.PhaseCompleted,
			Installation:     &v1alpha1.Installation{},
			Validate: func(t *testing.T, mock *marketplace.MarketplaceInterfaceMock) {
				if len(mock.InstallOperatorCalls()) != 1 {
					t.Fatalf("expected create subscription to be called once but was called %v", len(mock.InstallOperatorCalls()))
				}
				if len(mock.GetSubscriptionInstallPlansCalls()) != 1 {
					t.Fatalf("expected GetSubscriptionInstallPlansCalls to be called once but was called %v", len(mock.GetSubscriptionInstallPlansCalls()))
				}
			},
		},
		{
			Name:   "test reconcile subscription recreates subscription when installation plan not found completes successfully ",
			client: fakeclient.NewFakeClientWithScheme(scheme),
			FakeMPM: &marketplace.MarketplaceInterfaceMock{
				InstallOperatorFunc: func(ctx context.Context, serverClient client.Client, owner ownerutil.Owner, t marketplace.Target, operatorGroupNamespaces []string, approvalStrategy alpha1.Approval) error {

					return nil
				},
				GetSubscriptionInstallPlansFunc: func(ctx context.Context, serverClient client.Client, subName string, ns string) (plans *alpha1.InstallPlanList, subscription *alpha1.Subscription, e error) {
					return nil, &alpha1.Subscription{ObjectMeta: metav1.ObjectMeta{
						// simulate the time has passed
						CreationTimestamp: metav1.Time{Time: time.Now().AddDate(0, 0, -1)},
					}}, k8serr.NewNotFound(alpha1.Resource("installplan"), "my-install-plan")
				},
			},
			SubscriptionName: "something",
			ExpectedStatus:   v1alpha1.PhaseAwaitingOperator,
		},
		{
			Name: "test reconcile subscription returns waiting for operator when catalog source config not ready",
			client: fakeclient.NewFakeClientWithScheme(scheme, catalogSourceConfig, &alpha1.CatalogSourceList{
				Items: []alpha1.CatalogSource{
					alpha1.CatalogSource{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "test",
							Namespace: "test-ns",
						},
					},
				},
			}),
			SubscriptionName: "something",
			ExpectedStatus:   v1alpha1.PhaseFailed,
			FakeMPM:          marketplace.NewManager(),
			Installation:     ownerInstall,
			ExpectErr:        true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			reconciler := NewReconciler(
				tc.FakeMPM,
			)

			status, err := reconciler.ReconcileSubscription(context.TODO(), tc.Installation, marketplace.Target{Namespace: "test-ns", Channel: "integreatly", Pkg: tc.SubscriptionName}, "test-ns", tc.client)
			if tc.ExpectErr && err == nil {
				t.Fatal("expected an error but got none")
			}
			if !tc.ExpectErr && err != nil {
				t.Fatal("expected no error but got one ", err)
			}
			if tc.ExpectedStatus != status {
				t.Fatal("expected phase ", tc.ExpectedStatus, " but got ", status)
			}
			if tc.Validate != nil {
				tc.Validate(t, tc.FakeMPM.(*marketplace.MarketplaceInterfaceMock))
			}

		})
	}
}

func TestReconciler_reconcilePullSecret(t *testing.T) {
	scheme, err := buildScheme()
	if err != nil {
		t.Fatalf("error building scheme: %s", err.Error())
	}

	defPullSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DefaultOriginPullSecretName,
			Namespace: DefaultOriginPullSecretNamespace,
		},
		Data: map[string][]byte{
			"test": {'t', 'e', 's', 't'},
		},
	}

	customPullSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "test",
		},
		Data: map[string][]byte{
			"test": {'t', 'e', 's', 't'},
		},
	}

	cases := []struct {
		Name         string
		Client       client.Client
		Installation *v1alpha1.Installation
		Config       *config.ConfigReadWriterMock
		Validate     func(c client.Client) error
	}{
		{
			Name:   "test default pull secret details are used if not provided",
			Client: fakeclient.NewFakeClientWithScheme(scheme, defPullSecret),
			Installation: &v1alpha1.Installation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testinstall",
					Namespace: "testinstall",
				},
			},
			Config: basicConfigMock(),
			Validate: func(c client.Client) error {
				s := &corev1.Secret{}
				err := c.Get(context.TODO(), client.ObjectKey{Name: DefaultOriginPullSecretName, Namespace: DefaultOriginPullSecretNamespace}, s)
				if err != nil {
					return err
				}
				if bytes.Compare(s.Data["test"], customPullSecret.Data["test"]) != 0 {
					return errors.New(fmt.Sprintf("expected data %v, but got %v", customPullSecret.Data["test"], s.Data["test"]))
				}
				return nil
			},
		},
		{
			Name:   "test pull secret is reconciled successfully",
			Client: fakeclient.NewFakeClientWithScheme(scheme, defPullSecret, customPullSecret),
			Installation: &v1alpha1.Installation{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "testinstall",
					Namespace: "testinstall",
				},
				Spec: v1alpha1.InstallationSpec{
					PullSecret: v1alpha1.PullSecretSpec{
						Name:      "test",
						Namespace: "test",
					},
				},
			},
			Config: basicConfigMock(),
			Validate: func(c client.Client) error {
				s := &corev1.Secret{}
				err := c.Get(context.TODO(), client.ObjectKey{Name: "test", Namespace: "test"}, s)
				if err != nil {
					return err
				}
				if bytes.Compare(s.Data["test"], customPullSecret.Data["test"]) != 0 {
					return errors.New(fmt.Sprintf("expected data %v, but got %v", customPullSecret.Data["test"], s.Data["test"]))
				}
				return nil
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			testReconciler := NewReconciler(nil)
			_, err := testReconciler.ReconcilePullSecret(context.TODO(), "test", "new-pull-secret-name", tc.Installation, tc.Client)
			if err != nil {
				t.Fatal("failed to run pull secret reconcile: ", err)
			}
			if err = tc.Validate(tc.Client); err != nil {
				t.Fatal("test validation failed: ", err)
			}
		})
	}
}

func TestReconciler_ReconcileOauthClient(t *testing.T) {
	scheme, err := buildScheme()
	if err != nil {
		t.Fatalf("error building scheme: %s", err.Error())
	}
	existingClient := &oauthv1.OAuthClient{
		GrantMethod:  oauthv1.GrantHandlerAuto,
		Secret:       "test",
		RedirectURIs: []string{"http://test.com"},
	}
	cases := []struct {
		Name           string
		OauthClient    *oauthv1.OAuthClient
		ExpectErr      bool
		ExpectedStatus v1alpha1.StatusPhase
		client         client.Client
		Installation   *v1alpha1.Installation
	}{
		{
			Name: "test oauth client is reconciled correctly when it does not exist",
			OauthClient: &oauthv1.OAuthClient{
				GrantMethod:  oauthv1.GrantHandlerAuto,
				Secret:       "test",
				RedirectURIs: []string{"http://test.com"},
			},
			Installation: &v1alpha1.Installation{
				TypeMeta: metav1.TypeMeta{
					APIVersion: v1alpha1.SchemeGroupVersion.String(),
					Kind:       "Installation",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-install",
				},
			},
			client:         fakeclient.NewFakeClientWithScheme(scheme),
			ExpectedStatus: v1alpha1.PhaseCompleted,
		},
		{
			Name:        "test oauth client is reconciled correctly when it does exist",
			OauthClient: existingClient,
			Installation: &v1alpha1.Installation{
				TypeMeta: metav1.TypeMeta{
					APIVersion: v1alpha1.SchemeGroupVersion.String(),
					Kind:       "Installation",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-install",
				},
			},
			client:         fakeclient.NewFakeClientWithScheme(scheme, existingClient),
			ExpectedStatus: v1alpha1.PhaseCompleted,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			reconciler := NewReconciler(nil)
			phase, err := reconciler.ReconcileOauthClient(context.TODO(), tc.Installation, tc.OauthClient, tc.client)
			if tc.ExpectErr && err == nil {
				t.Fatal("expected an error but got none")
			}
			if !tc.ExpectErr && err != nil {
				t.Fatal("expected no error but got one ", err)
			}
			if tc.ExpectedStatus != phase {
				t.Fatal("expected phase ", tc.ExpectedStatus, " but got ", phase)
			}
		})
	}
}

func TestReconciler_ReconcileNamespace(t *testing.T) {
	nsName := "test-ns"
	installation := &v1alpha1.Installation{
		ObjectMeta: metav1.ObjectMeta{
			Name: "install",
			UID:  types.UID("xyz"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1alpha1.SchemeGroupVersion.String(),
		},
	}
	cases := []struct {
		Name           string
		client         client.Client
		Installation   *v1alpha1.Installation
		ExpectErr      bool
		ExpectedStatus v1alpha1.StatusPhase
		FakeMPM        *marketplace.MarketplaceInterfaceMock
	}{
		{
			Name: "Test namespace reconcile completes without error",
			client: fakeclient.NewFakeClient(&corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: nsName,
					Labels: map[string]string{
						OwnerLabelKey: string(installation.GetUID()),
					},
				},
				Status: corev1.NamespaceStatus{
					Phase: corev1.NamespaceActive,
				},
			}),
			Installation:   installation,
			ExpectedStatus: v1alpha1.PhaseCompleted,
		},
		{
			Name: "Test namespace reconcile returns waiting when ns not ready",
			client: fakeclient.NewFakeClient(&corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: nsName,
					Labels: map[string]string{
						OwnerLabelKey: string(installation.GetUID()),
					},
				},
				Status: corev1.NamespaceStatus{},
			}),
			Installation:   installation,
			ExpectedStatus: v1alpha1.PhaseInProgress,
		},
		{
			Name: "Test namespace reconcile returns waiting when ns is terminating",
			client: fakeclient.NewFakeClient(&corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: nsName,
					Labels: map[string]string{
						OwnerLabelKey: string(installation.GetUID()),
					},
				},
				Status: corev1.NamespaceStatus{
					Phase: corev1.NamespaceTerminating,
				},
			}),
			Installation:   &v1alpha1.Installation{},
			ExpectedStatus: v1alpha1.PhaseInProgress,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			reconciler := NewReconciler(
				tc.FakeMPM,
			)
			phase, err := reconciler.ReconcileNamespace(context.TODO(), "test-ns", tc.Installation, tc.client)
			if tc.ExpectErr && err == nil {
				t.Fatal("expected an error but got none")
			}
			if !tc.ExpectErr && err != nil {
				t.Fatal("expected no error but got one ", err)
			}
			if tc.ExpectedStatus != phase {
				t.Fatal("expected phase ", tc.ExpectedStatus, " but got ", phase)
			}
		})
	}
}
