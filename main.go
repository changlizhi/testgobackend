package main

import (
  "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
  "xm1shengcheng/ml2moxings"
  "xm1shengcheng/ml7yewus"
)
var app *iris.Application
func init(){
  app = iris.New()
  app.Use(Cors)
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
  common := app.Party("/")
  {
    common.Options("*", func(ctx iris.Context) {
      ctx.Next()
    })
  }
}

func HanFuXin(ctx iris.Context){
  canShu := ml2moxings.CanShu{}
  if err := ctx.ReadJSON(&canShu); err == nil {
    // 正确读取到参数之后直接丢给后面的操作，
    fanHui := ml7yewus.YeWuFenPei(canShu)
    ctx.JSON(fanHui)
  }else{
    app.Logger().Println("json数据错误",err)
    fanHui := ml2moxings.FanHui{
      BianMa:"ERR000000001",
      ZhuangTai:"01",
      MiaoShu:"参数格式错误",
      ShuJu:canShu.SHUJU,
    }
    ctx.JSON(fanHui)
  }

}
func Cors(ctx iris.Context) {
  ctx.Header("Access-Control-Allow-Origin", "*")
  if ctx.Request().Method == "OPTIONS" {
    ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
    ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
    ctx.StatusCode(204)
    return
  }
  ctx.Next()
}
func main() {
  app.Post("/hanfuxin",HanFuXin)
  app.Run(iris.Addr(":8080"), )
}
