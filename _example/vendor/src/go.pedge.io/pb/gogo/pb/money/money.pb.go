// Code generated by protoc-gen-gogo.
// source: pb/money/money.proto
// DO NOT EDIT!

package pbmoney

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Currency struct {
	Name        string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code        CurrencyCode `protobuf:"varint,2,opt,name=code,proto3,enum=pb.money.CurrencyCode" json:"code,omitempty"`
	NumericCode uint32       `protobuf:"varint,3,opt,name=numeric_code,proto3" json:"numeric_code,omitempty"`
	MinorUnit   uint32       `protobuf:"varint,4,opt,name=minor_unit,proto3" json:"minor_unit,omitempty"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}

type Money struct {
	CurrencyCode CurrencyCode `protobuf:"varint,1,opt,name=currency_code,proto3,enum=pb.money.CurrencyCode" json:"currency_code,omitempty"`
	ValueMicros  int64        `protobuf:"varint,2,opt,name=value_micros,proto3" json:"value_micros,omitempty"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}

type ExchangeRate struct {
	From        CurrencyCode `protobuf:"varint,1,opt,name=from,proto3,enum=pb.money.CurrencyCode" json:"from,omitempty"`
	To          CurrencyCode `protobuf:"varint,2,opt,name=to,proto3,enum=pb.money.CurrencyCode" json:"to,omitempty"`
	ValueMicros int64        `protobuf:"varint,3,opt,name=value_micros,proto3" json:"value_micros,omitempty"`
}

func (m *ExchangeRate) Reset()         { *m = ExchangeRate{} }
func (m *ExchangeRate) String() string { return proto.CompactTextString(m) }
func (*ExchangeRate) ProtoMessage()    {}

func init() {
	proto.RegisterType((*Currency)(nil), "pb.money.Currency")
	proto.RegisterType((*Money)(nil), "pb.money.Money")
	proto.RegisterType((*ExchangeRate)(nil), "pb.money.ExchangeRate")
}
