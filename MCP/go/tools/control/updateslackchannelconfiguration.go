package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-support-app/mcp-server/config"
	"github.com/aws-support-app/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateslackchannelconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/control/update-slack-channel-configuration", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UpdateSlackChannelConfigurationResult
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdateslackchannelconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_control_update-slack-channel-configuration",
		mcp.WithDescription("Updates the configuration for a Slack channel, such as case update notifications."),
		mcp.WithBoolean("notifyOnResolveCase", mcp.Description("Input parameter: Whether you want to get notified when a support case is resolved.")),
		mcp.WithString("teamId", mcp.Required(), mcp.Description("Input parameter: The team ID in Slack. This ID uniquely identifies a Slack workspace, such as <code>T012ABCDEFG</code>.")),
		mcp.WithString("channelId", mcp.Required(), mcp.Description("Input parameter: The channel ID in Slack. This ID identifies a channel within a Slack workspace.")),
		mcp.WithString("channelName", mcp.Description("Input parameter: The Slack channel name that you want to update.")),
		mcp.WithString("channelRoleArn", mcp.Description("Input parameter: The Amazon Resource Name (ARN) of an IAM role that you want to use to perform operations on Amazon Web Services. For more information, see <a href=\"https://docs.aws.amazon.com/awssupport/latest/user/support-app-permissions.html\">Managing access to the Amazon Web Services Support App</a> in the <i>Amazon Web Services Support User Guide</i>.")),
		mcp.WithBoolean("notifyOnAddCorrespondenceToCase", mcp.Description("Input parameter: Whether you want to get notified when a support case has a new correspondence.")),
		mcp.WithString("notifyOnCaseSeverity", mcp.Description("Input parameter: <p>The case severity for a support case that you want to receive notifications.</p> <p>If you specify <code>high</code> or <code>all</code>, at least one of the following parameters must be <code>true</code>:</p> <ul> <li> <p> <code>notifyOnAddCorrespondenceToCase</code> </p> </li> <li> <p> <code>notifyOnCreateOrReopenCase</code> </p> </li> <li> <p> <code>notifyOnResolveCase</code> </p> </li> </ul> <p>If you specify <code>none</code>, any of the following parameters that you specify in your request must be <code>false</code>:</p> <ul> <li> <p> <code>notifyOnAddCorrespondenceToCase</code> </p> </li> <li> <p> <code>notifyOnCreateOrReopenCase</code> </p> </li> <li> <p> <code>notifyOnResolveCase</code> </p> </li> </ul> <note> <p>If you don't specify these parameters in your request, the Amazon Web Services Support App uses the current values by default.</p> </note>")),
		mcp.WithBoolean("notifyOnCreateOrReopenCase", mcp.Description("Input parameter: Whether you want to get notified when a support case is created or reopened.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateslackchannelconfigurationHandler(cfg),
	}
}
