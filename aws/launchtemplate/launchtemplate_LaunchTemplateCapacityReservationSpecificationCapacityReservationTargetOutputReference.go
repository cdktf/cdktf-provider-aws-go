package launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/launchtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference interface {
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
	InternalValue() *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget
	SetInternalValue(val *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget)
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

// The jsii proxy struct for LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference
type jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) CapacityReservationResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) InternalValue() *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget {
	var returns *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference {
	_init_.Initialize()

	if err := validateNewLaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference_Override(l LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.launchTemplate.LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetCapacityReservationId(val *string) {
	if err := j.validateSetCapacityReservationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationId",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetCapacityReservationResourceGroupArn(val *string) {
	if err := j.validateSetCapacityReservationResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetInternalValue(val *LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ResetCapacityReservationId() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityReservationId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ResetCapacityReservationResourceGroupArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityReservationResourceGroupArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LaunchTemplateCapacityReservationSpecificationCapacityReservationTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

