package appflowflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v14/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v14/appflowflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference interface {
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
	DocumentType() *string
	SetDocumentType(val *string)
	DocumentTypeInput() *string
	// Experimental.
	Fqn() *string
	IncludeAllVersions() interface{}
	SetIncludeAllVersions(val interface{})
	IncludeAllVersionsInput() interface{}
	IncludeRenditions() interface{}
	SetIncludeRenditions(val interface{})
	IncludeRenditionsInput() interface{}
	IncludeSourceFiles() interface{}
	SetIncludeSourceFiles(val interface{})
	IncludeSourceFilesInput() interface{}
	InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva
	SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva)
	Object() *string
	SetObject(val *string)
	ObjectInput() *string
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
	ResetDocumentType()
	ResetIncludeAllVersions()
	ResetIncludeRenditions()
	ResetIncludeSourceFiles()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference
type jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) DocumentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) DocumentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeAllVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAllVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeAllVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAllVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeRenditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRenditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeRenditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRenditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeSourceFiles() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSourceFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) IncludeSourceFilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSourceFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) InternalValue() *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva {
	var returns *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) Object() *string {
	var returns *string
	_jsii_.Get(
		j,
		"object",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ObjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference {
	_init_.Initialize()

	if err := validateNewAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference_Override(a AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.appflowFlow.AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetDocumentType(val *string) {
	if err := j.validateSetDocumentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentType",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetIncludeAllVersions(val interface{}) {
	if err := j.validateSetIncludeAllVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAllVersions",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetIncludeRenditions(val interface{}) {
	if err := j.validateSetIncludeRenditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRenditions",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetIncludeSourceFiles(val interface{}) {
	if err := j.validateSetIncludeSourceFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSourceFiles",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetInternalValue(val *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetObject(val *string) {
	if err := j.validateSetObjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"object",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ResetDocumentType() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ResetIncludeAllVersions() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeAllVersions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ResetIncludeRenditions() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeRenditions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ResetIncludeSourceFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeSourceFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeevaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

