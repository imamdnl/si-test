package source_product

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pasztorpisti/qs"
	"go.uber.org/zap"
	"io"
	"net/http"
	"si-test/pkg/config"
	"si-test/pkg/utils"
	"strings"
	"time"
)

type ApiName string

const (
	SourceProduct ApiName = "source-product"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
	logger = config.Logger()
)

func init() {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	Client = &http.Client{Transport: customTransport, Timeout: 30 * time.Second}
	logger = logger.With(zap.String("service-name", "lumen"))

	config.Environment()
}

// Get sends a get request to the URL
func Get(apiName string, url string, body interface{}) ([]byte, error) {
	requestURL := "http://127.0.0.1:10495" + url
	fmt.Println("isi requestUrl", requestURL)

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("Accept", "application/json")

	if body != nil {
		queryStr, err := qs.Marshal(body)
		if err != nil {
			return nil, errors.New("can not marshal query param")
		}
		requestURL = requestURL + "?" + queryStr
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, errors.New("can not marshal JSON")
	}
	request, err := http.NewRequest(http.MethodGet, strings.TrimSpace(requestURL), nil)
	if err != nil {
		return nil, err
	}

	request.Header = headers
	res, err := Client.Do(request)
	if err != nil {
		logger.Error("error get: ", zap.Error(err))
	}
	requestBody, _ := io.ReadAll(bytes.NewBuffer(jsonBytes))
	logger.Info("requestURL", zap.String("requestURL", strings.TrimSpace(requestURL)))
	logger.Info("requestBody", zap.String("requestBody", string(requestBody)))

	responseBody, err := ManageResponseAndError(apiName, res, err)
	if err != nil {
		logger.Info("responseBody", zap.String("responseBody", string(responseBody)))
		logger.Error("error get", zap.Error(err))
		return nil, err
	}
	logger.Info("responseBody", zap.String("responseBody", string(responseBody)))

	return responseBody, nil
}

func ExceptionAPI(a ApiName) bool {
	switch a {
	case SourceProduct:
		return true
	}
	return false
}
func ManageResponseAndError(apiName string, response *http.Response, err error) ([]byte, error) {
	// check request success/failure
	if err != nil {
		return nil, errors.New(utils.ClassifyNetworkError(err))
	}

	// read response body
	res, err := io.ReadAll(response.Body)
	// leverage defer stack to defer closing of response body read operation
	// this will defer until this function is ready to return
	config.Logger().Info("response form Lumencl: ", zap.Any("response lumen:", string(res)))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if err != nil {
		config.Logger().Error("error response from lumen: ", zap.Error(err))
		return res, errors.New("invalid response body")
	}
	// check response status
	if response.StatusCode == 403 {
		config.Logger().Error("error response from lumen status 403: ", zap.Error(err))
		return res, errors.New("not authorized")
	}

	if response.StatusCode > 299 {
		if ExceptionAPI(ApiName(apiName)) && response.StatusCode == 400 {
			return res, nil
		}
		config.Logger().Error("error response from lumen code > 299: ", zap.Error(err))
		return res, errors.New("Lumen internal error")
	}

	return res, nil
}
