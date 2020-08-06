package ml4shujukustests
import(
  "testing"
  "xm1shengcheng/ml3shujukus"
  "log"
)

func TestZhiXingShanChu(t *testing.T){
  sql:=""
  values:= []string{
    "",
  }
  ret := ml3shujukus.ZhiXingShanChu(sql,values)
  log.Println("Wj3ZhiXingShanChu_test.go---TestZhiXingShanChu,ret",ret)
}