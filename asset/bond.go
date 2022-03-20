package asset

import "math"

type Bonds struct{
	t int
	Seg_No int
	Redemp_Y uint32
	Redemp_M uint32
	Coupon_Pc float64
	Coupon_Freq int
	Amort_Pc float64
	Face_Amt float64
}

func (bond *Bonds) Abv() float64 {
	mat := int((bond.Redemp_Y - Val_Y )* 12 + (bond.Redemp_M - Val_M)) - bond.t
	//Debug Print
	//fmt.Println(mat)
	var result float64
	for i:= 1; i<= mat ; i++ {
		if i < mat {
			if i % bond.Coupon_Freq == 0 {
				result += (bond.Coupon_Pc/100 * bond.Face_Amt) /math.Pow((1+bond.Amort_Pc/100), float64(i) / 12)
			}

		} else {
			result += (1 + bond.Coupon_Pc/100) * bond.Face_Amt /math.Pow((1+bond.Amort_Pc/100), float64(i) / 12)
		}
		//fmt.Println(result,math.Pow((1+bond.Amort_Pc/100), float64(i/12)), bond.Amort_Pc, i, float64(i)/12)
	}
	return result
}