// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package zipkincore

import (
	"bytes"
	"fmt"
	"github.com/uber/jaeger-client-go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ZipkinCollector interface {
	// Parameters:
	//  - Spans
	SubmitZipkinBatch(spans []*Span) (r []*Response, err error)
}

type ZipkinCollectorClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewZipkinCollectorClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ZipkinCollectorClient {
	return &ZipkinCollectorClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewZipkinCollectorClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ZipkinCollectorClient {
	return &ZipkinCollectorClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Spans
func (p *ZipkinCollectorClient) SubmitZipkinBatch(spans []*Span) (r []*Response, err error) {
	if err = p.sendSubmitZipkinBatch(spans); err != nil {
		return
	}
	return p.recvSubmitZipkinBatch()
}

func (p *ZipkinCollectorClient) sendSubmitZipkinBatch(spans []*Span) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("submitZipkinBatch", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ZipkinCollectorSubmitZipkinBatchArgs{
		Spans: spans,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ZipkinCollectorClient) recvSubmitZipkinBatch() (value []*Response, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "submitZipkinBatch" {
		err = thrift.NewTApplicationException(thrift.WRONG_MSEVOD_NAME, "submitZipkinBatch failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "submitZipkinBatch failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "submitZipkinBatch failed: invalid message type")
		return
	}
	result := ZipkinCollectorSubmitZipkinBatchResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ZipkinCollectorProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ZipkinCollector
}

func (p *ZipkinCollectorProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ZipkinCollectorProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ZipkinCollectorProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewZipkinCollectorProcessor(handler ZipkinCollector) *ZipkinCollectorProcessor {

	self4 := &ZipkinCollectorProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["submitZipkinBatch"] = &zipkinCollectorProcessorSubmitZipkinBatch{handler: handler}
	return self4
}

func (p *ZipkinCollectorProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_MSEVOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type zipkinCollectorProcessorSubmitZipkinBatch struct {
	handler ZipkinCollector
}

func (p *zipkinCollectorProcessorSubmitZipkinBatch) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ZipkinCollectorSubmitZipkinBatchArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("submitZipkinBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ZipkinCollectorSubmitZipkinBatchResult{}
	var retval []*Response
	var err2 error
	if retval, err2 = p.handler.SubmitZipkinBatch(args.Spans); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing submitZipkinBatch: "+err2.Error())
		oprot.WriteMessageBegin("submitZipkinBatch", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("submitZipkinBatch", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Spans
type ZipkinCollectorSubmitZipkinBatchArgs struct {
	Spans []*Span `thrift:"spans,1" json:"spans"`
}

func NewZipkinCollectorSubmitZipkinBatchArgs() *ZipkinCollectorSubmitZipkinBatchArgs {
	return &ZipkinCollectorSubmitZipkinBatchArgs{}
}

func (p *ZipkinCollectorSubmitZipkinBatchArgs) GetSpans() []*Span {
	return p.Spans
}
func (p *ZipkinCollectorSubmitZipkinBatchArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Span, 0, size)
	p.Spans = tSlice
	for i := 0; i < size; i++ {
		_elem6 := &Span{}
		if err := _elem6.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem6), err)
		}
		p.Spans = append(p.Spans, _elem6)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submitZipkinBatch_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("spans", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:spans: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Spans)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Spans {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:spans: ", p), err)
	}
	return err
}

func (p *ZipkinCollectorSubmitZipkinBatchArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ZipkinCollectorSubmitZipkinBatchArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ZipkinCollectorSubmitZipkinBatchResult struct {
	Success []*Response `thrift:"success,0" json:"success,omitempty"`
}

func NewZipkinCollectorSubmitZipkinBatchResult() *ZipkinCollectorSubmitZipkinBatchResult {
	return &ZipkinCollectorSubmitZipkinBatchResult{}
}

var ZipkinCollectorSubmitZipkinBatchResult_Success_DEFAULT []*Response

func (p *ZipkinCollectorSubmitZipkinBatchResult) GetSuccess() []*Response {
	return p.Success
}
func (p *ZipkinCollectorSubmitZipkinBatchResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ZipkinCollectorSubmitZipkinBatchResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchResult) readField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*Response, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem7 := &Response{}
		if err := _elem7.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem7), err)
		}
		p.Success = append(p.Success, _elem7)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("submitZipkinBatch_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ZipkinCollectorSubmitZipkinBatchResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ZipkinCollectorSubmitZipkinBatchResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ZipkinCollectorSubmitZipkinBatchResult(%+v)", *p)
}
