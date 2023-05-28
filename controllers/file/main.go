package file

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/thoas/go-funk"
	"go.uber.org/zap"
	"simple-mall/global"
	"simple-mall/utils"
	"time"
)

// Upload
//
//	@Summary		文件上传
//	@Description	文件上传
//	@Tags			file文件
//	@Accept			json
//	@Produce		json
//	@Param			file	formData	file	true	"请求对象"
//	@Success		200		{object}	utils.ResponseResultInfo
//	@Failure		500		{object}	utils.EmptyInfo
//	@Security		ApiKeyAuth
//	@Router			/file/upload [post]
func Upload(ctx *gin.Context) {
	files, _ := ctx.FormFile("file")
	fileName := files.Filename

	fileTypeWhitelist := []string{"jpg", "jpeg", "png", "webp"}
	fileSuffixName := utils.GetFileSuffixName(fileName)
	if !funk.Contains(fileTypeWhitelist, fileSuffixName) {
		utils.ResponseResultsError(ctx, "仅可以上传 jpg,jpeg,png,webp 格式文件")
		return
	}

	curTime := time.Now().UTC().Format("2006-01-02")
	key := fmt.Sprintf("simple-mall/%s/%s-%s", curTime, uuid.New(), fileName)
	filesReader, _ := files.Open()

	putPolicy := storage.PutPolicy{
		Scope: global.QiNiuConfig.Bucket,
	}
	mac := qbox.NewMac(global.QiNiuConfig.Access, global.QiNiuConfig.Secret)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadongZheJiang2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "simple-mall file",
		},
	}

	err := formUploader.Put(context.Background(), &ret, upToken, key, filesReader, files.Size, &putExtra)
	if err != nil {
		utils.ResponseResultsError(ctx, err.Error())
		return
	}
	fileUrl := global.QiNiuConfig.BaseUrl + key
	zap.S().Infof("文件访问路径:%s", fileUrl)
	utils.ResponseResultsSuccess(ctx, fileUrl)
}
