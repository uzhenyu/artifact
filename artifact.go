package z224

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

func Md5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
func Img(localFile, key string) {
	bucket := "j9s9q9"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := auth.New("JV1t0v6D3n_OSon-mflCCAkwVzAc9jEsQt6y3iS3", "nPTmBTZvJm4e0yZJYFb0szVcbi1Ci0lAkoUV7Bif")
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
