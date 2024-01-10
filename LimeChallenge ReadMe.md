# S3 Text File Searcher

This script searches through all text files in a specified AWS S3 bucket for a given substring.

## Setup

1. Install Go (version 1.x).
2. Install AWS SDK for Go v2: Run `go get github.com/aws/aws-sdk-go-v2/config` and `go get github.com/aws/aws-sdk-go-v2/service/s3`.

## Usage

1. Open `main.go`.
2. Replace `your-s3-bucket-name` with the name of your S3 bucket.
3. Replace `your-substring` with the substring you want to search for.
4. Run the script with `go run main.go`.

## Notes

- Ensure you have the necessary AWS credentials configured in your environment.
- This script only scans `.txt` files.
