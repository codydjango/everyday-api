package main

var addressNameLookup = make(map[string]string)

// GetAccountData is account data
func GetAccountData(account string) string {
	return addressNameLookup[account]
}

// SetAccountData is account data
func SetAccountData(account, data string) bool {
	addressNameLookup[account] = data
	return true
}
