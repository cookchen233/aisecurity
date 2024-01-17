package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"
)

func FilterZerosFromArray[T any](input []T) []T {
	var result []T
	for _, value := range input {
		// Handle different types using type assertions and reflection
		switch v := any(value).(type) {
		case int:
			if v != 0 {
				result = append(result, value)
			}
		case string:
			if v != "" {
				result = append(result, value)
			}
		default:
			// For types that are not int or string, use reflection
			// to check for the zero value of the type
			if !reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface()) {
				result = append(result, value)
			}
		}
	}
	return result
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// ErrorWrap adds a message to the error and ensures the stack trace is only added once.
func ErrorWrap(err error, message string) error {
	// Check if the error has a stack trace
	if _, ok := err.(stackTracer); ok {
		// The error already has a stack trace; add only the message.
		return errors.WithMessage(err, message)
	}

	// The error does not have a stack trace; add both message and stack trace.
	return errors.Wrap(err, message)
}

// ErrorWithStack adds a stack trace to the error if it doesn't have one already.
func ErrorWithStack(err error) error {
	// Check if the error has a stack trace
	if _, ok := err.(stackTracer); ok {
		// Error already has a stack trace, return it as is
		return err
	}

	// Error does not have a stack trace, add it
	return errors.WithStack(err)
}

func GetMethodName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func ReadFileBuffer(fh *multipart.FileHeader) ([]byte, error) {
	file, err := fh.Open()
	if err != nil {
		return nil, ErrorWrap(err, "文件识别发生错误")
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			Logger.Error("文件关闭时发生错误", zap.Error(err))
		}
	}(file)
	// Read the first 512 bytes to determine the file type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, ErrorWrap(err, "读取文件失败")
	}
	// Reset the file pointer
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, ErrorWrap(err, "文件重置失败")
	}
	return buffer, nil
}

func ProcessJSON(data []byte) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return processElement(result), nil
}

func processElement(element interface{}) interface{} {
	switch element := element.(type) {
	case map[string]interface{}:
		for key, value := range element {
			element[key] = processElement(value)
		}
	case []interface{}:
		for i, value := range element {
			element[i] = processElement(value)
		}
	case string:
		if len(element) > 64 {
			return element[:64]
		}
	}
	return element
}

func getFileNameWithoutSuffix(filePath string) string {
	// Use filepath.Base to get the base name of the file
	baseName := filepath.Base(filePath)

	// Use filepath.Ext to get the file extension
	fileExt := filepath.Ext(baseName)

	// Remove the extension from the base name
	fileNameWithoutSuffix := baseName[:len(baseName)-len(fileExt)]

	return fileNameWithoutSuffix
}

func Base64ToImage(basePath string, ImageData string) (string, error) {
	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(ImageData)
	if err != nil {
		return "", ErrorWrap(err, "failed decoding base64")
	}

	// Create a new seeded source
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	// Get the current date in Y-m and Ymd format
	now := time.Now()
	dirFormat := now.Format("2006-01")
	fileDateFormat := now.Format("200601021504")

	// Generate a 6-digit random number
	randomNumber := rnd.Intn(1000000)

	// Create the directory name and check if it exists
	dirName := filepath.Join(basePath, dirFormat)
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		// Create the directory if it does not exist
		err := os.Mkdir(dirName, 0755) // 0755 permissions
		if err != nil {
			return "", ErrorWrap(err, "failed creating directory")
		}
	}

	// Form the file name
	name := fmt.Sprintf("%s-%06d.jpg", fileDateFormat, randomNumber)
	filename := filepath.Join(dirName, name)

	// Write data to file
	err = os.WriteFile(filename, data, 0644) // 0644 are the file permissions
	if err != nil {
		return "", ErrorWrap(err, "failed writing to file")
	}
	return filename, nil
}
