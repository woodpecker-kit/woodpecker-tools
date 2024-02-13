package wd_steps_transfer_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"path/filepath"
	"testing"
)

func TestOutFast(t *testing.T) {
	// mock OutFast
	type outData struct {
		Foo string
		Bar int
	}

	type outDataArgs struct {
		root string
		info wd_info.WoodpeckerInfo
		mark string
		data interface{}
	}

	tests := []struct {
		name          string
		args          outDataArgs
		wantOutErr    bool
		appendArgs    *outDataArgs
		wantAppendErr bool
		compareOut    wd_steps_transfer.TransferObject
		wantInErr     bool
	}{
		{
			name: "sample", // testdata/TestOutFast/sample.golden
			args: outDataArgs{
				root: filepath.Join(testBaseFolderPath, "testdata", "sample"),
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithCiWorkspace(filepath.Join(testBaseFolderPath, "testdata", "sample")),
				),
				mark: "sample",
				data: outData{
					Foo: "foo",
					Bar: 2,
				},
			},
			compareOut: wd_steps_transfer.TransferObject{
				Mark:             "sample",
				CiWorkflowName:   "build",
				CiWorkflowNumber: "1",
				Data: outData{
					Foo: "foo",
					Bar: 2,
				},
			},
		},
		{
			name: "append", // testdata/TestOutFast/append.golden
			args: outDataArgs{
				root: filepath.Join(testBaseFolderPath, "testdata", "append"),
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithCiWorkspace(filepath.Join(testBaseFolderPath, "testdata", "append")),
				),
				mark: "append",
				data: outData{
					Foo: "foo",
					Bar: 1,
				},
			},
			appendArgs: &outDataArgs{
				root: filepath.Join(testBaseFolderPath, "testdata", "append"),
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithCiWorkspace(filepath.Join(testBaseFolderPath, "testdata", "append")),
				),
				mark: "append",
				data: outData{
					Foo: "foo",
					Bar: 11,
				},
			},
			compareOut: wd_steps_transfer.TransferObject{
				Mark:             "append",
				CiWorkflowName:   "build",
				CiWorkflowNumber: "1",
				Data: outData{
					Foo: "foo",
					Bar: 11,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do OutFast
			transferObj, gotErr := wd_steps_transfer.OutFast(tc.args.root, tc.args.info, tc.args.mark, tc.args.data)
			assert.Equal(t, tc.wantOutErr, gotErr != nil)
			if tc.wantOutErr {
				t.Logf("want out err: %v", gotErr)
				return
			}
			if tc.appendArgs == nil {
				assert.Equal(t, tc.compareOut, *transferObj)
			} else { // append
				transferAppendObj, gotAppendErr := wd_steps_transfer.OutFast(tc.appendArgs.root, tc.appendArgs.info, tc.appendArgs.mark, tc.appendArgs.data)
				assert.Equal(t, tc.wantAppendErr, gotAppendErr != nil)
				if tc.wantAppendErr {
					t.Logf("want out err: %v", gotAppendErr)
					return
				}
				assert.Equal(t, tc.compareOut, *transferAppendObj)
			}

			// verify OutFast
			var readData outData
			gotInErr := wd_steps_transfer.InFast(tc.args.root, tc.args.info, tc.args.mark, &readData)
			assert.Equal(t, tc.wantInErr, gotInErr != nil)
			if tc.wantInErr {
				t.Logf("want in err: %v", gotInErr)
				return
			}

			g.AssertJson(t, t.Name(), readData)
		})
	}
}
