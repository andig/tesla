package tesla

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// Contains the current charge states that exist within the vehicle
type ChargeState struct {
	ChargingState               string      `json:"charging_state"`
	ChargeLimitSoc              int         `json:"charge_limit_soc"`
	ChargeLimitSocStd           int         `json:"charge_limit_soc_std"`
	ChargeLimitSocMin           int         `json:"charge_limit_soc_min"`
	ChargeLimitSocMax           int         `json:"charge_limit_soc_max"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	MaxRangeChargeCounter       int         `json:"max_range_charge_counter"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	BatteryRange                float64     `json:"battery_range"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	BatteryLevel                int         `json:"battery_level"`
	UsableBatteryLevel          int         `json:"usable_battery_level"`
	BatteryCurrent              interface{} `json:"battery_current"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargerVoltage              interface{} `json:"charger_voltage"`
	ChargerPilotCurrent         interface{} `json:"charger_pilot_current"`
	ChargerActualCurrent        interface{} `json:"charger_actual_current"`
	ChargerPower                interface{} `json:"charger_power"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	TripCharging                interface{} `json:"trip_charging"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	MotorizedChargePort         bool        `json:"motorized_charge_port"`
	ScheduledChargingStartTime  interface{} `json:"scheduled_charging_start_time"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	UserChargeEnableRequest     interface{} `json:"user_charge_enable_request"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	EuVehicle                   bool        `json:"eu_vehicle"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeCurrentRequest        int         `json:"charge_current_request"`
	ChargeCurrentRequestMax     int         `json:"charge_current_request_max"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	ManagedChargingStartTime    interface{} `json:"managed_charging_start_time"`
	ChargePortcoldWeatherMode   bool        `json:"charge_port_cold_weather_mode"`
	ConnChargeCable             string      `json:"conn_charge_cable"`
	FastChargerBrand            string      `json:"fast_charger_brand"`
	MinutesToFullCharge         int         `json:"minutes_to_full_charge"`
}

// Contains the current climate states availale from the vehicle
type ClimateState struct {
	InsideTemp                 float64     `json:"inside_temp"`
	OutsideTemp                float64     `json:"outside_temp"`
	DriverTempSetting          float64     `json:"driver_temp_setting"`
	PassengerTempSetting       float64     `json:"passenger_temp_setting"`
	LeftTempDirection          float64     `json:"left_temp_direction"`
	RightTempDirection         float64     `json:"right_temp_direction"`
	IsAutoConditioningOn       bool        `json:"is_auto_conditioning_on"`
	IsFrontDefrosterOn         bool        `json:"is_front_defroster_on"`
	IsRearDefrosterOn          bool        `json:"is_rear_defroster_on"`
	FanStatus                  interface{} `json:"fan_status"`
	IsClimateOn                bool        `json:"is_climate_on"`
	MinAvailTemp               float64     `json:"min_avail_temp"`
	MaxAvailTemp               float64     `json:"max_avail_temp"`
	SeatHeaterLeft             int         `json:"seat_heater_left"`
	SeatHeaterRight            int         `json:"seat_heater_right"`
	SeatHeaterRearLeft         int         `json:"seat_heater_rear_left"`
	SeatHeaterRearRight        int         `json:"seat_heater_rear_right"`
	SeatHeaterRearCenter       int         `json:"seat_heater_rear_center"`
	SeatHeaterRearRightBack    int         `json:"seat_heater_rear_right_back"`
	SeatHeaterRearLeftBack     int         `json:"seat_heater_rear_left_back"`
	SmartPreconditioning       bool        `json:"smart_preconditioning"`
	BatteryHeater              bool        `json:"battery_heater"`
	BatteryHeaterNoPower       interface{} `json:"battery_heater_no_power"`
	ClimateKeeperMode          string      `json:"climate_keeper_mode"`
	DefrostMode                int         `json:"defrost_mode"`
	IsPreconditioning          bool        `json:"is_preconditioning"`
	RemoteHeaterControlEnabled bool        `json:"remote_heater_control_enabled"`
	SideMirrorHeaters          bool        `json:"side_mirror_heaters"`
	WiperBladeHeater           bool        `json:"wiper_blade_heater"`
}

// Contains the current drive state of the vehicle
type DriveState struct {
	ShiftState              interface{} `json:"shift_state"`
	Speed                   float64     `json:"speed"`
	Latitude                float64     `json:"latitude"`
	Longitude               float64     `json:"longitude"`
	Heading                 int         `json:"heading"`
	GpsAsOf                 int64       `json:"gps_as_of"`
	NativeLatitude          float64     `json:"native_latitude"`
	NativeLocationSupported int         `json:"native_location_supported"`
	NativeLongitude         float64     `json:"native_longitude"`
	NativeType              string      `json:"native_type"`
	Power                   int         `json:"power"`
}

