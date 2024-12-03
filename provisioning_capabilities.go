/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// CapabilityType defines model for CapabilityType.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilitytype
type CapabilityType string

const (
	// CapabilityTypeAccessWifiInformation is a capability type for AccessWifiInformation.
	CapabilityTypeAccessWifiInformation CapabilityType = "ACCESS_WIFI_INFORMATION"
	// CapabilityTypeAppleIDAuth is a capability type for AppleIDAuth.
	CapabilityTypeAppleIDAuth CapabilityType = "APPLE_ID_AUTH"
	// CapabilityTypeApplePay is a capability type for ApplePay.
	CapabilityTypeApplePay CapabilityType = "APPLE_PAY"
	// CapabilityTypeAppGroups is a capability type for AppGroups.
	CapabilityTypeAppGroups CapabilityType = "APP_GROUPS"
	// CapabilityTypeAssociatedDomains is a capability type for AssociatedDomains.
	CapabilityTypeAssociatedDomains CapabilityType = "ASSOCIATED_DOMAINS"
	// CapabilityTypeAutoFillCredentialProvider is a capability type for AutoFillCredentialProvider.
	CapabilityTypeAutoFillCredentialProvider CapabilityType = "AUTOFILL_CREDENTIAL_PROVIDER"
	// CapabilityTypeClassKit is a capability type for ClassKit.
	CapabilityTypeClassKit CapabilityType = "CLASSKIT"
	// CapabilityTypeCoreMediaHLSLowLatency is a capability type for CoreMediaHLSLowLatency.
	CapabilityTypeCoreMediaHLSLowLatency CapabilityType = "COREMEDIA_HLS_LOW_LATENCY"
	// CapabilityTypeDataProtection is a capability type for DataProtection.
	CapabilityTypeDataProtection CapabilityType = "DATA_PROTECTION"
	// CapabilityTypeGameCenter is a capability type for GameCenter.
	CapabilityTypeGameCenter CapabilityType = "GAME_CENTER"
	// CapabilityTypeHealthKit is a capability type for HealthKit.
	CapabilityTypeHealthKit CapabilityType = "HEALTHKIT"
	// CapabilityTypeHomeKit is a capability type for HomeKit.
	CapabilityTypeHomeKit CapabilityType = "HOMEKIT"
	// CapabilityTypeHotSpot is a capability type for HotSpot.
	CapabilityTypeHotSpot CapabilityType = "HOT_SPOT"
	// CapabilityTypeiCloud is a capability type for iCloud.
	CapabilityTypeiCloud CapabilityType = "ICLOUD"
	// CapabilityTypeInterAppAudio is a capability type for InterAppAudio.
	CapabilityTypeInterAppAudio CapabilityType = "INTER_APP_AUDIO"
	// CapabilityTypeInAppPurchase is a capability type for InAppPurchase.
	CapabilityTypeInAppPurchase CapabilityType = "IN_APP_PURCHASE"
	// CapabilityTypeMaps is a capability type for Maps.
	CapabilityTypeMaps CapabilityType = "MAPS"
	// CapabilityTypeMultipath is a capability type for Multipath.
	CapabilityTypeMultipath CapabilityType = "MULTIPATH"
	// CapabilityTypeNetworkCustomProtocol is a capability type for NetworkCustomProtocol.
	CapabilityTypeNetworkCustomProtocol CapabilityType = "NETWORK_CUSTOM_PROTOCOL"
	// CapabilityTypeNetworkExtensions is a capability type for NetworkExtensions.
	CapabilityTypeNetworkExtensions CapabilityType = "NETWORK_EXTENSIONS"
	// CapabilityTypeNFCTagReading is a capability type for NFCTagReading.
	CapabilityTypeNFCTagReading CapabilityType = "NFC_TAG_READING"
	// CapabilityTypePersonalVPN is a capability type for PersonalVPN.
	CapabilityTypePersonalVPN CapabilityType = "PERSONAL_VPN"
	// CapabilityTypePushNotifications is a capability type for PushNotifications.
	CapabilityTypePushNotifications CapabilityType = "PUSH_NOTIFICATIONS"
	// CapabilityTypeSiriKit is a capability type for SiriKit.
	CapabilityTypeSiriKit CapabilityType = "SIRIKIT"
	// CapabilityTypeSystemExtensionInstall is a capability type for SystemExtensionInstall.
	CapabilityTypeSystemExtensionInstall CapabilityType = "SYSTEM_EXTENSION_INSTALL"
	// CapabilityTypeUserManagement is a capability type for UserManagement.
	CapabilityTypeUserManagement CapabilityType = "USER_MANAGEMENT"
	// CapabilityTypeWallet is a capability type for Wallet.
	CapabilityTypeWallet CapabilityType = "WALLET"
	// CapabilityTypeWirelessAccessoryConfiguration is a capability type for WirelessAccessoryConfiguration.
	CapabilityTypeWirelessAccessoryConfiguration CapabilityType = "WIRELESS_ACCESSORY_CONFIGURATION"
	// CapabilityTypeExtendedVirtualAddressing is a capability type for ExtendedVirtualAddressing.
	CapabilityTypeExtendedVirtualAddressing CapabilityType = "EXTENDED_VIRTUAL_ADDRESSING"

	CapabilityTypeIncreasedMemoryLimit CapabilityType = "INCREASED_MEMORY_LIMIT"

	CapabilityTypeUserNotificationsCommunication CapabilityType = "USERNOTIFICATIONS_COMMUNICATION"

	CapabilityTypeWeatherKit CapabilityType = "WEATHERKIT"

	CapabilityTypeHealthKitRecalibrateEstimates CapabilityType = "HEALTHKIT_RECALIBRATE_ESTIMATES"
)

