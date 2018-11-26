// Code generated by thriftrw v1.14.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ThriftTest_TestSet_Args represents the arguments for the ThriftTest.testSet function.
//
// The arguments for testSet are sent and received over the wire as this struct.
type ThriftTest_TestSet_Args struct {
	Thing map[int32]struct{} `json:"thing,omitempty"`
}

// ToWire translates a ThriftTest_TestSet_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestSet_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Thing != nil {
		w, err = wire.NewValueSet(_Set_I32_ValueList(v.Thing)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestSet_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestSet_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestSet_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestSet_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TSet {
				v.Thing, err = _Set_I32_Read(field.Value.GetSet())
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestSet_Args
// struct.
func (v *ThriftTest_TestSet_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestSet_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestSet_Args match the
// provided ThriftTest_TestSet_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestSet_Args) Equals(rhs *ThriftTest_TestSet_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Thing == nil && rhs.Thing == nil) || (v.Thing != nil && rhs.Thing != nil && _Set_I32_Equals(v.Thing, rhs.Thing))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestSet_Args.
func (v *ThriftTest_TestSet_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Thing != nil {
		err = multierr.Append(err, enc.AddArray("thing", (_Set_I32_Zapper)(v.Thing)))
	}
	return err
}

// GetThing returns the value of Thing if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestSet_Args) GetThing() (o map[int32]struct{}) {
	if v.Thing != nil {
		return v.Thing
	}

	return
}

// IsSetThing returns true if Thing is not nil.
func (v *ThriftTest_TestSet_Args) IsSetThing() bool {
	return v.Thing != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testSet" for this struct.
func (v *ThriftTest_TestSet_Args) MethodName() string {
	return "testSet"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *ThriftTest_TestSet_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// ThriftTest_TestSet_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testSet
// function.
var ThriftTest_TestSet_Helper = struct {
	// Args accepts the parameters of testSet in-order and returns
	// the arguments struct for the function.
	Args func(
		thing map[int32]struct{},
	) *ThriftTest_TestSet_Args

	// IsException returns true if the given error can be thrown
	// by testSet.
	//
	// An error can be thrown by testSet only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for testSet
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// testSet into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by testSet
	//
	//   value, err := testSet(args)
	//   result, err := ThriftTest_TestSet_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from testSet: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(map[int32]struct{}, error) (*ThriftTest_TestSet_Result, error)

	// UnwrapResponse takes the result struct for testSet
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if testSet threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := ThriftTest_TestSet_Helper.UnwrapResponse(result)
	UnwrapResponse func(*ThriftTest_TestSet_Result) (map[int32]struct{}, error)
}{}

func init() {
	ThriftTest_TestSet_Helper.Args = func(
		thing map[int32]struct{},
	) *ThriftTest_TestSet_Args {
		return &ThriftTest_TestSet_Args{
			Thing: thing,
		}
	}

	ThriftTest_TestSet_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	ThriftTest_TestSet_Helper.WrapResponse = func(success map[int32]struct{}, err error) (*ThriftTest_TestSet_Result, error) {
		if err == nil {
			return &ThriftTest_TestSet_Result{Success: success}, nil
		}

		return nil, err
	}
	ThriftTest_TestSet_Helper.UnwrapResponse = func(result *ThriftTest_TestSet_Result) (success map[int32]struct{}, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// ThriftTest_TestSet_Result represents the result of a ThriftTest.testSet function call.
//
// The result of a testSet execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type ThriftTest_TestSet_Result struct {
	// Value returned by testSet after a successful execution.
	Success map[int32]struct{} `json:"success,omitempty"`
}

// ToWire translates a ThriftTest_TestSet_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestSet_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueSet(_Set_I32_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestSet_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestSet_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestSet_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestSet_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestSet_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TSet {
				v.Success, err = _Set_I32_Read(field.Value.GetSet())
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestSet_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestSet_Result
// struct.
func (v *ThriftTest_TestSet_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("ThriftTest_TestSet_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestSet_Result match the
// provided ThriftTest_TestSet_Result.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestSet_Result) Equals(rhs *ThriftTest_TestSet_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Set_I32_Equals(v.Success, rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ThriftTest_TestSet_Result.
func (v *ThriftTest_TestSet_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddArray("success", (_Set_I32_Zapper)(v.Success)))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestSet_Result) GetSuccess() (o map[int32]struct{}) {
	if v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *ThriftTest_TestSet_Result) IsSetSuccess() bool {
	return v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "testSet" for this struct.
func (v *ThriftTest_TestSet_Result) MethodName() string {
	return "testSet"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *ThriftTest_TestSet_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
