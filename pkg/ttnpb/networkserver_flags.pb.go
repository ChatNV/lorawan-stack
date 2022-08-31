// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v1.0.3
// - protoc              v3.9.1
// source: lorawan-stack/api/networkserver.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pflag "github.com/spf13/pflag"
)

// AddSetFlagsForGetDefaultMACSettingsRequest adds flags to select fields in GetDefaultMACSettingsRequest.
func AddSetFlagsForGetDefaultMACSettingsRequest(flags *pflag.FlagSet, prefix string, hidden bool) {
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("frequency-plan-id", prefix), "", flagsplugin.WithHidden(hidden)))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("lorawan-phy-version", prefix), flagsplugin.EnumValueDesc(PHYVersion_value, PHYVersion_customvalue), flagsplugin.WithHidden(hidden)))
}

// SetFromFlags sets the GetDefaultMACSettingsRequest message from flags.
func (m *GetDefaultMACSettingsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("frequency_plan_id", prefix)); err != nil {
		return nil, err
	} else if changed {
		m.FrequencyPlanId = val
		paths = append(paths, flagsplugin.Prefix("frequency_plan_id", prefix))
	}
	if val, changed, err := flagsplugin.GetString(flags, flagsplugin.Prefix("lorawan_phy_version", prefix)); err != nil {
		return nil, err
	} else if changed {
		enumValue, err := flagsplugin.SetEnumString(val, PHYVersion_value, PHYVersion_customvalue)
		if err != nil {
			return nil, err
		}
		m.LorawanPhyVersion = PHYVersion(enumValue)
		paths = append(paths, flagsplugin.Prefix("lorawan_phy_version", prefix))
	}
	return paths, nil
}
