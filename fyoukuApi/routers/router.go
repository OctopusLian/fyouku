/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-23 12:43:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-23 19:24:40
 */
// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fyoukuApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.UserController{})    //用户
	beego.Include(&controllers.VideoController{})   //视频
	beego.Include(&controllers.BaseController{})    //基础
	beego.Include(&controllers.CommentController{}) //评论
	beego.Include(&controllers.TopController{})     //排序
	beego.Include(&controllers.BarrageController{})
	beego.Include(&controllers.AliyunController{})
}
