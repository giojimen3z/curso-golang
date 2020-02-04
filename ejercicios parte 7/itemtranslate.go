package items

type Translation struct {
	TitleToTranslate string `json:"title_to_translate"`
	TranslateTo      []struct {
		Language string `json:"language"`
	} `json:"translate_to"`
}
