// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"sync"

	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
)

var (
	lockConfigReadWriterMockGetOauthClientsSecretName sync.RWMutex
	lockConfigReadWriterMockGetOperatorNamespace      sync.RWMutex
	lockConfigReadWriterMockReadAMQOnline             sync.RWMutex
	lockConfigReadWriterMockReadAMQStreams            sync.RWMutex
	lockConfigReadWriterMockReadCloudResources        sync.RWMutex
	lockConfigReadWriterMockReadCodeReady             sync.RWMutex
	lockConfigReadWriterMockReadFuse                  sync.RWMutex
	lockConfigReadWriterMockReadFuseOnOpenshift       sync.RWMutex
	lockConfigReadWriterMockReadMonitoring            sync.RWMutex
	lockConfigReadWriterMockReadProduct               sync.RWMutex
	lockConfigReadWriterMockReadRHSSO                 sync.RWMutex
	lockConfigReadWriterMockReadRHSSOUser             sync.RWMutex
	lockConfigReadWriterMockReadSolutionExplorer      sync.RWMutex
	lockConfigReadWriterMockReadThreeScale            sync.RWMutex
	lockConfigReadWriterMockReadUps                   sync.RWMutex
	lockConfigReadWriterMockWriteConfig               sync.RWMutex
	lockConfigReadWriterMockreadConfigForProduct      sync.RWMutex
)

// Ensure, that ConfigReadWriterMock does implement ConfigReadWriter.
// If this is not the case, regenerate this file with moq.
var _ ConfigReadWriter = &ConfigReadWriterMock{}

