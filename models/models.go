package models

import "go.mongodb.org/mongo-driver/bson"

type Name struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Category struct {
	Category string `json:"category"`
}

type PhysicalProp struct {
	Mass                     string `json:"mass,omitempty"`
	MassAtEndOfBoosterBurn   string `json:"massAtEndOfBoosterBurn,omitempty"`
	MassAtEndOfSustainerBurn string `json:"massAtEndOfSustainerBurn,omitempty"`
	Calibre                  string `json:"calibre,omitempty"`
	Length                   string `json:"length,omitempty"`
}

type EngineProp struct {
	ForceExertedByBooster      string `json:"forceExertedByBooster,omitempty"`
	BurnTimeOfBooster          string `json:"burnTimeOfBooster,omitempty"`
	RawAccelerationAtIgnition  string `json:"rawAccelerationAtIgnition,omitempty"`
	SpecificImpulseOfBooster   string `json:"specificImpulseOfBooster,omitempty"`
	DeltaSpeedOfBooster        string `json:"deltaSpeedOfBooster,omitempty"`
	BoosterStartDelay          string `json:"boosterStartDelay,omitempty"`
	ForceExertedBySustainer    string `json:"forceExertedBySustainer,omitempty"`
	BurnTimeOfSustainer        string `json:"burnTimeOfSustainer,omitempty"`
	SpecificImpulseOfSustainer string `json:"specificImpulseOfSustainer,omitempty"`
	DeltaSpeedOfSustainer      string `json:"deltaSpeedOfSustainer,omitempty"`
	TotalDeltaSpeed            string `json:"totalDeltaSpeed,omitempty"`
}

type FuseAndWarheadProp struct {
	ExplosiveMass                string `json:"explosiveMass,omitempty"`
	TandemCharge                 string `json:"tandemCharge,omitempty"`
	Penetration                  string `json:"penetration,omitempty"`
	ProximityFuse                string `json:"proximityFuse,omitempty"`
	ProximityFuseRange           string `json:"proximityFuseRange,omitempty"`
	ProximityFuseArmingDistance  string `json:"proximityFuseArmingDistance,omitempty"`
	ProximityFuseShellDetection  string `json:"proximityFuseShellDetection,omitempty"`
	ProximityFuseMinimumAltitude string `json:"proximityFuseMinimumAltitude,omitempty"`
	ProximityFuseDelay           string `json:"proximityFuseDelay,omitempty"`
}

