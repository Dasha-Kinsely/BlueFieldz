package validators

import "github.com/Dasha-Kinsely/leaveswears/models"

type ProfileUpdateValidator struct {
	Profile struct {
		PreferredColor []string `form:"color" json:"color" binding:"alphanum"`
		PreferredMaterial []string `form:"material" json:"material" binding:"alphanum"`
		ArmLength string `form:"armlength" json:"armlength" binding:"max=10"`
		TorsoLength string `form:"torsolength" json:"torsolength" binding:"max=10"`
		ShoulderSize string `form:"shouldersize" json:"shouldersize" binding:"max=10"`
		WaistSize string `form:"waistsize" json:"waistsize" binding:"max=10"`
		Allergy string `form:"allergy" json:"allergy" binding:"alphanum, max=255"`
	} `json:"profile"`
	ValidatedNewProfile models.Profile `json:"-"`
}

