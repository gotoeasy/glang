package cmn

import (
	"golang.org/x/sys/windows/registry"
)

// 通过注册表设定自定义URL协议，创建完成后用【appname://】可打开程序
func RegistrySetUrlProtocol(appname string, appNote string, exePath string) error {
	// 创建或打开 HKEY_CLASSES_ROOT\appname 子项
	appnameKey, err := registryOpenOrCreateKey(registry.CLASSES_ROOT, appname)
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
	shellKey, err := registryOpenOrCreateKey(appnameKey, "shell")
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell 子项`, err)
		return err
	}
	defer shellKey.Close()

	// 在 shell 子项下创建 open 子项
	openKey, err := registryOpenOrCreateKey(shellKey, "open")
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell\open 子项`, err)
		return err
	}
	defer openKey.Close()

	// 在 open 子项下创建 command 子项
	commandKey, err := registryOpenOrCreateKey(openKey, "command")
	if err != nil {
		Error(`无法创建 HKEY_CLASSES_ROOT\`+appname+`\shell\open\command 子项`, err)
		return err
	}
	defer openKey.Close()

	// 子项 command 中设置默认值（目标程序及参数）
	if err := commandKey.SetStringValue("", `"`+exePath+`" "%1"`); err != nil {
		Error(`无法设置 HKEY_CLASSES_ROOT\`+appname+`\shell\open\command 默认值`, err)
		return err
	}

	return nil
}

func registryOpenOrCreateKey(k registry.Key, path string) (registry.Key, error) {
	key, err := registry.OpenKey(k, path, registry.ALL_ACCESS)
	if err == nil {
		return key, err
	}

	key, _, err = registry.CreateKey(k, path, registry.ALL_ACCESS)
	return key, err
}
