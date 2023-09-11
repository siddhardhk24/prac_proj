package interfaces

import "prac_proj/model"

type IV_int interface{
	Create__IV(*model.IV_model)(string,error)
}