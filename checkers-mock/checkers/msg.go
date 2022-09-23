package checkers

type MsgCreateGame struct {
	Creator string
	Red     string
	Black   string
}

type MsgCreateGameResponse struct {
	IdValue string
}
