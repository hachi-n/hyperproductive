package hyperproductive

type HyperProductiveGroup struct {
	administrator *Administrator
}

func NewHyperProductiveGroup(numberOfWorker int, command interface{}, params ...interface{}) *HyperProductiveGroup {
	hyperProductiveGroup := new(HyperProductiveGroup)
	hyperProductiveGroup.administrator = NewAdministrator(numberOfWorker, command, params)
	return hyperProductiveGroup
}

func (h *HyperProductiveGroup) NotReportOperate() {
	h.administrator.TrustOrder()
}

func (h *HyperProductiveGroup) IndividualOperate() []interface{} {
	return h.administrator.ExpectOrder()
}

func (h *HyperProductiveGroup) PrudentOperate() []interface{} {
	return h.administrator.PrudentOrder()
}
