// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/metadata.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "github.com/gogo/protobuf/types"
)

// MarshalProtoJSON marshals the LocationSource to JSON.
func (x LocationSource) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), LocationSource_name)
}

// LocationSource_customvalue contains custom string values that extend LocationSource_value.
var LocationSource_customvalue = map[string]int32{
	"UNKNOWN":               0,
	"GPS":                   1,
	"REGISTRY":              3,
	"IP_GEOLOCATION":        4,
	"WIFI_RSSI_GEOLOCATION": 5,
	"BT_RSSI_GEOLOCATION":   6,
	"LORA_RSSI_GEOLOCATION": 7,
	"LORA_TDOA_GEOLOCATION": 8,
	"COMBINED_GEOLOCATION":  9,
}

// UnmarshalProtoJSON unmarshals the LocationSource from JSON.
func (x *LocationSource) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(LocationSource_value, LocationSource_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read LocationSource enum: %v", err)
		return
	}
	*x = LocationSource(v)
}

// MarshalProtoJSON marshals the RxMetadata message to JSON.
func (x *RxMetadata) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayIds != nil || s.HasField("gateway_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_ids")
		// NOTE: GatewayIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.GatewayIds)
	}
	if x.PacketBroker != nil || s.HasField("packet_broker") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("packet_broker")
		// NOTE: PacketBrokerMetadata does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.PacketBroker)
	}
	if x.AntennaIndex != 0 || s.HasField("antenna_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("antenna_index")
		s.WriteUint32(x.AntennaIndex)
	}
	if x.Time != nil || s.HasField("time") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("time")
		if x.Time == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.Time)
		}
	}
	if x.Timestamp != 0 || s.HasField("timestamp") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("timestamp")
		s.WriteUint32(x.Timestamp)
	}
	if x.FineTimestamp != 0 || s.HasField("fine_timestamp") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("fine_timestamp")
		s.WriteUint64(x.FineTimestamp)
	}
	if len(x.EncryptedFineTimestamp) > 0 || s.HasField("encrypted_fine_timestamp") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("encrypted_fine_timestamp")
		s.WriteBytes(x.EncryptedFineTimestamp)
	}
	if x.EncryptedFineTimestampKeyId != "" || s.HasField("encrypted_fine_timestamp_key_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("encrypted_fine_timestamp_key_id")
		s.WriteString(x.EncryptedFineTimestampKeyId)
	}
	if x.Rssi != 0 || s.HasField("rssi") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rssi")
		s.WriteFloat32(x.Rssi)
	}
	if x.SignalRssi != nil || s.HasField("signal_rssi") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("signal_rssi")
		if x.SignalRssi == nil {
			s.WriteNil()
		} else {
			s.WriteFloat32(x.SignalRssi.Value)
		}
	}
	if x.ChannelRssi != 0 || s.HasField("channel_rssi") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("channel_rssi")
		s.WriteFloat32(x.ChannelRssi)
	}
	if x.RssiStandardDeviation != 0 || s.HasField("rssi_standard_deviation") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rssi_standard_deviation")
		s.WriteFloat32(x.RssiStandardDeviation)
	}
	if x.Snr != 0 || s.HasField("snr") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("snr")
		s.WriteFloat32(x.Snr)
	}
	if x.FrequencyOffset != 0 || s.HasField("frequency_offset") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_offset")
		s.WriteInt64(x.FrequencyOffset)
	}
	if x.Location != nil || s.HasField("location") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("location")
		x.Location.MarshalProtoJSON(s.WithField("location"))
	}
	if x.DownlinkPathConstraint != 0 || s.HasField("downlink_path_constraint") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("downlink_path_constraint")
		x.DownlinkPathConstraint.MarshalProtoJSON(s)
	}
	if len(x.UplinkToken) > 0 || s.HasField("uplink_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("uplink_token")
		s.WriteBytes(x.UplinkToken)
	}
	if x.ChannelIndex != 0 || s.HasField("channel_index") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("channel_index")
		s.WriteUint32(x.ChannelIndex)
	}
	if x.HoppingWidth != 0 || s.HasField("hopping_width") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("hopping_width")
		s.WriteUint32(x.HoppingWidth)
	}
	if x.FrequencyDrift != 0 || s.HasField("frequency_drift") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_drift")
		s.WriteInt32(x.FrequencyDrift)
	}
	if x.Advanced != nil || s.HasField("advanced") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("advanced")
		if x.Advanced == nil {
			s.WriteNil()
		} else {
			gogo.MarshalStruct(s, x.Advanced)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the RxMetadata message from JSON.
func (x *RxMetadata) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_ids", "gatewayIds":
			s.AddField("gateway_ids")
			// NOTE: GatewayIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v GatewayIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.GatewayIds = &v
		case "packet_broker", "packetBroker":
			s.AddField("packet_broker")
			// NOTE: PacketBrokerMetadata does not seem to implement UnmarshalProtoJSON.
			var v PacketBrokerMetadata
			gogo.UnmarshalMessage(s, &v)
			x.PacketBroker = &v
		case "antenna_index", "antennaIndex":
			s.AddField("antenna_index")
			x.AntennaIndex = s.ReadUint32()
		case "time":
			s.AddField("time")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.Time = v
		case "timestamp":
			s.AddField("timestamp")
			x.Timestamp = s.ReadUint32()
		case "fine_timestamp", "fineTimestamp":
			s.AddField("fine_timestamp")
			x.FineTimestamp = s.ReadUint64()
		case "encrypted_fine_timestamp", "encryptedFineTimestamp":
			s.AddField("encrypted_fine_timestamp")
			x.EncryptedFineTimestamp = s.ReadBytes()
		case "encrypted_fine_timestamp_key_id", "encryptedFineTimestampKeyId":
			s.AddField("encrypted_fine_timestamp_key_id")
			x.EncryptedFineTimestampKeyId = s.ReadString()
		case "rssi":
			s.AddField("rssi")
			x.Rssi = s.ReadFloat32()
		case "signal_rssi", "signalRssi":
			s.AddField("signal_rssi")
			if !s.ReadNil() {
				v := s.ReadFloat32()
				if s.Err() != nil {
					return
				}
				x.SignalRssi = &types.FloatValue{Value: v}
			}
		case "channel_rssi", "channelRssi":
			s.AddField("channel_rssi")
			x.ChannelRssi = s.ReadFloat32()
		case "rssi_standard_deviation", "rssiStandardDeviation":
			s.AddField("rssi_standard_deviation")
			x.RssiStandardDeviation = s.ReadFloat32()
		case "snr":
			s.AddField("snr")
			x.Snr = s.ReadFloat32()
		case "frequency_offset", "frequencyOffset":
			s.AddField("frequency_offset")
			x.FrequencyOffset = s.ReadInt64()
		case "location":
			if !s.ReadNil() {
				x.Location = &Location{}
				x.Location.UnmarshalProtoJSON(s.WithField("location", true))
			}
		case "downlink_path_constraint", "downlinkPathConstraint":
			s.AddField("downlink_path_constraint")
			x.DownlinkPathConstraint.UnmarshalProtoJSON(s)
		case "uplink_token", "uplinkToken":
			s.AddField("uplink_token")
			x.UplinkToken = s.ReadBytes()
		case "channel_index", "channelIndex":
			s.AddField("channel_index")
			x.ChannelIndex = s.ReadUint32()
		case "hopping_width", "hoppingWidth":
			s.AddField("hopping_width")
			x.HoppingWidth = s.ReadUint32()
		case "frequency_drift", "frequencyDrift":
			s.AddField("frequency_drift")
			x.FrequencyDrift = s.ReadInt32()
		case "advanced":
			s.AddField("advanced")
			v := gogo.UnmarshalStruct(s)
			if s.Err() != nil {
				return
			}
			x.Advanced = v
		}
	})
}

// MarshalProtoJSON marshals the Location message to JSON.
func (x *Location) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Latitude != 0 || s.HasField("latitude") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("latitude")
		s.WriteFloat64(x.Latitude)
	}
	if x.Longitude != 0 || s.HasField("longitude") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("longitude")
		s.WriteFloat64(x.Longitude)
	}
	if x.Altitude != 0 || s.HasField("altitude") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("altitude")
		s.WriteInt32(x.Altitude)
	}
	if x.Accuracy != 0 || s.HasField("accuracy") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("accuracy")
		s.WriteInt32(x.Accuracy)
	}
	if x.Source != 0 || s.HasField("source") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("source")
		x.Source.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Location message from JSON.
func (x *Location) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "latitude":
			s.AddField("latitude")
			x.Latitude = s.ReadFloat64()
		case "longitude":
			s.AddField("longitude")
			x.Longitude = s.ReadFloat64()
		case "altitude":
			s.AddField("altitude")
			x.Altitude = s.ReadInt32()
		case "accuracy":
			s.AddField("accuracy")
			x.Accuracy = s.ReadInt32()
		case "source":
			s.AddField("source")
			x.Source.UnmarshalProtoJSON(s)
		}
	})
}
