package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference interface {
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
	CustomRequestHandling() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingOutputReference
	CustomRequestHandlingInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandling
	// Experimental.
	Fqn() *string
	InternalValue() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge
	SetInternalValue(val *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge)
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
	PutCustomRequestHandling(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandling)
	ResetCustomRequestHandling()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference
type jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) CustomRequestHandling() Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingOutputReference {
	var returns Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandlingOutputReference
	_jsii_.Get(
		j,
		"customRequestHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) CustomRequestHandlingInput() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandling {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandling
	_jsii_.Get(
		j,
		"customRequestHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) InternalValue() *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge {
	var returns *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference_Override(w Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference)SetInternalValue(val *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallenge) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) PutCustomRequestHandling(value *Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeCustomRequestHandling) {
	if err := w.validatePutCustomRequestHandlingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomRequestHandling",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) ResetCustomRequestHandling() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomRequestHandling",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleStatementRuleGroupReferenceStatementRuleActionOverrideActionToUseChallengeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

