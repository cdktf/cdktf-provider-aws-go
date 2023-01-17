package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Wafv2WebAclRuleActionOutputReference interface {
	cdktf.ComplexObject
	Allow() Wafv2WebAclRuleActionAllowOutputReference
	AllowInput() *Wafv2WebAclRuleActionAllow
	Block() Wafv2WebAclRuleActionBlockOutputReference
	BlockInput() *Wafv2WebAclRuleActionBlock
	Captcha() Wafv2WebAclRuleActionCaptchaOutputReference
	CaptchaInput() *Wafv2WebAclRuleActionCaptcha
	Challenge() Wafv2WebAclRuleActionChallengeOutputReference
	ChallengeInput() *Wafv2WebAclRuleActionChallenge
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
	Count() Wafv2WebAclRuleActionCountOutputReference
	CountInput() *Wafv2WebAclRuleActionCount
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *Wafv2WebAclRuleAction
	SetInternalValue(val *Wafv2WebAclRuleAction)
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
	PutAllow(value *Wafv2WebAclRuleActionAllow)
	PutBlock(value *Wafv2WebAclRuleActionBlock)
	PutCaptcha(value *Wafv2WebAclRuleActionCaptcha)
	PutChallenge(value *Wafv2WebAclRuleActionChallenge)
	PutCount(value *Wafv2WebAclRuleActionCount)
	ResetAllow()
	ResetBlock()
	ResetCaptcha()
	ResetChallenge()
	ResetCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Wafv2WebAclRuleActionOutputReference
type jsiiProxy_Wafv2WebAclRuleActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Allow() Wafv2WebAclRuleActionAllowOutputReference {
	var returns Wafv2WebAclRuleActionAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) AllowInput() *Wafv2WebAclRuleActionAllow {
	var returns *Wafv2WebAclRuleActionAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Block() Wafv2WebAclRuleActionBlockOutputReference {
	var returns Wafv2WebAclRuleActionBlockOutputReference
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) BlockInput() *Wafv2WebAclRuleActionBlock {
	var returns *Wafv2WebAclRuleActionBlock
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Captcha() Wafv2WebAclRuleActionCaptchaOutputReference {
	var returns Wafv2WebAclRuleActionCaptchaOutputReference
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) CaptchaInput() *Wafv2WebAclRuleActionCaptcha {
	var returns *Wafv2WebAclRuleActionCaptcha
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Challenge() Wafv2WebAclRuleActionChallengeOutputReference {
	var returns Wafv2WebAclRuleActionChallengeOutputReference
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ChallengeInput() *Wafv2WebAclRuleActionChallenge {
	var returns *Wafv2WebAclRuleActionChallenge
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Count() Wafv2WebAclRuleActionCountOutputReference {
	var returns Wafv2WebAclRuleActionCountOutputReference
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) CountInput() *Wafv2WebAclRuleActionCount {
	var returns *Wafv2WebAclRuleActionCount
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) InternalValue() *Wafv2WebAclRuleAction {
	var returns *Wafv2WebAclRuleAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWafv2WebAclRuleActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Wafv2WebAclRuleActionOutputReference {
	_init_.Initialize()

	if err := validateNewWafv2WebAclRuleActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAclRuleActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWafv2WebAclRuleActionOutputReference_Override(w Wafv2WebAclRuleActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAclRuleActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference)SetInternalValue(val *Wafv2WebAclRuleAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAclRuleActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) PutAllow(value *Wafv2WebAclRuleActionAllow) {
	if err := w.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAllow",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) PutBlock(value *Wafv2WebAclRuleActionBlock) {
	if err := w.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBlock",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) PutCaptcha(value *Wafv2WebAclRuleActionCaptcha) {
	if err := w.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) PutChallenge(value *Wafv2WebAclRuleActionChallenge) {
	if err := w.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putChallenge",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) PutCount(value *Wafv2WebAclRuleActionCount) {
	if err := w.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCount",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		w,
		"resetAllow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		w,
		"resetBlock",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		w,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		w,
		"resetChallenge",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		w,
		"resetCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_Wafv2WebAclRuleActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

