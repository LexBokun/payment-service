// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: enums/cryptocurrency/v1/crypto_currency.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Список доступных криптовалют(BTC, TON, ...)
type Currency int32

const (
	Currency_CURRENCY_CRYPTO_UNSPECIFIED Currency = 0
	Currency_CURRENCY_BTC                Currency = 1
	Currency_CURRENCY_TON                Currency = 2
	Currency_CURRENCY_USDT               Currency = 3
)

// Enum value maps for Currency.
var (
	Currency_name = map[int32]string{
		0: "CURRENCY_CRYPTO_UNSPECIFIED",
		1: "CURRENCY_BTC",
		2: "CURRENCY_TON",
		3: "CURRENCY_USDT",
	}
	Currency_value = map[string]int32{
		"CURRENCY_CRYPTO_UNSPECIFIED": 0,
		"CURRENCY_BTC":                1,
		"CURRENCY_TON":                2,
		"CURRENCY_USDT":               3,
	}
)

func (x Currency) Enum() *Currency {
	p := new(Currency)
	*p = x
	return p
}

func (x Currency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currency) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_cryptocurrency_v1_crypto_currency_proto_enumTypes[0].Descriptor()
}

func (Currency) Type() protoreflect.EnumType {
	return &file_enums_cryptocurrency_v1_crypto_currency_proto_enumTypes[0]
}

func (x Currency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currency.Descriptor instead.
func (Currency) EnumDescriptor() ([]byte, []int) {
	return file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescGZIP(), []int{0}
}

var File_enums_cryptocurrency_v1_crypto_currency_proto protoreflect.FileDescriptor

const file_enums_cryptocurrency_v1_crypto_currency_proto_rawDesc = "" +
	"\n" +
	"-enums/cryptocurrency/v1/crypto_currency.proto\x12\x17enums.cryptocurrency.v1*b\n" +
	"\bCurrency\x12\x1f\n" +
	"\x1bCURRENCY_CRYPTO_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fCURRENCY_BTC\x10\x01\x12\x10\n" +
	"\fCURRENCY_TON\x10\x02\x12\x11\n" +
	"\rCURRENCY_USDT\x10\x03BIZGgithub.com/LexBokun/transaction-service/pkg/api/enums/cryptocurrency/v1b\x06proto3"

var (
	file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescOnce sync.Once
	file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescData []byte
)

func file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescGZIP() []byte {
	file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescOnce.Do(func() {
		file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_enums_cryptocurrency_v1_crypto_currency_proto_rawDesc), len(file_enums_cryptocurrency_v1_crypto_currency_proto_rawDesc)))
	})
	return file_enums_cryptocurrency_v1_crypto_currency_proto_rawDescData
}

var file_enums_cryptocurrency_v1_crypto_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_cryptocurrency_v1_crypto_currency_proto_goTypes = []any{
	(Currency)(0), // 0: enums.cryptocurrency.v1.Currency
}
var file_enums_cryptocurrency_v1_crypto_currency_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_cryptocurrency_v1_crypto_currency_proto_init() }
func file_enums_cryptocurrency_v1_crypto_currency_proto_init() {
	if File_enums_cryptocurrency_v1_crypto_currency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_enums_cryptocurrency_v1_crypto_currency_proto_rawDesc), len(file_enums_cryptocurrency_v1_crypto_currency_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_cryptocurrency_v1_crypto_currency_proto_goTypes,
		DependencyIndexes: file_enums_cryptocurrency_v1_crypto_currency_proto_depIdxs,
		EnumInfos:         file_enums_cryptocurrency_v1_crypto_currency_proto_enumTypes,
	}.Build()
	File_enums_cryptocurrency_v1_crypto_currency_proto = out.File
	file_enums_cryptocurrency_v1_crypto_currency_proto_goTypes = nil
	file_enums_cryptocurrency_v1_crypto_currency_proto_depIdxs = nil
}
