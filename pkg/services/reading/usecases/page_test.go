package usecases

import (
	"micromango/pkg/common"
	"micromango/pkg/services/reading/entity"
	"micromango/pkg/services/reading/mock"
	"reflect"
	"testing"
)

func TestPage_AddPage(t *testing.T) {
	type fields struct {
		Repository PageRepository
		Static     StaticService
	}
	type args struct {
		dto entity.AddPageDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes entity.Page
		wantErr bool
	}{
		{
			name: "Add page without image",
			fields: fields{
				Repository: &mock.PageRepository{},
				Static:     &mock.StaticServiceMock{},
			},
			args: args{dto: entity.AddPageDto{MangaId: "1", ChapterId: "2", Number: 3}},
			wantRes: entity.Page{
				PageId:    "0",
				MangaId:   "1",
				ChapterId: "2",
				Number:    3,
				Image:     "",
			},
			wantErr: false,
		},
		{
			name: "Add page with image",
			fields: fields{
				Repository: &mock.PageRepository{},
				Static:     &mock.StaticServiceMock{},
			},
			args: args{dto: entity.AddPageDto{MangaId: "1", ChapterId: "2", Number: 3, Image: &common.File{}}},
			wantRes: entity.Page{
				PageId:    "0",
				MangaId:   "1",
				ChapterId: "2",
				Number:    3,
				Image:     "path/to/file",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Repository: tt.fields.Repository,
				Static:     tt.fields.Static,
			}
			gotRes, err := p.AddPage(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("AddPage() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestPage_GetPage(t *testing.T) {
	type fields struct {
		Repository PageRepository
		Static     StaticService
	}
	type args struct {
		pageId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Page
		wantErr bool
	}{
		{
			name: "GetRate page 1",
			fields: fields{
				Repository: &mock.PageRepository{},
				Static:     &mock.StaticServiceMock{},
			},
			args:    args{pageId: "1"},
			want:    entity.Page{PageId: "1", Number: 1, ChapterId: "1", MangaId: "1"},
			wantErr: false,
		},
		{
			name: "GetRate page not found",
			fields: fields{
				Repository: &mock.PageRepository{},
				Static:     &mock.StaticServiceMock{},
			},
			args:    args{pageId: "0"},
			want:    entity.Page{},
			wantErr: true,
		},
		{
			name: "GetRate page 2",
			fields: fields{
				Repository: &mock.PageRepository{},
				Static:     &mock.StaticServiceMock{},
			},
			args:    args{pageId: "2"},
			want:    entity.Page{PageId: "2", Number: 1, ChapterId: "1", MangaId: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Repository: tt.fields.Repository,
				Static:     tt.fields.Static,
			}
			got, err := p.GetPage(tt.args.pageId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
