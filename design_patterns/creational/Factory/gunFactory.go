package main

import "fmt"

// 工厂生产
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}