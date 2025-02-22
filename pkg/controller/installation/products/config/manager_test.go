package config

import (
	"context"
	"strings"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const (
	mockProductName   = "mock"
	mockConfigMapName = "test"
	mockNamespaceName = "test"
)

func TestWriteConfig(t *testing.T) {
	defaultProductConfig := ProductConfig{"testKey": "testVal"}
	defaultConfigReadable := &ConfigReadableMock{
		GetProductNameFunc: func() v1alpha1.ProductName {
			return mockProductName
		},
		ReadFunc: func() ProductConfig {
			return defaultProductConfig
		},
	}

	fakeInst := &v1alpha1.Installation{}

	tests := []struct {
		productName       string
		existingResources []runtime.Object
		toWrite           ConfigReadable
		expected          ProductConfig
	}{
		// Test basic adding config
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
			}},
			toWrite:  defaultConfigReadable,
			expected: defaultProductConfig,
		},
		// Test overwrite config
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
				Data: map[string]string{
					"testKey1": "testVal1",
					"testKey2": "testVal2",
				},
			}},
			toWrite:  defaultConfigReadable,
			expected: defaultProductConfig,
		},
		// Test create configmap if one doesn't exist
		{
			productName:       mockProductName,
			existingResources: []runtime.Object{},
			toWrite:           defaultConfigReadable,
			expected:          defaultProductConfig,
		},
	}
	for _, test := range tests {
		fakeClient := fake.NewFakeClient(test.existingResources...)

		mgr, err := NewManager(context.TODO(), fakeClient, mockNamespaceName, mockConfigMapName, fakeInst)
		if err != nil {
			t.Fatalf("could not create manager %v", err)
		}
		if err = mgr.WriteConfig(test.toWrite); err != nil {
			t.Fatalf("could not write config %v", err)
		}
		readCfgMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      mockConfigMapName,
				Namespace: mockNamespaceName,
			},
		}
		fakeClient.Get(context.TODO(), client.ObjectKey{Name: mockConfigMapName, Namespace: mockNamespaceName}, readCfgMap)

		decoder := yaml.NewDecoder(strings.NewReader(readCfgMap.Data[test.productName]))
		testCfg := map[string]string{}
		decoder.Decode(testCfg)

		for key, value := range test.expected {
			if strings.Compare(testCfg[key], value) != 0 {
				t.Fatalf("expected %s but got %s for key %s", value, testCfg[key], key)
			}
		}
	}
}

func TestReadConfigForProduct(t *testing.T) {
	fakeInst := &v1alpha1.Installation{}

	tests := []struct {
		productName       string
		existingResources []runtime.Object
		expected          ProductConfig
	}{
		{
			productName: mockProductName,
			existingResources: []runtime.Object{&corev1.ConfigMap{
				ObjectMeta: metav1.ObjectMeta{
					Name:      mockConfigMapName,
					Namespace: mockNamespaceName,
				},
				Data: map[string]string{
					"mock": "testKey: testVal",
				},
			}},
			expected: ProductConfig{"testKey": "testVal"},
		},
		{
			productName:       mockProductName,
			existingResources: []runtime.Object{},
			expected:          map[string]string{},
		},
	}

	for _, test := range tests {
		fakeClient := fake.NewFakeClient(test.existingResources...)
		mgr, err := NewManager(context.TODO(), fakeClient, mockNamespaceName, mockConfigMapName, fakeInst)
		if err != nil {
			t.Fatalf("could not create manager %v", err)
		}
		config, err := mgr.readConfigForProduct(mockProductName)
		if err != nil {
			t.Fatalf("could not read config %v", err)
		}
		for key, value := range test.expected {
			if strings.Compare(config[key], value) != 0 {
				t.Fatalf("expected %s but got %s for key %s", value, config[key], key)
			}
		}
	}

}
