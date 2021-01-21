package model

type RespGetFreebillAuth struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

type RespSetFreeBillRegister struct {
    Data string `json:"data"`
}

type RespGetFreeBillSearch struct {
    Data string `json:"data"`
}

type RespGetFreeBillSearchStatus struct {
    Data string `json:"data"`
}

type RespDelFreeBill struct {
    Data string `json:"data"`
}

type RespGetFreeBillView struct {
    Data string `json:"data"`
}

type RespDelFreeBillCancel struct {
    Data string `json:"data"`
}

type RespGetFreeBillPreviousRegister struct {
    Data string `json:"data"`
}

type RespSetFreeBillPublishNow struct {
    Data string `json:"data"`
}

type RespSetFreeBillSend struct {
    Data string `json:"data"`
}