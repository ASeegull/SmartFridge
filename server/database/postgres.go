package database

//ClientValidation checks login and pass for user
func ClientValidation(login string, pass string) error {

	return nil
}

//RegisterNewClient registers new client
func RegisterNewClient(clientID string) error {

	return nil
}

//ClientLogin checks login and pass for client
func ClientLogin(login string, pass string) error {

	return nil
}

//GetDefaultExpirationPeriod returns default period
func GetDefaultExpirationPeriod(productID int) (int, error) {

	days := 10
	return days, nil
}

//CheckAgent checks agent registration
func CheckAgent(agentID string) bool {
	return true
}

//RegisterNewAgent registers new agent
func RegisterNewAgent(clientID string, agentID string) error {

	return nil
}

//GetAllAgentsIDForClient returns all agent for clientID
func GetAllAgentsIDForClient(clientID string) ([]string, error) {

	return []string{}, nil
}
