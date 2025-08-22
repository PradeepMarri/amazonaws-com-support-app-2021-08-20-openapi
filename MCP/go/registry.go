package main

import (
	"github.com/aws-support-app/mcp-server/config"
	"github.com/aws-support-app/mcp-server/models"
	tools_control "github.com/aws-support-app/mcp-server/tools/control"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_control.CreateListslackworkspaceconfigurationsTool(cfg),
		tools_control.CreateDeleteslackworkspaceconfigurationTool(cfg),
		tools_control.CreateGetaccountaliasTool(cfg),
		tools_control.CreateListslackchannelconfigurationsTool(cfg),
		tools_control.CreateRegisterslackworkspacefororganizationTool(cfg),
		tools_control.CreateCreateslackchannelconfigurationTool(cfg),
		tools_control.CreatePutaccountaliasTool(cfg),
		tools_control.CreateUpdateslackchannelconfigurationTool(cfg),
		tools_control.CreateDeleteaccountaliasTool(cfg),
		tools_control.CreateDeleteslackchannelconfigurationTool(cfg),
	}
}