// ConfigReadWriterMock is a mock implementation of ConfigReadWriter.
//
//     func TestSomethingThatUsesConfigReadWriter(t *testing.T) {
//
//         // make and configure a mocked ConfigReadWriter
//         mockedConfigReadWriter := &ConfigReadWriterMock{
//             GetOauthClientsSecretNameFunc: func() string {
// 	               panic("mock out the GetOauthClientsSecretName method")
//             },
//             GetOperatorNamespaceFunc: func() string {
// 	               panic("mock out the GetOperatorNamespace method")
//             },
//             ReadAMQOnlineFunc: func() (*AMQOnline, error) {
// 	               panic("mock out the ReadAMQOnline method")
//             },
//             ReadAMQStreamsFunc: func() (*AMQStreams, error) {
// 	               panic("mock out the ReadAMQStreams method")
//             },
//             ReadCloudResourcesFunc: func() (*CloudResources, error) {
// 	               panic("mock out the ReadCloudResources method")
//             },
//             ReadCodeReadyFunc: func() (*CodeReady, error) {
// 	               panic("mock out the ReadCodeReady method")
//             },
//             ReadFuseFunc: func() (*Fuse, error) {
// 	               panic("mock out the ReadFuse method")
//             },
//             ReadFuseOnOpenshiftFunc: func() (*FuseOnOpenshift, error) {
// 	               panic("mock out the ReadFuseOnOpenshift method")
//             },
//             ReadMonitoringFunc: func() (*Monitoring, error) {
// 	               panic("mock out the ReadMonitoring method")
//             },
//             ReadProductFunc: func(product v1alpha1.ProductName) (ConfigReadable, error) {
// 	               panic("mock out the ReadProduct method")
//             },
//             ReadRHSSOFunc: func() (*RHSSO, error) {
// 	               panic("mock out the ReadRHSSO method")
//             },
//             ReadRHSSOUserFunc: func() (*RHSSOUser, error) {
// 	               panic("mock out the ReadRHSSOUser method")
//             },
//             ReadSolutionExplorerFunc: func() (*SolutionExplorer, error) {
// 	               panic("mock out the ReadSolutionExplorer method")
//             },
//             ReadThreeScaleFunc: func() (*ThreeScale, error) {
// 	               panic("mock out the ReadThreeScale method")
//             },
//             ReadUpsFunc: func() (*Ups, error) {
// 	               panic("mock out the ReadUps method")
//             },
//             WriteConfigFunc: func(config ConfigReadable) error {
// 	               panic("mock out the WriteConfig method")
//             },
//             readConfigForProductFunc: func(product v1alpha1.ProductName) (ProductConfig, error) {
// 	               panic("mock out the readConfigForProduct method")
//             },
//         }
//
//         // use mockedConfigReadWriter in code that requires ConfigReadWriter
//         // and then make assertions.
//
//     }
type ConfigReadWriterMock struct {
	// GetOauthClientsSecretNameFunc mocks the GetOauthClientsSecretName method.
	GetOauthClientsSecretNameFunc func() string

	// GetOperatorNamespaceFunc mocks the GetOperatorNamespace method.
	GetOperatorNamespaceFunc func() string

	// ReadAMQOnlineFunc mocks the ReadAMQOnline method.
	ReadAMQOnlineFunc func() (*AMQOnline, error)

	// ReadAMQStreamsFunc mocks the ReadAMQStreams method.
	ReadAMQStreamsFunc func() (*AMQStreams, error)

	// ReadCloudResourcesFunc mocks the ReadCloudResources method.
	ReadCloudResourcesFunc func() (*CloudResources, error)

	// ReadCodeReadyFunc mocks the ReadCodeReady method.
	ReadCodeReadyFunc func() (*CodeReady, error)

	// ReadFuseFunc mocks the ReadFuse method.
	ReadFuseFunc func() (*Fuse, error)

	// ReadFuseOnOpenshiftFunc mocks the ReadFuseOnOpenshift method.
	ReadFuseOnOpenshiftFunc func() (*FuseOnOpenshift, error)

	// ReadMonitoringFunc mocks the ReadMonitoring method.
	ReadMonitoringFunc func() (*Monitoring, error)

	// ReadProductFunc mocks the ReadProduct method.
	ReadProductFunc func(product v1alpha1.ProductName) (ConfigReadable, error)

	// ReadRHSSOFunc mocks the ReadRHSSO method.
	ReadRHSSOFunc func() (*RHSSO, error)

	// ReadRHSSOUserFunc mocks the ReadRHSSOUser method.
	ReadRHSSOUserFunc func() (*RHSSOUser, error)

	// ReadSolutionExplorerFunc mocks the ReadSolutionExplorer method.
	ReadSolutionExplorerFunc func() (*SolutionExplorer, error)

	// ReadThreeScaleFunc mocks the ReadThreeScale method.
	ReadThreeScaleFunc func() (*ThreeScale, error)

	// ReadUpsFunc mocks the ReadUps method.
	ReadUpsFunc func() (*Ups, error)

	// WriteConfigFunc mocks the WriteConfig method.
	WriteConfigFunc func(config ConfigReadable) error

	// readConfigForProductFunc mocks the readConfigForProduct method.
	readConfigForProductFunc func(product v1alpha1.ProductName) (ProductConfig, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetOauthClientsSecretName holds details about calls to the GetOauthClientsSecretName method.
		GetOauthClientsSecretName []struct {
		}
		// GetOperatorNamespace holds details about calls to the GetOperatorNamespace method.
		GetOperatorNamespace []struct {
		}
		// ReadAMQOnline holds details about calls to the ReadAMQOnline method.
		ReadAMQOnline []struct {
		}
		// ReadAMQStreams holds details about calls to the ReadAMQStreams method.
		ReadAMQStreams []struct {
		}
		// ReadCloudResources holds details about calls to the ReadCloudResources method.
		ReadCloudResources []struct {
		}
		// ReadCodeReady holds details about calls to the ReadCodeReady method.
		ReadCodeReady []struct {
		}
		// ReadFuse holds details about calls to the ReadFuse method.
		ReadFuse []struct {
		}
		// ReadFuseOnOpenshift holds details about calls to the ReadFuseOnOpenshift method.
		ReadFuseOnOpenshift []struct {
		}
		// ReadMonitoring holds details about calls to the ReadMonitoring method.
		ReadMonitoring []struct {
		}
		// ReadProduct holds details about calls to the ReadProduct method.
		ReadProduct []struct {
			// Product is the product argument value.
			Product v1alpha1.ProductName
		}
		// ReadRHSSO holds details about calls to the ReadRHSSO method.
		ReadRHSSO []struct {
		}
		// ReadRHSSOUser holds details about calls to the ReadRHSSOUser method.
		ReadRHSSOUser []struct {
		}
		// ReadSolutionExplorer holds details about calls to the ReadSolutionExplorer method.
		ReadSolutionExplorer []struct {
		}
		// ReadThreeScale holds details about calls to the ReadThreeScale method.
		ReadThreeScale []struct {
		}
		// ReadUps holds details about calls to the ReadUps method.
		ReadUps []struct {
		}
		// WriteConfig holds details about calls to the WriteConfig method.
		WriteConfig []struct {
			// Config is the config argument value.
			Config ConfigReadable
		}
		// readConfigForProduct holds details about calls to the readConfigForProduct method.
		readConfigForProduct []struct {
			// Product is the product argument value.
			Product v1alpha1.ProductName
		}
	}
}

