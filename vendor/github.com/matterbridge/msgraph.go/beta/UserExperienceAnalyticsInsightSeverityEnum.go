// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserExperienceAnalyticsInsightSeverity undocumented
type UserExperienceAnalyticsInsightSeverity int

const (
	// UserExperienceAnalyticsInsightSeverityVNone undocumented
	UserExperienceAnalyticsInsightSeverityVNone UserExperienceAnalyticsInsightSeverity = 0
	// UserExperienceAnalyticsInsightSeverityVInformational undocumented
	UserExperienceAnalyticsInsightSeverityVInformational UserExperienceAnalyticsInsightSeverity = 1
	// UserExperienceAnalyticsInsightSeverityVWarning undocumented
	UserExperienceAnalyticsInsightSeverityVWarning UserExperienceAnalyticsInsightSeverity = 2
	// UserExperienceAnalyticsInsightSeverityVError undocumented
	UserExperienceAnalyticsInsightSeverityVError UserExperienceAnalyticsInsightSeverity = 3
)

// UserExperienceAnalyticsInsightSeverityPNone returns a pointer to UserExperienceAnalyticsInsightSeverityVNone
func UserExperienceAnalyticsInsightSeverityPNone() *UserExperienceAnalyticsInsightSeverity {
	v := UserExperienceAnalyticsInsightSeverityVNone
	return &v
}

// UserExperienceAnalyticsInsightSeverityPInformational returns a pointer to UserExperienceAnalyticsInsightSeverityVInformational
func UserExperienceAnalyticsInsightSeverityPInformational() *UserExperienceAnalyticsInsightSeverity {
	v := UserExperienceAnalyticsInsightSeverityVInformational
	return &v
}

// UserExperienceAnalyticsInsightSeverityPWarning returns a pointer to UserExperienceAnalyticsInsightSeverityVWarning
func UserExperienceAnalyticsInsightSeverityPWarning() *UserExperienceAnalyticsInsightSeverity {
	v := UserExperienceAnalyticsInsightSeverityVWarning
	return &v
}

// UserExperienceAnalyticsInsightSeverityPError returns a pointer to UserExperienceAnalyticsInsightSeverityVError
func UserExperienceAnalyticsInsightSeverityPError() *UserExperienceAnalyticsInsightSeverity {
	v := UserExperienceAnalyticsInsightSeverityVError
	return &v
}