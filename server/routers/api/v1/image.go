package v1

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/athagi/hello-copilot/server/pkg/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type Image struct {
	Name         string
	LastModified time.Time
	Size         int64
}

type Page struct {
	Page   int
	Images []Image
}

func DeleteImage(c *gin.Context) {
	bucket := "255222094062-sample-images"
	fileName := c.Query("name")
	fmt.Println(fileName)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

func DownloadImage(c *gin.Context) {
	bucket := "255222094062-sample-images"
	objectKey := c.Query("key")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	// TODO, use stream not to create tmp file and not to use much memory.
	f, err := os.Create("tmp")
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	defer f.Close()

	client := s3.NewFromConfig(cfg)
	downloader := manager.NewDownloader(client)
	n, err := downloader.Download(context.TODO(), f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)

	_, err = io.Copy(c.Writer, f)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func ListImages(c *gin.Context) {
	var bucket string
	objectPrefix := ""
	objectDelimiter := ""
	maxKeys := 10

	bucket = "255222094062-sample-images"
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)
	// Set the parameters based on the CLI flag inputs.
	params := &s3.ListObjectsV2Input{
		Bucket: &bucket,
	}
	if len(objectPrefix) != 0 {
		params.Prefix = &objectPrefix
	}
	if len(objectDelimiter) != 0 {
		params.Delimiter = &objectDelimiter
	}

	// Create the Paginator for the ListObjectsV2 operation.
	p := s3.NewListObjectsV2Paginator(client, params, func(o *s3.ListObjectsV2PaginatorOptions) {
		if v := int32(maxKeys); v != 0 {
			o.Limit = v
		}
	})

	// Iterate through the S3 object pages, printing each object returned.
	var i int
	log.Println("Objects:")
	pages := []Page{}
	for p.HasMorePages() {
		i++

		// Next Page takes a new context for each page retrieval. This is where
		// you could add timeouts or deadlines.
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page %v, %v", i, err)
		}

		// Log the objects found
		images := make([]Image, len(page.Contents))
		for j, item := range page.Contents {
			key := *item.Key
			lastModified := *item.LastModified
			size := item.Size
			images[j] = Image{Name: key, LastModified: lastModified, Size: size}
			fmt.Println(images)
		}
		pages = append(pages, Page{Page: i, Images: images})
	}
	c.JSON(200, gin.H{
		"message": "get images",
		"items":   pages,
	})
}

func UploadImage(c *gin.Context) {
	bucket := "255222094062-sample-images"

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	uuid := util.GenerateUUID4()
	files, err := c.FormFile("file")
	ext := strings.Split(files.Filename, ".")[len(strings.Split(files.Filename, "."))-1]
	fileName := uuid + "." + ext

	uploadFile, err := files.Open()
	defer uploadFile.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   uploadFile,
	})

	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	log.Printf("successfully uploaded file to %s\n", result.Location)
	c.JSON(http.StatusOK, gin.H{"message": "success!!"})
}
