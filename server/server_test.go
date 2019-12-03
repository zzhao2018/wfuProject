package server

import "testing"

func TestAddUserMidWare(t *testing.T) {
	ip,err:=getLocalIp()
	if err!=nil {
		t.Fatalf("error:%+v\n",err)
		return
	}
	t.Logf("data::%+v\n",ip)
}