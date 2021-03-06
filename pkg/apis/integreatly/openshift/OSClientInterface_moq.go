// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package openshift

import (
	appsv1 "github.com/openshift/api/apps/v1"
	tmplv1 "github.com/openshift/api/template/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sync"
)

var (
	lockOSClientInterfaceMockDelete          sync.RWMutex
	lockOSClientInterfaceMockGetDC           sync.RWMutex
	lockOSClientInterfaceMockGetPod          sync.RWMutex
	lockOSClientInterfaceMockProcessTemplate sync.RWMutex
	lockOSClientInterfaceMockUpdateDC        sync.RWMutex
)

// Ensure, that OSClientInterfaceMock does implement OSClientInterface.
// If this is not the case, regenerate this file with moq.
var _ OSClientInterface = &OSClientInterfaceMock{}

// OSClientInterfaceMock is a mock implementation of OSClientInterface.
//
//     func TestSomethingThatUsesOSClientInterface(t *testing.T) {
//
//         // make and configure a mocked OSClientInterface
//         mockedOSClientInterface := &OSClientInterfaceMock{
//             DeleteFunc: func(ns string, label string) error {
// 	               panic("mock out the Delete method")
//             },
//             GetDCFunc: func(ns string, dcName string) (v1.DeploymentConfig, error) {
// 	               panic("mock out the GetDC method")
//             },
//             GetPodFunc: func(ns string, dc string) (v1.Pod, error) {
// 	               panic("mock out the GetPod method")
//             },
//             ProcessTemplateFunc: func(in1 *v1.Template, in2 map[string]string, in3 TemplateOpt) ([]runtime.RawExtension, error) {
// 	               panic("mock out the ProcessTemplate method")
//             },
//             UpdateDCFunc: func(ns string, dc *v1.DeploymentConfig) error {
// 	               panic("mock out the UpdateDC method")
//             },
//         }
//
//         // use mockedOSClientInterface in code that requires OSClientInterface
//         // and then make assertions.
//
//     }
type OSClientInterfaceMock struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ns string, label string) error

	// GetDCFunc mocks the GetDC method.
	GetDCFunc func(ns string, dcName string) (appsv1.DeploymentConfig, error)

	// GetPodFunc mocks the GetPod method.
	GetPodFunc func(ns string, dc string) (v1.Pod, error)

	// ProcessTemplateFunc mocks the ProcessTemplate method.
	ProcessTemplateFunc func(in1 *tmplv1.Template, in2 map[string]string, in3 TemplateOpt) ([]runtime.RawExtension, error)

	// UpdateDCFunc mocks the UpdateDC method.
	UpdateDCFunc func(ns string, dc *appsv1.DeploymentConfig) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ns is the ns argument value.
			Ns string
			// Label is the label argument value.
			Label string
		}
		// GetDC holds details about calls to the GetDC method.
		GetDC []struct {
			// Ns is the ns argument value.
			Ns string
			// DcName is the dcName argument value.
			DcName string
		}
		// GetPod holds details about calls to the GetPod method.
		GetPod []struct {
			// Ns is the ns argument value.
			Ns string
			// Dc is the dc argument value.
			Dc string
		}
		// ProcessTemplate holds details about calls to the ProcessTemplate method.
		ProcessTemplate []struct {
			// In1 is the in1 argument value.
			In1 *tmplv1.Template
			// In2 is the in2 argument value.
			In2 map[string]string
			// In3 is the in3 argument value.
			In3 TemplateOpt
		}
		// UpdateDC holds details about calls to the UpdateDC method.
		UpdateDC []struct {
			// Ns is the ns argument value.
			Ns string
			// Dc is the dc argument value.
			Dc *appsv1.DeploymentConfig
		}
	}
}

// Delete calls DeleteFunc.
func (mock *OSClientInterfaceMock) Delete(ns string, label string) error {
	if mock.DeleteFunc == nil {
		panic("OSClientInterfaceMock.DeleteFunc: method is nil but OSClientInterface.Delete was just called")
	}
	callInfo := struct {
		Ns    string
		Label string
	}{
		Ns:    ns,
		Label: label,
	}
	lockOSClientInterfaceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockOSClientInterfaceMockDelete.Unlock()
	return mock.DeleteFunc(ns, label)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedOSClientInterface.DeleteCalls())
func (mock *OSClientInterfaceMock) DeleteCalls() []struct {
	Ns    string
	Label string
} {
	var calls []struct {
		Ns    string
		Label string
	}
	lockOSClientInterfaceMockDelete.RLock()
	calls = mock.calls.Delete
	lockOSClientInterfaceMockDelete.RUnlock()
	return calls
}

