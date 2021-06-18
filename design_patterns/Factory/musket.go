package main

type musket struct {
	gun
}

// 具体枪支 - musket
func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
