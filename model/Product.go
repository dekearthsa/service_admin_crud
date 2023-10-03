package model

type Product struct {
	ProductId   string
	Category    string
	ProductName string
	ArrayImg    []ArrayImgPath
	Price       float64
	CreateDate  int
	UpdateDate  int
}

type ArrayImgPath struct {
	ImgPath string
}
