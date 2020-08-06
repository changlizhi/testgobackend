package ml3shujukus

import (
	"log"
  "strings"
	"database/sql"
	"xm1shengcheng/ml1changliangs"
	"xm1shengcheng/ml2moxings"
  "xm1shengcheng/ml00peizhis"
)
func ZhiXingChaXun(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  canShuTemp := &canShu
  canShuChaXun := *canShuTemp
  canShuChaXun.SHUJU[ml1changliangs.LeiBie] = ml1changliangs.ChaXunBiao
  biaosGuanLian := ml00peizhis.YeWuGuanLian(canShuChaXun)
  biaos := biaosGuanLian.ShuJuLieBiaos
  
  canShuChaXun.SHUJU[ml1changliangs.LeiBie] = ml1changliangs.ChaXunZiDuans
  liesGuanLian := ml00peizhis.YeWuGuanLian(canShuChaXun)
  lies := liesGuanLian.ShuJuLieBiaos
  
  sql := "SELECT " + strings.Join(lies, ",") + " FROM " + biaos[ml1changliangs.Sz0]
	log.Println("Wj1JiChuKu.go---ZhiXingSql,sql",sql)
	rows, err := LianJieChi().Raw(sql).Rows()
	ret := ml2moxings.FanHui{}
  shuJu := []map[string]string{}
	if err != nil || rows == nil {
		log.Println("Wj1JiChuKu.go---ZhiXingSql,sql,err,rows",sql,err,rows)
		return ret
	}
	shuJu = scanRet(lies,rows)
  ret.ShuJuDuiXiangs=shuJu
	return ret
}
func scanRet(lies []string,rows *sql.Rows)[]map[string]string{
	ret := []map[string]string{}
	for rows.Next() {//每一行
		tempLie := make([]string,len(lies))
    tempLieRef := make([]interface{},len(lies))

    for i := ml1changliangs.Sz0;i < len(lies); i++{
      tempLieRef[i] = &tempLie[i]
    }
    rows.Scan(
      tempLieRef...,
    )
		retOne := map[string]string{}
		for xiaBiao,lie := range lies{
      if tempLieRef[xiaBiao] !=  nil{
        retOne[lie]=*tempLieRef[xiaBiao].(*string)
      }
		}
		ret = append(ret,retOne)
	}
	return ret
}
