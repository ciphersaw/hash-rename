package action

import (
	"testing"
)

type args struct {
	Uppercase   bool
	Force       bool
	File        string
	Dir         string
	Suffix      string
	Hash        string
	Concurrency uint8
}

func importArgs(a *args) {
	argUppercase = a.Uppercase
	argForce = a.Force
	argFile = a.File
	argDir = a.Dir
	argSuffix = a.Suffix
	argHash = a.Hash
	argConcurrency = a.Concurrency
}

func createCaseRenameOneFile(t *testing.T, a *args) {
	importArgs(a)
	if hashFunc, err = setHashFunc(argHash); err != nil {
		t.Errorf("renameOneFile sets the hash function error: %v", err)
		return
	}
	if err = renameOneFile(); err != nil {
		t.Errorf("renameOneFile test fails: %v", err)
	} else {
		t.Logf("renameOneFile test successes.")
	}
}

func TestRenameOneFile(t *testing.T) {
	// Right test cases
	createCaseRenameOneFile(t, &args{
		File:      "../example/original.png",
		Hash:      "md5",
		Uppercase: false,
		Force:     false,
	})
	createCaseRenameOneFile(t, &args{
		File:      "../example/notocat",
		Hash:      "sha1",
		Uppercase: true,
		Force:     false,
	})
	createCaseRenameOneFile(t, &args{
		File:      "../example/CB5DCA692F6DB49A83C658F64E18E0005BE1A052",
		Hash:      "sha1",
		Uppercase: true,
		Force:     true,
	})

	// Wrong test cases
	createCaseRenameOneFile(t, &args{
		File:      "../example/CB5DCA692F6DB49A83C658F64E18E0005BE1A052",
		Hash:      "sha1",
		Uppercase: true,
		Force:     false,
	})
	createCaseRenameOneFile(t, &args{
		File:      "../example/unknown.png",
		Hash:      "md5",
		Uppercase: false,
		Force:     false,
	})
	createCaseRenameOneFile(t, &args{
		File:      "../example/original.png",
		Hash:      "unknown",
		Uppercase: false,
		Force:     false,
	})
}

func createCaseRenameBulkFiles(t *testing.T, a *args) {
	importArgs(a)
	if suffixConfig, err = setSuffixConfig(argSuffix); err != nil {
		t.Errorf("renameBulkFiles sets suffix config error: %v", err)
		return
	}
	if err = checkConcurrency(argConcurrency); err != nil {
		t.Errorf("renameBulkFiles sets goroutine concurrency error: %v", err)
		return
	}
	if hashFunc, err = setHashFunc(argHash); err != nil {
		t.Errorf("renameBulkFiles sets the hash function error: %v", err)
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
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "png,jpg",
		Hash:        "md5",
		Concurrency: 4,
		Uppercase:   false,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "null",
		Hash:        "sha1",
		Concurrency: 8,
		Uppercase:   false,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "all",
		Hash:        "sha256",
		Concurrency: 64,
		Uppercase:   true,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "all",
		Hash:        "sha256",
		Concurrency: 64,
		Uppercase:   true,
		Force:       true,
	})

	// Wrong test cases
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "all",
		Hash:        "sha256",
		Concurrency: 64,
		Uppercase:   true,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../unknown/",
		Suffix:      "all",
		Hash:        "md5",
		Concurrency: 8,
		Uppercase:   false,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "unknown",
		Hash:        "md5",
		Concurrency: 8,
		Uppercase:   false,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "all",
		Hash:        "unknown",
		Concurrency: 8,
		Uppercase:   false,
		Force:       false,
	})
	createCaseRenameBulkFiles(t, &args{
		Dir:         "../example/",
		Suffix:      "all",
		Hash:        "md5",
		Concurrency: 0,
		Uppercase:   false,
		Force:       false,
	})
}
