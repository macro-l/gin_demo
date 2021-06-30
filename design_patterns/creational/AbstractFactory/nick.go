package main

// 具体产品 - 耐克工厂类
type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nick",
			size: 14,
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nick",
			size: 14,
		},
	}
}