// GetDC calls GetDCFunc.
func (mock *OSClientInterfaceMock) GetDC(ns string, dcName string) (appsv1.DeploymentConfig, error) {
	if mock.GetDCFunc == nil {
		panic("OSClientInterfaceMock.GetDCFunc: method is nil but OSClientInterface.GetDC was just called")
	}
	callInfo := struct {
		Ns     string
		DcName string
	}{
		Ns:     ns,
		DcName: dcName,
	}
	lockOSClientInterfaceMockGetDC.Lock()
	mock.calls.GetDC = append(mock.calls.GetDC, callInfo)
	lockOSClientInterfaceMockGetDC.Unlock()
	return mock.GetDCFunc(ns, dcName)
}

// GetDCCalls gets all the calls that were made to GetDC.
// Check the length with:
//     len(mockedOSClientInterface.GetDCCalls())
func (mock *OSClientInterfaceMock) GetDCCalls() []struct {
	Ns     string
	DcName string
} {
	var calls []struct {
		Ns     string
		DcName string
	}
	lockOSClientInterfaceMockGetDC.RLock()
	calls = mock.calls.GetDC
	lockOSClientInterfaceMockGetDC.RUnlock()
	return calls
}

// GetPod calls GetPodFunc.
func (mock *OSClientInterfaceMock) GetPod(ns string, dc string) (v1.Pod, error) {
	if mock.GetPodFunc == nil {
		panic("OSClientInterfaceMock.GetPodFunc: method is nil but OSClientInterface.GetPod was just called")
	}
	callInfo := struct {
		Ns string
		Dc string
	}{
		Ns: ns,
		Dc: dc,
	}
	lockOSClientInterfaceMockGetPod.Lock()
	mock.calls.GetPod = append(mock.calls.GetPod, callInfo)
	lockOSClientInterfaceMockGetPod.Unlock()
	return mock.GetPodFunc(ns, dc)
}

// GetPodCalls gets all the calls that were made to GetPod.
// Check the length with:
//     len(mockedOSClientInterface.GetPodCalls())
func (mock *OSClientInterfaceMock) GetPodCalls() []struct {
	Ns string
	Dc string
} {
	var calls []struct {
		Ns string
		Dc string
	}
	lockOSClientInterfaceMockGetPod.RLock()
	calls = mock.calls.GetPod
	lockOSClientInterfaceMockGetPod.RUnlock()
	return calls
}

// ProcessTemplate calls ProcessTemplateFunc.
func (mock *OSClientInterfaceMock) ProcessTemplate(in1 *tmplv1.Template, in2 map[string]string, in3 TemplateOpt) ([]runtime.RawExtension, error) {
	if mock.ProcessTemplateFunc == nil {
		panic("OSClientInterfaceMock.ProcessTemplateFunc: method is nil but OSClientInterface.ProcessTemplate was just called")
	}
	callInfo := struct {
		In1 *tmplv1.Template
		In2 map[string]string
		In3 TemplateOpt
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	lockOSClientInterfaceMockProcessTemplate.Lock()
	mock.calls.ProcessTemplate = append(mock.calls.ProcessTemplate, callInfo)
	lockOSClientInterfaceMockProcessTemplate.Unlock()
	return mock.ProcessTemplateFunc(in1, in2, in3)
}

// ProcessTemplateCalls gets all the calls that were made to ProcessTemplate.
// Check the length with:
//     len(mockedOSClientInterface.ProcessTemplateCalls())
func (mock *OSClientInterfaceMock) ProcessTemplateCalls() []struct {
	In1 *tmplv1.Template
	In2 map[string]string
	In3 TemplateOpt
} {
	var calls []struct {
		In1 *tmplv1.Template
		In2 map[string]string
		In3 TemplateOpt
	}
	lockOSClientInterfaceMockProcessTemplate.RLock()
	calls = mock.calls.ProcessTemplate
	lockOSClientInterfaceMockProcessTemplate.RUnlock()
	return calls
}

// UpdateDC calls UpdateDCFunc.
func (mock *OSClientInterfaceMock) UpdateDC(ns string, dc *appsv1.DeploymentConfig) error {
	if mock.UpdateDCFunc == nil {
		panic("OSClientInterfaceMock.UpdateDCFunc: method is nil but OSClientInterface.UpdateDC was just called")
	}
	callInfo := struct {
		Ns string
		Dc *appsv1.DeploymentConfig
	}{
		Ns: ns,
		Dc: dc,
	}
	lockOSClientInterfaceMockUpdateDC.Lock()
	mock.calls.UpdateDC = append(mock.calls.UpdateDC, callInfo)
	lockOSClientInterfaceMockUpdateDC.Unlock()
	return mock.UpdateDCFunc(ns, dc)
}

// UpdateDCCalls gets all the calls that were made to UpdateDC.
// Check the length with:
//     len(mockedOSClientInterface.UpdateDCCalls())
func (mock *OSClientInterfaceMock) UpdateDCCalls() []struct {
	Ns string
	Dc *appsv1.DeploymentConfig
} {
	var calls []struct {
		Ns string
		Dc *appsv1.DeploymentConfig
	}
	lockOSClientInterfaceMockUpdateDC.RLock()
	calls = mock.calls.UpdateDC
	lockOSClientInterfaceMockUpdateDC.RUnlock()
	return calls
}
