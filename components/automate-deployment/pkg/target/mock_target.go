// Code generated by mockery v1.0.0. DO NOT EDIT.
package target

import cli "github.com/chef/automate/components/automate-deployment/pkg/cli"
import command "github.com/chef/automate/lib/platform/command"
import context "context"
import deployment "github.com/chef/automate/api/config/deployment"
import depot "github.com/chef/automate/components/automate-deployment/pkg/depot"
import habapi "github.com/chef/automate/components/automate-deployment/pkg/habapi"
import habpkg "github.com/chef/automate/components/automate-deployment/pkg/habpkg"
import interservicedeployment "github.com/chef/automate/api/interservice/deployment"
import manifest "github.com/chef/automate/components/automate-deployment/pkg/manifest"
import mock "github.com/stretchr/testify/mock"
import net "net"

// MockTarget is an autogenerated mock type for the Target type
type MockTarget struct {
	mock.Mock
}

// BinlinkPackage provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTarget) BinlinkPackage(_a0 context.Context, _a1 habpkg.VersionedPackage, _a2 string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage, string) string); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, habpkg.VersionedPackage, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CommandExecutor provides a mock function with given fields:
func (_m *MockTarget) CommandExecutor() command.Executor {
	ret := _m.Called()

	var r0 command.Executor
	if rf, ok := ret.Get(0).(func() command.Executor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(command.Executor)
		}
	}

	return r0
}

// DeployDeploymentService provides a mock function with given fields: ctx, cfg, releaseManifest, bootstrapBundlePath, writer
func (_m *MockTarget) DeployDeploymentService(ctx context.Context, cfg *deployment.ConfigRequest, releaseManifest manifest.ReleaseManifest, bootstrapBundlePath string, writer cli.BodyWriter) error {
	ret := _m.Called(ctx, cfg, releaseManifest, bootstrapBundlePath, writer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployment.ConfigRequest, manifest.ReleaseManifest, string, cli.BodyWriter) error); ok {
		r0 = rf(ctx, cfg, releaseManifest, bootstrapBundlePath, writer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeployedServices provides a mock function with given fields: ctx
func (_m *MockTarget) DeployedServices(ctx context.Context) (map[string]DeployedService, error) {
	ret := _m.Called(ctx)

	var r0 map[string]DeployedService
	if rf, ok := ret.Get(0).(func(context.Context) map[string]DeployedService); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]DeployedService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyData provides a mock function with given fields:
func (_m *MockTarget) DestroyData() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyPkgCache provides a mock function with given fields:
func (_m *MockTarget) DestroyPkgCache() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroySupervisor provides a mock function with given fields:
func (_m *MockTarget) DestroySupervisor() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disable provides a mock function with given fields:
func (_m *MockTarget) Disable() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureDisabled provides a mock function with given fields:
func (_m *MockTarget) EnsureDisabled() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EnsureStopped provides a mock function with given fields:
func (_m *MockTarget) EnsureStopped() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAutomateUnitFile provides a mock function with given fields:
func (_m *MockTarget) GetAutomateUnitFile() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSymlinkedHabSup provides a mock function with given fields:
func (_m *MockTarget) GetSymlinkedHabSup() (habpkg.HabPkg, error) {
	ret := _m.Called()

	var r0 habpkg.HabPkg
	if rf, ok := ret.Get(0).(func() habpkg.HabPkg); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(habpkg.HabPkg)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserToml provides a mock function with given fields: pkg
func (_m *MockTarget) GetUserToml(pkg habpkg.VersionedPackage) (string, error) {
	ret := _m.Called(pkg)

	var r0 string
	if rf, ok := ret.Get(0).(func(habpkg.VersionedPackage) string); ok {
		r0 = rf(pkg)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(habpkg.VersionedPackage) error); ok {
		r1 = rf(pkg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HabAPIClient provides a mock function with given fields:
func (_m *MockTarget) HabAPIClient() *habapi.Client {
	ret := _m.Called()

	var r0 *habapi.Client
	if rf, ok := ret.Get(0).(func() *habapi.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*habapi.Client)
		}
	}

	return r0
}

// HabCache provides a mock function with given fields:
func (_m *MockTarget) HabCache() depot.HabCache {
	ret := _m.Called()

	var r0 depot.HabCache
	if rf, ok := ret.Get(0).(func() depot.HabCache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(depot.HabCache)
		}
	}

	return r0
}

// HabSup provides a mock function with given fields:
func (_m *MockTarget) HabSup() HabSup {
	ret := _m.Called()

	var r0 HabSup
	if rf, ok := ret.Get(0).(func() HabSup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(HabSup)
		}
	}

	return r0
}

// HabSupRestart provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) HabSupRestart(_a0 context.Context, _a1 []string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HabSupRestartRequired provides a mock function with given fields: _a0
func (_m *MockTarget) HabSupRestartRequired(_a0 habpkg.HabPkg) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(habpkg.HabPkg) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(habpkg.HabPkg) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IPs provides a mock function with given fields:
func (_m *MockTarget) IPs() []net.IP {
	ret := _m.Called()

	var r0 []net.IP
	if rf, ok := ret.Get(0).(func() []net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]net.IP)
		}
	}

	return r0
}

