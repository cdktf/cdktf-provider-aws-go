package ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/ec2fleet/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2FleetOnDemandOptionsOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *Ec2FleetOnDemandOptions
	SetInternalValue(val *Ec2FleetOnDemandOptions)
	MaxTotalPrice() *string
	SetMaxTotalPrice(val *string)
	MaxTotalPriceInput() *string
	MinTargetCapacity() *float64
	SetMinTargetCapacity(val *float64)
	MinTargetCapacityInput() *float64
	SingleAvailabilityZone() interface{}
	SetSingleAvailabilityZone(val interface{})
	SingleAvailabilityZoneInput() interface{}
	SingleInstanceType() interface{}
	SetSingleInstanceType(val interface{})
	SingleInstanceTypeInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllocationStrategy()
	ResetMaxTotalPrice()
	ResetMinTargetCapacity()
	ResetSingleAvailabilityZone()
	ResetSingleInstanceType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2FleetOnDemandOptionsOutputReference
type jsiiProxy_Ec2FleetOnDemandOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) InternalValue() *Ec2FleetOnDemandOptions {
	var returns *Ec2FleetOnDemandOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) MaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) MaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) MinTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) MinTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) SingleAvailabilityZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) SingleAvailabilityZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) SingleInstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) SingleInstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2FleetOnDemandOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2FleetOnDemandOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2FleetOnDemandOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2FleetOnDemandOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetOnDemandOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2FleetOnDemandOptionsOutputReference_Override(e Ec2FleetOnDemandOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ec2Fleet.Ec2FleetOnDemandOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetInternalValue(val *Ec2FleetOnDemandOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetMaxTotalPrice(val *string) {
	if err := j.validateSetMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetMinTargetCapacity(val *float64) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetSingleAvailabilityZone(val interface{}) {
	if err := j.validateSetSingleAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetSingleInstanceType(val interface{}) {
	if err := j.validateSetSingleInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleInstanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ResetMaxTotalPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxTotalPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ResetSingleAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetSingleAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ResetSingleInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetSingleInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2FleetOnDemandOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

