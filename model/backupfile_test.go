// Copyright (c) 2020 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var _ = Describe("BackupFilename", func() {
	Describe("BuildBackupfileName", func() {
		It("builds the expected filename", func() {
			filename := BuildBackupfileName("myname", "/tmp", "mydb", time.Unix(1313123123, 0))
			Expect(filename.String()).To(Equal("/tmp/myname_mydb_2011-08-12.dump"))
		})
	})

	Describe("Exists", func() {
		It("returns true when the file exists and is not empty", func() {
			file, err := os.CreateTemp("", "backupfile")
			Expect(err).NotTo(HaveOccurred())
			defer func() { _ = os.Remove(file.Name()) }()

			_, err = file.WriteString("hello world")
			Expect(err).NotTo(HaveOccurred())
			file.Close()

			Expect(BackupFilename(file.Name()).Exists()).To(BeTrue())
		})

		It("returns false when the file exists but is empty", func() {
			file, err := os.CreateTemp("", "backupfile")
			Expect(err).NotTo(HaveOccurred())
			defer func() { _ = os.Remove(file.Name()) }()
			file.Close()

			Expect(BackupFilename(file.Name()).Exists()).To(BeFalse())
		})

		It("returns false when the file does not exist", func() {
			Expect(BackupFilename("/filedoesnotexists").Exists()).To(BeFalse())
		})
	})
})

func TestModel(t *testing.T) {
	time.Local = time.UTC
	format.TruncatedDiff = false
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}
