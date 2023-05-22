package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-fin-inst-account-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) finInstCodeExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	if &input.Header.FinInstCountry == nil ||
		&input.Header.FinInstCode == nil {
		*exconfErrMsg = "cannot specify null keys"
		return
	}
	res, err := c.finInstMasterExistenceConfRequest(*input.Header.FinInstCountry, *input.Header.FinInstCode, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}

func (c *ExistenceConf) finInstMasterExistenceConfRequest(finInstCountry string, finInstCode string, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	key := "FinInstMaster"
	keys := newResult(map[string]interface{}{
		"FinInstCountry": finInstCountry,
		"FinInstCode":    finInstCode,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	data := &FinInstMasterReturn{
		FinInstCountry: finInstCountry,
		FinInstCode:    finInstCode,
	}
	req.FinInstMasterReturn = data

	exist, err = c.exconfRequest(req, key, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}
