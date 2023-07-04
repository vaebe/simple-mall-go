package address

import (
	"github.com/gin-gonic/gin"
	"simple-mall/services/addressServices"
	"simple-mall/utils"
)

// GetAreasByParams
//
//	@Summary		根据参数获取区域数据
//	@Description	根据参数获取区域数据
//	@Tags				address地址管理
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	false	"区域id"
//	@Param			pid	query		int	false	"上级id"
//	@Param			deep	query		int	false	"层级"
//	@Success		200		{object}	utils.ResponseResultInfo{data=[]address.AreaInfo}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/address/getAreasByParams [get]
func GetAreasByParams(ctx *gin.Context) {
	id := ctx.Query("id")
	pid := ctx.Query("pid")
	deep := ctx.Query("deep")

	areas, err := addressServices.GetAreasByParams(id, pid, deep)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, areas)
}
