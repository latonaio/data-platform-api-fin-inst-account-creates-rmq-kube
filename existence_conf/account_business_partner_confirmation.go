package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-fin-inst-account-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *ExistenceConf) accountBusinessPartnerExistenceConf(input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	if input.Header.AccountBusinessPartner == nil {
		*exconfErrMsg = "cannot specify null keys"
		return
	}
	res, err := c.bpGeneralExistenceConfRequest(*input.Header.AccountBusinessPartner, input, existenceMap, mtx, log)
	if err != nil {
		mtx.Lock()
		*errs = append(*errs, err)
		mtx.Unlock()
	}
	if res != "" {
		*exconfErrMsg = res
	}
}