var entitlementToCapability = map[string]CapabilityType{
	"com.apple.developer.networking.wifi-info":                                 CapabilityTypeAccessWifiInformation,
	"com.apple.developer.appleid-auth":                                         CapabilityTypeAppleIDAuth, // [新增]
	"com.apple.developer.in-app-payments":                                      CapabilityTypeApplePay,
	"com.apple.security.application-groups":                                    CapabilityTypeAppGroups,
	"com.apple.developer.associated-domains":                                   CapabilityTypeAssociatedDomains,
	"com.apple.developer.authentication-services.autofill-credential-provider": CapabilityTypeAutoFillCredentialProvider,
	"com.apple.developer.ClassKit-environment":                                 CapabilityTypeClassKit,
	"com.apple.developer.coremedia.hls.low-latency":                            CapabilityTypeCoreMediaHLSLowLatency,
	"com.apple.developer.default-data-protection":                              CapabilityTypeDataProtection, // [新增]
	"com.apple.developer.game-center":                                          CapabilityTypeGameCenter,
	"com.apple.developer.healthkit":                                            CapabilityTypeHealthKit,
	"com.apple.developer.healthkit.recalibrate-estimates":                      CapabilityTypeHealthKitRecalibrateEstimates, // [新增]
	"com.apple.developer.homekit":                                              CapabilityTypeHomeKit,
	"com.apple.developer.hotspot":                                              CapabilityTypeHotSpot, // [新增]
	"com.apple.developer.icloud-services":                                      CapabilityTypeiCloud,
	"inter-app-audio":                                                          CapabilityTypeInterAppAudio,
	"com.apple.developer.in-app-purchase":                                      CapabilityTypeInAppPurchase, // [新增]
	"com.apple.developer.maps":                                                 CapabilityTypeMaps,          // [新增]
	"com.apple.developer.networking.multipath":                                 CapabilityTypeMultipath,
	"com.apple.developer.networking.custom-protocol":                           CapabilityTypeNetworkCustomProtocol, // [新增]
	"com.apple.developer.networking.networkextension":                          CapabilityTypeNetworkExtensions,
	"com.apple.developer.nfc.readersession.formats":                            CapabilityTypeNFCTagReading,
	"com.apple.developer.networking.vpn.api":                                   CapabilityTypePersonalVPN,
	"aps-environment":                                                          CapabilityTypePushNotifications,
	"com.apple.developer.siri":                                                 CapabilityTypeSiriKit,
	"com.apple.developer.system-extension.install":                             CapabilityTypeSystemExtensionInstall, // [新增]
	"com.apple.developer.user-management":                                      CapabilityTypeUserManagement,         // [新增]
	"com.apple.developer.pass-type-identifiers":                                CapabilityTypeWallet,
	"com.apple.external-accessory.wireless-configuration":                      CapabilityTypeWirelessAccessoryConfiguration,
	"com.apple.developer.kernel.extended-virtual-addressing":                   CapabilityTypeExtendedVirtualAddressing,
	"com.apple.developer.kernel.increased-memory-limit":                        CapabilityTypeIncreasedMemoryLimit,
	"com.apple.developer.usernotifications.communication":                      CapabilityTypeUserNotificationsCommunication, // [新增]
	"com.apple.developer.weatherkit":                                           CapabilityTypeWeatherKit,
}

