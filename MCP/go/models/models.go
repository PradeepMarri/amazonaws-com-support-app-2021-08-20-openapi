package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DeleteSlackWorkspaceConfigurationResult represents the DeleteSlackWorkspaceConfigurationResult schema from the OpenAPI specification
type DeleteSlackWorkspaceConfigurationResult struct {
}

// SlackChannelConfiguration represents the SlackChannelConfiguration schema from the OpenAPI specification
type SlackChannelConfiguration struct {
	Notifyoncaseseverity interface{} `json:"notifyOnCaseSeverity,omitempty"`
	Notifyoncreateorreopencase interface{} `json:"notifyOnCreateOrReopenCase,omitempty"`
	Notifyonresolvecase interface{} `json:"notifyOnResolveCase,omitempty"`
	Teamid interface{} `json:"teamId"`
	Channelid interface{} `json:"channelId"`
	Channelname interface{} `json:"channelName,omitempty"`
	Channelrolearn interface{} `json:"channelRoleArn,omitempty"`
	Notifyonaddcorrespondencetocase interface{} `json:"notifyOnAddCorrespondenceToCase,omitempty"`
}

// DeleteAccountAliasResult represents the DeleteAccountAliasResult schema from the OpenAPI specification
type DeleteAccountAliasResult struct {
}

// GetAccountAliasRequest represents the GetAccountAliasRequest schema from the OpenAPI specification
type GetAccountAliasRequest struct {
}

// DeleteAccountAliasRequest represents the DeleteAccountAliasRequest schema from the OpenAPI specification
type DeleteAccountAliasRequest struct {
}

// RegisterSlackWorkspaceForOrganizationResult represents the RegisterSlackWorkspaceForOrganizationResult schema from the OpenAPI specification
type RegisterSlackWorkspaceForOrganizationResult struct {
	Teamname interface{} `json:"teamName,omitempty"`
	Accounttype interface{} `json:"accountType,omitempty"`
	Teamid interface{} `json:"teamId,omitempty"`
}

// GetAccountAliasResult represents the GetAccountAliasResult schema from the OpenAPI specification
type GetAccountAliasResult struct {
	Accountalias interface{} `json:"accountAlias,omitempty"`
}

// UpdateSlackChannelConfigurationRequest represents the UpdateSlackChannelConfigurationRequest schema from the OpenAPI specification
type UpdateSlackChannelConfigurationRequest struct {
	Teamid interface{} `json:"teamId"`
	Channelid interface{} `json:"channelId"`
	Channelname interface{} `json:"channelName,omitempty"`
	Channelrolearn interface{} `json:"channelRoleArn,omitempty"`
	Notifyonaddcorrespondencetocase interface{} `json:"notifyOnAddCorrespondenceToCase,omitempty"`
	Notifyoncaseseverity interface{} `json:"notifyOnCaseSeverity,omitempty"`
	Notifyoncreateorreopencase interface{} `json:"notifyOnCreateOrReopenCase,omitempty"`
	Notifyonresolvecase interface{} `json:"notifyOnResolveCase,omitempty"`
}

// UpdateSlackChannelConfigurationResult represents the UpdateSlackChannelConfigurationResult schema from the OpenAPI specification
type UpdateSlackChannelConfigurationResult struct {
	Channelrolearn interface{} `json:"channelRoleArn,omitempty"`
	Notifyonaddcorrespondencetocase interface{} `json:"notifyOnAddCorrespondenceToCase,omitempty"`
	Notifyoncaseseverity interface{} `json:"notifyOnCaseSeverity,omitempty"`
	Notifyoncreateorreopencase interface{} `json:"notifyOnCreateOrReopenCase,omitempty"`
	Notifyonresolvecase interface{} `json:"notifyOnResolveCase,omitempty"`
	Teamid interface{} `json:"teamId,omitempty"`
	Channelid interface{} `json:"channelId,omitempty"`
	Channelname interface{} `json:"channelName,omitempty"`
}

// ListSlackChannelConfigurationsRequest represents the ListSlackChannelConfigurationsRequest schema from the OpenAPI specification
type ListSlackChannelConfigurationsRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// PutAccountAliasResult represents the PutAccountAliasResult schema from the OpenAPI specification
type PutAccountAliasResult struct {
}

// DeleteSlackWorkspaceConfigurationRequest represents the DeleteSlackWorkspaceConfigurationRequest schema from the OpenAPI specification
type DeleteSlackWorkspaceConfigurationRequest struct {
	Teamid interface{} `json:"teamId"`
}

// SlackWorkspaceConfiguration represents the SlackWorkspaceConfiguration schema from the OpenAPI specification
type SlackWorkspaceConfiguration struct {
	Alloworganizationmemberaccount interface{} `json:"allowOrganizationMemberAccount,omitempty"`
	Teamid interface{} `json:"teamId"`
	Teamname interface{} `json:"teamName,omitempty"`
}

// ListSlackWorkspaceConfigurationsResult represents the ListSlackWorkspaceConfigurationsResult schema from the OpenAPI specification
type ListSlackWorkspaceConfigurationsResult struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Slackworkspaceconfigurations interface{} `json:"slackWorkspaceConfigurations,omitempty"`
}

// ListSlackChannelConfigurationsResult represents the ListSlackChannelConfigurationsResult schema from the OpenAPI specification
type ListSlackChannelConfigurationsResult struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Slackchannelconfigurations interface{} `json:"slackChannelConfigurations"`
}

// CreateSlackChannelConfigurationResult represents the CreateSlackChannelConfigurationResult schema from the OpenAPI specification
type CreateSlackChannelConfigurationResult struct {
}

// RegisterSlackWorkspaceForOrganizationRequest represents the RegisterSlackWorkspaceForOrganizationRequest schema from the OpenAPI specification
type RegisterSlackWorkspaceForOrganizationRequest struct {
	Teamid interface{} `json:"teamId"`
}

// CreateSlackChannelConfigurationRequest represents the CreateSlackChannelConfigurationRequest schema from the OpenAPI specification
type CreateSlackChannelConfigurationRequest struct {
	Channelrolearn interface{} `json:"channelRoleArn"`
	Notifyonaddcorrespondencetocase interface{} `json:"notifyOnAddCorrespondenceToCase,omitempty"`
	Notifyoncaseseverity interface{} `json:"notifyOnCaseSeverity"`
	Notifyoncreateorreopencase interface{} `json:"notifyOnCreateOrReopenCase,omitempty"`
	Notifyonresolvecase interface{} `json:"notifyOnResolveCase,omitempty"`
	Teamid interface{} `json:"teamId"`
	Channelid interface{} `json:"channelId"`
	Channelname interface{} `json:"channelName,omitempty"`
}

// ListSlackWorkspaceConfigurationsRequest represents the ListSlackWorkspaceConfigurationsRequest schema from the OpenAPI specification
type ListSlackWorkspaceConfigurationsRequest struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteSlackChannelConfigurationRequest represents the DeleteSlackChannelConfigurationRequest schema from the OpenAPI specification
type DeleteSlackChannelConfigurationRequest struct {
	Teamid interface{} `json:"teamId"`
	Channelid interface{} `json:"channelId"`
}

// PutAccountAliasRequest represents the PutAccountAliasRequest schema from the OpenAPI specification
type PutAccountAliasRequest struct {
	Accountalias interface{} `json:"accountAlias"`
}

// DeleteSlackChannelConfigurationResult represents the DeleteSlackChannelConfigurationResult schema from the OpenAPI specification
type DeleteSlackChannelConfigurationResult struct {
}
