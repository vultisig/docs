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
	// PluginID defined in verifier database 'plugins'. You can find it in your local version while testing or request
	// from Vultisig team before going live
	PluginID = "<your-plugin-id>"
	// PluginName will be used as short description in policy
	PluginName = "<your-plugin-name>"
)

var supportedChains = []common.Chain{
	// If your app will trigger chain specific contracts you can choose exact one
	common.Ethereum,
	// Or you can expand supported chains with Metarules (https://github.com/vultisig/docs/blob/main/developer-docs/app-store/infrastructure-overview/metarules.md)
	//common.Arbitrum,
	//common.Avalanche,
	//common.BscChain,
	//common.Base,
	//common.Blast,
	//common.Optimism,
	//common.Polygon,
	//common.Bitcoin,
	//common.Solana,
	//common.XRP,
	//common.Zcash,
}

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

func (s *spec) GetRecipeSpecification() (*rtypes.RecipeSchema, error) {
	assetDef := rjsonschema.NewAsset()

	// Define your recipe properties there.
	// This will be called by verifier to show on UI
	cfg, err := plugin.RecipeConfiguration(map[string]any{
		"type":        "object",
		"definitions": rjsonschema.Definitions(),
		"properties": map[string]any{
			"asset": map[string]any{
				"$ref": "#/definitions/" + assetDef.Name(),
			},
			"fromAddress": map[string]any{
				"type": "string",
			},
			"toAddress": map[string]any{
				"type": "string",
			},
			"amount": map[string]any{
				"type": "string",
			},
		},
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

	// For better user experience you can prepare examples to make it easier for them to create automations
	cfgExample, err := plugin.RecipeConfiguration(map[string]any{
		"asset": map[string]any{
			"chain": "Ethereum",
			"token": "",
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
		SupportedResources:   s.buildSupportedResources(),
		Configuration:        cfg,
		ConfigurationExample: cfgExamples,
		Requirements: &rtypes.PluginRequirements{
			MinVultisigVersion: 1,
			SupportedChains:    getSupportedChainStrings(),
		},
	}, nil
}

func (s *spec) ValidatePluginPolicy(policy types.PluginPolicy) error {
	// Here you can validate the data that the user has filled in UI (amounts, addresses, parameters, etc.).
	spc, err := s.GetRecipeSpecification()
	if err != nil {
		return fmt.Errorf("failed to get recipe spec: %w", err)
	}
	return plugin.ValidatePluginPolicy(policy, spc)
}

func (s *spec) Suggest(cfg map[string]any) (*rtypes.PolicySuggest, error) {
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

	rule, err := s.createSendMetaRule(cfg, chainTyped)
	if err != nil {
		return nil, fmt.Errorf("failed to create send meta rule: %w", err)
	}

	return &rtypes.PolicySuggest{
		// Variables for configuring available tx per period (in seconds) count
		RateLimitWindow: conv.Ptr(uint32(60)),
		MaxTxsPerWindow: conv.Ptr(uint32(2)),
		Rules:           []*rtypes.Rule{rule},
	}, nil
}

func (s *spec) createSendMetaRule(cfg map[string]any, chainTyped common.Chain) (*rtypes.Rule, error) {
	chainLowercase := strings.ToLower(chainTyped.String())

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

	return &rtypes.Rule{
		Resource: chainLowercase + ".send",
		Effect:   rtypes.Effect_EFFECT_ALLOW,
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

func GetStr(cfg map[string]any, key string) string {
	var res string
	if val, ok := cfg[key]; ok && val != nil {
		res, _ = val.(string)
	}
	return res
}

func (*spec) buildSupportedResources() []*rtypes.ResourcePattern {
	var resources []*rtypes.ResourcePattern

	for _, chain := range supportedChains {
		chainNameLower := strings.ToLower(chain.String())

		resources = append(resources, &rtypes.ResourcePattern{
			ResourcePath: &rtypes.ResourcePath{
				ChainId:    chainNameLower,
				ProtocolId: "send",
				FunctionId: "Access to transaction signing",
				Full:       chainNameLower + ".send",
			},
			Target: rtypes.TargetType_TARGET_TYPE_UNSPECIFIED,
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
