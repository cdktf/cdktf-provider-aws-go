package emr


type EmrClusterEc2Attributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#instance_profile EmrCluster#instance_profile}.
	InstanceProfile *string `field:"required" json:"instanceProfile" yaml:"instanceProfile"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#additional_master_security_groups EmrCluster#additional_master_security_groups}.
	AdditionalMasterSecurityGroups *string `field:"optional" json:"additionalMasterSecurityGroups" yaml:"additionalMasterSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#additional_slave_security_groups EmrCluster#additional_slave_security_groups}.
	AdditionalSlaveSecurityGroups *string `field:"optional" json:"additionalSlaveSecurityGroups" yaml:"additionalSlaveSecurityGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#emr_managed_master_security_group EmrCluster#emr_managed_master_security_group}.
	EmrManagedMasterSecurityGroup *string `field:"optional" json:"emrManagedMasterSecurityGroup" yaml:"emrManagedMasterSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#emr_managed_slave_security_group EmrCluster#emr_managed_slave_security_group}.
	EmrManagedSlaveSecurityGroup *string `field:"optional" json:"emrManagedSlaveSecurityGroup" yaml:"emrManagedSlaveSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#key_name EmrCluster#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#service_access_security_group EmrCluster#service_access_security_group}.
	ServiceAccessSecurityGroup *string `field:"optional" json:"serviceAccessSecurityGroup" yaml:"serviceAccessSecurityGroup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#subnet_id EmrCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/emr_cluster#subnet_ids EmrCluster#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

