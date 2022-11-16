package ecrpublicrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/ecrpublicrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcrpublicRepositoryCatalogDataOutputReference interface {
	cdktf.ComplexObject
	AboutText() *string
	SetAboutText(val *string)
	AboutTextInput() *string
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *EcrpublicRepositoryCatalogData
	SetInternalValue(val *EcrpublicRepositoryCatalogData)
	LogoImageBlob() *string
	SetLogoImageBlob(val *string)
	LogoImageBlobInput() *string
	OperatingSystems() *[]*string
	SetOperatingSystems(val *[]*string)
	OperatingSystemsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageText() *string
	SetUsageText(val *string)
	UsageTextInput() *string
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
	ResetAboutText()
	ResetArchitectures()
	ResetDescription()
	ResetLogoImageBlob()
	ResetOperatingSystems()
	ResetUsageText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcrpublicRepositoryCatalogDataOutputReference
type jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) AboutText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) AboutTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aboutTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) InternalValue() *EcrpublicRepositoryCatalogData {
	var returns *EcrpublicRepositoryCatalogData
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) LogoImageBlob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImageBlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) LogoImageBlobInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoImageBlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) OperatingSystems() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) OperatingSystemsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatingSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) UsageText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) UsageTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageTextInput",
		&returns,
	)
	return returns
}


func NewEcrpublicRepositoryCatalogDataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EcrpublicRepositoryCatalogDataOutputReference {
	_init_.Initialize()

	if err := validateNewEcrpublicRepositoryCatalogDataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecrpublicRepository.EcrpublicRepositoryCatalogDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcrpublicRepositoryCatalogDataOutputReference_Override(e EcrpublicRepositoryCatalogDataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecrpublicRepository.EcrpublicRepositoryCatalogDataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetAboutText(val *string) {
	if err := j.validateSetAboutTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aboutText",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetInternalValue(val *EcrpublicRepositoryCatalogData) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetLogoImageBlob(val *string) {
	if err := j.validateSetLogoImageBlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoImageBlob",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetOperatingSystems(val *[]*string) {
	if err := j.validateSetOperatingSystemsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatingSystems",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference)SetUsageText(val *string) {
	if err := j.validateSetUsageTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageText",
		val,
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetAboutText() {
	_jsii_.InvokeVoid(
		e,
		"resetAboutText",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetArchitectures() {
	_jsii_.InvokeVoid(
		e,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetLogoImageBlob() {
	_jsii_.InvokeVoid(
		e,
		"resetLogoImageBlob",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetOperatingSystems() {
	_jsii_.InvokeVoid(
		e,
		"resetOperatingSystems",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ResetUsageText() {
	_jsii_.InvokeVoid(
		e,
		"resetUsageText",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrpublicRepositoryCatalogDataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

