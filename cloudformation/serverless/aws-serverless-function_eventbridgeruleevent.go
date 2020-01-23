package serverless

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// Function_EventBridgeRuleEvent AWS CloudFormation Resource (AWS::Serverless::Function.EventBridgeRuleEvent)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#eventbridgerule
type Function_EventBridgeRuleEvent struct {

	// Input AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#eventbridgerule
	Input string `json:"Input,omitempty"`

	// InputPath AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#eventbridgerule
	InputPath string `json:"InputPath,omitempty"`

	// Pattern AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#eventbridgerule
	Pattern interface{} `json:"Pattern,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Function_EventBridgeRuleEvent) AWSCloudFormationType() string {
	return "AWS::Serverless::Function.EventBridgeRuleEvent"
}