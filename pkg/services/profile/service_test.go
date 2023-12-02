package profile

import (
	"context"
	"gorm.io/gorm"
	pb "micromango/pkg/grpc/profile"
	"micromango/pkg/grpc/static"
	"reflect"
	"testing"
)

func Test_service_AddToList(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		in0 context.Context
		req *pb.AddToListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.AddToList(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddToList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddToList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateProfile(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		in0 context.Context
		req *pb.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.CreateProfile(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Get(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		in0 context.Context
		req *pb.GetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.Get(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetList(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		in0 context.Context
		req *pb.GetListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.GetList(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_RemoveFromList(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		in0 context.Context
		req *pb.RemoveFromListRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.RemoveFromList(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveFromList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFromList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateProfile(t *testing.T) {
	type fields struct {
		UnimplementedProfileServer pb.UnimplementedProfileServer
		db                         *gorm.DB
		static                     static.StaticClient
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				UnimplementedProfileServer: tt.fields.UnimplementedProfileServer,
				db:                         tt.fields.db,
				static:                     tt.fields.static,
			}
			got, err := s.UpdateProfile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
