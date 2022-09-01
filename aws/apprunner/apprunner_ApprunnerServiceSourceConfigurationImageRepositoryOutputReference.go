package apprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/apprunner/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApprunnerServiceSourceConfigurationImageRepositoryOutputReference interface {
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
	ImageConfiguration() ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference
	ImageConfigurationInput() *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration
	ImageIdentifier() *string
	SetImageIdentifier(val *string)
	ImageIdentifierInput() *string
	ImageRepositoryType() *string
	SetImageRepositoryType(val *string)
	ImageRepositoryTypeInput() *string
	InternalValue() *ApprunnerServiceSourceConfigurationImageRepository
	SetInternalValue(val *ApprunnerServiceSourceConfigurationImageRepository)
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
	PutImageConfiguration(value *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration)
	ResetImageConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApprunnerServiceSourceConfigurationImageRepositoryOutputReference
type jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageConfiguration() ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference {
	var returns ApprunnerServiceSourceConfigurationImageRepositoryImageConfigurationOutputReference
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageConfigurationInput() *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration {
	var returns *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration
	_jsii_.Get(
		j,
		"imageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageRepositoryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ImageRepositoryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRepositoryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) InternalValue() *ApprunnerServiceSourceConfigurationImageRepository {
	var returns *ApprunnerServiceSourceConfigurationImageRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApprunnerServiceSourceConfigurationImageRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApprunnerServiceSourceConfigurationImageRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewApprunnerServiceSourceConfigurationImageRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.apprunner.ApprunnerServiceSourceConfigurationImageRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApprunnerServiceSourceConfigurationImageRepositoryOutputReference_Override(a ApprunnerServiceSourceConfigurationImageRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.apprunner.ApprunnerServiceSourceConfigurationImageRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetImageIdentifier(val *string) {
	if err := j.validateSetImageIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageIdentifier",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetImageRepositoryType(val *string) {
	if err := j.validateSetImageRepositoryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageRepositoryType",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetInternalValue(val *ApprunnerServiceSourceConfigurationImageRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) PutImageConfiguration(value *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration) {
	if err := a.validatePutImageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ResetImageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApprunnerServiceSourceConfigurationImageRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

