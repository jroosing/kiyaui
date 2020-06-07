package kiyaui

import (
	"path"
	"sort"

	cloudstore "cloud.google.com/go/storage"
	"github.com/fsnotify/fsnotify"
	kiyalib "github.com/kramphub/kiya"
	"github.com/wailsapp/wails"
	"google.golang.org/api/cloudkms/v1"
)

type Kiya struct {
	profilesFileLoc string
	runtime *wails.Runtime
	logger *wails.CustomLogger
	profilesWatcher *fsnotify.Watcher
	gcpStorage *cloudstore.Client
	kmsService *cloudkms.Service

	Profiles map[string]kiyalib.Profile
}

func NewKiya(gcpStorage *cloudstore.Client, kmsService *cloudkms.Service) *Kiya {
	kiya := &Kiya{
		gcpStorage: gcpStorage,
		kmsService: kmsService,
	}
	return kiya
}

func (k *Kiya) WailsInit(runtime *wails.Runtime) error {
	k.runtime = runtime
	k.logger = k.runtime.Log.New("KiyaUI")
	homedir, err := runtime.FileSystem.HomeDir()
	if err != nil {
		return err
	}
	k.profilesFileLoc = path.Join(homedir, ".kiya")
	k.runtime.Window.SetTitle(k.profilesFileLoc)
	ensureFileExists(k.profilesFileLoc)

	kiyalib.LoadConfiguration(k.profilesFileLoc)
	k.Profiles = kiyalib.Profiles
	return k.startWatcher()
}

func (k *Kiya) reloadConfig() {
	kiyalib.LoadConfiguration(k.profilesFileLoc)
	k.Profiles = kiyalib.Profiles
}

func (k *Kiya) ProfileNames() (profNames []string) {
	for k, _ := range k.Profiles {
		profNames = append(profNames, k)
	}
	sort.Strings(profNames)
	return profNames
}

func (k *Kiya) startWatcher() error {
	k.logger.Info("Starting Watcher")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	k.profilesWatcher = watcher

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op & fsnotify.Write == fsnotify.Write {
					k.logger.Infof("modified file: %s", event.Name)
					k.runtime.Events.Emit("filemodified")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				k.logger.Error(err.Error())
			}
		}
	}()

	err = watcher.Add(k.profilesFileLoc)
	if err != nil {
		return err
	}
	return nil
}
