//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package secretsmanagersecret

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsmanagerSecretReplicaList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecretsmanagerSecretReplicaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretReplicaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretReplicaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretReplicaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecretsmanagerSecretReplicaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecretsmanagerSecretReplicaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