var capabilityToChineseMap = map[CapabilityType]string{
	CapabilityTypeAccessWifiInformation:          "Wi-Fi 信息访问",
	CapabilityTypeAppleIDAuth:                    "Apple ID 认证", // [新增]
	CapabilityTypeApplePay:                       "Apple Pay 支付",
	CapabilityTypeAppGroups:                      "应用组共享",
	CapabilityTypeAssociatedDomains:              "关联域名",
	CapabilityTypeAutoFillCredentialProvider:     "自动填充凭据",
	CapabilityTypeClassKit:                       "ClassKit 支持",
	CapabilityTypeCoreMediaHLSLowLatency:         "低延迟 HLS 流媒体",
	CapabilityTypeDataProtection:                 "数据保护", // [新增]
	CapabilityTypeGameCenter:                     "Game Center 支持",
	CapabilityTypeHealthKit:                      "健康数据",
	CapabilityTypeHealthKitRecalibrateEstimates:  "健康数据校准", // [新增]
	CapabilityTypeHomeKit:                        "家庭自动化",
	CapabilityTypeHotSpot:                        "个人热点：不允许",
	CapabilityTypeiCloud:                         "iCloud 支持",
	CapabilityTypeInterAppAudio:                  "应用间音频",
	CapabilityTypeInAppPurchase:                  "应用内购买：不允许",
	CapabilityTypeMaps:                           "地图服务", // [新增]
	CapabilityTypeMultipath:                      "多路径传输",
	CapabilityTypeNetworkCustomProtocol:          "自定义网络协议", // [新增]
	CapabilityTypeNetworkExtensions:              "网络扩展功能",
	CapabilityTypeNFCTagReading:                  "NFC 标签读取",
	CapabilityTypePersonalVPN:                    "个人 VPN",
	CapabilityTypePushNotifications:              "推送通知",
	CapabilityTypeSiriKit:                        "Siri 支持",
	CapabilityTypeSystemExtensionInstall:         "系统扩展安装", // [新增]
	CapabilityTypeUserManagement:                 "用户管理",   // [新增]
	CapabilityTypeWallet:                         "Wallet 支持",
	CapabilityTypeWirelessAccessoryConfiguration: "无线配件配置",
	CapabilityTypeExtendedVirtualAddressing:      "扩展虚拟地址支持",
	CapabilityTypeIncreasedMemoryLimit:           "增加内存限制",
	CapabilityTypeUserNotificationsCommunication: "用户通知通信", // [新增]
	CapabilityTypeWeatherKit:                     "天气服务",
}

func GetCapabilityForEntitlement(entitlement string) (CapabilityType, bool) {
	capability, exists := entitlementToCapability[entitlement]
	return capability, exists
}

func GetCapabilityChinese(entitlement string) string {
	capability, exists := entitlementToCapability[entitlement]
	if !exists {
		// 如果 entitlement 不存在，直接返回 entitlement
		return entitlement
	}
	chinese, exists := capabilityToChineseMap[capability]
	if !exists {
		// 如果没有定义中文描述，返回 capability
		return string(capability)
	}
	return chinese
}

// BundleIDCapability defines model for BundleIdCapability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapability
type BundleIDCapability struct {
	Attributes *BundleIDCapabilityAttributes `json:"attributes,omitempty"`
	ID         string                        `json:"id"`
	Links      ResourceLinks                 `json:"links"`
	Type       string                        `json:"type"`
}

// BundleIDCapabilityAttributes defines model for BundleIdCapability.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapability/attributes
type BundleIDCapabilityAttributes struct {
	CapabilityType *CapabilityType     `json:"capabilityType,omitempty"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
}

// bundleIDCapabilityCreateRequest defines model for BundleIdCapabilityCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data
type bundleIDCapabilityCreateRequest struct {
	Attributes    bundleIDCapabilityCreateRequestAttributes    `json:"attributes"`
	Relationships bundleIDCapabilityCreateRequestRelationships `json:"relationships"`
	Type          string                                       `json:"type"`
}

// bundleIDCapabilityCreateRequestAttributes are attributes for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/attributes
type bundleIDCapabilityCreateRequestAttributes struct {
	CapabilityType CapabilityType      `json:"capabilityType"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
}

