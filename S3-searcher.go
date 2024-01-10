package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
    "strings"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
    // Replace with your S3 bucket name and substring
    bucketName := "your-s3-bucket-name"
    searchSubstring := "your-substring"

    // Convert search substring to lower case for case-insensitive search
    searchSubstring = strings.ToLower(searchSubstring)

    // Load AWS configuration
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatalf("Error loading S3 configuration: %v", err)
    }

    client := s3.NewFromConfig(cfg)

    // Create a paginator to iterate through the bucket objects
    paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
        Bucket: aws.String(bucketName),
    })

    foundFiles := []string{}
    for paginator.HasMorePages() {
        page, err := paginator.NextPage(context.TODO())
        if err != nil {
            log.Fatalf("Failed to get next page: %v", err)
        }

        for _, object := range page.Contents {
            key := *object.Key
            fileExtension := filepath.Ext(key)

            if fileExtension == ".txt" {
                obj, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
                    Bucket: aws.String(bucketName),
                    Key:    object.Key,
                })
                if err != nil {
                    log.Printf("Failed to get S3 object: %v", err)
                    continue
                }

                defer obj.Body.Close()

                fileContent, err := ioutil.ReadAll(obj.Body)
                if err != nil {
                    log.Printf("Failed to read S3 object: %v", err)
                    continue
                }

                // Perform case-insensitive search
                if strings.Contains(strings.ToLower(string(fileContent)), searchSubstring) {
                    fileLocation := "s3://" + bucketName + "/" + key
                    foundFiles = append(foundFiles, fileLocation)
                }
            }
        }
    }

    fmt.Println("Files containing substring:", foundFiles)
}
