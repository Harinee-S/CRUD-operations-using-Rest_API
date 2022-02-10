package construct

type Users struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"num"`
	PromoCode   string `json:"promo_code"`
	Reference   string `json:"refer"`
}
