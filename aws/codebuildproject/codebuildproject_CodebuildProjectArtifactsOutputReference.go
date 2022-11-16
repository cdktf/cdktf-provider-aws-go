package codebuildproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v11/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v11/codebuildproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildProjectArtifactsOutputReference interface {
	cdktf.ComplexObject
	ArtifactIdentifier() *string
	SetArtifactIdentifier(val *string)
	ArtifactIdentifierInput() *string
	BucketOwnerAccess() *string
	SetBucketOwnerAccess(val *string)
	BucketOwnerAccessInput() *string
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
	EncryptionDisabled() interface{}
	SetEncryptionDisabled(val interface{})
	EncryptionDisabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CodebuildProjectArtifacts
	SetInternalValue(val *CodebuildProjectArtifacts)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceType() *string
	SetNamespaceType(val *string)
	NamespaceTypeInput() *string
	OverrideArtifactName() interface{}
	SetOverrideArtifactName(val interface{})
	OverrideArtifactNameInput() interface{}
	Packaging() *string
	SetPackaging(val *string)
	PackagingInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetArtifactIdentifier()
	ResetBucketOwnerAccess()
	ResetEncryptionDisabled()
	ResetLocation()
	ResetName()
	ResetNamespaceType()
	ResetOverrideArtifactName()
	ResetPackaging()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodebuildProjectArtifactsOutputReference
type jsiiProxy_CodebuildProjectArtifactsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ArtifactIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ArtifactIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) BucketOwnerAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) BucketOwnerAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) EncryptionDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) EncryptionDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionDisabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) InternalValue() *CodebuildProjectArtifacts {
	var returns *CodebuildProjectArtifacts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NamespaceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) NamespaceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) OverrideArtifactName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) OverrideArtifactNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideArtifactNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Packaging() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) PackagingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packagingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCodebuildProjectArtifactsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodebuildProjectArtifactsOutputReference {
	_init_.Initialize()

	if err := validateNewCodebuildProjectArtifactsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodebuildProjectArtifactsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodebuildProjectArtifactsOutputReference_Override(c CodebuildProjectArtifactsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetArtifactIdentifier(val *string) {
	if err := j.validateSetArtifactIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactIdentifier",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetBucketOwnerAccess(val *string) {
	if err := j.validateSetBucketOwnerAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketOwnerAccess",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetEncryptionDisabled(val interface{}) {
	if err := j.validateSetEncryptionDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionDisabled",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetInternalValue(val *CodebuildProjectArtifacts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetNamespaceType(val *string) {
	if err := j.validateSetNamespaceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceType",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetOverrideArtifactName(val interface{}) {
	if err := j.validateSetOverrideArtifactNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideArtifactName",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetPackaging(val *string) {
	if err := j.validateSetPackagingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packaging",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectArtifactsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetArtifactIdentifier() {
	_jsii_.InvokeVoid(
		c,
		"resetArtifactIdentifier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetBucketOwnerAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetBucketOwnerAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetEncryptionDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetNamespaceType() {
	_jsii_.InvokeVoid(
		c,
		"resetNamespaceType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetOverrideArtifactName() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideArtifactName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetPackaging() {
	_jsii_.InvokeVoid(
		c,
		"resetPackaging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		c,
		"resetPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodebuildProjectArtifactsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

