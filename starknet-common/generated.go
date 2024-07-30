//go:build tinygo || wasip1

package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"

	pbstarknet "github.com/streamingfast/substreams-foundational-modules/starknet-common/pb/sf/starknet/type/v1"
	v1 "github.com/streamingfast/substreams-foundational-modules/starknet-common/pb/sf/substreams/starknet/type/v1"
)

//go:generate substreams-starknet protogen ./substreams-starknet.yaml --with-tinygo-maps // creates genre substreams-starknet.gen.go

// Dans WASI: _start
func main() {}

func panic(a any) {
	os.Exit(2) //fail-safe so we know someone call panic
}

//export output
func _output(ptr, len int32)

//go:wasm-module logger
//export println
func _log(ptr int32, len int32)

// Log a line to the Substreams engine
func Logf(message string, args ...any) {
	_log(stringToPtr(fmt.Sprintf(message, args...)))
}

// Output the serialized protobuf byte array to the Substreams engine
func output(out []byte) {
	_output(byteArrayToPtr(out))
}

type OutputError struct {
	msg string
}

func (o OutputError) Error() string {
	return o.msg
}

func outputVT(out vtMessage) error {
	if out == nil || reflect.ValueOf(out).IsNil() {
		return nil
	}
	b, err := out.MarshalVT()
	if err != nil {
		return fmt.Errorf("marshalling output: %w", err)
	}
	if len(b) == 0 {
		return nil
	}
	output(b)
	return nil
}

//export all_transactions
func _all_transactions(blockPtr, blockLen int32) (retval int32) {
	block := &pbstarknet.Block{}
	if err := unmarshalVT(blockPtr, blockLen, block); err != nil {
		Logf("failed unmarshal block: %s", err)
		return 1
	}

	output, err := AllTransactions(block)
	if err != nil {
		Logf("executing all_transactions: %s", err)
		return 1
	}

	if err := outputVT(output); err != nil {
		Logf("outputing vt message: %s", err)
		return 1
	}

	return 0
}

//export index_transactions
func _index_transactions(ptr, len int32) (retval int32) {
	transactions := &v1.Transactions{}
	if err := unmarshalVT(ptr, len, transactions); err != nil {
		Logf("unmarshalling transactions: %s", err)
		return 1
	}

	output, err := IndexTransactions(transactions)
	if err != nil {
		Logf("executing index_transactions: %s", err)
		return 1

	}
	if err := outputVT(output); err != nil {
		Logf("outputing vt message: %s", err)
		return 1
	}
	return 0
}

//export filtered_transactions
func _filtered_transactions(queryPtr, queryLen int32, transactionsPtr, transactionsLen int32) (ret int32) {
	query := ptrToString(queryPtr, queryLen)

	txs := &v1.Transactions{}
	if err := unmarshalVT(transactionsPtr, transactionsLen, txs); err != nil {
		Logf("failed unmarshal transactions: %s", err)
		return 1
	}

	output, err := FilteredTransactions(query, txs)
	if err != nil {
		Logf("calling filtered_transactions: %s", err)
		return 1
	}

	if err := outputVT(output); err != nil {
		Logf("outputing vt message: %s", err)
		return 1
	}
	return 0
}

// ptrToString returns a string from WebAssembly compatible numeric types
// representing its pointer and length.
func ptrToString(ptr int32, size int32) string {
	// Get a slice view of the underlying bytes in the stream. We use SliceHeader, not StringHeader
	// as it allows us to fix the capacity to what was allocated.
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(size), // Tinygo requires these as uintptrs even if they are int fields.
		Cap:  int(size), // ^^ See https://github.com/tinygo-org/tinygo/issues/1284
	}))
}

func ptrToBytes(ptr int32, size int32) []byte {
	// Get a slice view of the underlying bytes in the stream. We use SliceHeader, not StringHeader
	// as it allows us to fix the capacity to what was allocated.
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  int(size), // Tinygo requires these as uintptrs even if they are int fields.
		Cap:  int(size), // ^^ See https://github.com/tinygo-org/tinygo/issues/1284
	}))
}

// stringToPtr returns a pointer and size pair for the given string in a way
// compatible with WebAssembly numeric types.
func stringToPtr(s string) (int32, int32) {
	buf := []byte(s)
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return int32(unsafePtr), int32(len(buf))
}

// byteArrayToPtr returns a pointer and size pair for the given byte array, for WASM compat.
func byteArrayToPtr(buf []byte) (int32, int32) {
	ptr := &buf[0]
	unsafePtr := uintptr(unsafe.Pointer(ptr))
	return int32(unsafePtr), int32(len(buf))
}

type vtMessage interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

func unmarshalVT(prt, len int32, dest vtMessage) error {
	b := ptrToBytes(prt, len)
	return dest.UnmarshalVT(b)
}