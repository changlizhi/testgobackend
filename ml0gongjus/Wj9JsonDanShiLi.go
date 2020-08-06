package ml0gongjus

//这里想把json工具做成单实例的，不想在每个方法中都声明一次

import (
  "github.com/json-iterator/go"   // 引入
)

func JsonGongJu() jsoniter.API{
  ret := jsoniter.ConfigCompatibleWithStandardLibrary
  return ret
}
