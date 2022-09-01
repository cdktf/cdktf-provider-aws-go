package apigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/apigateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiGatewayStageCanarySettingsOutputReference interface {
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
	InternalValue() *ApiGatewayStageCanarySettings
	SetInternalValue(val *ApiGatewayStageCanarySettings)
	PercentTraffic() *float64
	SetPercentTraffic(val *float64)
	PercentTrafficInput() *float64
	StageVariableOverrides() *map[string]*string
	SetStageVariableOverrides(val *map[string]*string)
	StageVariableOverridesInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseStageCache() interface{}
	SetUseStageCache(val interface{})
	UseStageCacheInput() interface{}
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
	ResetPercentTraffic()
	ResetStageVariableOverrides()
	ResetUseStageCache()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiGatewayStageCanarySettingsOutputReference
type jsiiProxy_ApiGatewayStageCanarySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) InternalValue() *ApiGatewayStageCanarySettings {
	var returns *ApiGatewayStageCanarySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) PercentTraffic() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) PercentTrafficInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) StageVariableOverrides() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) StageVariableOverridesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"stageVariableOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) UseStageCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) UseStageCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useStageCacheInput",
		&returns,
	)
	return returns
}


func NewApiGatewayStageCanarySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiGatewayStageCanarySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewApiGatewayStageCanarySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayStageCanarySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.apigateway.ApiGatewayStageCanarySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiGatewayStageCanarySettingsOutputReference_Override(a ApiGatewayStageCanarySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apigateway.ApiGatewayStageCanarySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetInternalValue(val *ApiGatewayStageCanarySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetPercentTraffic(val *float64) {
	if err := j.validateSetPercentTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentTraffic",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetStageVariableOverrides(val *map[string]*string) {
	if err := j.validateSetStageVariableOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageVariableOverrides",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference)SetUseStageCache(val interface{}) {
	if err := j.validateSetUseStageCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useStageCache",
		val,
	)
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ResetPercentTraffic() {
	_jsii_.InvokeVoid(
		a,
		"resetPercentTraffic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ResetStageVariableOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetStageVariableOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ResetUseStageCache() {
	_jsii_.InvokeVoid(
		a,
		"resetUseStageCache",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiGatewayStageCanarySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

