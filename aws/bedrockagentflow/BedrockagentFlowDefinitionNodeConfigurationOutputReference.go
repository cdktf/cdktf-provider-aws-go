// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockagentFlowDefinitionNodeConfigurationOutputReference interface {
	cdktf.ComplexObject
	Agent() BedrockagentFlowDefinitionNodeConfigurationAgentList
	AgentInput() interface{}
	Collector() BedrockagentFlowDefinitionNodeConfigurationCollectorList
	CollectorInput() interface{}
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
	Condition() BedrockagentFlowDefinitionNodeConfigurationConditionList
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InlineCode() BedrockagentFlowDefinitionNodeConfigurationInlineCodeList
	InlineCodeInput() interface{}
	Input() BedrockagentFlowDefinitionNodeConfigurationInputList
	InputInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iterator() BedrockagentFlowDefinitionNodeConfigurationIteratorList
	IteratorInput() interface{}
	KnowledgeBase() BedrockagentFlowDefinitionNodeConfigurationKnowledgeBaseList
	KnowledgeBaseInput() interface{}
	LambdaFunction() BedrockagentFlowDefinitionNodeConfigurationLambdaFunctionList
	LambdaFunctionInput() interface{}
	Lex() BedrockagentFlowDefinitionNodeConfigurationLexList
	LexInput() interface{}
	Output() BedrockagentFlowDefinitionNodeConfigurationOutputList
	OutputInput() interface{}
	Prompt() BedrockagentFlowDefinitionNodeConfigurationPromptList
	PromptInput() interface{}
	Retrieval() BedrockagentFlowDefinitionNodeConfigurationRetrievalList
	RetrievalInput() interface{}
	Storage() BedrockagentFlowDefinitionNodeConfigurationStorageList
	StorageInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAgent(value interface{})
	PutCollector(value interface{})
	PutCondition(value interface{})
	PutInlineCode(value interface{})
	PutInput(value interface{})
	PutIterator(value interface{})
	PutKnowledgeBase(value interface{})
	PutLambdaFunction(value interface{})
	PutLex(value interface{})
	PutOutput(value interface{})
	PutPrompt(value interface{})
	PutRetrieval(value interface{})
	PutStorage(value interface{})
	ResetAgent()
	ResetCollector()
	ResetCondition()
	ResetInlineCode()
	ResetInput()
	ResetIterator()
	ResetKnowledgeBase()
	ResetLambdaFunction()
	ResetLex()
	ResetOutput()
	ResetPrompt()
	ResetRetrieval()
	ResetStorage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockagentFlowDefinitionNodeConfigurationOutputReference
type jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Agent() BedrockagentFlowDefinitionNodeConfigurationAgentList {
	var returns BedrockagentFlowDefinitionNodeConfigurationAgentList
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) AgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Collector() BedrockagentFlowDefinitionNodeConfigurationCollectorList {
	var returns BedrockagentFlowDefinitionNodeConfigurationCollectorList
	_jsii_.Get(
		j,
		"collector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) CollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Condition() BedrockagentFlowDefinitionNodeConfigurationConditionList {
	var returns BedrockagentFlowDefinitionNodeConfigurationConditionList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InlineCode() BedrockagentFlowDefinitionNodeConfigurationInlineCodeList {
	var returns BedrockagentFlowDefinitionNodeConfigurationInlineCodeList
	_jsii_.Get(
		j,
		"inlineCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InlineCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Input() BedrockagentFlowDefinitionNodeConfigurationInputList {
	var returns BedrockagentFlowDefinitionNodeConfigurationInputList
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Iterator() BedrockagentFlowDefinitionNodeConfigurationIteratorList {
	var returns BedrockagentFlowDefinitionNodeConfigurationIteratorList
	_jsii_.Get(
		j,
		"iterator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) IteratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iteratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) KnowledgeBase() BedrockagentFlowDefinitionNodeConfigurationKnowledgeBaseList {
	var returns BedrockagentFlowDefinitionNodeConfigurationKnowledgeBaseList
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) KnowledgeBaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) LambdaFunction() BedrockagentFlowDefinitionNodeConfigurationLambdaFunctionList {
	var returns BedrockagentFlowDefinitionNodeConfigurationLambdaFunctionList
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) LambdaFunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Lex() BedrockagentFlowDefinitionNodeConfigurationLexList {
	var returns BedrockagentFlowDefinitionNodeConfigurationLexList
	_jsii_.Get(
		j,
		"lex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) LexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Output() BedrockagentFlowDefinitionNodeConfigurationOutputList {
	var returns BedrockagentFlowDefinitionNodeConfigurationOutputList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Prompt() BedrockagentFlowDefinitionNodeConfigurationPromptList {
	var returns BedrockagentFlowDefinitionNodeConfigurationPromptList
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Retrieval() BedrockagentFlowDefinitionNodeConfigurationRetrievalList {
	var returns BedrockagentFlowDefinitionNodeConfigurationRetrievalList
	_jsii_.Get(
		j,
		"retrieval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) RetrievalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrievalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Storage() BedrockagentFlowDefinitionNodeConfigurationStorageList {
	var returns BedrockagentFlowDefinitionNodeConfigurationStorageList
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) StorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockagentFlowDefinitionNodeConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BedrockagentFlowDefinitionNodeConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockagentFlowDefinitionNodeConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentFlow.BedrockagentFlowDefinitionNodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBedrockagentFlowDefinitionNodeConfigurationOutputReference_Override(b BedrockagentFlowDefinitionNodeConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentFlow.BedrockagentFlowDefinitionNodeConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutAgent(value interface{}) {
	if err := b.validatePutAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgent",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutCollector(value interface{}) {
	if err := b.validatePutCollectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCollector",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutCondition(value interface{}) {
	if err := b.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putCondition",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutInlineCode(value interface{}) {
	if err := b.validatePutInlineCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInlineCode",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutInput(value interface{}) {
	if err := b.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putInput",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutIterator(value interface{}) {
	if err := b.validatePutIteratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putIterator",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutKnowledgeBase(value interface{}) {
	if err := b.validatePutKnowledgeBaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putKnowledgeBase",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutLambdaFunction(value interface{}) {
	if err := b.validatePutLambdaFunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLambdaFunction",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutLex(value interface{}) {
	if err := b.validatePutLexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putLex",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutOutput(value interface{}) {
	if err := b.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putOutput",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutPrompt(value interface{}) {
	if err := b.validatePutPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPrompt",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutRetrieval(value interface{}) {
	if err := b.validatePutRetrievalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetrieval",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) PutStorage(value interface{}) {
	if err := b.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putStorage",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetCollector() {
	_jsii_.InvokeVoid(
		b,
		"resetCollector",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		b,
		"resetCondition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetInlineCode() {
	_jsii_.InvokeVoid(
		b,
		"resetInlineCode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		b,
		"resetInput",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetIterator() {
	_jsii_.InvokeVoid(
		b,
		"resetIterator",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetKnowledgeBase() {
	_jsii_.InvokeVoid(
		b,
		"resetKnowledgeBase",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetLambdaFunction() {
	_jsii_.InvokeVoid(
		b,
		"resetLambdaFunction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetLex() {
	_jsii_.InvokeVoid(
		b,
		"resetLex",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		b,
		"resetOutput",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetPrompt() {
	_jsii_.InvokeVoid(
		b,
		"resetPrompt",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetRetrieval() {
	_jsii_.InvokeVoid(
		b,
		"resetRetrieval",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ResetStorage() {
	_jsii_.InvokeVoid(
		b,
		"resetStorage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

