//go:build !no_runtime_type_checking

package kinesisfirehosedeliverystream

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validatePutProcessorsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessors:
		value := value.(*[]*KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessors)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessors:
		value_ := value.([]*KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessors)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationProcessors; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetEnabledParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetInternalValueParameters(val *KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfiguration) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_KinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisFirehoseDeliveryStreamHttpEndpointConfigurationProcessingConfigurationOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

