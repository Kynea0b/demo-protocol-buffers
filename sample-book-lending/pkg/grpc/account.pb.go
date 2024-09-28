// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.3
// source: account.proto

// packageの宣言

package grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 図書館のアカウント
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 借りる人の名前
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 貸し出し状況
type AccountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 借りた人の名前
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 借りている本一覧
	Books *Books `protobuf:"bytes,2,opt,name=books,proto3" json:"books,omitempty"`
}

func (x *AccountInfo) Reset() {
	*x = AccountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountInfo) ProtoMessage() {}

func (x *AccountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountInfo.ProtoReflect.Descriptor instead.
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountInfo) GetBooks() *Books {
	if x != nil {
		return x.Books
	}
	return nil
}

// 貸し出しリクエスト
type BorrowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 　本を借りるアカウント
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// 　借りたい本
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *BorrowRequest) Reset() {
	*x = BorrowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowRequest) ProtoMessage() {}

func (x *BorrowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowRequest.ProtoReflect.Descriptor instead.
func (*BorrowRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *BorrowRequest) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *BorrowRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

type BorrrowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 　本を借りたアカウント
	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// 　借りた本
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *BorrrowResponse) Reset() {
	*x = BorrrowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrrowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrrowResponse) ProtoMessage() {}

func (x *BorrrowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrrowResponse.ProtoReflect.Descriptor instead.
func (*BorrrowResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *BorrrowResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *BorrrowResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1d, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3f, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x22, 0x54, 0x0a, 0x0d, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x56, 0x0a, 0x0f, 0x42, 0x6f, 0x72, 0x72, 0x72, 0x6f,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x79, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0xac,
	0x01, 0x0a, 0x13, 0x4c, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x12, 0x14, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x42, 0x6f, 0x72,
	0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x79, 0x61,
	0x70, 0x70, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x79, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_account_proto_goTypes = []interface{}{
	(*Account)(nil),         // 0: myapp.Account
	(*AccountInfo)(nil),     // 1: myapp.AccountInfo
	(*BorrowRequest)(nil),   // 2: myapp.BorrowRequest
	(*BorrrowResponse)(nil), // 3: myapp.BorrrowResponse
	(*Books)(nil),           // 4: Books
	(*Book)(nil),            // 5: Book
}
var file_account_proto_depIdxs = []int32{
	4, // 0: myapp.AccountInfo.books:type_name -> Books
	0, // 1: myapp.BorrowRequest.account:type_name -> myapp.Account
	5, // 2: myapp.BorrowRequest.book:type_name -> Book
	0, // 3: myapp.BorrrowResponse.account:type_name -> myapp.Account
	5, // 4: myapp.BorrrowResponse.book:type_name -> Book
	2, // 5: myapp.LendingBooksService.SendBorrow:input_type -> myapp.BorrowRequest
	0, // 6: myapp.LendingBooksService.ShowAccountInfo:input_type -> myapp.Account
	3, // 7: myapp.LendingBooksService.SendBorrow:output_type -> myapp.BorrrowResponse
	1, // 8: myapp.LendingBooksService.ShowAccountInfo:output_type -> myapp.AccountInfo
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_book_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrrowResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}
