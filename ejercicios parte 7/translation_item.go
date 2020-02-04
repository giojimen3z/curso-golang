package item

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mercadolibre/fury_items-translation-api/src/api/app/requester/infrastructure/adapter/restclient"
	"github.com/mercadolibre/fury_items-translation-api/src/api/app/requester/infrastructure/domain/model/items"

	"github.com/mercadolibre/go-meli-toolkit/goutils/apierrors"
)

//Post to get  the trantaleed text
func PostAwsTranslateApi(jsonLanguage string) (string, apierrors.ApiError) {

	restClient := &restclient.RestClientDefault
	var translate items.Translation
	jsonConvert := []byte(jsonLanguage)
	languages := ""
	var bodyPost string

	json.Unmarshal(jsonConvert, &translate)

	bodyPost = translate.TitleToTranslate
	for _, v := range translate.TranslateTo {
		languages += v.Language + ","
	}

	languages = strings.TrimSuffix(languages, ",")

	url := fmt.Sprint(os.Getenv("POST_TRANSLATION_API"), languages, os.Getenv("TRANSLATION_CLIENT_ID"))

	translation := restClient.Post(url, bodyPost)

	return translation.String(), nil

}
