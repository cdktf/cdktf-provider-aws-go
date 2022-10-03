package appmeshroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/appmeshroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference interface {
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
	Exact() *string
	SetExact(val *string)
	ExactInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AppmeshRouteSpecGrpcRouteMatchMetadataMatch
	SetInternalValue(val *AppmeshRouteSpecGrpcRouteMatchMetadataMatch)
	Prefix() *string
	SetPrefix(val *string)
	PrefixInput() *string
	Range() AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference
	RangeInput() *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange
	Regex() *string
	SetRegex(val *string)
	RegexInput() *string
	Suffix() *string
	SetSuffix(val *string)
	SuffixInput() *string
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
	PutRange(value *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange)
	ResetExact()
	ResetPrefix()
	ResetRange()
	ResetRegex()
	ResetSuffix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference
type jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Exact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ExactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) InternalValue() *AppmeshRouteSpecGrpcRouteMatchMetadataMatch {
	var returns *AppmeshRouteSpecGrpcRouteMatchMetadataMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Range() AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference {
	var returns AppmeshRouteSpecGrpcRouteMatchMetadataMatchRangeOutputReference
	_jsii_.Get(
		j,
		"range",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) RangeInput() *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange {
	var returns *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange
	_jsii_.Get(
		j,
		"rangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Regex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) RegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Suffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) SuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference {
	_init_.Initialize()

	if err := validateNewAppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference_Override(a AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appmeshRoute.AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetExact(val *string) {
	if err := j.validateSetExactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exact",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetInternalValue(val *AppmeshRouteSpecGrpcRouteMatchMetadataMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetRegex(val *string) {
	if err := j.validateSetRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regex",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetSuffix(val *string) {
	if err := j.validateSetSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suffix",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) PutRange(value *AppmeshRouteSpecGrpcRouteMatchMetadataMatchRange) {
	if err := a.validatePutRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ResetExact() {
	_jsii_.InvokeVoid(
		a,
		"resetExact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ResetRange() {
	_jsii_.InvokeVoid(
		a,
		"resetRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ResetSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppmeshRouteSpecGrpcRouteMatchMetadataMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