// bundleIDCapabilityCreateRequestRelationships are relationships for BundleIDCapabilityCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitycreaterequest/data/relationships
type bundleIDCapabilityCreateRequestRelationships struct {
	BundleID relationshipDeclaration `json:"bundleId"`
}

// BundleIDCapabilityUpdateRequest defines model for BundleIdCapabilityUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest/data
type bundleIDCapabilityUpdateRequest struct {
	Attributes *bundleIDCapabilityUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                     `json:"id"`
	Type       string                                     `json:"type"`
}

// BundleIDCapabilityUpdateRequestAttributes are attributes for BundleIDCapabilityUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityupdaterequest/data/attributes
type bundleIDCapabilityUpdateRequestAttributes struct {
	CapabilityType *CapabilityType     `json:"capabilityType,omitempty"`
	Settings       []CapabilitySetting `json:"settings,omitempty"`
}

// BundleIDCapabilityResponse defines model for BundleIdCapabilityResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilityresponse
type BundleIDCapabilityResponse struct {
	Data  BundleIDCapability `json:"data"`
	Links DocumentLinks      `json:"links"`
}

// BundleIDCapabilitiesResponse defines model for BundleIdCapabilitiesResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/bundleidcapabilitiesresponse
type BundleIDCapabilitiesResponse struct {
	Data  []BundleIDCapability `json:"data"`
	Links PagedDocumentLinks   `json:"links"`
	Meta  *PagingInformation   `json:"meta,omitempty"`
}

// CapabilityOption defines model for CapabilityOption.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilityoption
type CapabilityOption struct {
	Description      *string `json:"description,omitempty"`
	Enabled          *bool   `json:"enabled,omitempty"`
	EnabledByDefault *bool   `json:"enabledByDefault,omitempty"`
	Key              *string `json:"key,omitempty"`
	Name             *string `json:"name,omitempty"`
	SupportsWildcard *bool   `json:"supportsWildcard,omitempty"`
}

// CapabilitySetting defines model for CapabilitySetting.
//
// https://developer.apple.com/documentation/appstoreconnectapi/capabilitysetting
type CapabilitySetting struct {
	AllowedInstances *string            `json:"allowedInstances,omitempty"`
	Description      *string            `json:"description,omitempty"`
	EnabledByDefault *bool              `json:"enabledByDefault,omitempty"`
	Key              *string            `json:"key,omitempty"`
	MinInstances     *int               `json:"minInstances,omitempty"`
	Name             *string            `json:"name,omitempty"`
	Options          []CapabilityOption `json:"options,omitempty"`
	Visible          *bool              `json:"visible,omitempty"`
}

// EnableCapability enables a capability for a bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/enable_a_capability
func (s *ProvisioningService) EnableCapability(ctx context.Context, capabilityType CapabilityType, capabilitySettings []CapabilitySetting, bundleIDRelationship string) (*BundleIDCapabilityResponse, *Response, error) {
	req := bundleIDCapabilityCreateRequest{
		Attributes: bundleIDCapabilityCreateRequestAttributes{
			CapabilityType: capabilityType,
			Settings:       capabilitySettings,
		},
		Relationships: bundleIDCapabilityCreateRequestRelationships{
			BundleID: relationshipDeclaration{
				Data: RelationshipData{
					ID:   bundleIDRelationship,
					Type: "bundleIds",
				},
			},
		},
		Type: "bundleIdCapabilities",
	}
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.post(ctx, "bundleIdCapabilities", newRequestBody(req), res)

	return res, resp, err
}

// DisableCapability disables a capability for a bundle ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/disable_a_capability
func (s *ProvisioningService) DisableCapability(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("bundleIdCapabilities/%s", id)

	return s.client.delete(ctx, url, nil)
}

// UpdateCapability updates the configuration of a specific capability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_capability_configuration
func (s *ProvisioningService) UpdateCapability(ctx context.Context, id string, capabilityType *CapabilityType, settings []CapabilitySetting) (*BundleIDCapabilityResponse, *Response, error) {
	req := bundleIDCapabilityUpdateRequest{
		ID:   id,
		Type: "bundleIdCapabilities",
	}

	if capabilityType != nil || settings != nil {
		req.Attributes = &bundleIDCapabilityUpdateRequestAttributes{
			CapabilityType: capabilityType,
			Settings:       settings,
		}
	}

	url := fmt.Sprintf("bundleIdCapabilities/%s", id)
	res := new(BundleIDCapabilityResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
