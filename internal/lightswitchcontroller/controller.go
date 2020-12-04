package lightswitchcontroller

func start_lightswitch_controller() {
	s := edisonLightSwitch{
		lightBulb: edisonLightbulb{},
	}

	for {
		if s.IsOn() {
			s.lightBulb.TurnOn()
		} else {
			s.lightBulb.TurnOff()
		}
	}
}
