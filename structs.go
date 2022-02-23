package main

import "net"

type Counters struct {
	InOctets         int64 `xml:"in-octets" json:"in_octets"`
	InPkts           int64 `xml:"in-pkts" json:"in_pkts"`
	InDiscards       int64 `xml:"in-discards" json:"in_discards"`
	InErrors         int64 `xml:"in-errors" json:"in_errors"`
	InNoRuleDiscards int64 `xml:"in-no-rule-discards" json:"in_no_rule_discards"`
	OutOctets        int64 `xml:"out-octets" json:"out_octets"`
	OutPkts          int64 `xml:"out-pkts" json:"out_pkts"`
	OutDiscards      int64 `xml:"out-discards" json:"out_discards"`
	OutErrors        int64 `xml:"out-errors" json:"out_errors"`
}

type RadioProfile struct {
	Polarity     string `xml:"polarity" json:"polarity"`
	Frequency    string `xml:"frequency" json:"frequency"`
	TxGolayIndex string `xml:"tx-golay-index" json:"tx_golay_index"`
	RxGolayIndex string `xml:"rx-golay-index" json:"rx_golay_index"`
}

type Data struct {
	DatabaseVersion struct {
		DatabaseVersionNumber string `xml:"database-version-number" json:"database_version_number"`
	} `xml:"database-version" json:"database_version"`

	// Gui struct {
	// 	Icon struct {
	// 	} `xml:"icon"`
	// } `xml:"gui"`

	Interfaces struct {
		Host struct {
			State struct {
				MacAddress string   `xml:"mac-address" json:"mac_address"`
				Counters   Counters `xml:"counters" json:"counters"`
			} `xml:"state" json:"state"`
		} `xml:"host" json:"host"`
		Ports []struct {
			AdminStatus          string `xml:"admin-status" json:"admin_status"`
			CopperSpecificConfig struct {
				AutoNegotiate string `xml:"auto-negotiate" json:"auto_negotiate"`
				PseOut        string `xml:"pse-out" json:"pse_out"`
			} `xml:"copper-specific-config" json:"copper_specific_config"`
			ConnectorType string `xml:"connector-type" json:"connector_type"`
			PortSpeed     string `xml:"port-speed" json:"port_speed"`
			State         struct {
				OperStatus       string   `xml:"oper-status" json:"oper_status"`
				Counters         Counters `xml:"counters" json:"counters"`
				ActualDuplexMode string   `xml:"actual-duplex-mode" json:"actual_duplex_mode"`
				ActualPortSpeed  string   `xml:"actual-port-speed" json:"actual_port_speed"`
			} `xml:"state" json:"state"`
		} `xml:"ports" json:"ports"`
		RfInterface struct {
			Name  string `xml:"name" json:"name"`
			State struct {
				OperStatus string   `xml:"oper-status" json:"oper_status"`
				Counters   Counters `xml:"counters" json:"counters"`
			} `xml:"state" json:"state"`
		} `xml:"rf-interface" json:"rf_interface"`
	} `xml:"interfaces" json:"interfaces"`

	Inventory struct {
		Component []struct {
			Name         string `xml:"name" json:"name"`
			Parent       string `xml:"parent" json:"parent"`
			ParentRelPos string `xml:"parent-rel-pos" json:"parent_rel_pos"`
			Child        []struct {
				Name string `xml:"name" json:"name"`
			} `xml:"child" json:"child"`
			HardwareRev string `xml:"hardware-rev" json:"hardware_rev"`
			SerialNum   string `xml:"serial-num" json:"serial_num"`
			IsFru       string `xml:"is-fru" json:"is_fru"`
			FirmwareRev string `xml:"firmware-rev" json:"firmware_rev"`
			Description string `xml:"description" json:"description"`
			SoftwareRev string `xml:"software-rev" json:"software_rev"`
			MfgName     string `xml:"mfg-name" json:"mfg_name"`
			ModelName   string `xml:"model-name" json:"model_name"`
		} `xml:"component" json:"component"`
	} `xml:"inventory" json:"inventory"`

	IP struct {
		IPv4 struct {
			Address []struct {
				IP           net.IP `xml:"ip" json:"ip"`
				PrefixLength uint8  `xml:"prefix-length" json:"prefix_length"`
				CVlan        uint16 `xml:"c-vlan" json:"c_vlan"`
			} `xml:"address" json:"address"`
			DefaultGateway string `xml:"default-gateway" json:"default_gateway"`
		} `xml:"ipv4" json:"ipv4"`
		IPV6 struct {
			LinkLocal net.IP `xml:"link-local" json:"link_local"`
		} `xml:"ipv6" json:"ipv6"`
	} `xml:"ip" json:"ip"`

	RadioCommon struct {
		NodeConfig struct {
			DefaultSsidProfile struct {
				Ssid     string `xml:"ssid" json:"ssid"`
				Password string `xml:"password" json:"password"`
			} `xml:"default-ssid-profile" json:"default_ssid_profile"`
			OperationMode          string `xml:"operation-mode" json:"operation_mode"`
			LinkDistance           string `xml:"link-distance" json:"link_distance"`
			AvailableOperationMode struct {
				OperationMode string `xml:"operation-mode" json:"operation_mode"`
			} `xml:"available-operation-mode" json:"available_operation_mode"`
			TxPowerControl string `xml:"tx-power-control" json:"tx_power_control"`
		} `xml:"node-config" json:"node_config"`
		SectorsConfig struct {
			Sector []struct {
				Index       string `xml:"index" json:"index"`
				Alias       string `xml:"alias" json:"alias"`
				AdminStatus string `xml:"admin-status" json:"admin_status"`
				State       struct {
					MacAddr      string `xml:"mac-addr" json:"mac_addr"`
					Temperatures struct {
						ModemTemperature string `xml:"modem-temperature" json:"modem_temperature"`
						Rf               []struct {
							Index         string `xml:"index" json:"index"`
							RfTemperature string `xml:"rf-temperature" json:"rf_temperature"`
						} `xml:"rf" json:"rf"`
					} `xml:"temperatures" json:"temperatures"`
				} `xml:"state" json:"state"`
			} `xml:"sector" json:"sector"`
		} `xml:"sectors-config" json:"sectors-config"`
		Links struct {
			Active struct {
				RemoteAssignedName      string `xml:"remote-assigned-name" json:"remote_assigned_name"`
				ActualRemoteSectorIndex string `xml:"actual-remote-sector-index" json:"actual_remote_sector_index"`
				ActualLocalSectorIndex  string `xml:"actual-local-sector-index" json:"actual_local_sector_index"`
				RemoteMacAddr           string `xml:"remote-mac-addr" json:"remote_mac_addr"`
				LocalRole               string `xml:"local-role" json:"local_role"`
				LinkUptime              string `xml:"link-uptime" json:"link_uptime"`
				Rssi                    string `xml:"rssi" json:"rssi"`
				Snr                     string `xml:"snr" json:"snr"`
				McsRx                   string `xml:"mcs-rx" json:"mcs_rx"`
				McsTx                   string `xml:"mcs-tx" json:"mcs_tx"`
				TxPer                   string `xml:"tx-per" json:"tx_per"`
				RxPer                   string `xml:"rx-per" json:"rx_per"`
				TxPowerIndex            string `xml:"tx-power-index" json:"tx_power_index"`
				SpeedRx                 string `xml:"speed-rx" json:"speed_rx"`
				SpeedTx                 string `xml:"speed-tx" json:"speed_tx"`
				TxBeam                  struct {
					BeamIndex     string `xml:"beam-index" json:"beam_index"`
					BeamAzimuth   string `xml:"beam-azimuth" json:"beam_azimuth"`
					BeamElevation string `xml:"beam-elevation" json:"beam_elevation"`
				} `xml:"tx-beam" json:"tx_beam"`
				RxBeam struct {
					BeamIndex     string `xml:"beam-index" json:"beam_index"`
					BeamAzimuth   string `xml:"beam-azimuth" json:"beam_azimuth"`
					BeamElevation string `xml:"beam-elevation" json:"beam_elevation"`
				} `xml:"rx-beam" json:"rx_beam"`
				Counters struct {
					RxOk                 int64 `xml:"rx-ok" json:"rx_ok"`
					TxOk                 int64 `xml:"tx-ok" json:"tx_ok"`
					TxFail               int64 `xml:"tx-fail" json:"tx_fail"`
					RxFail               int64 `xml:"rx-fail" json:"rx_fail"`
					RxHcsFail            int64 `xml:"rx-hcs-fail" json:"rx_hcs_fail"`
					TxFailures           int64 `xml:"tx-failures" json:"tx_failures"`
					RxFailures           int64 `xml:"rx-failures" json:"rx_failures"`
					RxDropBufSize        int64 `xml:"rx-drop-buf-size" json:"rx_drop_buf_size"`
					RxDropEncryptionFail int64 `xml:"rx-drop-encryption-fail" json:"rx_drop_encryption_fail"`
					RxDropRaMismatch     int64 `xml:"rx-drop-ra-mismatch" json:"rx_drop_ra_mismatch"`
					RxDropUnexpected     int64 `xml:"rx-drop-unexpected" json:"rx_drop_unexpected"`
				} `xml:"counters" json:"counters"`
			} `xml:"active" json:"active"`
		} `xml:"links" json:"links"`
	} `xml:"radio-common" json:"radio_common"`

	RadioDn struct {
		NodeConfig struct {
			DefaultRadioProfile RadioProfile `xml:"default-radio-profile" json:"default_radio_profile"`
			IsPopDn             bool         `xml:"is-pop-dn" json:"is_pop_dn"`
			IgnoreGps           string       `xml:"ignore-gps" json:"ignore_gps"`
		} `xml:"node-config" json:"node_config"`
		SectorsConfig struct {
			Sector []struct {
				Index        uint8        `xml:"index" json:"index"`
				RadioProfile RadioProfile `xml:"radio-profile" json:"radio_profile`
			} `xml:"sector" json:"sector"`
		} `xml:"sectors-config" json:"sectors_config"`
		Links struct {
			Configured struct {
				RemoteAssignedName string `xml:"remote-assigned-name" json:"remote_assigned_name"`
				RemoteSector       struct {
					Index string `xml:"index" json:"index"`
				} `xml:"remote-sector" json:"remote_sector"`
				LocalSector struct {
					Index string `xml:"index"`
				} `xml:"local-sector" json:"local_sector"`
				ControlSuperframe string `xml:"control-superframe" json:"control_superframe"`
				ResponderNodeType string `xml:"responder-node-type" json:"responder_node_type"`
				AdminStatus       string `xml:"admin-status" json:"admin_status"`
				State             string `xml:"state" json:"state"`
			} `xml:"configured" json:"configured"`
		} `xml:"links" json:"links"`
	} `xml:"radio-dn" json:"radio_dn"`

	System struct {
		Name         string `xml:"name" json:"name"`
		RebootNeeded struct {
		} `xml:"reboot-needed" json:"reboot_needed"` // TODO
		Control struct {
			SoftWatchdogEnabled string `xml:"soft-watchdog-enabled" json:"soft_watchdog_enabled"`
		} `xml:"control"`
		State struct {
			Product     string `xml:"product" json:"product"`
			DateAndTime string `xml:"date-and-time" json:"date_and_time"`
			Uptime      string `xml:"uptime" json:"uptime"`
			BanksInfo   struct {
				Banks []struct {
					Number            string `xml:"number" json:"number"`
					SoftwareVersion   string `xml:"software-version" json:"software_version"`
					Status            string `xml:"status" json:"status"`
					ScheduledToSwitch string `xml:"scheduled-to-switch" json:"scheduled_to_switch"`
				} `xml:"banks"`
			} `xml:"banks-info"`
			SwUpgradeInfo struct {
				DownloadAndBurningState string `xml:"download-and-burning-state" json:"download_and_burning_state"`
			} `xml:"sw-upgrade-info" json:"sw_upgrade_info"`
			Gps struct {
				FixMode             string `xml:"fix-mode" json:"fix_mode"`
				FixSatellitesNumber string `xml:"fix-satellites-number" json:"fix_satellites_number"`
			} `xml:"gps" json:"gps"`
		} `xml:"state" json:"state"`
	} `xml:"system" json:"system"`

	UserBridge struct {
		Bridge struct {
			BridgeID   string `xml:"bridge-id" json:"bridge_id"`
			BridgePort []struct {
				BridgePortID   string `xml:"bridge-port-id" json:"bridge_port_id"`
				Interface      string `xml:"interface" json:"interface"`
				BridgePortType string `xml:"bridge-port-type" json:"bridge_port_type"`
			} `xml:"bridge-port" json:"bridge_port"`
		} `xml:"bridge" json:"bridge"`
	} `xml:"user-bridge" json:"user_bridge"`

	UserManagement struct {
		User struct {
			Username string `xml:"username" json:"username"`
			Password string `xml:"password" json:"password"`
		} `xml:"user" json:"user"`
	} `xml:"user-management" json:"user_management"`
}
