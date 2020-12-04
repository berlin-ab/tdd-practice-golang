package lightswitchcontroller

type edisonLightSwitch struct {
	lightBulb edisonLightbulb
}

func (s edisonLightSwitch) IsOn() bool {
	panic("edison light switch sdk requires license")
}
