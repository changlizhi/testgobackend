package ml0gongjustests

import(
  "log"
  "testing"
  "xm1shengcheng/ml0gongjus"
)
func TestBase64(t *testing.T){
  log.Println(ml0gongjus.Base64BianMa("abcdefghijklmn"))
}
func TestBase64JieMa(t *testing.T){
  log.Println(ml0gongjus.Base64JieMa("YWJjZGVmZ2hpamtsbW4="))
}
func TestBase64UrlBianMa(t *testing.T){
  log.Println(ml0gongjus.Base64UrlBianMa("http://www.hanfuxin.com?abc=123"))
}
func TestBase64UrlJiema(t *testing.T){
  log.Println(ml0gongjus.Base64UrlJiema("aHR0cDovL3d3dy5oYW5mdXhpbi5jb20_YWJjPTEyMw=="))
}
