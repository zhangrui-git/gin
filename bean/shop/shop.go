package shop

type Pk struct {
	Val uint `uri:"id" form:"id" binding:"required,number,gt=0"`
}

type PkSet struct {
	Val []uint `json:"ids" form:"ids" binding:"required"`
}