// InstallDeploymentService provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTarget) InstallDeploymentService(_a0 context.Context, _a1 *deployment.ConfigRequest, _a2 manifest.ReleaseManifest) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployment.ConfigRequest, manifest.ReleaseManifest) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstallHabitat provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTarget) InstallHabitat(_a0 context.Context, _a1 manifest.ReleaseManifest, _a2 cli.BodyWriter) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, manifest.ReleaseManifest, cli.BodyWriter) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InstallService provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTarget) InstallService(_a0 context.Context, _a1 habpkg.Installable, _a2 string) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.Installable, string) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsBinlinked provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) IsBinlinked(_a0 habpkg.VersionedPackage, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(habpkg.VersionedPackage, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(habpkg.VersionedPackage, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsInstalled provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) IsInstalled(_a0 context.Context, _a1 habpkg.VersionedPackage) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadDeploymentService provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) LoadDeploymentService(_a0 context.Context, _a1 habpkg.VersionedPackage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadService provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTarget) LoadService(_a0 context.Context, _a1 habpkg.VersionedPackage, _a2 ...LoadOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage, ...LoadOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveService provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) RemoveService(_a0 context.Context, _a1 habpkg.VersionedPackage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenderAutomateUnitFile provides a mock function with given fields: proxyConfig, habP, habLauncherP
func (_m *MockTarget) RenderAutomateUnitFile(proxyConfig string, habP habpkg.HabPkg, habLauncherP habpkg.HabPkg) (string, error) {
	ret := _m.Called(proxyConfig, habP, habLauncherP)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, habpkg.HabPkg, habpkg.HabPkg) string); ok {
		r0 = rf(proxyConfig, habP, habLauncherP)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, habpkg.HabPkg, habpkg.HabPkg) error); ok {
		r1 = rf(proxyConfig, habP, habLauncherP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetHabitatEnvironment provides a mock function with given fields: _a0
func (_m *MockTarget) SetHabitatEnvironment(_a0 manifest.ReleaseManifest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(manifest.ReleaseManifest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUserToml provides a mock function with given fields: name, config
func (_m *MockTarget) SetUserToml(name string, config string) error {
	ret := _m.Called(name, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(name, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetupSupervisor provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockTarget) SetupSupervisor(_a0 context.Context, _a1 *deployment.ConfigRequest, _a2 manifest.ReleaseManifest, _a3 cli.FormatWriter) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *deployment.ConfigRequest, manifest.ReleaseManifest, cli.FormatWriter) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartService provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) StartService(_a0 context.Context, _a1 habpkg.VersionedPackage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Status provides a mock function with given fields: ctx, serviceNames
func (_m *MockTarget) Status(ctx context.Context, serviceNames []string) *interservicedeployment.ServiceStatus {
	ret := _m.Called(ctx, serviceNames)

	var r0 *interservicedeployment.ServiceStatus
	if rf, ok := ret.Get(0).(func(context.Context, []string) *interservicedeployment.ServiceStatus); ok {
		r0 = rf(ctx, serviceNames)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interservicedeployment.ServiceStatus)
		}
	}

	return r0
}

// Stop provides a mock function with given fields: ctx
func (_m *MockTarget) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopService provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) StopService(_a0 context.Context, _a1 habpkg.VersionedPackage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SymlinkHabSup provides a mock function with given fields: habSupP
func (_m *MockTarget) SymlinkHabSup(habSupP habpkg.HabPkg) error {
	ret := _m.Called(habSupP)

	var r0 error
	if rf, ok := ret.Get(0).(func(habpkg.HabPkg) error); ok {
		r0 = rf(habSupP)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemdReload provides a mock function with given fields:
func (_m *MockTarget) SystemdReload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SystemdReloadRequired provides a mock function with given fields:
func (_m *MockTarget) SystemdReloadRequired() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemdRestart provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) SystemdRestart(_a0 context.Context, _a1 []string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, []string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SystemdRunning provides a mock function with given fields:
func (_m *MockTarget) SystemdRunning() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnloadService provides a mock function with given fields: _a0, _a1
func (_m *MockTarget) UnloadService(_a0 context.Context, _a1 habpkg.VersionedPackage) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habpkg.VersionedPackage) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteAutomateUnitFile provides a mock function with given fields: _a0
func (_m *MockTarget) WriteAutomateUnitFile(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
