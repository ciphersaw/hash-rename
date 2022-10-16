package action

import (
	"testing"
)

type args struct {
	File        string
	Dir         string
	Suffix      string
	Hash        string
	Concurrency uint8
}

func importArgs(a *args) {
	argFile = a.File
	argDir = a.Dir
	argSuffix = a.Suffix
	argHash = a.Hash
	argConcurrency = a.Concurrency
}

func createCaseRenameOneFile(t *testing.T, a *args) {
	importArgs(a)
	if err = renameOneFile(); err != nil {
		t.Errorf("renameOneFile test fails: %v", err)
	} else {
		t.Logf("renameOneFile test successes.")
	}
}

func TestRenameOneFile(t *testing.T) {
	// Right test cases
	createCaseRenameOneFile(t, &args{File: "../example/original.png", Hash: "md5"})
	createCaseRenameOneFile(t, &args{File: "../example/notocat", Hash: "sha1"})
	// Wrong test cases
	createCaseRenameOneFile(t, &args{File: "../example/unknown.png", Hash: "md5"})
	createCaseRenameOneFile(t, &args{File: "../example/original.png", Hash: "unknown"})
}

func createCaseRenameBulkFiles(t *testing.T, a *args) {
	importArgs(a)
	suffixConfig, err = setSuffixConfig(argSuffix)
	if err != nil {
		t.Errorf("renameBulkFiles sets suffix config error: %v", err)
		return
	}
	err = checkConcurrency(argConcurrency)
	if err != nil {
		t.Errorf("renameBulkFiles sets goroutine concurrency error: %v", err)
		return
	}
	if err = renameBulkFiles(); err != nil {
		t.Errorf("renameBulkFiles test fails: %v", err)
	} else {
		t.Logf("renameBulkFiles test successes.")
	}
}

func TestRenameBulkFiles(t *testing.T) {
	// Right test cases
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "png,jpg", Hash: "md5", Concurrency: 4})
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "null", Hash: "sha1", Concurrency: 8})
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "all", Hash: "sha256", Concurrency: 64})
	// Wrong test cases
	createCaseRenameBulkFiles(t, &args{Dir: "../unknown/", Suffix: "all", Hash: "md5", Concurrency: 8})
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "unknown", Hash: "md5", Concurrency: 8})
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "all", Hash: "unknown", Concurrency: 8})
	createCaseRenameBulkFiles(t, &args{Dir: "../example/", Suffix: "all", Hash: "md5", Concurrency: 0})
}
