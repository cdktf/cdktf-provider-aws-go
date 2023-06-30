package codebuildproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/codebuildproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildProjectBuildBatchConfigRestrictionsOutputReference interface {
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
	ComputeTypesAllowed() *[]*string
	SetComputeTypesAllowed(val *[]*string)
	ComputeTypesAllowedInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CodebuildProjectBuildBatchConfigRestrictions
	SetInternalValue(val *CodebuildProjectBuildBatchConfigRestrictions)
	MaximumBuildsAllowed() *float64
	SetMaximumBuildsAllowed(val *float64)
	MaximumBuildsAllowedInput() *float64
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
	ResetComputeTypesAllowed()
	ResetMaximumBuildsAllowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodebuildProjectBuildBatchConfigRestrictionsOutputReference
type jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComputeTypesAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeTypesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComputeTypesAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"computeTypesAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InternalValue() *CodebuildProjectBuildBatchConfigRestrictions {
	var returns *CodebuildProjectBuildBatchConfigRestrictions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) MaximumBuildsAllowed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBuildsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) MaximumBuildsAllowedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBuildsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCodebuildProjectBuildBatchConfigRestrictionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodebuildProjectBuildBatchConfigRestrictionsOutputReference {
	_init_.Initialize()

	if err := validateNewCodebuildProjectBuildBatchConfigRestrictionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectBuildBatchConfigRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodebuildProjectBuildBatchConfigRestrictionsOutputReference_Override(c CodebuildProjectBuildBatchConfigRestrictionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectBuildBatchConfigRestrictionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetComputeTypesAllowed(val *[]*string) {
	if err := j.validateSetComputeTypesAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeTypesAllowed",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetInternalValue(val *CodebuildProjectBuildBatchConfigRestrictions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetMaximumBuildsAllowed(val *float64) {
	if err := j.validateSetMaximumBuildsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBuildsAllowed",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ResetComputeTypesAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetComputeTypesAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ResetMaximumBuildsAllowed() {
	_jsii_.InvokeVoid(
		c,
		"resetMaximumBuildsAllowed",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodebuildProjectBuildBatchConfigRestrictionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

