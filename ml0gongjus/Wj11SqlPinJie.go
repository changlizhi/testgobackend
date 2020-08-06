package ml0gongjus

import(
  "xm1shengcheng/ml1changliangs"
  "xm1shengcheng/ml2moxings"
)
//基础的四个sql，增删改查，sql结构定死之后传入不同的部分放入结构中的不同位置，尽量让所有的业务都组装成这四个基本结构的方式
//查询: SELECT $lies FROM $biaos（里面可以是leftjion的东西或者字查询） WHERE $tiaoJians = $tiaoJianZhis(且关系用小括号包起来传入) $juhes
//增加: INSERT INTO $biaos[1]($lies) values ($zhis)
//删除: DELETE FROM $biaos[1] WHERE $tiaoJians=$tiaoJianZhis
//修改: UPDATE $biaos[1] SET $lies=$zhis(多个用逗号分割) WHERE $tiaoJians=$tiaoJianZhis

//所以传入这个文件的参数至少需要一个这样的结构体：
//1. 类型：说明这次是增删改查中的哪种
//2. Lies
//3. Biaos
//4. Zhis
//5. TiaoJians
//6. TiaoJianZhis
//7. JuHes
//其中lies和zhis在增加和修改的时候必须成对出现，且顺序必须一致，而tiaojian和tiaojianzhi也必须成对出现并且必须顺序一致
//尽量一个sql搞定一个交易，所有的交易必然至少有一个sql，所以这里拼装的sql要根据交易区分开，根据传入的参数判定是什么基础类型的sql

func PinJieSql(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}
func PinJieSelectSql(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}
func PinJieInsertSql(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}
func PinJieUpdateSql(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}
func PinJieDeleteSql(canShu ml2moxings.CanShu)ml2moxings.FanHui{
  ret := ml2moxings.FanHui{
    ZhuangTai:ml1changliangs.CHENGGONG,
    BianMa:ml1changliangs.HFX000000000,
    BiaoZhi:ml1changliangs.LieBiaosDuiXiangs,
  }
  return ret
}


