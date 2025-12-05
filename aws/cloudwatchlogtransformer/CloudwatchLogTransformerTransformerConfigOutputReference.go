// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogtransformer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/cloudwatchlogtransformer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchLogTransformerTransformerConfigOutputReference interface {
	cdktf.ComplexObject
	AddKeys() CloudwatchLogTransformerTransformerConfigAddKeysList
	AddKeysInput() interface{}
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
	CopyValue() CloudwatchLogTransformerTransformerConfigCopyValueList
	CopyValueInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csv() CloudwatchLogTransformerTransformerConfigCsvList
	CsvInput() interface{}
	DateTimeConverter() CloudwatchLogTransformerTransformerConfigDateTimeConverterList
	DateTimeConverterInput() interface{}
	DeleteKeys() CloudwatchLogTransformerTransformerConfigDeleteKeysList
	DeleteKeysInput() interface{}
	// Experimental.
	Fqn() *string
	Grok() CloudwatchLogTransformerTransformerConfigGrokList
	GrokInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListToMap() CloudwatchLogTransformerTransformerConfigListToMapList
	ListToMapInput() interface{}
	LowerCaseString() CloudwatchLogTransformerTransformerConfigLowerCaseStringList
	LowerCaseStringInput() interface{}
	MoveKeys() CloudwatchLogTransformerTransformerConfigMoveKeysList
	MoveKeysInput() interface{}
	ParseCloudfront() CloudwatchLogTransformerTransformerConfigParseCloudfrontList
	ParseCloudfrontInput() interface{}
	ParseJson() CloudwatchLogTransformerTransformerConfigParseJsonList
	ParseJsonInput() interface{}
	ParseKeyValue() CloudwatchLogTransformerTransformerConfigParseKeyValueList
	ParseKeyValueInput() interface{}
	ParsePostgres() CloudwatchLogTransformerTransformerConfigParsePostgresList
	ParsePostgresInput() interface{}
	ParseRoute53() CloudwatchLogTransformerTransformerConfigParseRoute53List
	ParseRoute53Input() interface{}
	ParseToOcsf() CloudwatchLogTransformerTransformerConfigParseToOcsfList
	ParseToOcsfInput() interface{}
	ParseVpc() CloudwatchLogTransformerTransformerConfigParseVpcList
	ParseVpcInput() interface{}
	ParseWaf() CloudwatchLogTransformerTransformerConfigParseWafList
	ParseWafInput() interface{}
	RenameKeys() CloudwatchLogTransformerTransformerConfigRenameKeysList
	RenameKeysInput() interface{}
	SplitString() CloudwatchLogTransformerTransformerConfigSplitStringList
	SplitStringInput() interface{}
	SubstituteString() CloudwatchLogTransformerTransformerConfigSubstituteStringList
	SubstituteStringInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrimString() CloudwatchLogTransformerTransformerConfigTrimStringList
	TrimStringInput() interface{}
	TypeConverter() CloudwatchLogTransformerTransformerConfigTypeConverterList
	TypeConverterInput() interface{}
	UpperCaseString() CloudwatchLogTransformerTransformerConfigUpperCaseStringList
	UpperCaseStringInput() interface{}
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
	PutAddKeys(value interface{})
	PutCopyValue(value interface{})
	PutCsv(value interface{})
	PutDateTimeConverter(value interface{})
	PutDeleteKeys(value interface{})
	PutGrok(value interface{})
	PutListToMap(value interface{})
	PutLowerCaseString(value interface{})
	PutMoveKeys(value interface{})
	PutParseCloudfront(value interface{})
	PutParseJson(value interface{})
	PutParseKeyValue(value interface{})
	PutParsePostgres(value interface{})
	PutParseRoute53(value interface{})
	PutParseToOcsf(value interface{})
	PutParseVpc(value interface{})
	PutParseWaf(value interface{})
	PutRenameKeys(value interface{})
	PutSplitString(value interface{})
	PutSubstituteString(value interface{})
	PutTrimString(value interface{})
	PutTypeConverter(value interface{})
	PutUpperCaseString(value interface{})
	ResetAddKeys()
	ResetCopyValue()
	ResetCsv()
	ResetDateTimeConverter()
	ResetDeleteKeys()
	ResetGrok()
	ResetListToMap()
	ResetLowerCaseString()
	ResetMoveKeys()
	ResetParseCloudfront()
	ResetParseJson()
	ResetParseKeyValue()
	ResetParsePostgres()
	ResetParseRoute53()
	ResetParseToOcsf()
	ResetParseVpc()
	ResetParseWaf()
	ResetRenameKeys()
	ResetSplitString()
	ResetSubstituteString()
	ResetTrimString()
	ResetTypeConverter()
	ResetUpperCaseString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudwatchLogTransformerTransformerConfigOutputReference
type jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) AddKeys() CloudwatchLogTransformerTransformerConfigAddKeysList {
	var returns CloudwatchLogTransformerTransformerConfigAddKeysList
	_jsii_.Get(
		j,
		"addKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) AddKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) CopyValue() CloudwatchLogTransformerTransformerConfigCopyValueList {
	var returns CloudwatchLogTransformerTransformerConfigCopyValueList
	_jsii_.Get(
		j,
		"copyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) CopyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) Csv() CloudwatchLogTransformerTransformerConfigCsvList {
	var returns CloudwatchLogTransformerTransformerConfigCsvList
	_jsii_.Get(
		j,
		"csv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) CsvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"csvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) DateTimeConverter() CloudwatchLogTransformerTransformerConfigDateTimeConverterList {
	var returns CloudwatchLogTransformerTransformerConfigDateTimeConverterList
	_jsii_.Get(
		j,
		"dateTimeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) DateTimeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) DeleteKeys() CloudwatchLogTransformerTransformerConfigDeleteKeysList {
	var returns CloudwatchLogTransformerTransformerConfigDeleteKeysList
	_jsii_.Get(
		j,
		"deleteKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) DeleteKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) Grok() CloudwatchLogTransformerTransformerConfigGrokList {
	var returns CloudwatchLogTransformerTransformerConfigGrokList
	_jsii_.Get(
		j,
		"grok",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GrokInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grokInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ListToMap() CloudwatchLogTransformerTransformerConfigListToMapList {
	var returns CloudwatchLogTransformerTransformerConfigListToMapList
	_jsii_.Get(
		j,
		"listToMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ListToMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listToMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) LowerCaseString() CloudwatchLogTransformerTransformerConfigLowerCaseStringList {
	var returns CloudwatchLogTransformerTransformerConfigLowerCaseStringList
	_jsii_.Get(
		j,
		"lowerCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) LowerCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowerCaseStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) MoveKeys() CloudwatchLogTransformerTransformerConfigMoveKeysList {
	var returns CloudwatchLogTransformerTransformerConfigMoveKeysList
	_jsii_.Get(
		j,
		"moveKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) MoveKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"moveKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseCloudfront() CloudwatchLogTransformerTransformerConfigParseCloudfrontList {
	var returns CloudwatchLogTransformerTransformerConfigParseCloudfrontList
	_jsii_.Get(
		j,
		"parseCloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseCloudfrontInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseCloudfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseJson() CloudwatchLogTransformerTransformerConfigParseJsonList {
	var returns CloudwatchLogTransformerTransformerConfigParseJsonList
	_jsii_.Get(
		j,
		"parseJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseJsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseKeyValue() CloudwatchLogTransformerTransformerConfigParseKeyValueList {
	var returns CloudwatchLogTransformerTransformerConfigParseKeyValueList
	_jsii_.Get(
		j,
		"parseKeyValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseKeyValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseKeyValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParsePostgres() CloudwatchLogTransformerTransformerConfigParsePostgresList {
	var returns CloudwatchLogTransformerTransformerConfigParsePostgresList
	_jsii_.Get(
		j,
		"parsePostgres",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParsePostgresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parsePostgresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseRoute53() CloudwatchLogTransformerTransformerConfigParseRoute53List {
	var returns CloudwatchLogTransformerTransformerConfigParseRoute53List
	_jsii_.Get(
		j,
		"parseRoute53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseRoute53Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseRoute53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseToOcsf() CloudwatchLogTransformerTransformerConfigParseToOcsfList {
	var returns CloudwatchLogTransformerTransformerConfigParseToOcsfList
	_jsii_.Get(
		j,
		"parseToOcsf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseToOcsfInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseToOcsfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseVpc() CloudwatchLogTransformerTransformerConfigParseVpcList {
	var returns CloudwatchLogTransformerTransformerConfigParseVpcList
	_jsii_.Get(
		j,
		"parseVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseVpcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseVpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseWaf() CloudwatchLogTransformerTransformerConfigParseWafList {
	var returns CloudwatchLogTransformerTransformerConfigParseWafList
	_jsii_.Get(
		j,
		"parseWaf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ParseWafInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseWafInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) RenameKeys() CloudwatchLogTransformerTransformerConfigRenameKeysList {
	var returns CloudwatchLogTransformerTransformerConfigRenameKeysList
	_jsii_.Get(
		j,
		"renameKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) RenameKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renameKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) SplitString() CloudwatchLogTransformerTransformerConfigSplitStringList {
	var returns CloudwatchLogTransformerTransformerConfigSplitStringList
	_jsii_.Get(
		j,
		"splitString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) SplitStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) SubstituteString() CloudwatchLogTransformerTransformerConfigSubstituteStringList {
	var returns CloudwatchLogTransformerTransformerConfigSubstituteStringList
	_jsii_.Get(
		j,
		"substituteString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) SubstituteStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"substituteStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TrimString() CloudwatchLogTransformerTransformerConfigTrimStringList {
	var returns CloudwatchLogTransformerTransformerConfigTrimStringList
	_jsii_.Get(
		j,
		"trimString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TrimStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trimStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TypeConverter() CloudwatchLogTransformerTransformerConfigTypeConverterList {
	var returns CloudwatchLogTransformerTransformerConfigTypeConverterList
	_jsii_.Get(
		j,
		"typeConverter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) TypeConverterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeConverterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) UpperCaseString() CloudwatchLogTransformerTransformerConfigUpperCaseStringList {
	var returns CloudwatchLogTransformerTransformerConfigUpperCaseStringList
	_jsii_.Get(
		j,
		"upperCaseString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) UpperCaseStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upperCaseStringInput",
		&returns,
	)
	return returns
}


func NewCloudwatchLogTransformerTransformerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudwatchLogTransformerTransformerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudwatchLogTransformerTransformerConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchLogTransformer.CloudwatchLogTransformerTransformerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudwatchLogTransformerTransformerConfigOutputReference_Override(c CloudwatchLogTransformerTransformerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudwatchLogTransformer.CloudwatchLogTransformerTransformerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutAddKeys(value interface{}) {
	if err := c.validatePutAddKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAddKeys",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutCopyValue(value interface{}) {
	if err := c.validatePutCopyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCopyValue",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutCsv(value interface{}) {
	if err := c.validatePutCsvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCsv",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutDateTimeConverter(value interface{}) {
	if err := c.validatePutDateTimeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDateTimeConverter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutDeleteKeys(value interface{}) {
	if err := c.validatePutDeleteKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeleteKeys",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutGrok(value interface{}) {
	if err := c.validatePutGrokParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrok",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutListToMap(value interface{}) {
	if err := c.validatePutListToMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putListToMap",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutLowerCaseString(value interface{}) {
	if err := c.validatePutLowerCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLowerCaseString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutMoveKeys(value interface{}) {
	if err := c.validatePutMoveKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMoveKeys",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseCloudfront(value interface{}) {
	if err := c.validatePutParseCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseCloudfront",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseJson(value interface{}) {
	if err := c.validatePutParseJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseJson",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseKeyValue(value interface{}) {
	if err := c.validatePutParseKeyValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseKeyValue",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParsePostgres(value interface{}) {
	if err := c.validatePutParsePostgresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParsePostgres",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseRoute53(value interface{}) {
	if err := c.validatePutParseRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseRoute53",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseToOcsf(value interface{}) {
	if err := c.validatePutParseToOcsfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseToOcsf",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseVpc(value interface{}) {
	if err := c.validatePutParseVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseVpc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutParseWaf(value interface{}) {
	if err := c.validatePutParseWafParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putParseWaf",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutRenameKeys(value interface{}) {
	if err := c.validatePutRenameKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRenameKeys",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutSplitString(value interface{}) {
	if err := c.validatePutSplitStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSplitString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutSubstituteString(value interface{}) {
	if err := c.validatePutSubstituteStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSubstituteString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutTrimString(value interface{}) {
	if err := c.validatePutTrimStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrimString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutTypeConverter(value interface{}) {
	if err := c.validatePutTypeConverterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTypeConverter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) PutUpperCaseString(value interface{}) {
	if err := c.validatePutUpperCaseStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpperCaseString",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetAddKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetAddKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetCopyValue() {
	_jsii_.InvokeVoid(
		c,
		"resetCopyValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetCsv() {
	_jsii_.InvokeVoid(
		c,
		"resetCsv",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetDateTimeConverter() {
	_jsii_.InvokeVoid(
		c,
		"resetDateTimeConverter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetDeleteKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetDeleteKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetGrok() {
	_jsii_.InvokeVoid(
		c,
		"resetGrok",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetListToMap() {
	_jsii_.InvokeVoid(
		c,
		"resetListToMap",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetLowerCaseString() {
	_jsii_.InvokeVoid(
		c,
		"resetLowerCaseString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetMoveKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetMoveKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseCloudfront() {
	_jsii_.InvokeVoid(
		c,
		"resetParseCloudfront",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseJson() {
	_jsii_.InvokeVoid(
		c,
		"resetParseJson",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseKeyValue() {
	_jsii_.InvokeVoid(
		c,
		"resetParseKeyValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParsePostgres() {
	_jsii_.InvokeVoid(
		c,
		"resetParsePostgres",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseRoute53() {
	_jsii_.InvokeVoid(
		c,
		"resetParseRoute53",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseToOcsf() {
	_jsii_.InvokeVoid(
		c,
		"resetParseToOcsf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseVpc() {
	_jsii_.InvokeVoid(
		c,
		"resetParseVpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetParseWaf() {
	_jsii_.InvokeVoid(
		c,
		"resetParseWaf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetRenameKeys() {
	_jsii_.InvokeVoid(
		c,
		"resetRenameKeys",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetSplitString() {
	_jsii_.InvokeVoid(
		c,
		"resetSplitString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetSubstituteString() {
	_jsii_.InvokeVoid(
		c,
		"resetSubstituteString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetTrimString() {
	_jsii_.InvokeVoid(
		c,
		"resetTrimString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetTypeConverter() {
	_jsii_.InvokeVoid(
		c,
		"resetTypeConverter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ResetUpperCaseString() {
	_jsii_.InvokeVoid(
		c,
		"resetUpperCaseString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudwatchLogTransformerTransformerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

