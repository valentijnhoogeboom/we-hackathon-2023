package models

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	MeterID                   uint    `json:"meter_id"`
	Meter                     Meter   `json:"-"`
	WifiSsid                  string  `json:"wifi_ssid"`
	WifiStrength              int     `json:"wifi_strength"`
	SmrVersion                int     `json:"smr_version"`
	MeterModel                string  `json:"meter_model"`
	UniqueId                  string  `json:"unique_id"`
	ActiveTariff              int     `json:"active_tariff"`
	TotalPowerImportKwh       float64 `json:"total_power_import_kwh"`
	TotalPowerImportT1Kwh     float64 `json:"total_power_import_t1_kwh"`
	TotalPowerImportT2Kwh     float64 `json:"total_power_import_t2_kwh"`
	TotalPowerExportKwh       int     `json:"total_power_export_kwh"`
	TotalPowerExportT1Kwh     int     `json:"total_power_export_t1_kwh"`
	TotalPowerExportT2Kwh     int     `json:"total_power_export_t2_kwh"`
	ActivePowerW              int     `json:"active_power_w"`
	ActivePowerL1W            int     `json:"active_power_l1_w"`
	ActivePowerL2W            int     `json:"active_power_l2_w"`
	ActivePowerL3W            int     `json:"active_power_l3_w"`
	ActiveCurrentL1A          int     `json:"active_current_l1_a"`
	ActiveCurrentL2A          int     `json:"active_current_l2_a"`
	ActiveCurrentL3A          int     `json:"active_current_l3_a"`
	VoltageSagL1Count         int     `json:"voltage_sag_l1_count"`
	VoltageSagL2Count         int     `json:"voltage_sag_l2_count"`
	VoltageSagL3Count         int     `json:"voltage_sag_l3_count"`
	VoltageSwellL1Count       int     `json:"voltage_swell_l1_count"`
	VoltageSwellL2Count       int     `json:"voltage_swell_l2_count"`
	VoltageSwellL3Count       int     `json:"voltage_swell_l3_count"`
	AnyPowerFailCount         int     `json:"any_power_fail_count"`
	LongPowerFailCount        int     `json:"long_power_fail_count"`
	GasUniqueId               string  `json:"gas_unique_id"`
	ActivePowerAverageW       float64 `json:"active_power_average_w"`
	MonthlyPowerPeakW         float64 `json:"montly_power_peak_w"`
	MonthlyPowerPeakTimestamp int64   `json:"montly_power_peak_timestamp"`
}
