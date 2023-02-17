package codepipelinecustomactiontype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v12/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v12/codepipelinecustomactiontype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodepipelineCustomActionTypeSettingsOutputReference interface {
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
	EntityUrlTemplate() *string
	SetEntityUrlTemplate(val *string)
	EntityUrlTemplateInput() *string
	ExecutionUrlTemplate() *string
	SetExecutionUrlTemplate(val *string)
	ExecutionUrlTemplateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *CodepipelineCustomActionTypeSettings
	SetInternalValue(val *CodepipelineCustomActionTypeSettings)
	RevisionUrlTemplate() *string
	SetRevisionUrlTemplate(val *string)
	RevisionUrlTemplateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThirdPartyConfigurationUrl() *string
	SetThirdPartyConfigurationUrl(val *string)
	ThirdPartyConfigurationUrlInput() *string
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
	ResetEntityUrlTemplate()
	ResetExecutionUrlTemplate()
	ResetRevisionUrlTemplate()
	ResetThirdPartyConfigurationUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CodepipelineCustomActionTypeSettingsOutputReference
type jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) EntityUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) EntityUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ExecutionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ExecutionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) InternalValue() *CodepipelineCustomActionTypeSettings {
	var returns *CodepipelineCustomActionTypeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) RevisionUrlTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionUrlTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) RevisionUrlTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revisionUrlTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ThirdPartyConfigurationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thirdPartyConfigurationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ThirdPartyConfigurationUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thirdPartyConfigurationUrlInput",
		&returns,
	)
	return returns
}


func NewCodepipelineCustomActionTypeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CodepipelineCustomActionTypeSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewCodepipelineCustomActionTypeSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.codepipelineCustomActionType.CodepipelineCustomActionTypeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCodepipelineCustomActionTypeSettingsOutputReference_Override(c CodepipelineCustomActionTypeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.codepipelineCustomActionType.CodepipelineCustomActionTypeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetEntityUrlTemplate(val *string) {
	if err := j.validateSetEntityUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetExecutionUrlTemplate(val *string) {
	if err := j.validateSetExecutionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetInternalValue(val *CodepipelineCustomActionTypeSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetRevisionUrlTemplate(val *string) {
	if err := j.validateSetRevisionUrlTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revisionUrlTemplate",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference)SetThirdPartyConfigurationUrl(val *string) {
	if err := j.validateSetThirdPartyConfigurationUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thirdPartyConfigurationUrl",
		val,
	)
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ResetEntityUrlTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetEntityUrlTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ResetExecutionUrlTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetExecutionUrlTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ResetRevisionUrlTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetRevisionUrlTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ResetThirdPartyConfigurationUrl() {
	_jsii_.InvokeVoid(
		c,
		"resetThirdPartyConfigurationUrl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CodepipelineCustomActionTypeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

