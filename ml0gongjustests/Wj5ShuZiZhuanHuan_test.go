package ml0gongjustests
import(
  "log"
  "testing"
  "xm1shengcheng/ml0gongjus"
)
func TestString2Float64(t *testing.T){
  ret:=ml0gongjus.String2Float64("1.111")
  log.Println(ret)
}
func TestInt642String(t *testing.T){
  ret:=ml0gongjus.Int642String(1111111111111111111)
  log.Println(ret)
}
func TestString2Int64(t *testing.T){
  ret:=ml0gongjus.String2Int64("2222222222222222222")
  log.Println(ret)
}
func TestString2Int(t *testing.T){
  ret:=ml0gongjus.String2Int("33333333")
  log.Println(ret)
}
func TestFloat642String(t *testing.T){
  ret:=ml0gongjus.Float642String(1.111111111111111111111111)
  log.Println(ret)
}
