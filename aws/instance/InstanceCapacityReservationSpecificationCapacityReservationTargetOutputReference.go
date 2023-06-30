package instance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/instance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference interface {
	cdktf.ComplexObject
	CapacityReservationId() *string
	SetCapacityReservationId(val *string)
	CapacityReservationIdInput() *string
	CapacityReservationResourceGroupArn() *string
	SetCapacityReservationResourceGroupArn(val *string)
	CapacityReservationResourceGroupArnInput() *string
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
	InternalValue() *InstanceCapacityReservationSpecificationCapacityReservationTarget
	SetInternalValue(val *InstanceCapacityReservationSpecificationCapacityReservationTarget)
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
	ResetCapacityReservationId()
	ResetCapacityReservationResourceGroupArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference
type jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) InternalValue() *InstanceCapacityReservationSpecificationCapacityReservationTarget {
	var returns *InstanceCapacityReservationSpecificationCapacityReservationTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference {
	_init_.Initialize()

	if err := validateNewInstanceCapacityReservationSpecificationCapacityReservationTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.instance.InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference_Override(i InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.instance.InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetCapacityReservationId(val *string) {
	if err := j.validateSetCapacityReservationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationId",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetCapacityReservationResourceGroupArn(val *string) {
	if err := j.validateSetCapacityReservationResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetInternalValue(val *InstanceCapacityReservationSpecificationCapacityReservationTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ResetCapacityReservationId() {
	_jsii_.InvokeVoid(
		i,
		"resetCapacityReservationId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ResetCapacityReservationResourceGroupArn() {
	_jsii_.InvokeVoid(
		i,
		"resetCapacityReservationResourceGroupArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InstanceCapacityReservationSpecificationCapacityReservationTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

