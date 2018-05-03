package routers

import (
	"class-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	AdminRun()
	beego.Router("/info/teacher/index", &controllers.InfoController{}, "*:GetTeaList")
	beego.Router("/info/teacher/saveTea", &controllers.InfoController{}, "*:SaveTea")
	beego.Router("/info/teacher/delTea", &controllers.InfoController{}, "*:DelTea")
	beego.Router("/info/student/index", &controllers.InfoController{}, "*:GetStuList")
	beego.Router("/info/student/saveStu", &controllers.InfoController{}, "*:SaveStu")
	beego.Router("/info/student/delStu", &controllers.InfoController{}, "*:DelStu")
	beego.Router("/info/lesson/index", &controllers.InfoController{}, "*:GetLesList")
	beego.Router("/info/lesson/saveLes", &controllers.InfoController{}, "*:SaveLes")
	beego.Router("/info/lesson/delLes", &controllers.InfoController{}, "*:DelLes")
	beego.Router("/info/lecture/index", &controllers.InfoController{}, "*:GetLecList")
	beego.Router("/info/lecture/saveLec", &controllers.InfoController{}, "*:SaveLec")
	beego.Router("/info/lecture/delLec", &controllers.InfoController{}, "*:DelLec")
	beego.Router("/info/class/index", &controllers.InfoController{}, "*:GetClassList")
	beego.Router("/info/class/save", &controllers.InfoController{}, "*:SaveClass")
	beego.Router("/info/class/del", &controllers.InfoController{}, "*:DelClass")
	beego.Router("/info/lecture/lectureByTea", &controllers.InfoController{}, "*:GetLecByTea")

	beego.Router("/video/index", &controllers.VideoController{}, "*:GetVideoList")
	beego.Router("/video/saveVideo", &controllers.VideoController{}, "*:SaveVideo")
	beego.Router("/video/delVideo", &controllers.VideoController{}, "*:DelVideo")

	beego.Router("/notice/index", &controllers.NoticeController{}, "*:GetNoticeList")
	beego.Router("/notice/save", &controllers.NoticeController{}, "*:Save")
	beego.Router("/notice/del", &controllers.NoticeController{}, "*:Del")

	beego.Router("/grade/index", &controllers.GradeController{}, "*:GetClassList")

	beego.SetStaticPath("/file", "file")
}
