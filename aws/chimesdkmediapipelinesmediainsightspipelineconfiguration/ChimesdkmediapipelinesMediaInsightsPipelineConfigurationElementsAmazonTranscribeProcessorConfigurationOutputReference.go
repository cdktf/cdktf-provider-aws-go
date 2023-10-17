// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/chimesdkmediapipelinesmediainsightspipelineconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference interface {
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
	ContentIdentificationType() *string
	SetContentIdentificationType(val *string)
	ContentIdentificationTypeInput() *string
	ContentRedactionType() *string
	SetContentRedactionType(val *string)
	ContentRedactionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnablePartialResultsStabilization() interface{}
	SetEnablePartialResultsStabilization(val interface{})
	EnablePartialResultsStabilizationInput() interface{}
	FilterPartialResults() interface{}
	SetFilterPartialResults(val interface{})
	FilterPartialResultsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration
	SetInternalValue(val *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration)
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	LanguageModelName() *string
	SetLanguageModelName(val *string)
	LanguageModelNameInput() *string
	PartialResultsStability() *string
	SetPartialResultsStability(val *string)
	PartialResultsStabilityInput() *string
	PiiEntityTypes() *string
	SetPiiEntityTypes(val *string)
	PiiEntityTypesInput() *string
	ShowSpeakerLabel() interface{}
	SetShowSpeakerLabel(val interface{})
	ShowSpeakerLabelInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VocabularyFilterMethod() *string
	SetVocabularyFilterMethod(val *string)
	VocabularyFilterMethodInput() *string
	VocabularyFilterName() *string
	SetVocabularyFilterName(val *string)
	VocabularyFilterNameInput() *string
	VocabularyName() *string
	SetVocabularyName(val *string)
	VocabularyNameInput() *string
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
	ResetContentIdentificationType()
	ResetContentRedactionType()
	ResetEnablePartialResultsStabilization()
	ResetFilterPartialResults()
	ResetLanguageModelName()
	ResetPartialResultsStability()
	ResetPiiEntityTypes()
	ResetShowSpeakerLabel()
	ResetVocabularyFilterMethod()
	ResetVocabularyFilterName()
	ResetVocabularyName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference
type jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentIdentificationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentIdentificationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdentificationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentRedactionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ContentRedactionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentRedactionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) EnablePartialResultsStabilization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) EnablePartialResultsStabilizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePartialResultsStabilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) FilterPartialResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) FilterPartialResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterPartialResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InternalValue() *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration {
	var returns *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) LanguageModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PartialResultsStability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PartialResultsStabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partialResultsStabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PiiEntityTypes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) PiiEntityTypesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"piiEntityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ShowSpeakerLabel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ShowSpeakerLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"showSpeakerLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) VocabularyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vocabularyNameInput",
		&returns,
	)
	return returns
}


func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference_Override(c ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.chimesdkmediapipelinesMediaInsightsPipelineConfiguration.ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetContentIdentificationType(val *string) {
	if err := j.validateSetContentIdentificationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentIdentificationType",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetContentRedactionType(val *string) {
	if err := j.validateSetContentRedactionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentRedactionType",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetEnablePartialResultsStabilization(val interface{}) {
	if err := j.validateSetEnablePartialResultsStabilizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePartialResultsStabilization",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetFilterPartialResults(val interface{}) {
	if err := j.validateSetFilterPartialResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterPartialResults",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetInternalValue(val *ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetLanguageModelName(val *string) {
	if err := j.validateSetLanguageModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageModelName",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetPartialResultsStability(val *string) {
	if err := j.validateSetPartialResultsStabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partialResultsStability",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetPiiEntityTypes(val *string) {
	if err := j.validateSetPiiEntityTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"piiEntityTypes",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetShowSpeakerLabel(val interface{}) {
	if err := j.validateSetShowSpeakerLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"showSpeakerLabel",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyFilterMethod(val *string) {
	if err := j.validateSetVocabularyFilterMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterMethod",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyFilterName(val *string) {
	if err := j.validateSetVocabularyFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyFilterName",
		val,
	)
}

func (j *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference)SetVocabularyName(val *string) {
	if err := j.validateSetVocabularyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vocabularyName",
		val,
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetContentIdentificationType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentIdentificationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetContentRedactionType() {
	_jsii_.InvokeVoid(
		c,
		"resetContentRedactionType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetEnablePartialResultsStabilization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePartialResultsStabilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetFilterPartialResults() {
	_jsii_.InvokeVoid(
		c,
		"resetFilterPartialResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetLanguageModelName() {
	_jsii_.InvokeVoid(
		c,
		"resetLanguageModelName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetPartialResultsStability() {
	_jsii_.InvokeVoid(
		c,
		"resetPartialResultsStability",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetPiiEntityTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetPiiEntityTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetShowSpeakerLabel() {
	_jsii_.InvokeVoid(
		c,
		"resetShowSpeakerLabel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyFilterMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyFilterName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyFilterName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ResetVocabularyName() {
	_jsii_.InvokeVoid(
		c,
		"resetVocabularyName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeProcessorConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

