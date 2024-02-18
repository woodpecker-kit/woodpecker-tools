package wd_steps_transfer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"path/filepath"
)

var ErrInEmptyData = errors.New("in data with empty")
var ErrInNotMatchSameCiWorkflowNumber = errors.New("in data not match same ci workflow number")
var ErrInNotExists = errors.New("in data not exists")

// InFast
// read data from method Out or OutFast by use DefaultKitStepsFileName.
//
//	warning: if you want to use this method, you should use Out or OutFast to save data
//
// read data from file
// root is the root path is runner workspace.
// info is the woodpecker wd_info.WoodpeckerInfo.
// mark is the mark key with key
// data is the data you want to read
//
//	out
//
// if success, will return nil.
// if failed, will return error.
//
//	error is the error if read failed, include:
//
// ErrInEmptyData is the error if in data with empty
// ErrInNotMatchSameCiWorkflowNumber is the error if in data not match same ci workflow number
// ErrInNotExists is the error if in data not exists
//
//	code:
//	var readData outData
//	gotInErr := wd_steps_transfer.InFast(root, info, mark, &readData)
func InFast(root string, info wd_info.WoodpeckerInfo, mark string, data interface{}) error {
	return In(root, DefaultKitStepsFileName, info, mark, data)
}

// In
// read data from method Out or OutFast
//
//	warning: if you want to use this method, you should use Out or OutFast to save data
//
// read data from file
// root is the root path is runner workspace.
// filename will join as filepath you can use DefaultKitStepsFileName.
// info is the woodpecker wd_info.WoodpeckerInfo.
// mark is the mark key with key
// data is the data you want to read
//
//	out
//
// if success, will return nil.
// if failed, will return error.
//
//	error is the error if read failed, include:
//
// ErrInEmptyData is the error if in data with empty.
// ErrInNotMatchSameCiWorkflowNumber is the error if in data not match same ci workflow number, this error will return data, but you must check!
// ErrInNotExists is the error if in data not exists.
//
//	code:
//	var readData outData
//	gotInErr := wd_steps_transfer.In(root, filename, info, mark, &readData)
func In(root, filename string, info wd_info.WoodpeckerInfo, mark string, data interface{}) error {
	targetSave := filepath.Join(root, filename)
	if !pathExistsFast(targetSave) {
		return fmt.Errorf("want read path not exist %v", targetSave)
	}
	var tsObjs []TransferObject
	errReadAsJson := inAsJson(root, filename, &tsObjs)
	if errReadAsJson != nil {
		return fmt.Errorf("failed to read transfer path: %s\n: %v", targetSave, errReadAsJson)
	}
	if len(tsObjs) == 0 {
		return ErrInEmptyData
	}
	for _, obj := range tsObjs {
		if obj.Mark == mark {
			marshalData, _ := json.Marshal(obj.Data)
			errDataFromJson := json.Unmarshal(marshalData, data)
			if errDataFromJson != nil {
				return fmt.Errorf("failed to unmarshal data from json: %v", errDataFromJson)
			}
			if obj.CiWorkflowNumber != info.CurrentInfo.CurrentWorkflowInfo.CiWorkflowNumber {
				return ErrInNotMatchSameCiWorkflowNumber
			}
			return nil
		}
	}

	return ErrInNotExists
}

func inAsJson(root, filename string, inRead interface{}) error {
	fileAsByte, err := inAsByte(root, filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}
	errJson := json.Unmarshal(fileAsByte, inRead)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data from json: %v", errJson)
	}
	return nil
}

func inAsByte(root string, filename string) ([]byte, error) {
	targetPath := filepath.Join(root, filename)
	if !pathExistsFast(targetPath) {
		return nil, fmt.Errorf("path not exist %v", targetPath)
	}
	return readFileAsByte(targetPath)
}
