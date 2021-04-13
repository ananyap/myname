package main

import (
	"database/sql"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/myname/managers"
	"github.com/myname/names"
	"net/http"
)

func main() {
	db, err := DBConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := gin.Default()
	router.HTMLRender = ginview.Default()
	router.Static("/views/static", "./views/static")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/namesatsha/:myname", func(ctx *gin.Context) {
		myName := ctx.Param("myname")
		nameManager := managers.NewNameManager(myName)

		nameManager.SetReangName()
		nameReang := nameManager.GetReangName()

		nameManager.SetLeksat()
		leksat := nameManager.GetLeksat()

		nameManager.SetShadow()
		shadow := nameManager.GetShadow()

		realName := names.RealName{
			ThaiName:   nameManager.GetThaiName(),
			ReangThai:  nameReang,
			LeksatThai: leksat,
			Shadow:     shadow,
		}

		dbManager := managers.NewDbManager(db)
		dbManager.SetRealNames()
		dbManager.SetNumbersMiracle()

		pairManager := managers.NewNamePairManager(realName, dbManager.GetRealNames(), dbManager.GetNumberMiracle())
		pairManager.SetListSatDShaD()

		listSatDShaD := pairManager.GetListSatDShaD()

		pageManager := managers.NewNamePageManager(len(listSatDShaD), 10)

		pageManager.SetTotalPage()

		ctx.JSON(200, gin.H{
			"myName":    realName,
			"nameDPlus": listSatDShaD,
		})
	})

	router.GET("/namesat/:myname", func(ctx *gin.Context) {
		myName := ctx.Param("myname")
		nameManager := managers.NewNameManager(myName)
		nameManager.SetReangName()
		nameReang := nameManager.GetReangName()
		nameManager.SetLeksat()
		leksat := nameManager.GetLeksat()
		nameManager.SetShadow()
		shadow := nameManager.GetShadow()

		realName := names.RealName{
			ThaiName:   nameManager.GetThaiName(),
			ReangThai:  nameReang,
			LeksatThai: leksat,
			Shadow:     shadow,
		}

		dbManager := managers.NewDbManager(db)
		dbManager.SetRealNames()
		dbManager.SetNumbersMiracle()

		pairManager := managers.NewNamePairManager(realName, dbManager.GetRealNames(), dbManager.GetNumberMiracle())
		pairManager.SetPlusTitleNumD()

		ctx.JSON(200, gin.H{
			"myName":    realName,
			"nameDPlus": pairManager.GetPlusTitleNumD(),
		})
	})

	router.GET("/namesha/:myname", func(ctx *gin.Context) {
		myName := ctx.Param("myname")
		nameManager := managers.NewNameManager(myName)
		nameManager.SetReangName()
		nameReang := nameManager.GetReangName()
		nameManager.SetLeksat()
		leksat := nameManager.GetLeksat()
		nameManager.SetShadow()
		shadow := nameManager.GetShadow()

		realName := names.RealName{
			ThaiName:   nameManager.GetThaiName(),
			ReangThai:  nameReang,
			LeksatThai: leksat,
			Shadow:     shadow,
		}

		dbManager := managers.NewDbManager(db)
		dbManager.SetRealNames()
		dbManager.SetNumbersMiracle()

		pairManager := managers.NewNamePairManager(realName, dbManager.GetRealNames(), dbManager.GetNumberMiracle())
		pairManager.SetPlusShaNumD()

		ctx.JSON(200, gin.H{
			"myName":    realName,
			"nameDPlus": pairManager.GetPlusShaNumD(),
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	_ = router.Run(":8080")

	//http.ListenAndServe(":8080", nil)
}

func DBConnect() (*sql.DB, error) {
	//db, err := sql.Open("mysql", "ananya:In^telliP24.pa@tcp(localhost:3306)/ananyadb")
	db, err := sql.Open("mysql", "root:@IntelliP24@tcp(localhost:3306)/ananyadb")
	return db, err
}
