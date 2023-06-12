package role

import (
	"github.com/gin-gonic/gin"
	"simple-mall/services/roleServices"
	"simple-mall/utils"
)

// GetRoleList
//
//	@Summary		获取角色列表
//	@Description	获取角色列表
//	@Tags				role角色
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	utils.ResponseResultInfo{data=[]role.Role}
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/role/getRoleList [get]
func GetRoleList(ctx *gin.Context) {
	roleList, err := roleServices.GetRoleList()
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}

	utils.ResponseResultsSuccess(ctx, roleList)
}