// GetOauthClientsSecretName calls GetOauthClientsSecretNameFunc.
func (mock *ConfigReadWriterMock) GetOauthClientsSecretName() string {
	if mock.GetOauthClientsSecretNameFunc == nil {
		panic("ConfigReadWriterMock.GetOauthClientsSecretNameFunc: method is nil but ConfigReadWriter.GetOauthClientsSecretName was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockGetOauthClientsSecretName.Lock()
	mock.calls.GetOauthClientsSecretName = append(mock.calls.GetOauthClientsSecretName, callInfo)
	lockConfigReadWriterMockGetOauthClientsSecretName.Unlock()
	return mock.GetOauthClientsSecretNameFunc()
}

// GetOauthClientsSecretNameCalls gets all the calls that were made to GetOauthClientsSecretName.
// Check the length with:
//     len(mockedConfigReadWriter.GetOauthClientsSecretNameCalls())
func (mock *ConfigReadWriterMock) GetOauthClientsSecretNameCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockGetOauthClientsSecretName.RLock()
	calls = mock.calls.GetOauthClientsSecretName
	lockConfigReadWriterMockGetOauthClientsSecretName.RUnlock()
	return calls
}

// GetOperatorNamespace calls GetOperatorNamespaceFunc.
func (mock *ConfigReadWriterMock) GetOperatorNamespace() string {
	if mock.GetOperatorNamespaceFunc == nil {
		panic("ConfigReadWriterMock.GetOperatorNamespaceFunc: method is nil but ConfigReadWriter.GetOperatorNamespace was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockGetOperatorNamespace.Lock()
	mock.calls.GetOperatorNamespace = append(mock.calls.GetOperatorNamespace, callInfo)
	lockConfigReadWriterMockGetOperatorNamespace.Unlock()
	return mock.GetOperatorNamespaceFunc()
}

// GetOperatorNamespaceCalls gets all the calls that were made to GetOperatorNamespace.
// Check the length with:
//     len(mockedConfigReadWriter.GetOperatorNamespaceCalls())
func (mock *ConfigReadWriterMock) GetOperatorNamespaceCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockGetOperatorNamespace.RLock()
	calls = mock.calls.GetOperatorNamespace
	lockConfigReadWriterMockGetOperatorNamespace.RUnlock()
	return calls
}

// ReadAMQOnline calls ReadAMQOnlineFunc.
func (mock *ConfigReadWriterMock) ReadAMQOnline() (*AMQOnline, error) {
	if mock.ReadAMQOnlineFunc == nil {
		panic("ConfigReadWriterMock.ReadAMQOnlineFunc: method is nil but ConfigReadWriter.ReadAMQOnline was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadAMQOnline.Lock()
	mock.calls.ReadAMQOnline = append(mock.calls.ReadAMQOnline, callInfo)
	lockConfigReadWriterMockReadAMQOnline.Unlock()
	return mock.ReadAMQOnlineFunc()
}

// ReadAMQOnlineCalls gets all the calls that were made to ReadAMQOnline.
// Check the length with:
//     len(mockedConfigReadWriter.ReadAMQOnlineCalls())
func (mock *ConfigReadWriterMock) ReadAMQOnlineCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadAMQOnline.RLock()
	calls = mock.calls.ReadAMQOnline
	lockConfigReadWriterMockReadAMQOnline.RUnlock()
	return calls
}

// ReadAMQStreams calls ReadAMQStreamsFunc.
func (mock *ConfigReadWriterMock) ReadAMQStreams() (*AMQStreams, error) {
	if mock.ReadAMQStreamsFunc == nil {
		panic("ConfigReadWriterMock.ReadAMQStreamsFunc: method is nil but ConfigReadWriter.ReadAMQStreams was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadAMQStreams.Lock()
	mock.calls.ReadAMQStreams = append(mock.calls.ReadAMQStreams, callInfo)
	lockConfigReadWriterMockReadAMQStreams.Unlock()
	return mock.ReadAMQStreamsFunc()
}

// ReadAMQStreamsCalls gets all the calls that were made to ReadAMQStreams.
// Check the length with:
//     len(mockedConfigReadWriter.ReadAMQStreamsCalls())
func (mock *ConfigReadWriterMock) ReadAMQStreamsCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadAMQStreams.RLock()
	calls = mock.calls.ReadAMQStreams
	lockConfigReadWriterMockReadAMQStreams.RUnlock()
	return calls
}

// ReadCloudResources calls ReadCloudResourcesFunc.
func (mock *ConfigReadWriterMock) ReadCloudResources() (*CloudResources, error) {
	if mock.ReadCloudResourcesFunc == nil {
		panic("ConfigReadWriterMock.ReadCloudResourcesFunc: method is nil but ConfigReadWriter.ReadCloudResources was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadCloudResources.Lock()
	mock.calls.ReadCloudResources = append(mock.calls.ReadCloudResources, callInfo)
	lockConfigReadWriterMockReadCloudResources.Unlock()
	return mock.ReadCloudResourcesFunc()
}

// ReadCloudResourcesCalls gets all the calls that were made to ReadCloudResources.
// Check the length with:
//     len(mockedConfigReadWriter.ReadCloudResourcesCalls())
func (mock *ConfigReadWriterMock) ReadCloudResourcesCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadCloudResources.RLock()
	calls = mock.calls.ReadCloudResources
	lockConfigReadWriterMockReadCloudResources.RUnlock()
	return calls
}

// ReadCodeReady calls ReadCodeReadyFunc.
func (mock *ConfigReadWriterMock) ReadCodeReady() (*CodeReady, error) {
	if mock.ReadCodeReadyFunc == nil {
		panic("ConfigReadWriterMock.ReadCodeReadyFunc: method is nil but ConfigReadWriter.ReadCodeReady was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadCodeReady.Lock()
	mock.calls.ReadCodeReady = append(mock.calls.ReadCodeReady, callInfo)
	lockConfigReadWriterMockReadCodeReady.Unlock()
	return mock.ReadCodeReadyFunc()
}

// ReadCodeReadyCalls gets all the calls that were made to ReadCodeReady.
// Check the length with:
//     len(mockedConfigReadWriter.ReadCodeReadyCalls())
func (mock *ConfigReadWriterMock) ReadCodeReadyCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadCodeReady.RLock()
	calls = mock.calls.ReadCodeReady
	lockConfigReadWriterMockReadCodeReady.RUnlock()
	return calls
}

// ReadFuse calls ReadFuseFunc.
func (mock *ConfigReadWriterMock) ReadFuse() (*Fuse, error) {
	if mock.ReadFuseFunc == nil {
		panic("ConfigReadWriterMock.ReadFuseFunc: method is nil but ConfigReadWriter.ReadFuse was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadFuse.Lock()
	mock.calls.ReadFuse = append(mock.calls.ReadFuse, callInfo)
	lockConfigReadWriterMockReadFuse.Unlock()
	return mock.ReadFuseFunc()
}

// ReadFuseCalls gets all the calls that were made to ReadFuse.
// Check the length with:
//     len(mockedConfigReadWriter.ReadFuseCalls())
func (mock *ConfigReadWriterMock) ReadFuseCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadFuse.RLock()
	calls = mock.calls.ReadFuse
	lockConfigReadWriterMockReadFuse.RUnlock()
	return calls
}

// ReadFuseOnOpenshift calls ReadFuseOnOpenshiftFunc.
func (mock *ConfigReadWriterMock) ReadFuseOnOpenshift() (*FuseOnOpenshift, error) {
	if mock.ReadFuseOnOpenshiftFunc == nil {
		panic("ConfigReadWriterMock.ReadFuseOnOpenshiftFunc: method is nil but ConfigReadWriter.ReadFuseOnOpenshift was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadFuseOnOpenshift.Lock()
	mock.calls.ReadFuseOnOpenshift = append(mock.calls.ReadFuseOnOpenshift, callInfo)
	lockConfigReadWriterMockReadFuseOnOpenshift.Unlock()
	return mock.ReadFuseOnOpenshiftFunc()
}

// ReadFuseOnOpenshiftCalls gets all the calls that were made to ReadFuseOnOpenshift.
// Check the length with:
//     len(mockedConfigReadWriter.ReadFuseOnOpenshiftCalls())
func (mock *ConfigReadWriterMock) ReadFuseOnOpenshiftCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadFuseOnOpenshift.RLock()
	calls = mock.calls.ReadFuseOnOpenshift
	lockConfigReadWriterMockReadFuseOnOpenshift.RUnlock()
	return calls
}

// ReadMonitoring calls ReadMonitoringFunc.
func (mock *ConfigReadWriterMock) ReadMonitoring() (*Monitoring, error) {
	if mock.ReadMonitoringFunc == nil {
		panic("ConfigReadWriterMock.ReadMonitoringFunc: method is nil but ConfigReadWriter.ReadMonitoring was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadMonitoring.Lock()
	mock.calls.ReadMonitoring = append(mock.calls.ReadMonitoring, callInfo)
	lockConfigReadWriterMockReadMonitoring.Unlock()
	return mock.ReadMonitoringFunc()
}

// ReadMonitoringCalls gets all the calls that were made to ReadMonitoring.
// Check the length with:
//     len(mockedConfigReadWriter.ReadMonitoringCalls())
func (mock *ConfigReadWriterMock) ReadMonitoringCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadMonitoring.RLock()
	calls = mock.calls.ReadMonitoring
	lockConfigReadWriterMockReadMonitoring.RUnlock()
	return calls
}

// ReadProduct calls ReadProductFunc.
func (mock *ConfigReadWriterMock) ReadProduct(product v1alpha1.ProductName) (ConfigReadable, error) {
	if mock.ReadProductFunc == nil {
		panic("ConfigReadWriterMock.ReadProductFunc: method is nil but ConfigReadWriter.ReadProduct was just called")
	}
	callInfo := struct {
		Product v1alpha1.ProductName
	}{
		Product: product,
	}
	lockConfigReadWriterMockReadProduct.Lock()
	mock.calls.ReadProduct = append(mock.calls.ReadProduct, callInfo)
	lockConfigReadWriterMockReadProduct.Unlock()
	return mock.ReadProductFunc(product)
}

// ReadProductCalls gets all the calls that were made to ReadProduct.
// Check the length with:
//     len(mockedConfigReadWriter.ReadProductCalls())
func (mock *ConfigReadWriterMock) ReadProductCalls() []struct {
	Product v1alpha1.ProductName
} {
	var calls []struct {
		Product v1alpha1.ProductName
	}
	lockConfigReadWriterMockReadProduct.RLock()
	calls = mock.calls.ReadProduct
	lockConfigReadWriterMockReadProduct.RUnlock()
	return calls
}

// ReadRHSSO calls ReadRHSSOFunc.
func (mock *ConfigReadWriterMock) ReadRHSSO() (*RHSSO, error) {
	if mock.ReadRHSSOFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOFunc: method is nil but ConfigReadWriter.ReadRHSSO was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadRHSSO.Lock()
	mock.calls.ReadRHSSO = append(mock.calls.ReadRHSSO, callInfo)
	lockConfigReadWriterMockReadRHSSO.Unlock()
	return mock.ReadRHSSOFunc()
}

// ReadRHSSOCalls gets all the calls that were made to ReadRHSSO.
// Check the length with:
//     len(mockedConfigReadWriter.ReadRHSSOCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadRHSSO.RLock()
	calls = mock.calls.ReadRHSSO
	lockConfigReadWriterMockReadRHSSO.RUnlock()
	return calls
}

// ReadRHSSOUser calls ReadRHSSOUserFunc.
func (mock *ConfigReadWriterMock) ReadRHSSOUser() (*RHSSOUser, error) {
	if mock.ReadRHSSOUserFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOUserFunc: method is nil but ConfigReadWriter.ReadRHSSOUser was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadRHSSOUser.Lock()
	mock.calls.ReadRHSSOUser = append(mock.calls.ReadRHSSOUser, callInfo)
	lockConfigReadWriterMockReadRHSSOUser.Unlock()
	return mock.ReadRHSSOUserFunc()
}

// ReadRHSSOUserCalls gets all the calls that were made to ReadRHSSOUser.
// Check the length with:
//     len(mockedConfigReadWriter.ReadRHSSOUserCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOUserCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadRHSSOUser.RLock()
	calls = mock.calls.ReadRHSSOUser
	lockConfigReadWriterMockReadRHSSOUser.RUnlock()
	return calls
}

// ReadSolutionExplorer calls ReadSolutionExplorerFunc.
func (mock *ConfigReadWriterMock) ReadSolutionExplorer() (*SolutionExplorer, error) {
	if mock.ReadSolutionExplorerFunc == nil {
		panic("ConfigReadWriterMock.ReadSolutionExplorerFunc: method is nil but ConfigReadWriter.ReadSolutionExplorer was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadSolutionExplorer.Lock()
	mock.calls.ReadSolutionExplorer = append(mock.calls.ReadSolutionExplorer, callInfo)
	lockConfigReadWriterMockReadSolutionExplorer.Unlock()
	return mock.ReadSolutionExplorerFunc()
}

// ReadSolutionExplorerCalls gets all the calls that were made to ReadSolutionExplorer.
// Check the length with:
//     len(mockedConfigReadWriter.ReadSolutionExplorerCalls())
func (mock *ConfigReadWriterMock) ReadSolutionExplorerCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadSolutionExplorer.RLock()
	calls = mock.calls.ReadSolutionExplorer
	lockConfigReadWriterMockReadSolutionExplorer.RUnlock()
	return calls
}

// ReadThreeScale calls ReadThreeScaleFunc.
func (mock *ConfigReadWriterMock) ReadThreeScale() (*ThreeScale, error) {
	if mock.ReadThreeScaleFunc == nil {
		panic("ConfigReadWriterMock.ReadThreeScaleFunc: method is nil but ConfigReadWriter.ReadThreeScale was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadThreeScale.Lock()
	mock.calls.ReadThreeScale = append(mock.calls.ReadThreeScale, callInfo)
	lockConfigReadWriterMockReadThreeScale.Unlock()
	return mock.ReadThreeScaleFunc()
}

// ReadThreeScaleCalls gets all the calls that were made to ReadThreeScale.
// Check the length with:
//     len(mockedConfigReadWriter.ReadThreeScaleCalls())
func (mock *ConfigReadWriterMock) ReadThreeScaleCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadThreeScale.RLock()
	calls = mock.calls.ReadThreeScale
	lockConfigReadWriterMockReadThreeScale.RUnlock()
	return calls
}

// ReadUps calls ReadUpsFunc.
func (mock *ConfigReadWriterMock) ReadUps() (*Ups, error) {
	if mock.ReadUpsFunc == nil {
		panic("ConfigReadWriterMock.ReadUpsFunc: method is nil but ConfigReadWriter.ReadUps was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadUps.Lock()
	mock.calls.ReadUps = append(mock.calls.ReadUps, callInfo)
	lockConfigReadWriterMockReadUps.Unlock()
	return mock.ReadUpsFunc()
}

// ReadUpsCalls gets all the calls that were made to ReadUps.
// Check the length with:
//     len(mockedConfigReadWriter.ReadUpsCalls())
func (mock *ConfigReadWriterMock) ReadUpsCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadUps.RLock()
	calls = mock.calls.ReadUps
	lockConfigReadWriterMockReadUps.RUnlock()
	return calls
}

// WriteConfig calls WriteConfigFunc.
func (mock *ConfigReadWriterMock) WriteConfig(config ConfigReadable) error {
	if mock.WriteConfigFunc == nil {
		panic("ConfigReadWriterMock.WriteConfigFunc: method is nil but ConfigReadWriter.WriteConfig was just called")
	}
	callInfo := struct {
		Config ConfigReadable
	}{
		Config: config,
	}
	lockConfigReadWriterMockWriteConfig.Lock()
	mock.calls.WriteConfig = append(mock.calls.WriteConfig, callInfo)
	lockConfigReadWriterMockWriteConfig.Unlock()
	return mock.WriteConfigFunc(config)
}

// WriteConfigCalls gets all the calls that were made to WriteConfig.
// Check the length with:
//     len(mockedConfigReadWriter.WriteConfigCalls())
func (mock *ConfigReadWriterMock) WriteConfigCalls() []struct {
	Config ConfigReadable
} {
	var calls []struct {
		Config ConfigReadable
	}
	lockConfigReadWriterMockWriteConfig.RLock()
	calls = mock.calls.WriteConfig
	lockConfigReadWriterMockWriteConfig.RUnlock()
	return calls
}

// readConfigForProduct calls readConfigForProductFunc.
func (mock *ConfigReadWriterMock) readConfigForProduct(product v1alpha1.ProductName) (ProductConfig, error) {
	if mock.readConfigForProductFunc == nil {
		panic("ConfigReadWriterMock.readConfigForProductFunc: method is nil but ConfigReadWriter.readConfigForProduct was just called")
	}
	callInfo := struct {
		Product v1alpha1.ProductName
	}{
		Product: product,
	}
	lockConfigReadWriterMockreadConfigForProduct.Lock()
	mock.calls.readConfigForProduct = append(mock.calls.readConfigForProduct, callInfo)
	lockConfigReadWriterMockreadConfigForProduct.Unlock()
	return mock.readConfigForProductFunc(product)
}

// readConfigForProductCalls gets all the calls that were made to readConfigForProduct.
// Check the length with:
//     len(mockedConfigReadWriter.readConfigForProductCalls())
func (mock *ConfigReadWriterMock) readConfigForProductCalls() []struct {
	Product v1alpha1.ProductName
} {
	var calls []struct {
		Product v1alpha1.ProductName
	}
	lockConfigReadWriterMockreadConfigForProduct.RLock()
	calls = mock.calls.readConfigForProduct
	lockConfigReadWriterMockreadConfigForProduct.RUnlock()
	return calls
}
