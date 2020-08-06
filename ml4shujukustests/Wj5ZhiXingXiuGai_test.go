package ml4shujukustests
import(
  "testing"
  "xm1shengcheng/ml3shujukus"
  "log"
)

func TestZhiXingXiuGai(t *testing.T){
  sql:=""
  values:=[]string{
    "",
  }
  ret := ml3shujukus.ZhiXingXiuGai(sql,values)
  log.Println("Wj3ZhiXingXiuGai_test.go---TestZhiXingXiuGai,ret",ret)
}