type GuidanceProp struct {
	Zoom                                          string `json:"zoom,omitempty"`
	GuidanceType                                  string `json:"guidanceType,omitempty"`
	GuidanceStartDelay                            string `json:"guidanceStartDelay,omitempty"`
	GuidanceDuration                              string `json:"guidanceDuration,omitempty"`
	GuidanceRange                                 string `json:"guidanceRange,omitempty"`
	LaunchSector                                  string `json:"launchSector,omitempty"`
	ControlConeFOV                                string `json:"controlConeFOV,omitempty"`
	AimTrackingSensitivity                        string `json:"aimTrackingSensitivity,omitempty"`
	MaximumAngleAllowedBetweenMissileAndCrosshair string `json:"maximumAngleAllowedBetweenMissileAndCrosshair,omitempty"`
	SeekerWarmUpTime                              string `json:"seekerWarmUpTime,omitempty"`
	SeekerSearchDuration                          string `json:"seekerSearchDuration,omitempty"`
	FieldOfView                                   string `json:"fieldOfView,omitempty"`
	OpticSightFieldOfView                         string `json:"opticSightFieldOfView,omitempty"`
	GimbalLimit                                   string `json:"gimbalLimit,omitempty"`
	TrackRate                                     string `json:"trackRate,omitempty"`
	UncageSeekerBeforeLaunch                      string `json:"uncageSeekerBeforeLaunch,omitempty"`
	MaximumLockAngleBeforeLaunch                  string `json:"maximumLockAngleBeforeLaunch,omitempty"`
	MinimumAngleBetweenSeekerAndSunForNotCapture  string `json:"minimumAngleBetweenSeekerAndSunForNotCapture,omitempty"`
	CanLockGround                                 string `json:"canLockGround,omitempty"`
	LockOnRangeGround                             string `json:"lockOnRangeGround,omitempty"`
	LockOnRangeVehicle                            string `json:"lockOnRangeVehicle,omitempty"`
	LockOnRangeFromRearAspect                     string `json:"lockOnRangeFromRearAspect,omitempty"`
	FlareDetectionRange                           string `json:"flareDetectionRange,omitempty"`
	IRCMDetectionRange                            string `json:"IRCMDetectionRange,omitempty"`
	DIRCMDetectionRange                           string `json:"DIRCMDetectionRange,omitempty"`
	HeadOnLockOnRangeAgainstAfterburnerTarget     string `json:"headOnLockOnRangeAgainstAfterburnerTarget,omitempty"`
	IRCCM                                         string `json:"IRCCM,omitempty"`
	IRCCMType                                     string `json:"IRCCMType,omitempty"`
	IRCCMFieldOfView                              string `json:"IRCCMFieldOfView,omitempty"`
	IRCCMRejectionThreshold                       string `json:"IRCCMRejectionThreshold,omitempty"`
	IRCCMReactionTime                             string `json:"IRCCMReactionTime,omitempty"`
	LockOnRangeFromAllAspect                      string `json:"lockOnRangeFromAllAspect,omitempty"`
	MaximumBreakLockTime                          string `json:"maximumBreakLockTime,omitempty"`
	CanBeSlavedToRadar                            string `json:"canBeSlavedToRadar,omitempty"`
	CanLockAfterLaunch                            string `json:"canLockAfterLaunch,omitempty"`
	Band                                          string `json:"band,omitempty"`
	AngularSpeedRejectionThreshold                string `json:"angularSpeedRejectionThreshold,omitempty"`
	AccelerationRejectionThresholdRange           string `json:"accelerationRejectionThresholdRange,omitempty"`
	SidelobeAttenuation                           string `json:"sidelobeAttenuation,omitempty"`
	TransmitterPower                              string `json:"transmitterPower,omitempty"`
	TransmitterAngleOfHalfSensitivity             string `json:"transmitterAngleOfHalfSensitivity,omitempty"`
	TransmitterSidelobeSensitivity                string `json:"transmitterSidelobeSensitivity,omitempty"`
	ReceiverAngleOfHalfSensitivity                string `json:"receiverAngleOfHalfSensitivity,omitempty"`
	ReceiverSidelobeSensitivity                   string `json:"receiverSidelobeSensitivity,omitempty"`
	DistanceMinimumValue                          string `json:"distanceMinimumValue,omitempty"`
	DistanceMaximumValue                          string `json:"distanceMaximumValue,omitempty"`
	DistanceWidth                                 string `json:"distanceWidth,omitempty"`
	DistanceMinimumSignalGate                     string `json:"distanceMinimumSignalGate,omitempty"`
	DistanceRefWidth                              string `json:"distanceRefWidth,omitempty"`
	DopplerSpeedMinimumValue                      string `json:"dopplerSpeedMinimumValue,omitempty"`
	DopplerSpeedMaximumValue                      string `json:"dopplerSpeedMaximumValue,omitempty"`
	DopplerSpeedWidth                             string `json:"dopplerSpeedWidth,omitempty"`
	DopplerSpeedRefWidth                          string `json:"dopplerSpeedRefWidth,omitempty"`
	DopplerSpeedMinimumSignalGate                 string `json:"dopplerSpeedMinimumSignalGate,omitempty"`
	DopplerSpeedGateSearchRange                   string `json:"dopplerSpeedGateSearchRange,omitempty"`
	DopplerSpeedGateAlphaFilter                   string `json:"dopplerSpeedGateAlphaFilter,omitempty"`
	DopplerSpeedGateBetaFilter                    string `json:"dopplerSpeedGateBetaFilter,omitempty"`
	ProportionalNavigationMultiplier              string `json:"proportionalNavigationMultiplier,omitempty"`
	BaseIndicatedAirSpeed                         string `json:"baseIndicatedAirSpeed,omitempty"`
	PIDProportionalTerm                           string `json:"PIDProportionalTerm,omitempty"`
	PIDIntegralTerm                               string `json:"PIDIntegralTerm,omitempty"`
	PIDIntegralTermLimit                          string `json:"PIDIntegralTermLimit,omitempty"`
	PIDDerivativeTerm                             string `json:"PIDDerivativeTerm,omitempty"`
	InertialGuidanceDriftSpeed                    string `json:"inertialGuidanceDriftSpeed,omitempty"`
	InertialNavigation                            string `json:"inertialNavigation,omitempty"`
	InertialNavigationDriftSpeed                  string `json:"inertialNavigationDriftSpeed,omitempty"`
}

