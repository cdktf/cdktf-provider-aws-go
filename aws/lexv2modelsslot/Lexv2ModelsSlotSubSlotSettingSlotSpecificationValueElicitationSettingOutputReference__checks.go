// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package lexv2modelsslot

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validatePutDefaultValueSpecificationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification:
		value := value.(*[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification:
		value_ := value.([]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingDefaultValueSpecification; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validatePutPromptSpecificationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification:
		value := value.(*[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification:
		value_ := value.([]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingPromptSpecification; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validatePutSampleUtteranceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance:
		value := value.(*[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance:
		value_ := value.([]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingSampleUtterance; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validatePutWaitAndContinueSpecificationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification:
		value := value.(*[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification:
		value_ := value.([]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingWaitAndContinueSpecification; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (l *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting:
		val := val.(*Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting:
		val_ := val.(Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSetting; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Lexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewLexv2ModelsSlotSubSlotSettingSlotSpecificationValueElicitationSettingOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

