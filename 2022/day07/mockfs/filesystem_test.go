package mockfs

import (
	"strings"
	"testing"
)

func TestFileSystem(t *testing.T) {
	t.Run("Move", func(t *testing.T) {
		t.Run("Forward", func(t *testing.T) {
			t1 := CreateDirectory("t1")
			fs := CreateFileSystem()
			fs.CurrentDir().AddDirectory(t1)

			if err := fs.Move("t1"); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != t1 {
				t.Error("Expected dir", t1, "got", r1)
			}
		})

		t.Run("ForwardMultiple", func(t *testing.T) {
			t1 := CreateDirectory("t1")
			t2 := CreateDirectory("t2")
			t1.AddDirectory(t2)

			fs := CreateFileSystem()
			fs.CurrentDir().AddDirectory(t1)

			if err := fs.Move("t1/t2"); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != t2 {
				t.Error("Expected dir", t2, "got", r1)
			}
		})

		t.Run("Back", func(t *testing.T) {
			t1 := CreateDirectory("t1")
			fs := CreateFileSystem()
			root := fs.CurrentDir()
			root.AddDirectory(t1)

			if err := fs.Move("t1"); err != nil {
				t.Error(err)
			}

			if err := fs.Move(".."); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != root {
				t.Error("Expected dir", root, "got", r1)
			}
		})

		t.Run("Nowhere", func(t *testing.T) {
			t1 := CreateDirectory("t1")
			fs := CreateFileSystem()
			fs.CurrentDir().AddDirectory(t1)

			if err := fs.Move("t1"); err != nil {
				t.Error(err)
			}

			if err := fs.Move(""); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != t1 {
				t.Error("Expected dir", t1, "got", r1)
			}
		})

		t.Run("FromRoot", func(t *testing.T) {
			t1 := CreateDirectory("t1")
			fs := CreateFileSystem()
			root := fs.CurrentDir()
			root.AddDirectory(t1)

			if err := fs.Move("t1"); err != nil {
				t.Error(err)
			}

			if err := fs.Move("/"); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != root {
				t.Error("Expected dir", root, "got", r1)
			}

			if err := fs.Move("t1"); err != nil {
				t.Error(err)
			}
			if err := fs.Move("/t1"); err != nil {
				t.Error(err)
			}

			if r1 := fs.CurrentDir(); r1 != t1 {
				t.Error("Expected dir", t1, "got", r1)
			}
		})
	})

	t.Run("FindAll", func(t *testing.T) {
		t1 := CreateDirectory("a1")
		t2 := CreateDirectory("t2")
		t21 := CreateDirectory("t21")
		t22 := CreateDirectory("a22")
		fs := CreateFileSystem()
		fs.CurrentDir().AddDirectory(t1)
		fs.CurrentDir().AddDirectory(t2)
		t2.AddDirectory(t21)
		t2.AddDirectory(t22)

		results := fs.FindAll(func(d Directory) bool {
			return strings.Contains(d.GetName(), "a")
		})

		for r := range results {
			if r != t1 && r != t22 {
				t.Error("Expected", t1, "or", t22, "got", r)
			}
		}
	})

}
