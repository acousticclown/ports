package ports

func IsPortRunning(target string) (bool, *PortInfo, error) {
	ports, err := ListPorts()
	if err != nil {
		return false, nil, err
	}

	for _, p := range ports {
		if p.Port == target {
			return true, &p, nil
		}
	}
	return false, nil, nil
}
