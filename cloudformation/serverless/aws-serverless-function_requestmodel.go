package serverless

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Function_RequestModel AWS CloudFormation Resource (AWS::Serverless::Function.RequestModel)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#function-request-model-object
type Function_RequestModel struct {

	// Model AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#function-request-model-object
	Model string `json:"Model,omitempty"`

	// Required AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#function-request-model-object
	Required bool `json:"Required,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_RequestModel) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.RequestModel"
}
