package access

import "time"

// Controller - controller attributes
type Controller struct {
	Timeout    time.Duration
	RateLimit  string
	RateBurst  string
	Percentage string
	Code       string
}

func UpdateDefaults(c *Controller) {
	if c.RateLimit == "" {
		c.RateLimit = "-1"
	}
	if c.RateBurst == "" {
		c.RateBurst = "-1"
	}
	if c.Percentage == "" {
		c.Percentage = "0"
	}

}

// NilController - used when Controller is not applicable
func NilController() Controller {
	return Controller{RateLimit: "-1", RateBurst: "-1", Percentage: "0"}
}
