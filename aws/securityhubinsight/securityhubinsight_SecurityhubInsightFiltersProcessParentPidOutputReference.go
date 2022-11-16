package securityhubinsight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/securityhubinsight/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubInsightFiltersProcessParentPidOutputReference interface {
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
	Eq() *string
	SetEq(val *string)
	EqInput() *string
	// Experimental.
	Fqn() *string
	Gte() *string
	SetGte(val *string)
	GteInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lte() *string
	SetLte(val *string)
	LteInput() *string
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
	ResetEq()
	ResetGte()
	ResetLte()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubInsightFiltersProcessParentPidOutputReference
type jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) Eq() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) EqInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) Gte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) Lte() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) LteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubInsightFiltersProcessParentPidOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityhubInsightFiltersProcessParentPidOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubInsightFiltersProcessParentPidOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersProcessParentPidOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityhubInsightFiltersProcessParentPidOutputReference_Override(s SecurityhubInsightFiltersProcessParentPidOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.securityhubInsight.SecurityhubInsightFiltersProcessParentPidOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetEq(val *string) {
	if err := j.validateSetEqParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eq",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetGte(val *string) {
	if err := j.validateSetGteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gte",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetLte(val *string) {
	if err := j.validateSetLteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lte",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ResetEq() {
	_jsii_.InvokeVoid(
		s,
		"resetEq",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ResetGte() {
	_jsii_.InvokeVoid(
		s,
		"resetGte",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ResetLte() {
	_jsii_.InvokeVoid(
		s,
		"resetLte",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubInsightFiltersProcessParentPidOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

