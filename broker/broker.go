package broker

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

var (
	truePointer  bool = true
	falsePointer bool = false
)

//Broker - Broker Struct
type Broker struct {
}

// Services - Service Catalog
func (b Broker) Services(ctx context.Context) ([]brokerapi.Service, error) {

	smallPlan := brokerapi.ServicePlan{
		Description: "small-plan",
		Free:        &truePointer,
		Name:        "small",
		Bindable:    &truePointer,
		ID:          "small-id",
	}

	s := brokerapi.Service{
		ID:            "Pivotal",
		Name:          "Pivotal Service Plans",
		Description:   "Service Plan which gives us Pivotal",
		Bindable:      true,
		PlanUpdatable: true,
		Tags:          []string{},
		Plans:         []brokerapi.ServicePlan{smallPlan},
	}
	return []brokerapi.Service{s}, nil
}

// Provision
func (b Broker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {

	returnObject := brokerapi.ProvisionedServiceSpec{}

	return returnObject, nil
}

// Deprovision
func (b Broker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {

	returnObject := brokerapi.DeprovisionServiceSpec{}

	return returnObject, nil
}

// GetInstance
func (b Broker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {

	returnObject := brokerapi.GetInstanceDetailsSpec{}
	return returnObject, nil
}

// LastOperation
func (b Broker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {

	returnObject := brokerapi.LastOperation{}

	return returnObject, nil
}

// Update
func (b Broker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {

	returnObject := brokerapi.UpdateServiceSpec{}

	return returnObject, nil
}

// Bind
func (b Broker) Bind(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {

	returnObject := brokerapi.Binding{}

	return returnObject, nil
}

// GetBinding
func (b Broker) GetBinding(ctx context.Context, instanceID, bindingID string) (brokerapi.GetBindingSpec, error) {

	returnObject := brokerapi.GetBindingSpec{}

	return returnObject, nil
}

// LastBindingOperation
func (b Broker) LastBindingOperation(ctx context.Context, instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {

	returnObject := brokerapi.LastOperation{}

	return returnObject, nil
}

// Unbind
func (b Broker) Unbind(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {

	returnObject := brokerapi.UnbindSpec{}

	return returnObject, nil

}
