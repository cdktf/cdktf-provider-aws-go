package cloudsearchdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/cloudsearchdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudsearchDomainScalingParametersOutputReference interface {
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
	DesiredInstanceType() *string
	SetDesiredInstanceType(val *string)
	DesiredInstanceTypeInput() *string
	DesiredPartitionCount() *float64
	SetDesiredPartitionCount(val *float64)
	DesiredPartitionCountInput() *float64
	DesiredReplicationCount() *float64
	SetDesiredReplicationCount(val *float64)
	DesiredReplicationCountInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *CloudsearchDomainScalingParameters
	SetInternalValue(val *CloudsearchDomainScalingParameters)
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
	ResetDesiredInstanceType()
	ResetDesiredPartitionCount()
	ResetDesiredReplicationCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudsearchDomainScalingParametersOutputReference
type jsiiProxy_CloudsearchDomainScalingParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredPartitionCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredPartitionCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredPartitionCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredPartitionCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredReplicationCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredReplicationCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) DesiredReplicationCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredReplicationCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) InternalValue() *CloudsearchDomainScalingParameters {
	var returns *CloudsearchDomainScalingParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudsearchDomainScalingParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudsearchDomainScalingParametersOutputReference {
	_init_.Initialize()

	if err := validateNewCloudsearchDomainScalingParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudsearchDomainScalingParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudsearchDomain.CloudsearchDomainScalingParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudsearchDomainScalingParametersOutputReference_Override(c CloudsearchDomainScalingParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudsearchDomain.CloudsearchDomainScalingParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetDesiredInstanceType(val *string) {
	if err := j.validateSetDesiredInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredInstanceType",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetDesiredPartitionCount(val *float64) {
	if err := j.validateSetDesiredPartitionCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredPartitionCount",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetDesiredReplicationCount(val *float64) {
	if err := j.validateSetDesiredReplicationCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredReplicationCount",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetInternalValue(val *CloudsearchDomainScalingParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudsearchDomainScalingParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ResetDesiredInstanceType() {
	_jsii_.InvokeVoid(
		c,
		"resetDesiredInstanceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ResetDesiredPartitionCount() {
	_jsii_.InvokeVoid(
		c,
		"resetDesiredPartitionCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ResetDesiredReplicationCount() {
	_jsii_.InvokeVoid(
		c,
		"resetDesiredReplicationCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudsearchDomainScalingParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

