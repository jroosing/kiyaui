package main

import (
	"context"
	"flag"
	"log"

	cloudkms "cloud.google.com/go/kms/apiv1"
	cloudstore "cloud.google.com/go/storage"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/jroosing/kiyaui"
)

var (
	gcpStorage *cloudstore.Client
)

func main() {
	flag.Parse()

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	// Create the KMS client.
	kmsCtx := context.Background()
	creds, err := google.FindDefaultCredentials(kmsCtx, cloudkms.DefaultAuthScopes()...)
	if err != nil{
		log.Fatal(err)
	}
	kmsService, err := cloudkms.NewKeyManagementClient(kmsCtx, option.WithCredentials(creds))
	if err != nil {
	 log.Fatal(err)
	}
	// Create the Bucket client
	storageService, err := cloudstore.NewClient(context.Background())
	if err != nil {
		log.Fatalf("failed to create client [%v]", err)
	}
	gcpStorage = storageService
	defer gcpStorage.Close()

	kiya := kiyaui.NewKiya(gcpStorage, kmsService)

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "kiyaui",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(kiya)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
