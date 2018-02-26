package tempfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTempFileUniqueness(t *testing.T) {
	filenames := make(map[string]struct{})
	directories := make(map[string]struct{})
	bases := make(map[string]struct{})

	defer func() {
		for filename := range filenames {
			if err := os.Remove(filename); err != nil {
				t.Error("Unable to clean up file:", filename)
			}
		}
	}()

	for i := 0; i < 50000; i++ {
		file, err := TempFile("", "myPrefix", "mySuffix")
		if err != nil {
			t.Error("Unable to generate a unique file:", err)
			continue
		}
		if _, err := fmt.Fprintf(file, "Some content"); err != nil {
			t.Error("Expected to be able to write to file:", err)
		}
		if err := file.Close(); err != nil {
			// Important to close all the files here to now run out of file
			// descriptors.
			t.Error("Could not close file:", err)
		}
		base := filepath.Base(file.Name())

		if !strings.HasPrefix(base, "myPrefix") {
			t.Error("Expected 'myPrefix' prefix. File:", base)
		}
		if !strings.HasSuffix(base, "mySuffix") {
			t.Error("Expected 'mySuffix' suffix. File:", base)
		}

		filenames[file.Name()] = struct{}{}
		bases[base] = struct{}{}
		directories[filepath.Dir(file.Name())] = struct{}{}
	}

	if len(directories) != 1 {
		t.Error("Expected all files to end up in the same temporary directory.")
	}

}
