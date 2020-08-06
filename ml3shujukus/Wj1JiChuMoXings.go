package ml3shujukus

//由于这个模型会调用Ku中的读取数据最大值情况，而且所有的表其实不过5个，所以放到这里也算正常的
//moxing中就放业务模型就行了。

type JieGou struct{
  BianMa string `gorm:"column:BianMa;not null;type:text;default:'hfxmoren'"`
  MingCheng string `gorm:"column:MingCheng;not null;type:text;default:'hfxmoren'"`
  BsjBiao string `gorm:"column:BsjBiao;not null;type:text;default:'hfxmoren'"`//有了这个都不用再要计数表了，直接拆分表名
  ZdsjBiao string `gorm:"column:ZdsjBiao;not null;type:text;default:'hfxmoren'"`
  ZhengZe string `gorm:"column:ZhengZe;not null;type:text;default:'hfxmoren'"`
  ShiJian  string `gorm:"column:ShiJian;not null;type:text;default:'hfxmoren'"`
}

func (obj JieGou) TableName() string {
  return "JieGou"
}

type ZongLing struct{
  BiaoMing string `gorm:"column:BiaoMing;not null;type:text;default:'hfxmoren'"`
  ZiDuan string `gorm:"column:ZiDuan;not null;type:text;default:'hfxmoren'"`
}
func (obj ZongLing) TableName() string {
  return "ZongLing"
}

type RiZhi struct{
  JiaoYiMa string `gorm:"column:JiaoYiMa;not null;type:text;default:'hfxmoren'"`
  CanShu string `gorm:"column:CanShu;not null;type:text;default:'hfxmoren'"`
  ShiJian string `gorm:"column:ShiJian;not null;type:text;default:'hfxmoren'"`
}
func (obj RiZhi) TableName() string {
  return "RiZhi"
}
//在这里进行表数据和字段数据段表创建，调用库业务进行创建，生成表名的方式也是在这里调用库进行查询和加一后返回，得到真实的数据表进行创建，然后给出一个表名，然后进行增删改查的操作。
type Bsjs struct{
  ZhuJian string `gorm:"column:ZhuJian;not null;type:text;default:'hfxmoren'"`
}
func (obj Bsjs) TableName() string {
  // 读取最大的表计数，加一然后根据传入的字段如果不存在这个表就要新增
  return "Bsjs"
}

type Zdsjs struct{
  BiaoZhuJian string `gorm:"column:BiaoZhuJian;not null;type:text;default:'hfxmoren'"`
  Zhi string `gorm:"column:Zhi;not null;type:text;default:'hfxmoren'"`
}
func (obj Zdsjs) TableName() string {
  //读取最大的ZongLing表中的最大数，然后加一之后新增这个表，必须是有表和字段的时候才会增加这个表和ZongLing数据
  return "Zdsjs"
}

