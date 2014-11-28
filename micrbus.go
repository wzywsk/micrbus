package micrbus

import (
	"easy/bus"
	"easy/esc"
	"errors"
	"sync"
)

type Micrbus struct {
	mu        sync.Mutex
	readpoint map[string]interface{}
}

func NewMicrbus() (micrbus *Micrbus) {
	micrbus = new(Micrbus)
	micrbus.readpoint = make(map[string]interface{})
	return
}

//增加读节点
func (this *Micrbus) Add(name string, value interface{}) (err error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if _, ok := this.readpoint[name]; !ok {
		this.readpoint[name] = value
		err = nil
	} else {
		err = errors.New("读节点值已经存在")
	}
	return
}

//获取写节点的值
func (this *Micrbus) Get(name string) (value interface{}, err error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if nodevalue, ok := this.readpoint[name]; ok {
		value = nodevalue
		err = nil
	} else {
		err = errors.New("读节点值未找到")
	}
	return
}

//设置读节点值
func (this *Micrbus) Set(name string, value interface{}) (err error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	this.readpoint[name] = value
	return nil
}

//读取所有读节点的值
func (this *Micrbus) GetAllName(name []string) {
	i := 0
	for newname := range this.readpoint {
		name[i] = newname
		i++
	}
	return
}

type NodeInf struct {
	m *Micrbus
}

func NewNodeInf() *NodeInf {
	return &NodeInf{
		m: NewMicrbus(),
	}
}
func (n *NodeInf) GetName() string {
	return ""
}

func (n *NodeInf) Count(context *esc.EsContext) int {
	return 1
}

func (n *NodeInf) Echo(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Add(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Adds(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVars []*bus.EsTagVar) (tagRets []*bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Get(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tag string) (tagRet *bus.EsTagRet, ret bus.OpRet) {

	tagRet = new(bus.EsTagRet)
	tagRet.Tag = tag
	value, _ := n.m.Get(tag)
	tagRet.Value = value.(string)
	ret.Ok = true
	ret.Err = nil
	return
}

func (n *NodeInf) Gets(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tags []string) (tagRets []*bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Set(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	n.m.Set(tagVar.Tag, tagVar.Value)
	tagRet = new(bus.EsTagRet)
	tagRet.Tag = tagVar.Tag
	tagRet.Value = tagVar.Value
	ret.Ok = true
	ret.Err = nil
	return
}

func (n *NodeInf) Sets(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVars []*bus.EsTagVar) (tagRets []*bus.EsTagRet, ret bus.OpRet) {
	for _, value := range tagVars {
		n.m.Set(value.Tag, value.Value)
	}
	ret.Ok = true
	ret.Err = nil
	return
}

func (n *NodeInf) SetV(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Inc(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) IncStep(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Dec(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) DecStep(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Del(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVar *bus.EsTagVar) (tagRet *bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) Dels(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, tagVars []*bus.EsTagVar) (tagRets []*bus.EsTagRet, ret bus.OpRet) {
	return
}

func (n *NodeInf) NodeSvrConnected(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, svrConnCnt int32) (ret bus.OpRet) {
	return
}

func (n *NodeInf) NodeSvrDisconnected(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, svrConnCnt int32) (ret bus.OpRet) {
	return
}

func (n *NodeInf) NodeCltConnected(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, cltConnCnt int32) (ret bus.OpRet) {
	return
}

func (n *NodeInf) NodeCltDisconnected(context *esc.EsContext, adpter *bus.EsAdpterInf,
	node *bus.EsArgNode, cltConnCnt int32) (ret bus.OpRet) {
	return
}
