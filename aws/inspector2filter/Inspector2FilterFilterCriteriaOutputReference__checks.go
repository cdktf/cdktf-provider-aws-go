// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package inspector2filter

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutAwsAccountIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaAwsAccountId:
		value := value.(*[]*Inspector2FilterFilterCriteriaAwsAccountId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaAwsAccountId:
		value_ := value.([]*Inspector2FilterFilterCriteriaAwsAccountId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaAwsAccountId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityDetectorNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorName:
		value := value.(*[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorName:
		value_ := value.([]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityDetectorTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTags:
		value := value.(*[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTags:
		value_ := value.([]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityDetectorTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityFilePathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityFilePath:
		value := value.(*[]*Inspector2FilterFilterCriteriaCodeVulnerabilityFilePath)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaCodeVulnerabilityFilePath:
		value_ := value.([]*Inspector2FilterFilterCriteriaCodeVulnerabilityFilePath)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaCodeVulnerabilityFilePath; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutComponentIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaComponentId:
		value := value.(*[]*Inspector2FilterFilterCriteriaComponentId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaComponentId:
		value_ := value.([]*Inspector2FilterFilterCriteriaComponentId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaComponentId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutComponentTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaComponentType:
		value := value.(*[]*Inspector2FilterFilterCriteriaComponentType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaComponentType:
		value_ := value.([]*Inspector2FilterFilterCriteriaComponentType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaComponentType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEc2InstanceImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEc2InstanceImageId:
		value := value.(*[]*Inspector2FilterFilterCriteriaEc2InstanceImageId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEc2InstanceImageId:
		value_ := value.([]*Inspector2FilterFilterCriteriaEc2InstanceImageId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEc2InstanceImageId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEc2InstanceSubnetIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEc2InstanceSubnetId:
		value := value.(*[]*Inspector2FilterFilterCriteriaEc2InstanceSubnetId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEc2InstanceSubnetId:
		value_ := value.([]*Inspector2FilterFilterCriteriaEc2InstanceSubnetId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEc2InstanceSubnetId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEc2InstanceVpcIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEc2InstanceVpcId:
		value := value.(*[]*Inspector2FilterFilterCriteriaEc2InstanceVpcId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEc2InstanceVpcId:
		value_ := value.([]*Inspector2FilterFilterCriteriaEc2InstanceVpcId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEc2InstanceVpcId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImageArchitectureParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImageArchitecture:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImageArchitecture)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImageArchitecture:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImageArchitecture)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImageArchitecture; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImageHashParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImageHash:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImageHash)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImageHash:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImageHash)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImageHash; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImagePushedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImagePushedAt:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImagePushedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImagePushedAt:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImagePushedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImagePushedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImageRegistryParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImageRegistry:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImageRegistry)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImageRegistry:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImageRegistry)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImageRegistry; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImageRepositoryNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImageRepositoryName:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImageRepositoryName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImageRepositoryName:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImageRepositoryName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImageRepositoryName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEcrImageTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEcrImageTags:
		value := value.(*[]*Inspector2FilterFilterCriteriaEcrImageTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEcrImageTags:
		value_ := value.([]*Inspector2FilterFilterCriteriaEcrImageTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEcrImageTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutEpssScoreParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaEpssScore:
		value := value.(*[]*Inspector2FilterFilterCriteriaEpssScore)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaEpssScore:
		value_ := value.([]*Inspector2FilterFilterCriteriaEpssScore)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaEpssScore; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutExploitAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaExploitAvailable:
		value := value.(*[]*Inspector2FilterFilterCriteriaExploitAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaExploitAvailable:
		value_ := value.([]*Inspector2FilterFilterCriteriaExploitAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaExploitAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutFindingArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaFindingArn:
		value := value.(*[]*Inspector2FilterFilterCriteriaFindingArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaFindingArn:
		value_ := value.([]*Inspector2FilterFilterCriteriaFindingArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaFindingArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutFindingStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaFindingStatus:
		value := value.(*[]*Inspector2FilterFilterCriteriaFindingStatus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaFindingStatus:
		value_ := value.([]*Inspector2FilterFilterCriteriaFindingStatus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaFindingStatus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutFindingTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaFindingType:
		value := value.(*[]*Inspector2FilterFilterCriteriaFindingType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaFindingType:
		value_ := value.([]*Inspector2FilterFilterCriteriaFindingType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaFindingType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutFirstObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaFirstObservedAt:
		value := value.(*[]*Inspector2FilterFilterCriteriaFirstObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaFirstObservedAt:
		value_ := value.([]*Inspector2FilterFilterCriteriaFirstObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaFirstObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutFixAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaFixAvailable:
		value := value.(*[]*Inspector2FilterFilterCriteriaFixAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaFixAvailable:
		value_ := value.([]*Inspector2FilterFilterCriteriaFixAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaFixAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutInspectorScoreParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaInspectorScore:
		value := value.(*[]*Inspector2FilterFilterCriteriaInspectorScore)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaInspectorScore:
		value_ := value.([]*Inspector2FilterFilterCriteriaInspectorScore)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaInspectorScore; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionExecutionRoleArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArn:
		value := value.(*[]*Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArn:
		value_ := value.([]*Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLambdaFunctionExecutionRoleArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionLastModifiedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAt:
		value := value.(*[]*Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAt:
		value_ := value.([]*Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLambdaFunctionLastModifiedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionLayersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLambdaFunctionLayers:
		value := value.(*[]*Inspector2FilterFilterCriteriaLambdaFunctionLayers)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLambdaFunctionLayers:
		value_ := value.([]*Inspector2FilterFilterCriteriaLambdaFunctionLayers)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLambdaFunctionLayers; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLambdaFunctionName:
		value := value.(*[]*Inspector2FilterFilterCriteriaLambdaFunctionName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLambdaFunctionName:
		value_ := value.([]*Inspector2FilterFilterCriteriaLambdaFunctionName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLambdaFunctionName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionRuntimeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLambdaFunctionRuntime:
		value := value.(*[]*Inspector2FilterFilterCriteriaLambdaFunctionRuntime)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLambdaFunctionRuntime:
		value_ := value.([]*Inspector2FilterFilterCriteriaLambdaFunctionRuntime)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLambdaFunctionRuntime; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaLastObservedAt:
		value := value.(*[]*Inspector2FilterFilterCriteriaLastObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaLastObservedAt:
		value_ := value.([]*Inspector2FilterFilterCriteriaLastObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaLastObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutNetworkProtocolParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaNetworkProtocol:
		value := value.(*[]*Inspector2FilterFilterCriteriaNetworkProtocol)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaNetworkProtocol:
		value_ := value.([]*Inspector2FilterFilterCriteriaNetworkProtocol)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaNetworkProtocol; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutPortRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaPortRange:
		value := value.(*[]*Inspector2FilterFilterCriteriaPortRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaPortRange:
		value_ := value.([]*Inspector2FilterFilterCriteriaPortRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaPortRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutRelatedVulnerabilitiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaRelatedVulnerabilities:
		value := value.(*[]*Inspector2FilterFilterCriteriaRelatedVulnerabilities)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaRelatedVulnerabilities:
		value_ := value.([]*Inspector2FilterFilterCriteriaRelatedVulnerabilities)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaRelatedVulnerabilities; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutResourceIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaResourceId:
		value := value.(*[]*Inspector2FilterFilterCriteriaResourceId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaResourceId:
		value_ := value.([]*Inspector2FilterFilterCriteriaResourceId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaResourceId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutResourceTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaResourceTags:
		value := value.(*[]*Inspector2FilterFilterCriteriaResourceTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaResourceTags:
		value_ := value.([]*Inspector2FilterFilterCriteriaResourceTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaResourceTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutResourceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaResourceType:
		value := value.(*[]*Inspector2FilterFilterCriteriaResourceType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaResourceType:
		value_ := value.([]*Inspector2FilterFilterCriteriaResourceType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaResourceType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutSeverityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaSeverity:
		value := value.(*[]*Inspector2FilterFilterCriteriaSeverity)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaSeverity:
		value_ := value.([]*Inspector2FilterFilterCriteriaSeverity)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaSeverity; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutTitleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaTitle:
		value := value.(*[]*Inspector2FilterFilterCriteriaTitle)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaTitle:
		value_ := value.([]*Inspector2FilterFilterCriteriaTitle)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaTitle; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaUpdatedAt:
		value := value.(*[]*Inspector2FilterFilterCriteriaUpdatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaUpdatedAt:
		value_ := value.([]*Inspector2FilterFilterCriteriaUpdatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaUpdatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutVendorSeverityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaVendorSeverity:
		value := value.(*[]*Inspector2FilterFilterCriteriaVendorSeverity)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaVendorSeverity:
		value_ := value.([]*Inspector2FilterFilterCriteriaVendorSeverity)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaVendorSeverity; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutVulnerabilityIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaVulnerabilityId:
		value := value.(*[]*Inspector2FilterFilterCriteriaVulnerabilityId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaVulnerabilityId:
		value_ := value.([]*Inspector2FilterFilterCriteriaVulnerabilityId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaVulnerabilityId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutVulnerabilitySourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaVulnerabilitySource:
		value := value.(*[]*Inspector2FilterFilterCriteriaVulnerabilitySource)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaVulnerabilitySource:
		value_ := value.([]*Inspector2FilterFilterCriteriaVulnerabilitySource)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaVulnerabilitySource; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validatePutVulnerablePackagesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Inspector2FilterFilterCriteriaVulnerablePackages:
		value := value.(*[]*Inspector2FilterFilterCriteriaVulnerablePackages)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspector2FilterFilterCriteriaVulnerablePackages:
		value_ := value.([]*Inspector2FilterFilterCriteriaVulnerablePackages)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Inspector2FilterFilterCriteriaVulnerablePackages; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Inspector2FilterFilterCriteria:
		val := val.(*Inspector2FilterFilterCriteria)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Inspector2FilterFilterCriteria:
		val_ := val.(Inspector2FilterFilterCriteria)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Inspector2FilterFilterCriteria; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspector2FilterFilterCriteriaOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewInspector2FilterFilterCriteriaOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

