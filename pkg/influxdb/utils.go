package influxdb

func checkEmpty(input string) string {
	if input == "" {
		return "\"\""
	}

	return input
}
