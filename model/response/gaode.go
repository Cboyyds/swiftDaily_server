package response

type IPResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	InfoCode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}
