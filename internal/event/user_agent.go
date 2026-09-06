package event

import "github.com/mileusna/useragent"

type UserAgentInfo struct {
	DeviceType string
	OS         string
	Browser    string
}

func ParseUserAgent(raw string) UserAgentInfo {
	ua := useragent.Parse(raw)

	deviceType := "desktop"

	if ua.Mobile {
		deviceType = "mobile"
	} else if ua.Tablet {
		deviceType = "tablet"
	}

	return UserAgentInfo{
		DeviceType: deviceType,
		OS:         ua.OS,
		Browser:    ua.Name,
	}
}