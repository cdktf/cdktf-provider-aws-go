package codebuildproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/codebuildproject/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodebuildProjectSecondarySourcesOutputReference interface {
	cdktf.ComplexObject
	Buildspec() *string
	SetBuildspec(val *string)
	BuildspecInput() *string
	BuildStatusConfig() CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference
	BuildStatusConfigInput() *CodebuildProjectSecondarySourcesBuildStatusConfig
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
	GitCloneDepth() *float64
	SetGitCloneDepth(val *float64)
	GitCloneDepthInput() *float64
	GitSubmodulesConfig() CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference
	GitSubmodulesConfigInput() *CodebuildProjectSecondarySourcesGitSubmodulesConfig
	InsecureSsl() interface{}
	SetInsecureSsl(val interface{})
	InsecureSslInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ReportBuildStatus() interface{}
	SetReportBuildStatus(val interface{})
	ReportBuildStatusInput() interface{}
	SourceIdentifier() *string
	SetSourceIdentifier(val *string)
	SourceIdentifierInput() *string
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
	PutBuildStatusConfig(value *CodebuildProjectSecondarySourcesBuildStatusConfig)
	PutGitSubmodulesConfig(value *CodebuildProjectSecondarySourcesGitSubmodulesConfig)
	ResetBuildspec()
	ResetBuildStatusConfig()
	ResetGitCloneDepth()
	ResetGitSubmodulesConfig()
	ResetInsecureSsl()
	ResetLocation()
	ResetReportBuildStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodebuildProjectSecondarySourcesOutputReference
type jsiiProxy_CodebuildProjectSecondarySourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) Buildspec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) BuildspecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildspecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) BuildStatusConfig() CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference {
	var returns CodebuildProjectSecondarySourcesBuildStatusConfigOutputReference
	_jsii_.Get(
		j,
		"buildStatusConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) BuildStatusConfigInput() *CodebuildProjectSecondarySourcesBuildStatusConfig {
	var returns *CodebuildProjectSecondarySourcesBuildStatusConfig
	_jsii_.Get(
		j,
		"buildStatusConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GitCloneDepth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GitCloneDepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gitCloneDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GitSubmodulesConfig() CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference {
	var returns CodebuildProjectSecondarySourcesGitSubmodulesConfigOutputReference
	_jsii_.Get(
		j,
		"gitSubmodulesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GitSubmodulesConfigInput() *CodebuildProjectSecondarySourcesGitSubmodulesConfig {
	var returns *CodebuildProjectSecondarySourcesGitSubmodulesConfig
	_jsii_.Get(
		j,
		"gitSubmodulesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) InsecureSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) InsecureSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ReportBuildStatus() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ReportBuildStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reportBuildStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) SourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) SourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewCodebuildProjectSecondarySourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CodebuildProjectSecondarySourcesOutputReference {
	_init_.Initialize()

	if err := validateNewCodebuildProjectSecondarySourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodebuildProjectSecondarySourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectSecondarySourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCodebuildProjectSecondarySourcesOutputReference_Override(c CodebuildProjectSecondarySourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codebuildProject.CodebuildProjectSecondarySourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetBuildspec(val *string) {
	if err := j.validateSetBuildspecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildspec",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetGitCloneDepth(val *float64) {
	if err := j.validateSetGitCloneDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitCloneDepth",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetInsecureSsl(val interface{}) {
	if err := j.validateSetInsecureSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSsl",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetReportBuildStatus(val interface{}) {
	if err := j.validateSetReportBuildStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportBuildStatus",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetSourceIdentifier(val *string) {
	if err := j.validateSetSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) PutBuildStatusConfig(value *CodebuildProjectSecondarySourcesBuildStatusConfig) {
	if err := c.validatePutBuildStatusConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBuildStatusConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) PutGitSubmodulesConfig(value *CodebuildProjectSecondarySourcesGitSubmodulesConfig) {
	if err := c.validatePutGitSubmodulesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitSubmodulesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetBuildspec() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildspec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetBuildStatusConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildStatusConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetGitCloneDepth() {
	_jsii_.InvokeVoid(
		c,
		"resetGitCloneDepth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetGitSubmodulesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGitSubmodulesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetInsecureSsl() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecureSsl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ResetReportBuildStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetReportBuildStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodebuildProjectSecondarySourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

