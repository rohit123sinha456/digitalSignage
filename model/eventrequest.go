package DataModel

type EventStreamRequest struct {
	Screencode string               `form:"screencode" json:"message" binding:"required,max=100"`
	Userinfo   UserSystemIdentifeir `json:"userinfo"`
}
