// Code generated by mockery v1.0.0
package servicecatalog

import gqlschema "github.com/kyma-project/kyma/components/ui-api-layer/internal/gqlschema"
import mock "github.com/stretchr/testify/mock"
import status "github.com/kyma-project/kyma/components/ui-api-layer/internal/domain/servicecatalog/status"
import v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"

// mockGqlInstanceConverter is an autogenerated mock type for the gqlInstanceConverter type
type mockGqlInstanceConverter struct {
	mock.Mock
}

// GQLCreateInputToInstanceCreateParameters provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) GQLCreateInputToInstanceCreateParameters(in *gqlschema.ServiceInstanceCreateInput) *instanceCreateParameters {
	ret := _m.Called(in)

	var r0 *instanceCreateParameters
	if rf, ok := ret.Get(0).(func(*gqlschema.ServiceInstanceCreateInput) *instanceCreateParameters); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*instanceCreateParameters)
		}
	}

	return r0
}

// GQLStatusToServiceStatus provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) GQLStatusToServiceStatus(in *gqlschema.ServiceInstanceStatus) *status.ServiceInstanceStatus {
	ret := _m.Called(in)

	var r0 *status.ServiceInstanceStatus
	if rf, ok := ret.Get(0).(func(*gqlschema.ServiceInstanceStatus) *status.ServiceInstanceStatus); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*status.ServiceInstanceStatus)
		}
	}

	return r0
}

// GQLStatusTypeToServiceStatusType provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) GQLStatusTypeToServiceStatusType(in gqlschema.InstanceStatusType) status.ServiceInstanceStatusType {
	ret := _m.Called(in)

	var r0 status.ServiceInstanceStatusType
	if rf, ok := ret.Get(0).(func(gqlschema.InstanceStatusType) status.ServiceInstanceStatusType); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(status.ServiceInstanceStatusType)
	}

	return r0
}

// ServiceStatusToGQLStatus provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) ServiceStatusToGQLStatus(in status.ServiceInstanceStatus) gqlschema.ServiceInstanceStatus {
	ret := _m.Called(in)

	var r0 gqlschema.ServiceInstanceStatus
	if rf, ok := ret.Get(0).(func(status.ServiceInstanceStatus) gqlschema.ServiceInstanceStatus); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(gqlschema.ServiceInstanceStatus)
	}

	return r0
}

// ServiceStatusTypeToGQLStatusType provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) ServiceStatusTypeToGQLStatusType(in status.ServiceInstanceStatusType) gqlschema.InstanceStatusType {
	ret := _m.Called(in)

	var r0 gqlschema.InstanceStatusType
	if rf, ok := ret.Get(0).(func(status.ServiceInstanceStatusType) gqlschema.InstanceStatusType); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(gqlschema.InstanceStatusType)
	}

	return r0
}

// ToGQL provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) ToGQL(in *v1beta1.ServiceInstance) (*gqlschema.ServiceInstance, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ServiceInstance
	if rf, ok := ret.Get(0).(func(*v1beta1.ServiceInstance) *gqlschema.ServiceInstance); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1beta1.ServiceInstance) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *mockGqlInstanceConverter) ToGQLs(in []*v1beta1.ServiceInstance) ([]gqlschema.ServiceInstance, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ServiceInstance
	if rf, ok := ret.Get(0).(func([]*v1beta1.ServiceInstance) []gqlschema.ServiceInstance); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1beta1.ServiceInstance) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
