package cmn

import (
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// 生成指定文件的快捷文件
func CreateShortcutFile(targetFile string, shortcutLnkFile string) error {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}
	defer oleShellObject.Release()

	shellObject, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	shortcut, err := oleutil.CallMethod(shellObject, "CreateShortcut", shortcutLnkFile)
	if err != nil {
		return err
	}

	shortcutObject := shortcut.ToIDispatch()
	defer shortcutObject.Release()

	_, err = oleutil.PutProperty(shortcutObject, "TargetPath", targetFile)
	if err != nil {
		return err
	}
	_, err = oleutil.PutProperty(shortcutObject, "WorkingDirectory", filepath.Dir(targetFile))
	if err != nil {
		return err
	}
	_, err = oleutil.CallMethod(shortcutObject, "Save")
	if err != nil {
		return err
	}

	return nil
}