// Contains the current GUI settings of the vehicle
type GuiSettings struct {
	GuiDistanceUnits    string `json:"gui_distance_units"`
	GuiTemperatureUnits string `json:"gui_temperature_units"`
	GuiChargeRateUnits  string `json:"gui_charge_rate_units"`
	Gui24HourTime       bool   `json:"gui_24_hour_time"`
	GuiRangeDisplay     string `json:"gui_range_display"`
	ShowRangeUnits      bool   `json:"show_range_units"`
}

// Contains the current state of the vehicle
type VehicleState struct {
	APIVersion              int     `json:"api_version"`
	AutoParkState           string  `json:"autopark_state"`
	AutoParkStateV2         string  `json:"autopark_state_v2"`
	CalendarSupported       bool    `json:"calendar_supported"`
	CarType                 string  `json:"car_type"`
	CarVersion              string  `json:"car_version"`
	CenterDisplayState      int     `json:"center_display_state"`
	DarkRims                bool    `json:"dark_rims"`
	Df                      int     `json:"df"`
	Dr                      int     `json:"dr"`
	ExteriorColor           string  `json:"exterior_color"`
	Ft                      int     `json:"ft"`
	HasSpoiler              bool    `json:"has_spoiler"`
	Locked                  bool    `json:"locked"`
	NotificationsSupported  bool    `json:"notifications_supported"`
	Odometer                float64 `json:"odometer"`
	ParsedCalendarSupported bool    `json:"parsed_calendar_supported"`
	PerfConfig              string  `json:"perf_config"`
	Pf                      int     `json:"pf"`
	Pr                      int     `json:"pr"`
	RearSeatHeaters         int     `json:"rear_seat_heaters"`
	RemoteStart             bool    `json:"remote_start"`
	RemoteStartSupported    bool    `json:"remote_start_supported"`
	Rhd                     bool    `json:"rhd"`
	RoofColor               string  `json:"roof_color"`
	Rt                      int     `json:"rt"`
	SentryMode              bool    `json:"sentry_mode"`
	SentryModeAvailable     bool    `json:"sentry_mode_available"`
	SeatType                int     `json:"seat_type"`
	SpoilerType             string  `json:"spoiler_type"`
	SunRoofInstalled        int     `json:"sun_roof_installed"`
	SunRoofPercentOpen      int     `json:"sun_roof_percent_open"`
	SunRoofState            string  `json:"sun_roof_state"`
	ThirdRowSeats           string  `json:"third_row_seats"`
	ValetMode               bool    `json:"valet_mode"`
	VehicleName             string  `json:"vehicle_name"`
	WheelType               string  `json:"wheel_type"`
	FdWindow                int     `json:"fd_window"`
	FpWindow                int     `json:"fp_window"`
	RdWindow                int     `json:"rd_window"`
	RpWindow                int     `json:"rp_window"`
	IsUserPresent           bool    `json:"is_user_present"`
	RemoteStartEnabled      bool    `json:"remote_start_enabled"`
	ValetPinNeeded          bool    `json:"valet_pin_needed"`
	MediaState              struct {
		RemoteControlEnabled bool `json:"remote_control_enabled"`
	} `json:"media_state"`
	SoftwareUpdate struct {
		DownloadPerc        int    `json:"download_perc"`
		ExpectedDurationSec int    `json:"expected_duration_sec"`
		InstallPerc         int    `json:"install_perc"`
		Status              string `json:"status"`
		Version             string `json:"version"`
	} `json:"software_update" `
	SpeedLimitMode struct {
		Active          bool    `json:"active"`
		CurrentLimitMph float64 `json:"current_limit_mph"`
		MaxLimitMph     int     `json:"max_limit_mph"`
		MinLimitMph     int     `json:"min_limit_mph"`
		PinCodeSet      bool    `json:"pin_code_set"`
	} `json:"speed_limit_mode"`
}

type ServiceData struct {
	ServiceETC    time.Time `json:"service_etc"`
	ServiceStatus string    `json:"service_status"`
}

// Represents the request to get the states of the vehicle
type StateRequest struct {
	Response struct {
		Timestamp time.Time `json:"timestamp"`
		*ChargeState
		*ClimateState
		*DriveState
		*GuiSettings
		*VehicleState
		*ServiceData
	} `json:"response"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// The response when a state is requested
type MobileEnabledResponse struct {
	Bool bool `json:"response"`
}

// MobileEnabled returns if the vehicle is mobile enabled for Tesla API control
func (v *Vehicle) MobileEnabled() (bool, error) {
	r := &MobileEnabledResponse{}
	if err := v.c.getJSON(v.c.BaseURL+"/vehicles/"+strconv.FormatInt(v.ID, 10)+"/mobile_enabled", r); err != nil {
		return false, err
	}
	return r.Bool, nil
}

type timeSecs struct {
	time.Time
}

func (t *timeSecs) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	t.Time = (time.Unix(i, 0))
	return nil
}

type timeUsecs struct {
	time.Time
}

func (t *timeUsecs) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	t.Time = (time.Unix(i/1000, i%1000))
	return nil
}

type NearbyChargingSitesResponse struct {
	Response struct {
		CongestionSyncTimeUtcSecs timeSecs `json:"congestion_sync_time_utc_secs"`
		DestinationCharging       []struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"long"`
			} `json:"location"`
			Name          string  `json:"name"`
			Type          string  `json:"type"`
			DistanceMiles float64 `json:"distance_miles"`
		} `json:"destination_charging"`
		Superchargers []struct {
			Location struct {
				Lat  float64 `json:"lat"`
				Long float64 `json:"long"`
			} `json:"location"`
			Name            string  `json:"name"`
			Type            string  `json:"type"`
			DistanceMiles   float64 `json:"distance_miles"`
			AvailableStalls int     `json:"available_stalls"`
			TotalStalls     int     `json:"total_stalls"`
			SiteClosed      bool    `json:"site_closed"`
		} `json:"superchargers"`
		Timestamp timeUsecs `json:"timestamp"`
	} `json:"response"`
}

