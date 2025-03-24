package access

import "time"

// Controller - controller attributes
type Controller struct {
	RateLimit string
	RateBurst string
	Redirect  string // redirect percentage
	Timeout   time.Duration
}

func UpdateDefaults(c *Controller) {
	if c.RateLimit == "" {
		c.RateLimit = "-1"
	}
	if c.RateBurst == "" {
		c.RateBurst = "-1"
	}
	if c.Redirect == "" {
		c.Redirect = "-1"
	}

}

// NilController - used when Controller is not applicable
func NilController() Controller {
	return Controller{RateLimit: "-1", RateBurst: "-1", Redirect: "-1", Timeout: -1}
}
