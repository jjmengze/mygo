package observer

import (
	"github.com/jjmengze/mygo/pkg/design-pattern/behavioral/observer/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

// define mock type
type SubjecterMock struct {
	mock.Mock
}

//func TestNewReader(t *testing.T) {
//	type args struct {
//		name string
//	}
//	tests := []struct {
//		name string
//		args args
//		want *Reader
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewReader(tt.args.name); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewReader() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestNewSubject(t *testing.T) {
//	tests := []struct {
//		name string
//		want *Subject
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewSubject(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewSubject() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestReader_Update(t *testing.T) {
//	type fields struct {
//		name string
//	}
//	type args struct {
//		msg string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			r := &Reader{
//				name: tt.fields.name,
//			}
//		})
//	}
//}

//func TestSubject_Attach(t *testing.T) {
//	type fields struct {
//		observers []Observer
//	}
//	type args struct {
//		o Observer
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Subject{
//				observers: tt.fields.observers,
//			}
//		})
//	}
//}

//func TestSubject_UpdateMsg(t *testing.T) {
//	s := &Subject{}
//	type fields struct {
//		observers []Observer
//	}
//	type args struct {
//		msg string
//	}
//	tests := []struct {
//		name   string
//		args   args
//	}{
//		{
//			name:   "",
//			args:   args{},
//		},
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s.Attach()
//			s.UpdateMsg(tt.args.msg)
//		})
//	}
//}

//func TestSubject_notify(t *testing.T) {
//	type fields struct {
//		observers []Observer
//	}
//	type args struct {
//		msg string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Subject{
//				observers: tt.fields.observers,
//			}
//		})
//	}
//}

func TestSubject_UpdateMsg(t *testing.T) {
	mockObserver := &mocks.Observer{}
	mockObserver.On("Update", mock.Anything)
	type fields struct {
		observers []Observer
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				[]Observer{mockObserver},
			},
			args: args{
				msg: "hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{
				observers: tt.fields.observers,
			}
			s.UpdateMsg(tt.args.msg)
			mockObserver.AssertExpectations(t)
		})
	}
}

func TestSubject_Attach(t *testing.T) {
	mockObserver := &mocks.Observer{}
	type fields struct {
		observers []Observer
	}
	type args struct {
		o Observer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{observers: []Observer{mockObserver}},
			args:   args{mockObserver},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject{}
			s.Attach(tt.args.o)
			if ok := reflect.DeepEqual(s.observers, tt.fields.observers); !ok {
				t.Errorf("Happy() = %v, want %v", s.observers, tt.fields.observers)
			}
		})
	}
}
