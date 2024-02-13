package wd_steps_transfer

import (
	"encoding/json"
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"path/filepath"
)

// OutFast
// save data to file use DefaultKitStepsFileName
//
//	if file exist, will append data to file latest data with mark
//	if file not exist, will create file and save data
//
// root is the root path is runner workspace.
// info is the woodpecker wd_info.WoodpeckerInfo.
// mark is the mark key with key
// data is the data you want to save
//
//	Out
//
// TransferObject is the data you save to file
// error is the error if save failed
//
//	code:
//	var data outData
//	errOut := wd_steps_transfer.OutFast(root, info, mark, data)
func OutFast(root string, info wd_info.WoodpeckerInfo, mark string, data interface{}) (*TransferObject, error) {
	return Out(root, DefaultKitStepsFileName, info, mark, data)
}

// Out
// save data to file
//
//	if file exist, will append data to file latest data with mark
//	if file not exist, will create file and save data
//
// root is the root path is runner workspace.
// filename will join as filepath you can use DefaultKitStepsFileName.
// info is the woodpecker wd_info.WoodpeckerInfo.
// mark is the mark key with key
// data is the data you want to save
//
//	Out
//
// TransferObject is the data you save to file
// error is the error if save failed
//
//	code:
//	var data outData
//	errOut := wd_steps_transfer.Out(root, filename, info, mark, data)
func Out(root, filename string, info wd_info.WoodpeckerInfo, mark string, data interface{}) (*TransferObject, error) {
	targetSave := filepath.Join(root, filename)
	var tsObjs []TransferObject
	if pathExistsFast(targetSave) {
		errReadOld := inAsJson(root, filename, &tsObjs)
		if errReadOld != nil {
			return nil, fmt.Errorf("failed to read old transfer path: %s\n: %v", targetSave, errReadOld)
		}
		var tempTsObjs []TransferObject
		for _, obj := range tsObjs { // remove old mark
			if obj.Mark == mark {
				continue
			}
			tempTsObjs = append(tempTsObjs, obj)
		}
		tsObjs = tempTsObjs
	} else {
		tsObjs = make([]TransferObject, 0)
	}
	transferObject := TransferObject{
		Mark:             mark,
		CiWorkflowName:   info.CurrentInfo.CurrentWorkflowInfo.CiWorkflowName,
		CiWorkflowNumber: info.CurrentInfo.CurrentWorkflowInfo.CiWorkflowNumber,
		Data:             data,
	}
	tsObjs = append(tsObjs, transferObject)
	errOut := outAsJson(root, filename, tsObjs)
	if errOut != nil {
		return nil, fmt.Errorf("failed to out transfer path: %s\n: %v", targetSave, errOut)
	}
	return &transferObject, errOut
}

func outAsJson(root, filename string, data interface{}) error {
	marshalData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}
	return outAsByte(root, filename, marshalData, true)
}

func outAsByte(root string, filename string, data []byte, cover bool) error {
	targetSave := filepath.Join(root, filename)
	if pathExistsFast(targetSave) && !cover {
		return fmt.Errorf("not cover, which path exist %v", targetSave)
	}
	return writeFileFast(targetSave, data, cover)
}
