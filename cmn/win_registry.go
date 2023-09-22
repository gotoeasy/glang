package cmn

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

// 创建自定义URL协议，创建完成后用【appname://】可打开程序，暂不支持更新
func CreateRegistry4UrlProtocol(appname string, appNote string, exePath string) error {
	// 创建或打开 HKEY_CLASSES_ROOT\appname 子项
	_, err := registry.OpenKey(registry.CLASSES_ROOT, appname, registry.ALL_ACCESS)
	if err == nil {
		return nil
	}

	// 如果出错，创建新的子项 appname
	appnameKey, _, err := registry.CreateKey(registry.CLASSES_ROOT, appname, registry.ALL_ACCESS)
	if err != nil {
		Error(`无法创建或打开 HKEY_CLASSES_ROOT\`+appname+` 子项`, err)
		return err
	}
	defer appnameKey.Close()

	// 子项 appname 子项中设置默认值（说明）
	if err := appnameKey.SetStringValue("", appNote); err != nil {
		Error(err)
		return err
	}
	// 子项 appname 子项中设置执行文件
	if err := appnameKey.SetStringValue("URL Protocol", exePath); err != nil {
		Error(err)
		return err
	}

	// 在 appname 子项下创建 shell 子项
	shellKey, _, err := registry.CreateKey(appnameKey, "shell", registry.ALL_ACCESS)
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell 子项`, err)
		return err
	}
	defer shellKey.Close()

	// 在 shell 子项下创建 open 子项
	openKey, _, err := registry.CreateKey(shellKey, "open", registry.ALL_ACCESS)
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell\open 子项`, err)
		return err
	}
	defer openKey.Close()

	// 在 open 子项下创建 command 子项
	commandKey, _, err := registry.CreateKey(openKey, "command", registry.ALL_ACCESS)
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell\open\command 子项`, err)
		return err
	}
	defer openKey.Close()

	// 子项 command 中设置默认值（目标程序及参数）
	if err := commandKey.SetStringValue("", `"`+exePath+`" "%1"`); err != nil {
		Error(`无法设置 HKEY_CLASSES_ROOT\`+appname+`\shell\open\command 默认值`, err)
		fmt.Println("无法默认值:", err)
		return err
	}

	return nil
}