type FlightProp struct {
	MaximumLaunchAngleHorizontalVertical               string `json:"maximumLaunchAngleHorizontalVertical,omitempty"`
	AimSensitivity                                     string `json:"aimSensitivity,omitempty"`
	MaximumAxisValues                                  string `json:"maximumAxisValues,omitempty"`
	MaximumFinAngleOfAttack                            string `json:"maximumFinAngleOfAttack,omitempty"`
	FinsLateralAcceleration                            string `json:"finsLateralAcceleration,omitempty"`
	MaximumAOA                                         string `json:"maximumAOA,omitempty"`
	MaximumFinLateralAcceleration                      string `json:"maximumFinLateralAcceleration,omitempty"`
	WingAreaMultiplier                                 string `json:"wingAreaMultiplier,omitempty"`
	MaximumLateralAcceleration                         string `json:"maximumLateralAcceleration,omitempty"`
	StartSpeed                                         string `json:"startSpeed,omitempty"`
	MaximumSpeed                                       string `json:"maximumSpeed,omitempty"`
	MinimumRange                                       string `json:"minimumRange,omitempty"`
	MaximumFlightRange                                 string `json:"maximumFlightRange,omitempty"`
	Tracer                                             string `json:"tracer,omitempty"`
	LoadFactorLimitAtLaunch                            string `json:"loadFactorLimitAtLaunch,omitempty"`
	MaximumOverLoad                                    string `json:"maximumOverLoad,omitempty"`
	SeaSkimming                                        string `json:"seaSkimming,omitempty"`
	FlightTimeUntilGuidanceStarts                      string `json:"flightTimeUntilGuidanceStarts,omitempty"`
	FlightTimeWhenPullLimit30                          string `json:"flightTimeWhenPullLimit30%,omitempty"`
	FlightTimeWhenPullLimit40                          string `json:"flightTimeWhenPullLimit40%,omitempty"`
	FlightTimeWhenPullLimit100                         string `json:"flightTimeWhenPullLimit100%,omitempty"`
	Loft                                               string `json:"loft,omitempty"`
	LoftAngle                                          string `json:"loftAngle,omitempty"`
	TargetElevation                                    string `json:"targetElevation,omitempty"`
	MaximumTargetAngularChange                         string `json:"maximumTargetAngularChange"`
	ThrustVectoring                                    string `json:"thrustVectoring,omitempty"`
	ThrustVectoringAngle                               string `json:"thrustVectoringAngle,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage30  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage30%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage50  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage50%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage80  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage80%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage90  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage90%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage100 string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage100%,omitempty"`
}

type Params struct {
	Category           string `json:"category"`
	Name               string `json:"name"`
	PhysicalProp       `json:"physicalProp"`
	EngineProp         `json:"engineProp"`
	FuseAndWarheadProp `json:"fuseAndWarheadProp"`
	GuidanceProp       `json:"guidanceProp"`
	FlightProp         `json:"flightProp"`
}

