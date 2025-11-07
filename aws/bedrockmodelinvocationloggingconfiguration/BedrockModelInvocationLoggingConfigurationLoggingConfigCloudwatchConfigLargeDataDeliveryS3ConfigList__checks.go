// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package bedrockmodelinvocationloggingconfiguration

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config:
		val := val.(*[]*BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config:
		val_ := val.([]*BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3Config; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockModelInvocationLoggingConfigurationLoggingConfigCloudwatchConfigLargeDataDeliveryS3ConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

