package main

type ak47 struct {
	gun
}

// 具体枪支 - ak47
func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "Ak47 gun",
			power: 4,
		},
	}
}