func (v *Vehicle) NearbyChargingSites() (*NearbyChargingSitesResponse, error) {
	resp := &NearbyChargingSitesResponse{}
	if err := v.c.getJSON(v.c.BaseURL+"/vehicles/"+strconv.FormatInt(v.ID, 10)+"/nearby_charging_sites", resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ChargeState returns the charge state of the vehicle
func (v *Vehicle) ChargeState() (*ChargeState, error) {
	stateRequest, err := v.c.fetchState("/charge_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.ChargeState, nil
}

// ClimateState returns the climate state of the vehicle
func (v Vehicle) ClimateState() (*ClimateState, error) {
	stateRequest, err := v.c.fetchState("/climate_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.ClimateState, nil
}

func (v Vehicle) DriveState() (*DriveState, error) {
	stateRequest, err := v.c.fetchState("/drive_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.DriveState, nil
}

// GuiSettings returns the GUI settings of the vehicle
func (v Vehicle) GuiSettings() (*GuiSettings, error) {
	stateRequest, err := v.c.fetchState("/gui_settings", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.GuiSettings, nil
}

func (v Vehicle) VehicleState() (*VehicleState, error) {
	stateRequest, err := v.c.fetchState("/vehicle_state", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.VehicleState, nil
}

func (v Vehicle) ServiceData() (*ServiceData, error) {
	stateRequest, err := v.c.fetchState("/service_data", v.ID)
	if err != nil {
		return nil, err
	}
	return stateRequest.Response.ServiceData, nil
}

func stateError(sr *StateRequest) error {
	if sr.Error == "" {
		return nil
	}

	if sr.ErrorDescription != "" {
		return fmt.Errorf("%s: %s", sr.Error, sr.ErrorDescription)
	}
	return fmt.Errorf("%s", sr.Error)
}

// A utility function to fetch the appropriate state of the vehicle
func (c *Client) fetchState(resource string, id int64) (*StateRequest, error) {
	stateRequest := &StateRequest{}
	if err := c.getJSON(c.BaseURL+"/vehicles/"+strconv.FormatInt(id, 10)+"/data_request"+resource, stateRequest); err != nil {
		return nil, err
	}
	if err := stateError(stateRequest); err != nil {
		return nil, err
	}
	return stateRequest, nil
}

// Data : Get data of the vehicle (calling this will not permit the car to sleep)
func (v Vehicle) Data(vid int64) (*StateRequest, error) {
	log.Println("Retreiving vehicle data")
	stateRequest := &StateRequest{}

	/*log.Println(v.c.URL + "/vehicles/" + strconv.FormatInt(vid, 10) + "/vehicle_data")
	body, err := v.c.get(v.c.URL + "/vehicles/" + strconv.FormatInt(vid, 10) + "/vehicle_data")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, stateRequest); err != nil {
		return nil, err
	}*/

	// climate_state
	stateRequestClimate, err := v.c.fetchState("/climate_state", v.ID)
	if err != nil {
		log.Println("Error getting climate_state")
		return nil, err
	}
	stateRequest.Response.ClimateState = stateRequestClimate.Response.ClimateState

	// drive_state
	stateRequestGui, err := v.c.fetchState("/drive_state", v.ID)
	if err != nil {
		log.Println("Error getting drive_state")
		return nil, err
	}
	stateRequest.Response.DriveState = stateRequestGui.Response.DriveState

	// gui_settings
	stateRequestSettings, err := v.c.fetchState("/gui_settings", v.ID)
	if err != nil {
		log.Println("Error getting gui_settings")
		return nil, err
	}
	stateRequest.Response.GuiSettings = stateRequestSettings.Response.GuiSettings

	// vehicle_state
	stateRequestVehicle, err := v.c.fetchState("/vehicle_state", v.ID)
	if err != nil {
		log.Println("Error getting vehicle_state")
		return nil, err
	}
	stateRequest.Response.VehicleState = stateRequestVehicle.Response.VehicleState

	// charge_state
	stateRequestCharge, err := v.c.fetchState("/charge_state", v.ID)
	if err != nil {
		log.Println("Error getting charge_state")
		return nil, err
	}
	stateRequest.Response.ChargeState = stateRequestCharge.Response.ChargeState

	return stateRequest, nil
}
