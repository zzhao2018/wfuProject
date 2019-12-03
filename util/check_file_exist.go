package util

import "os"

//判断文件是否存在
func CheckFileExist(filepathS string)bool{
	_,err:=os.Stat(filepathS)
	if err!=nil {
		if os.IsExist(err)==true {
			return true
		}
		return false
	}
	return true
}
