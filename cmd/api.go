package main

import "encoding/hex"

// ===== Account ===== //
func GetAccount(address string) (map[string]interface{}, error) {

	payload := map[string]any{
		"address": address,
		"visible": true,
	}

	return post("/wallet/getaccount", payload)
}

func GetAccountBalance(address string) (int, error) {

	result, err := GetAccount(address)
	if err != nil {
		return 0, nil
	}

	balance, ok := result["balance"]
	if !ok {
		return 0, nil
	}

	return convertToInt(balance), nil
}

// ===== Account Resource ===== //

func GetAccountResource(address string) (int, int, int, error) {

	payload := map[string]any{
		"address": address,
		"visible": true,
	}

	result, err := post("/wallet/getaccountresource", payload)
	if err != nil {
		return 0, 0, 0, err
	}

	totalEneryLimit, ok := result["TotalEnergyLimit"]
	if !ok {
		return 0, 0, 0, err
	}

	totalNetLimit, ok := result["TotalNetLimit"]
	if !ok {
		return 0, 0, 0, err
	}

	freeNetLimit, ok := result["freeNetLimit"]
	if !ok {
		return 0, 0, 0, err
	}

	return convertToInt(totalEneryLimit), convertToInt(totalNetLimit), convertToInt(freeNetLimit), nil
}

func GetAccountNet(address string) (int, int, error) {

	payload := map[string]any{
		"address": address,
		"visible": true,
	}

	result, err := post("/wallet/getaccountnet", payload)
	if err != nil {
		return 0, 0, err
	}

	totalNetLimit, ok := result["TotalNetLimit"]
	if !ok {
		return 0, 0, err
	}

	freeNetLimit, ok := result["freeNetLimit"]
	if !ok {
		return 0, 0, err
	}

	return convertToInt(totalNetLimit), convertToInt(freeNetLimit), nil
}

// ===== Network ===== //

func GetNodeInfo() (map[string]any, error) {

	return get("/wallet/getnodeinfo")
}

// ===== TRC10 ===== //

func GetAccountTRC10(address string) (map[string]any, error) {

	payload := map[string]any{
		"address": address,
		"visible": true,
	}

	result, err := post("/wallet/getassetissuebyaccount", payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateTRC10Token(address string, name string) (map[string]any, error) {

	tokenName := hex.EncodeToString([]byte(name))

	payload := map[string]any{
		"owner_address": address,
		"name":          tokenName,
		"abbr":          abbr,
	}

	result, err := post("/wallet/getassetissuebyaccount", payload)
	if err != nil {
		return nil, err
	}

	return result, nil
}
