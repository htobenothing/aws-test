package main

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws-demo/core"

	"github.com/gin-gonic/gin"
)

func main() {
	router := initRouter()
	router.Run("0.0.0.0:8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	// router.Use(static.Serve("/html", static.LocalFile("./html/", true)))
	router.Static("/html", "./html")
	router.GET("/xsstest", func(c *gin.Context) {
		name := c.Query("name")
		c.String(200, "hello "+name)
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "welcome to aws test",
		})
	})

	router.POST("/upload", uploadFile)
	return router
}

// func UploadFileMultiplePart(c *gin.Context) {
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		fmt.Printf("error %v", err.Error())
// 	}
// 	files :=form.File["file"]

// }
func uploadFileStream(c *gin.Context) {
	// req := c.Request
	// // p := make([]byte, 1024)
	// buf := new(bytes.Buffer)
	// buf.ReadFrom(req.Body)
	// readCloser, _ := c.Request.GetBody()

}

func S3Upload(sess *session.Session, file *multipart.File, fileHead *multipart.FileHeader) (result *s3manager.UploadOutput, err error) {

	upload := s3manager.NewUploader(sess)
	uploadInput := s3manager.UploadInput{
		Bucket: aws.String("testing-nothing"),
		Key:    aws.String(fileHead.Filename),
		Body:   *file,
	}

	result, err = upload.Upload(&uploadInput)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func uploadFile(c *gin.Context) {

	fmt.Printf("start uploading ...\n")
	start := time.Now()
	c.Request.ParseMultipartForm(8 << 20)
	file, fileHead, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Printf("error %v", file)
	}

	defer file.Close()
	sess := core.AwsSession()
	result, err := S3Upload(sess, &file, fileHead)

	if err != nil {
		fmt.Errorf("failed to upload file, %v", err)
	}
	size := fileHead.Size / 1024 / 1024
	elapsed := time.Since(start)
	speed := float64(size) / elapsed.Seconds()
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	fmt.Printf("total time %v, file size %v, speed %v\n", elapsed, size, speed)
	c.JSON(http.StatusOK, gin.H{
		"size":      size,
		"totaltime": elapsed,
		"speed":     speed,
		"location":  result.Location,
	})

}
