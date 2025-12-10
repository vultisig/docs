package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/structpb"

	rjsonschema "github.com/vultisig/recipes/jsonschema"
	rtypes "github.com/vultisig/recipes/types"
	"github.com/vultisig/verifier/plugin"
	"github.com/vultisig/verifier/plugin/tx_indexer/pkg/conv"
	"github.com/vultisig/verifier/types"
	"github.com/vultisig/vultisig-go/common"
)

const (
	// PluginID is the unique identifier for this plugin within the Vultisig Verifier database.
	// DEVELOPMENT: You can find this in your local database 'plugins' table.
	// PRODUCTION: Request this ID from the Vultisig team before deploying.
	PluginID = "<your-plugin-id>"

	// PluginName is the human-readable name of the plugin.
	// This name appears in the policy details and UI descriptions.
	PluginName = "<your-plugin-name>"
)

// supportedChains defines the list of blockchains this plugin is authorized to interact with.
// This limits the scope of the plugin to prevent unintended operations on other chains.
var supportedChains = []common.Chain{
	// Specific Chain: Use this if your app triggers contracts specific to Ethereum.
	common.Ethereum,

	// Expand Support: You can uncomment other chains here if your logic is generic enough
	// to handle them (e.g., EVM-compatible chains).
	// See Metarules docs: https://github.com/vultisig/docs/blob/main/developer-docs/app-store/infrastructure-overview/metarules.md
	//common.Arbitrum,
	//common.Avalanche,
	//common.BscChain,
	//common.Base,
	//common.Bitcoin,
	//common.Solana,
}

// getSupportedChainStrings converts the list of chain objects into their string representation
// required by the plugin requirements specification.
func getSupportedChainStrings() []string {
	var cc []string
	for _, c := range supportedChains {
		cc = append(cc, c.String())
	}
	return cc
}

type spec struct {
}

func newSpec() *spec {
	return &spec{}
}

