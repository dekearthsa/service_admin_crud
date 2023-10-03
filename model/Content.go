package model

type Content struct {
	ContentId       string
	CategoryContent string
	Title           string
	Content         string
	ArrayImg        []ArrayImgPath
	CreateDate      int
	UpdateDate      int
}
