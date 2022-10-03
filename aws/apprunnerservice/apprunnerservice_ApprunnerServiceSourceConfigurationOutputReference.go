package apprunnerservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/apprunnerservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApprunnerServiceSourceConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfiguration() ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() *ApprunnerServiceSourceConfigurationAuthenticationConfiguration
	AutoDeploymentsEnabled() interface{}
	SetAutoDeploymentsEnabled(val interface{})
	AutoDeploymentsEnabledInput() interface{}
	CodeRepository() ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference
	CodeRepositoryInput() *ApprunnerServiceSourceConfigurationCodeRepository
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
	ImageRepository() ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
	ImageRepositoryInput() *ApprunnerServiceSourceConfigurationImageRepository
	InternalValue() *ApprunnerServiceSourceConfiguration
	SetInternalValue(val *ApprunnerServiceSourceConfiguration)
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
	PutAuthenticationConfiguration(value *ApprunnerServiceSourceConfigurationAuthenticationConfiguration)
	PutCodeRepository(value *ApprunnerServiceSourceConfigurationCodeRepository)
	PutImageRepository(value *ApprunnerServiceSourceConfigurationImageRepository)
	ResetAuthenticationConfiguration()
	ResetAutoDeploymentsEnabled()
	ResetCodeRepository()
	ResetImageRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AuthenticationConfiguration() ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AuthenticationConfigurationInput() *ApprunnerServiceSourceConfigurationAuthenticationConfiguration {
	var returns *ApprunnerServiceSourceConfigurationAuthenticationConfiguration
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AutoDeploymentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) AutoDeploymentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeploymentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) CodeRepository() ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference {
	var returns ApprunnerServiceSourceConfigurationCodeRepositoryOutputReference
	_jsii_.Get(
		j,
		"codeRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) CodeRepositoryInput() *ApprunnerServiceSourceConfigurationCodeRepository {
	var returns *ApprunnerServiceSourceConfigurationCodeRepository
	_jsii_.Get(
		j,
		"codeRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ImageRepository() ApprunnerServiceSourceConfigurationImageRepositoryOutputReference {
	var returns ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
	_jsii_.Get(
		j,
		"imageRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ImageRepositoryInput() *ApprunnerServiceSourceConfigurationImageRepository {
	var returns *ApprunnerServiceSourceConfigurationImageRepository
	_jsii_.Get(
		j,
		"imageRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) InternalValue() *ApprunnerServiceSourceConfiguration {
	var returns *ApprunnerServiceSourceConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApprunnerServiceSourceConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApprunnerServiceSourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApprunnerServiceSourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.apprunnerService.ApprunnerServiceSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationOutputReference_Override(a ApprunnerServiceSourceConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apprunnerService.ApprunnerServiceSourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetAutoDeploymentsEnabled(val interface{}) {
	if err := j.validateSetAutoDeploymentsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDeploymentsEnabled",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetInternalValue(val *ApprunnerServiceSourceConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutAuthenticationConfiguration(value *ApprunnerServiceSourceConfigurationAuthenticationConfiguration) {
	if err := a.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutCodeRepository(value *ApprunnerServiceSourceConfigurationCodeRepository) {
	if err := a.validatePutCodeRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) PutImageRepository(value *ApprunnerServiceSourceConfigurationImageRepository) {
	if err := a.validatePutImageRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageRepository",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetAutoDeploymentsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoDeploymentsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetCodeRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ResetImageRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetImageRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

