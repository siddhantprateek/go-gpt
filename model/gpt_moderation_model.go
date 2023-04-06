package model

type GPTModerationModel struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type GPTModerationModelResponse struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Results []struct {
		Flagged    bool `json:"flagged"`
		Categories struct {
			Sexual          bool `json:"sexual"`
			Hate            bool `json:"hate"`
			Violence        bool `json:"violence"`
			SelfHarm        bool `json:"self-harm"`
			SexualMinors    bool `json:"sexual/minors"`
			HateThreatening bool `json:"hate/threatening"`
			ViolenceGraphic bool `json:"violence/graphic"`
		} `json:"categories"`
		CategoryScores struct {
			Sexual          float64 `json:"sexual"`
			Hate            float64 `json:"hate"`
			Violence        float64 `json:"violence"`
			SelfHarm        float64 `json:"self-harm"`
			SexualMinors    float64 `json:"sexual/minors"`
			HateThreatening float64 `json:"hate/threatening"`
			ViolenceGraphic float64 `json:"violence/graphic"`
		} `json:"category_scores"`
	} `json:"results"`
}
