package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ProcessManager manges the file watching process.
// As soon as all files in a subdirectory of <CMD arg -src>
// (or a file directly in <CMD arg -src>) are
// not changed for almost exactly <CMD -duration> seconds,
// the subdirectory will be pushed into the channel 'done_files'.
type ProcessManager struct {
	args               *Args
	done_files         chan string
	done_flat_prefixes map[string]bool
}

func (m ProcessManager) collectTarPrefixes() {
	entries, err := ReadDir(FlatTarTempPath)
	if err != nil {
		log.Fatal(err)
	}
	for prefix := range m.done_flat_prefixes {
		delete(m.done_flat_prefixes, prefix)
	}

	for _, v := range entries {
		if v.IsDir() {
			continue
		}
		m.done_flat_prefixes[v.Name()[0:m.args.commonPrefixLength]] = true
	}
}

func (m ProcessManager) processTarPrefixes() {
	entries, err := ReadDir(FlatTarTempPath)
	if err != nil {
		log.Fatal(err)
	}

	for prefix := range m.done_flat_prefixes {
		dirName := filepath.Join(FlatTarTempPath, prefix)
		err := os.MkdirAll(dirName, os.ModeDir|os.ModePerm)
		if err != nil {
			ErrorLogger.Println(err)
			continue
		}
		for _, v := range entries {
			if v.IsDir() {
				continue
			}
			if strings.HasPrefix(v.Name(), prefix) {
				sourcePath := filepath.Join(FlatTarTempPath, v.Name())
				err := Copy(sourcePath, filepath.Join(dirName, v.Name()))
				if err != nil {
					ErrorLogger.Println(err)
				}
				err = os.Remove(sourcePath)
			}
		}

		tarPaht, err := tarFolder(dirName)
		if err != nil {
			ErrorLogger.Println(err)
		}

		err = Copy(tarPaht, filepath.Join(dirName, filepath.Base(tarPaht)))
		if err != nil {
			ErrorLogger.Println(err)
		}
		m.done_files <- tarPaht
		err = os.RemoveAll(dirName)
		if err != nil {
			ErrorLogger.Println(err)
			continue
		}

		delete(m.done_flat_prefixes, prefix)
	}
}

// doWork runs in a endless loop. It watches the files in the <CMD arg -src> directory.
// It terminates as soon as a value is pushed into quit. Run in extra goroutine.
func (m ProcessManager) doWork(quit chan int) {
	InfoLogger.Println("Started watch process.")
	for {
		select {
		case <-quit:
			return
		default:
			now := time.Now()
			done_folders := make(map[string]bool)
			// Checking all files in <CMD arg -src>.
			err := filepath.Walk(m.args.src,
				func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}
					if !info.IsDir() {
						modifiedtime := info.ModTime()
						diff := now.Sub(modifiedtime)
						if diff < 2*m.args.duration {
							if relpath, err := filepath.Rel(m.args.src, path); err == nil {
								folder := relpath
								if m.args.sendType != "file" && m.args.sendType != "flat_tar" {
									folder = getRootDir(relpath)
								}

								if _, ok := done_folders[folder]; !ok {
									done_folders[folder] = true
								}
								if diff <= m.args.duration {
									done_folders[folder] = false
								}
							} else {
								ErrorLogger.Println(err)
							}
						}
					}
					return nil
				})

			if err != nil {
				ErrorLogger.Println(err)
			}

			// Pushing all complete subdirectory into done_files channel.
			if m.args.sendType == "flat_tar" {
				m.collectTarPrefixes()
			}

			for k, v := range done_folders {
				if v {
					InfoLogger.Println("Folder/File ready to send: ", k)
					if m.args.sendType == "flat_tar" {
						src := filepath.Join(m.args.src, k)
						dst := filepath.Join(FlatTarTempPath, filepath.Base(k))
						err := Copy(src, dst)
						if err != nil {
							ErrorLogger.Println(err)
						}
					} else {
						m.done_files <- filepath.Join(m.args.src, k)
					}
				}
			}
			if m.args.sendType == "flat_tar" {
				m.processTarPrefixes()
			}

			time.Sleep(m.args.duration - time.Since(now))
		}
	}
}

// newProcessManager factory for ProcessManager struct
func newProcessManager(args *Args, done_files chan string) ProcessManager {
	return ProcessManager{args: args, done_files: done_files, done_flat_prefixes: make(map[string]bool)}
}
