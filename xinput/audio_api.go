package xinput

func GetAudioDeviceIds(controllerIndex ControllerIndex) (outputDeviceId, inputDeviceId []byte, err error) {
	return getAudioDeviceIds(controllerIndex)
}