func NewWeapon(params *Params) *Params {
	return &Params{
		Category: params.Category,
		Name:     params.Name,
		PhysicalProp: PhysicalProp{
			Mass:                     params.Mass,
			MassAtEndOfBoosterBurn:   params.MassAtEndOfBoosterBurn,
			MassAtEndOfSustainerBurn: params.MassAtEndOfSustainerBurn,
			Calibre:                  params.Calibre,
			Length:                   params.Length,
		},
		EngineProp: EngineProp{
			ForceExertedByBooster:      params.ForceExertedByBooster,
			BurnTimeOfBooster:          params.BurnTimeOfBooster,
			RawAccelerationAtIgnition:  params.RawAccelerationAtIgnition,
			SpecificImpulseOfBooster:   params.SpecificImpulseOfBooster,
			DeltaSpeedOfBooster:        params.DeltaSpeedOfBooster,
			BoosterStartDelay:          params.BoosterStartDelay,
			ForceExertedBySustainer:    params.ForceExertedBySustainer,
			BurnTimeOfSustainer:        params.BurnTimeOfSustainer,
			SpecificImpulseOfSustainer: params.SpecificImpulseOfSustainer,
			DeltaSpeedOfSustainer:      params.DeltaSpeedOfSustainer,
			TotalDeltaSpeed:            params.TotalDeltaSpeed,
		},
		FuseAndWarheadProp: FuseAndWarheadProp{
			ExplosiveMass:                params.ExplosiveMass,
			TandemCharge:                 params.TandemCharge,
			Penetration:                  params.Penetration,
			ProximityFuse:                params.ProximityFuse,
			ProximityFuseRange:           params.ProximityFuseRange,
			ProximityFuseArmingDistance:  params.ProximityFuseArmingDistance,
			ProximityFuseShellDetection:  params.ProximityFuseShellDetection,
			ProximityFuseMinimumAltitude: params.ProximityFuseMinimumAltitude,
			ProximityFuseDelay:           params.ProximityFuseDelay,
		},
		GuidanceProp: GuidanceProp{
			Zoom:                   params.Zoom,
			GuidanceType:           params.GuidanceType,
			GuidanceStartDelay:     params.GuidanceStartDelay,
			GuidanceDuration:       params.GuidanceDuration,
			GuidanceRange:          params.GuidanceRange,
			LaunchSector:           params.LaunchSector,
			ControlConeFOV:         params.ControlConeFOV,
			AimTrackingSensitivity: params.AimTrackingSensitivity,
			MaximumAngleAllowedBetweenMissileAndCrosshair: params.MaximumAngleAllowedBetweenMissileAndCrosshair,
			SeekerWarmUpTime:             params.SeekerWarmUpTime,
			SeekerSearchDuration:         params.SeekerSearchDuration,
			FieldOfView:                  params.FieldOfView,
			OpticSightFieldOfView:        params.OpticSightFieldOfView,
			GimbalLimit:                  params.GimbalLimit,
			TrackRate:                    params.TrackRate,
			UncageSeekerBeforeLaunch:     params.UncageSeekerBeforeLaunch,
			MaximumLockAngleBeforeLaunch: params.MaximumLockAngleBeforeLaunch,
			MinimumAngleBetweenSeekerAndSunForNotCapture: params.MinimumAngleBetweenSeekerAndSunForNotCapture,
			CanLockGround:                             params.CanLockGround,
			LockOnRangeGround:                         params.LockOnRangeGround,
			LockOnRangeVehicle:                        params.LockOnRangeVehicle,
			LockOnRangeFromRearAspect:                 params.LockOnRangeFromRearAspect,
			FlareDetectionRange:                       params.FlareDetectionRange,
			IRCMDetectionRange:                        params.IRCMDetectionRange,
			DIRCMDetectionRange:                       params.DIRCMDetectionRange,
			HeadOnLockOnRangeAgainstAfterburnerTarget: params.HeadOnLockOnRangeAgainstAfterburnerTarget,
			IRCCM:                               params.IRCCM,
			IRCCMType:                           params.IRCCMType,
			IRCCMFieldOfView:                    params.IRCCMFieldOfView,
			IRCCMRejectionThreshold:             params.IRCCMRejectionThreshold,
			IRCCMReactionTime:                   params.IRCCMReactionTime,
			LockOnRangeFromAllAspect:            params.LockOnRangeFromAllAspect,
			MaximumBreakLockTime:                params.MaximumBreakLockTime,
			CanBeSlavedToRadar:                  params.CanBeSlavedToRadar,
			CanLockAfterLaunch:                  params.CanLockAfterLaunch,
			Band:                                params.Band,
			AngularSpeedRejectionThreshold:      params.AngularSpeedRejectionThreshold,
			AccelerationRejectionThresholdRange: params.AccelerationRejectionThresholdRange,
			SidelobeAttenuation:                 params.SidelobeAttenuation,
			TransmitterPower:                    params.TransmitterPower,
			TransmitterAngleOfHalfSensitivity:   params.TransmitterAngleOfHalfSensitivity,
			TransmitterSidelobeSensitivity:      params.TransmitterSidelobeSensitivity,
			ReceiverAngleOfHalfSensitivity:      params.ReceiverAngleOfHalfSensitivity,
			ReceiverSidelobeSensitivity:         params.ReceiverSidelobeSensitivity,
			DistanceMinimumValue:                params.DistanceMinimumValue,
			DistanceMaximumValue:                params.DistanceMaximumValue,
			DistanceWidth:                       params.DistanceWidth,
			DistanceMinimumSignalGate:           params.DistanceMinimumSignalGate,
			DistanceRefWidth:                    params.DistanceRefWidth,
			DopplerSpeedMinimumValue:            params.DopplerSpeedMinimumValue,
			DopplerSpeedMaximumValue:            params.DopplerSpeedMaximumValue,
			DopplerSpeedWidth:                   params.DopplerSpeedWidth,
			DopplerSpeedRefWidth:                params.DopplerSpeedRefWidth,
			DopplerSpeedMinimumSignalGate:       params.DopplerSpeedMinimumSignalGate,
			DopplerSpeedGateSearchRange:         params.DopplerSpeedGateSearchRange,
			DopplerSpeedGateAlphaFilter:         params.DopplerSpeedGateAlphaFilter,
			DopplerSpeedGateBetaFilter:          params.DopplerSpeedGateBetaFilter,
			ProportionalNavigationMultiplier:    params.ProportionalNavigationMultiplier,
			BaseIndicatedAirSpeed:               params.BaseIndicatedAirSpeed,
			PIDProportionalTerm:                 params.PIDProportionalTerm,
			PIDIntegralTerm:                     params.PIDIntegralTerm,
			PIDIntegralTermLimit:                params.PIDIntegralTermLimit,
			PIDDerivativeTerm:                   params.PIDDerivativeTerm,
			InertialGuidanceDriftSpeed:          params.InertialGuidanceDriftSpeed,
			InertialNavigation:                  params.InertialNavigation,
			InertialNavigationDriftSpeed:        params.InertialNavigationDriftSpeed,
		},
		FlightProp: FlightProp{
			MaximumLaunchAngleHorizontalVertical: params.MaximumLaunchAngleHorizontalVertical,
			AimSensitivity:                       params.AimSensitivity,
			MaximumAxisValues:                    params.MaximumAxisValues,
			MaximumFinAngleOfAttack:              params.MaximumFinAngleOfAttack,
			FinsLateralAcceleration:              params.FinsLateralAcceleration,
			MaximumAOA:                           params.MaximumAOA,
			MaximumFinLateralAcceleration:        params.MaximumFinLateralAcceleration,
			WingAreaMultiplier:                   params.WingAreaMultiplier,
			MaximumLateralAcceleration:           params.MaximumLateralAcceleration,
			StartSpeed:                           params.StartSpeed,
			MaximumSpeed:                         params.MaximumSpeed,
			MinimumRange:                         params.MinimumRange,
			MaximumFlightRange:                   params.MaximumFlightRange,
			Tracer:                               params.Tracer,
			LoadFactorLimitAtLaunch:              params.LoadFactorLimitAtLaunch,
			MaximumOverLoad:                      params.MaximumOverLoad,
			SeaSkimming:                          params.SeaSkimming,
			FlightTimeUntilGuidanceStarts:        params.FlightTimeUntilGuidanceStarts,
			FlightTimeWhenPullLimit30:            params.FlightTimeWhenPullLimit30,
			FlightTimeWhenPullLimit40:            params.FlightTimeWhenPullLimit40,
			FlightTimeWhenPullLimit100:           params.FlightTimeWhenPullLimit100,
			Loft:                                 params.Loft,
			LoftAngle:                            params.LoftAngle,
			TargetElevation:                      params.TargetElevation,
			MaximumTargetAngularChange:           params.MaximumTargetAngularChange,
			ThrustVectoring:                      params.ThrustVectoring,
			ThrustVectoringAngle:                 params.ThrustVectoringAngle,
			ETAToImpactWhenPropMultiplierReachesXPercentage30:  params.ETAToImpactWhenPropMultiplierReachesXPercentage30,
			ETAToImpactWhenPropMultiplierReachesXPercentage50:  params.ETAToImpactWhenPropMultiplierReachesXPercentage50,
			ETAToImpactWhenPropMultiplierReachesXPercentage80:  params.ETAToImpactWhenPropMultiplierReachesXPercentage80,
			ETAToImpactWhenPropMultiplierReachesXPercentage90:  params.ETAToImpactWhenPropMultiplierReachesXPercentage90,
			ETAToImpactWhenPropMultiplierReachesXPercentage100: params.ETAToImpactWhenPropMultiplierReachesXPercentage100,
		},
	}
}

