package construct

type JSONResponse struct {
	Type    string  `json:"type"`
	Data    []Users `json:"data"`
	Message string  `json:"message"`
}
