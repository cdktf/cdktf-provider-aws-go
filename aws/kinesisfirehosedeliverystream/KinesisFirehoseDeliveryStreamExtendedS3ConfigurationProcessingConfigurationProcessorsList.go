// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/kinesisfirehosedeliverystream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList
type jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList{}

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList_Override(k KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.kinesisFirehoseDeliveryStream.KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := k.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		k,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) Get(index *float64) KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamExtendedS3ConfigurationProcessingConfigurationProcessorsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