func UpdateWeaponParams(params *Params) bson.M {
	return bson.M{
		"Category": params.Category,
		"Name":     params.Name,
		"PhysicalProp": PhysicalProp{
			Mass:                     params.Mass,
			MassAtEndOfBoosterBurn:   params.MassAtEndOfBoosterBurn,
			MassAtEndOfSustainerBurn: params.MassAtEndOfSustainerBurn,
			Calibre:                  params.Calibre,
			Length:                   params.Length,
		},
		"EngineProp": EngineProp{
			ForceExertedByBooster:      params.ForceExertedByBooster,
			BurnTimeOfBooster:          params.BurnTimeOfBooster,
			RawAccelerationAtIgnition:  params.RawAccelerationAtIgnition,
			SpecificImpulseOfBooster:   params.SpecificImpulseOfBooster,
			DeltaSpeedOfBooster:        params.DeltaSpeedOfBooster,
			BoosterStartDelay:          params.BoosterStartDelay,
			ForceExertedBySustainer:    params.ForceExertedBySustainer,
			BurnTimeOfSustainer:        params.BurnTimeOfSustainer,
			SpecificImpulseOfSustainer: params.SpecificImpulseOfSustainer,
			DeltaSpeedOfSustainer:      params.DeltaSpeedOfSustainer,
			TotalDeltaSpeed:            params.TotalDeltaSpeed,
		},
		"FuseAndWarheadProp": FuseAndWarheadProp{
			ExplosiveMass:                params.ExplosiveMass,
			TandemCharge:                 params.TandemCharge,
			Penetration:                  params.Penetration,
			ProximityFuse:                params.ProximityFuse,
			ProximityFuseRange:           params.ProximityFuseRange,
			ProximityFuseArmingDistance:  params.ProximityFuseArmingDistance,
			ProximityFuseShellDetection:  params.ProximityFuseShellDetection,
			ProximityFuseMinimumAltitude: params.ProximityFuseMinimumAltitude,
			ProximityFuseDelay:           params.ProximityFuseDelay,
		},
		"GuidanceProp": GuidanceProp{
			Zoom:                   params.Zoom,
			GuidanceType:           params.GuidanceType,
			GuidanceStartDelay:     params.GuidanceStartDelay,
			GuidanceDuration:       params.GuidanceDuration,
			GuidanceRange:          params.GuidanceRange,
			LaunchSector:           params.LaunchSector,
			ControlConeFOV:         params.ControlConeFOV,
			AimTrackingSensitivity: params.AimTrackingSensitivity,
			MaximumAngleAllowedBetweenMissileAndCrosshair: params.MaximumAngleAllowedBetweenMissileAndCrosshair,
			SeekerWarmUpTime:             params.SeekerWarmUpTime,
			SeekerSearchDuration:         params.SeekerSearchDuration,
			FieldOfView:                  params.FieldOfView,
			OpticSightFieldOfView:        params.OpticSightFieldOfView,
			GimbalLimit:                  params.GimbalLimit,
			TrackRate:                    params.TrackRate,
			UncageSeekerBeforeLaunch:     params.UncageSeekerBeforeLaunch,
			MaximumLockAngleBeforeLaunch: params.MaximumLockAngleBeforeLaunch,
			MinimumAngleBetweenSeekerAndSunForNotCapture: params.MinimumAngleBetweenSeekerAndSunForNotCapture,
			CanLockGround:                             params.CanLockGround,
			LockOnRangeGround:                         params.LockOnRangeGround,
			LockOnRangeVehicle:                        params.LockOnRangeVehicle,
			LockOnRangeFromRearAspect:                 params.LockOnRangeFromRearAspect,
			FlareDetectionRange:                       params.FlareDetectionRange,
			IRCMDetectionRange:                        params.IRCMDetectionRange,
			DIRCMDetectionRange:                       params.DIRCMDetectionRange,
			HeadOnLockOnRangeAgainstAfterburnerTarget: params.HeadOnLockOnRangeAgainstAfterburnerTarget,
			IRCCM:                               params.IRCCM,
			IRCCMType:                           params.IRCCMType,
			IRCCMFieldOfView:                    params.IRCCMFieldOfView,
			IRCCMRejectionThreshold:             params.IRCCMRejectionThreshold,
			IRCCMReactionTime:                   params.IRCCMReactionTime,
			LockOnRangeFromAllAspect:            params.LockOnRangeFromAllAspect,
			MaximumBreakLockTime:                params.MaximumBreakLockTime,
			CanBeSlavedToRadar:                  params.CanBeSlavedToRadar,
			CanLockAfterLaunch:                  params.CanLockAfterLaunch,
			Band:                                params.Band,
			AngularSpeedRejectionThreshold:      params.AngularSpeedRejectionThreshold,
			AccelerationRejectionThresholdRange: params.AccelerationRejectionThresholdRange,
			SidelobeAttenuation:                 params.SidelobeAttenuation,
			TransmitterPower:                    params.TransmitterPower,
			TransmitterAngleOfHalfSensitivity:   params.TransmitterAngleOfHalfSensitivity,
			TransmitterSidelobeSensitivity:      params.TransmitterSidelobeSensitivity,
			ReceiverAngleOfHalfSensitivity:      params.ReceiverAngleOfHalfSensitivity,
			ReceiverSidelobeSensitivity:         params.ReceiverSidelobeSensitivity,
			DistanceMinimumValue:                params.DistanceMinimumValue,
			DistanceMaximumValue:                params.DistanceMaximumValue,
			DistanceWidth:                       params.DistanceWidth,
			DistanceMinimumSignalGate:           params.DistanceMinimumSignalGate,
			DistanceRefWidth:                    params.DistanceRefWidth,
			DopplerSpeedMinimumValue:            params.DopplerSpeedMinimumValue,
			DopplerSpeedMaximumValue:            params.DopplerSpeedMaximumValue,
			DopplerSpeedWidth:                   params.DopplerSpeedWidth,
			DopplerSpeedRefWidth:                params.DopplerSpeedRefWidth,
			DopplerSpeedMinimumSignalGate:       params.DopplerSpeedMinimumSignalGate,
			DopplerSpeedGateSearchRange:         params.DopplerSpeedGateSearchRange,
			DopplerSpeedGateAlphaFilter:         params.DopplerSpeedGateAlphaFilter,
			DopplerSpeedGateBetaFilter:          params.DopplerSpeedGateBetaFilter,
			ProportionalNavigationMultiplier:    params.ProportionalNavigationMultiplier,
			BaseIndicatedAirSpeed:               params.BaseIndicatedAirSpeed,
			PIDProportionalTerm:                 params.PIDProportionalTerm,
			PIDIntegralTerm:                     params.PIDIntegralTerm,
			PIDIntegralTermLimit:                params.PIDIntegralTermLimit,
			PIDDerivativeTerm:                   params.PIDDerivativeTerm,
			InertialGuidanceDriftSpeed:          params.InertialGuidanceDriftSpeed,
			InertialNavigation:                  params.InertialNavigation,
			InertialNavigationDriftSpeed:        params.InertialNavigationDriftSpeed,
		},
		"FlightProp": FlightProp{
			MaximumLaunchAngleHorizontalVertical: params.MaximumLaunchAngleHorizontalVertical,
			AimSensitivity:                       params.AimSensitivity,
			MaximumAxisValues:                    params.MaximumAxisValues,
			MaximumFinAngleOfAttack:              params.MaximumFinAngleOfAttack,
			FinsLateralAcceleration:              params.FinsLateralAcceleration,
			MaximumAOA:                           params.MaximumAOA,
			MaximumFinLateralAcceleration:        params.MaximumFinLateralAcceleration,
			WingAreaMultiplier:                   params.WingAreaMultiplier,
			MaximumLateralAcceleration:           params.MaximumLateralAcceleration,
			StartSpeed:                           params.StartSpeed,
			MaximumSpeed:                         params.MaximumSpeed,
			MinimumRange:                         params.MinimumRange,
			MaximumFlightRange:                   params.MaximumFlightRange,
			Tracer:                               params.Tracer,
			LoadFactorLimitAtLaunch:              params.LoadFactorLimitAtLaunch,
			MaximumOverLoad:                      params.MaximumOverLoad,
			SeaSkimming:                          params.SeaSkimming,
			FlightTimeUntilGuidanceStarts:        params.FlightTimeUntilGuidanceStarts,
			FlightTimeWhenPullLimit30:            params.FlightTimeWhenPullLimit30,
			FlightTimeWhenPullLimit40:            params.FlightTimeWhenPullLimit40,
			FlightTimeWhenPullLimit100:           params.FlightTimeWhenPullLimit100,
			Loft:                                 params.Loft,
			LoftAngle:                            params.LoftAngle,
			TargetElevation:                      params.TargetElevation,
			MaximumTargetAngularChange:           params.MaximumTargetAngularChange,
			ThrustVectoring:                      params.ThrustVectoring,
			ThrustVectoringAngle:                 params.ThrustVectoringAngle,
			ETAToImpactWhenPropMultiplierReachesXPercentage30:  params.ETAToImpactWhenPropMultiplierReachesXPercentage30,
			ETAToImpactWhenPropMultiplierReachesXPercentage50:  params.ETAToImpactWhenPropMultiplierReachesXPercentage50,
			ETAToImpactWhenPropMultiplierReachesXPercentage80:  params.ETAToImpactWhenPropMultiplierReachesXPercentage80,
			ETAToImpactWhenPropMultiplierReachesXPercentage90:  params.ETAToImpactWhenPropMultiplierReachesXPercentage90,
			ETAToImpactWhenPropMultiplierReachesXPercentage100: params.ETAToImpactWhenPropMultiplierReachesXPercentage100,
		},
	}
}
