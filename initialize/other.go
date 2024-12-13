package initialize

import "gin-api/pkg/utils"

func OtherInit() {
	// 初始化Encrypt对象, 加载AES对称加密密钥
	utils.InitEncryption()
}
