package autoscalinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/autoscalinggroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *AutoscalingGroupMixedInstancesPolicyInstancesDistribution
	SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyInstancesDistribution)
	OnDemandAllocationStrategy() *string
	SetOnDemandAllocationStrategy(val *string)
	OnDemandAllocationStrategyInput() *string
	OnDemandBaseCapacity() *float64
	SetOnDemandBaseCapacity(val *float64)
	OnDemandBaseCapacityInput() *float64
	OnDemandPercentageAboveBaseCapacity() *float64
	SetOnDemandPercentageAboveBaseCapacity(val *float64)
	OnDemandPercentageAboveBaseCapacityInput() *float64
	SpotAllocationStrategy() *string
	SetSpotAllocationStrategy(val *string)
	SpotAllocationStrategyInput() *string
	SpotInstancePools() *float64
	SetSpotInstancePools(val *float64)
	SpotInstancePoolsInput() *float64
	SpotMaxPrice() *string
	SetSpotMaxPrice(val *string)
	SpotMaxPriceInput() *string
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
	ResetOnDemandAllocationStrategy()
	ResetOnDemandBaseCapacity()
	ResetOnDemandPercentageAboveBaseCapacity()
	ResetSpotAllocationStrategy()
	ResetSpotInstancePools()
	ResetSpotMaxPrice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference
type jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InternalValue() *AutoscalingGroupMixedInstancesPolicyInstancesDistribution {
	var returns *AutoscalingGroupMixedInstancesPolicyInstancesDistribution
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandPercentageAboveBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) OnDemandPercentageAboveBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotInstancePools() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotInstancePoolsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotMaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) SpotMaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference {
	_init_.Initialize()

	if err := validateNewAutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference_Override(a AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.autoscalingGroup.AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetInternalValue(val *AutoscalingGroupMixedInstancesPolicyInstancesDistribution) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetOnDemandPercentageAboveBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandPercentageAboveBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandPercentageAboveBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotAllocationStrategy(val *string) {
	if err := j.validateSetSpotAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotInstancePools(val *float64) {
	if err := j.validateSetSpotInstancePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePools",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetSpotMaxPrice(val *string) {
	if err := j.validateSetSpotMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetOnDemandPercentageAboveBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandPercentageAboveBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotInstancePools() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotInstancePools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ResetSpotMaxPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotMaxPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoscalingGroupMixedInstancesPolicyInstancesDistributionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

