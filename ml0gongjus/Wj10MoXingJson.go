package ml0gongjus

import(
  "os"
  "io/ioutil"
  "log"
)

func duQuWenJian(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
  log.Println("f.Name()",f.Name())
	return ioutil.ReadAll(f)
}
func DuQuMoXingJson()map[string][]string{
  var ret map[string][]string
  byteJson, err:= duQuWenJian("../JieGou.json")
  if err != nil{
    log.Println("Wj10MoXingJson.go---DuQuMoXingJson(),err1",err)
    return ret
  }
  log.Println("Wj10MoXingJson.go---DuQuMoXingJson(),string(byteJson)",string(byteJson))
  err = JsonGongJu().Unmarshal(byteJson,&ret)
  log.Println("DuQuMoXingJson(),err1",err)
  return ret
}