// GetRecipeSpecification defines the "Recipe" - the configuration interface for the Vultisig UI.
// It returns a JSON Schema that dictates what fields users must fill out to create an automation.
func (s *spec) GetRecipeSpecification() (*rtypes.RecipeSchema, error) {
	// rjsonschema.NewAsset() creates a standard schema definition for selecting assets (Chain + Token).
	assetDef := rjsonschema.NewAsset()

	// Define the properties users will see in the UI.
	// Here we define a simple transfer form: Asset, Sender, Receiver, Amount.
	cfg, err := plugin.RecipeConfiguration(map[string]any{
		"type":        "object",
		"definitions": rjsonschema.Definitions(), // Standard definitions helper
		"properties": map[string]any{
			"asset": map[string]any{
				"$ref": "#/definitions/" + assetDef.Name(), // Reference the standard Asset picker
			},
			"fromAddress": map[string]any{
				"type": "string", // Input field for sender address
			},
			"toAddress": map[string]any{
				"type": "string", // Input field for receiver address
			},
			"amount": map[string]any{
				"type": "string", // Input field for amount (handled as string to avoid precision loss)
			},
		},
		// List fields that are mandatory for the policy to be valid.
		"required": []any{
			"asset",
			"fromAddress",
			"toAddress",
			"amount",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build pb recipe config: %w", err)
	}

	// Provide a default example configuration.
	// This helps the UI pre-fill data or show users what valid input looks like.
	cfgExample, err := plugin.RecipeConfiguration(map[string]any{
		"asset": map[string]any{
			"chain": "Ethereum",
			"token": "", // Empty token usually implies the native coin (ETH)
		},
		"amount": "10000000",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build pb recipe config example: %w", err)
	}
	cfgExamples := []*structpb.Struct{cfgExample}

	return &rtypes.RecipeSchema{
		Version:              1,
		PluginId:             PluginID,
		PluginName:           PluginName,
		PluginVersion:        1,
		SupportedResources:   s.buildSupportedResources(), // Defines capabilities (what logic allows)
		Configuration:        cfg,                         // Defines UI (what user sees)
		ConfigurationExample: cfgExamples,
		Requirements: &rtypes.PluginRequirements{
			MinVultisigVersion: 1,
			SupportedChains:    getSupportedChainStrings(),
		},
	}, nil
}

// ValidatePluginPolicy validates user input against the schema defined in GetRecipeSpecification.
// This is a hook called by the Verifier before saving a policy to ensure data integrity.
func (s *spec) ValidatePluginPolicy(policy types.PluginPolicy) error {
	spc, err := s.GetRecipeSpecification()
	if err != nil {
		return fmt.Errorf("failed to get recipe spec: %w", err)
	}
	return plugin.ValidatePluginPolicy(policy, spc)
}

// Suggest translates the user's high-level configuration (UI inputs) into specific low-level execution rules.
// This is where you define exactly what the plugin is allowed to sign.
func (s *spec) Suggest(cfg map[string]any) (*rtypes.PolicySuggest, error) {
	// 1. Extract and validate raw inputs from the config map
	assetMap, ok := cfg["asset"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("'asset' must be an object")
	}

	chainStr, ok := assetMap["chain"].(string)
	if !ok {
		return nil, fmt.Errorf("'asset.chain' could not be empty")
	}

	chainTyped, err := common.FromString(chainStr)
	if err != nil {
		return nil, fmt.Errorf("unsupported chain: %s", chainStr)
	}

	// 2. Create the specific Rule for this policy (e.g., "Allow Ethereum Send")
	rule, err := s.createSendMetaRule(cfg, chainTyped)
	if err != nil {
		return nil, fmt.Errorf("failed to create send meta rule: %w", err)
	}

	// 3. Return the complete Policy Suggestion
	// This includes Rate Limits (e.g., 2 transactions per 60 seconds) to prevent spam/draining.
	return &rtypes.PolicySuggest{
		RateLimitWindow: conv.Ptr(uint32(60)), // Window size in seconds
		MaxTxsPerWindow: conv.Ptr(uint32(2)),  // Max transactions allowed in that window
		Rules:           []*rtypes.Rule{rule}, // The actual permission rules
	}, nil
}

// createSendMetaRule constructs a specific Rule object.
// It maps the user's input values to "Fixed" constraints.
// Meaning: "This policy ONLY allows signing IF amount == X, to == Y, from == Z".
func (s *spec) createSendMetaRule(cfg map[string]any, chainTyped common.Chain) (*rtypes.Rule, error) {
	chainLowercase := strings.ToLower(chainTyped.String())

	// Extract values to lock them into the rule
	assetMap, ok := cfg["asset"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("'asset' must be an object")
	}

	fromAddressStr, ok := cfg["fromAddress"].(string)
	if !ok || fromAddressStr == "" {
		return nil, fmt.Errorf("'fromAddress' could not be empty")
	}

	toAddressStr, ok := cfg["toAddress"].(string)
	if !ok || toAddressStr == "" {
		return nil, fmt.Errorf("'toAddress' could not be empty")
	}

	amountStr := GetStr(cfg, "amount")
	if amountStr == "" {
		return nil, fmt.Errorf("'amount' could not be empty")
	}

	tokenStr := GetStr(assetMap, "token")

	// Construct the Rule
	return &rtypes.Rule{
		// Resource identifies the action: e.g., "ethereum.send"
		Resource: chainLowercase + ".send",
		Effect:   rtypes.Effect_EFFECT_ALLOW,
		// ParameterConstraints enforce security.
		// CONSTRAINT_TYPE_FIXED means the value in the signing request MUST match the value in the policy.
		ParameterConstraints: []*rtypes.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &rtypes.Constraint{
					Type: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &rtypes.Constraint_FixedValue{
						FixedValue: tokenStr,
					},
					Required: false,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &rtypes.Constraint{
					Type: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &rtypes.Constraint_FixedValue{
						FixedValue: fromAddressStr,
					},
					Required: true,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &rtypes.Constraint{
					Type: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &rtypes.Constraint_FixedValue{
						FixedValue: amountStr,
					},
					Required: true,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &rtypes.Constraint{
					Type: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &rtypes.Constraint_FixedValue{
						FixedValue: toAddressStr,
					},
					Required: true,
				},
			},
			// Memo is optional and set to ANY, allowing dynamic memos.
			{
				ParameterName: "memo",
				Constraint: &rtypes.Constraint{
					Type:     rtypes.ConstraintType_CONSTRAINT_TYPE_ANY,
					Required: false,
				},
			},
		},
		Target: &rtypes.Target{
			TargetType: rtypes.TargetType_TARGET_TYPE_UNSPECIFIED,
		},
	}, nil
}

// GetStr safely extracts a string value from a map[string]any
func GetStr(cfg map[string]any, key string) string {
	var res string
	if val, ok := cfg[key]; ok && val != nil {
		res, _ = val.(string)
	}
	return res
}

// buildSupportedResources generates a manifest of ALL actions this plugin supports.
// This tells the Verifier: "I know how to handle 'send' operations on these chains."
func (*spec) buildSupportedResources() []*rtypes.ResourcePattern {
	var resources []*rtypes.ResourcePattern

	for _, chain := range supportedChains {
		chainNameLower := strings.ToLower(chain.String())

		resources = append(resources, &rtypes.ResourcePattern{
			ResourcePath: &rtypes.ResourcePath{
				ChainId:    chainNameLower,
				ProtocolId: "send", // Standard identifier for native transfers
				FunctionId: "Access to transaction signing",
				Full:       chainNameLower + ".send", // e.g., "ethereum.send"
			},
			Target: rtypes.TargetType_TARGET_TYPE_UNSPECIFIED,
			// Capabilities define which parameters MUST be present for this resource to be valid.
			ParameterCapabilities: []*rtypes.ParameterConstraintCapability{
				{
					ParameterName:  "asset",
					SupportedTypes: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Required:       false,
				},
				{
					ParameterName:  "from_address",
					SupportedTypes: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Required:       true,
				},
				{
					ParameterName:  "amount",
					SupportedTypes: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Required:       true,
				},
				{
					ParameterName:  "to_address",
					SupportedTypes: rtypes.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Required:       true,
				},
			},
			Required: true,
		})
	}

	return resources
}
