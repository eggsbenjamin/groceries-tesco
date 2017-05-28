package tesco

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/eggsbenjamin/groceries-tesco/domain"
	"github.com/spf13/viper"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type noopCloser struct {
	io.Reader
}

func (n *noopCloser) Close() error {
	return nil
}

func CreateMockResponse(sc int, bStr string) *http.Response {
	return &http.Response{
		StatusCode: sc,
		Body:       &noopCloser{bytes.NewBufferString(bStr)},
	}
}

type TescoService struct {
	client HTTPClient
}

//	Constructor
func NewTescoService(client HTTPClient) *TescoService {
	return &TescoService{
		client: client,
	}
}

type ProductsResponse struct {
	Products []*domain.Product
}

func (t *TescoService) GetProductsByGTIN(gtin string) ([]*domain.Product, error) {
	url := fmt.Sprintf(
		"%s/product/?gtin=%s",
		viper.GetString("base_url"),
		gtin,
	)
	rq, err := http.NewRequest("GET", url, nil)
	rq.Header.Add(
		viper.GetString("auth_header_name"),
		viper.GetString("api_key"),
	)
	rErr := fmt.Errorf(`unable to get products with GTIN : '%s'`, gtin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create request - %v", err)
		return nil, rErr
	}
	rsp, err := t.client.Do(rq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to send request - %v", err)
		return nil, rErr
	}
	pRes := &ProductsResponse{}
	raw, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read response - %v", err)
		return nil, rErr
	}
	err = json.Unmarshal(raw, &pRes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to marshal response - %v\nresponse - %s", err, string(raw))
		return nil, rErr
	}
	return pRes.Products, nil
}
