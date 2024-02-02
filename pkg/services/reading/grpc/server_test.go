package grpc

import (
	"context"
	"github.com/golang/protobuf/proto"
	"micromango/pkg/common/utils"
	"micromango/pkg/grpc/reading"
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading/mock"
	"micromango/pkg/services/reading/usecases"
	"testing"
)

// Mock chapter use case
var ChapterCaseMock = usecases.Chapter{
	Repository: &mock.ChapterRepository{},
	Activity:   &mock.ActivityServiceMock{},
}

var PageCaseMock = usecases.Page{
	Repository: &mock.PageRepository{},
	Static:     &mock.StaticServiceMock{},
}

func TestServer_AddChapter(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.AddChapterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ChapterResponse
		wantErr bool
	}{
		{
			name: "Add chapter",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.AddChapterRequest{
					MangaId: "1",
					Number:  1,
					Title:   "test",
				},
			},
			want: &pb.ChapterResponse{
				ChapterId: "1",
				MangaId:   "1",
				Number:    1,
				Title:     "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.AddChapter(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("AddChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_AddPage(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.AddPageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.PageResponse
		wantErr bool
	}{
		{
			name: "Add page",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.AddPageRequest{
					ChapterId: "1",
					Number:    1,
				},
			},
			want: &pb.PageResponse{
				PageId:    "0",
				ChapterId: "1",
				Number:    1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.AddPage(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("AddPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetChapter(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.ChapterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ChapterResponse
		wantErr bool
	}{
		{
			name: "GetRate chapter",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.ChapterRequest{
					ChapterId: "1",
				},
			},
			want: &pb.ChapterResponse{
				ChapterId: "1",
				MangaId:   "1",
				Number:    1,
				Title:     "title",
				Pages:     []*pb.ChapterResponse_PageHead{{PageId: "1", Number: 1}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.GetChapter(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("GetChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetMangaContent(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.MangaContentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.MangaContentResponse
		wantErr bool
	}{
		{
			name: "GetRate manga content",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.MangaContentRequest{
					MangaId: "1",
				},
			},
			want: &pb.MangaContentResponse{
				//MangaId: "1",
				Chapters: []*pb.MangaContentResponse_ChapterHead{
					{ChapterId: "1", Number: 1, Title: "title", Pages: 1, CreatedAt: "0001-01-01 00:00:00 +0000 UTC"},
					{ChapterId: "2", Number: 2, Title: "title2", Pages: 1, CreatedAt: "0001-01-01 00:00:00 +0000 UTC"},
					{ChapterId: "3", Number: 3, Title: "title3", Pages: 1, CreatedAt: "0001-01-01 00:00:00 +0000 UTC"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.GetMangaContent(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMangaContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("GetMangaContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetPage(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.PageRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.PageResponse
		wantErr bool
	}{
		{
			name: "GetRate page",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.PageRequest{
					PageId: "1",
				},
			},
			want: &pb.PageResponse{
				PageId:    "1",
				ChapterId: "1",
				Number:    1,
			},
			wantErr: false,
		},
		{
			name: "Page not found",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.PageRequest{
					PageId: "0",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.GetPage(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("GetPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateChapter(t *testing.T) {
	type fields struct {
		ChapterCase                usecases.Chapter
		PageCase                   usecases.Page
		UnimplementedReadingServer reading.UnimplementedReadingServer
	}
	type args struct {
		in0 context.Context
		req *pb.UpdateChapterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ChapterResponse
		wantErr bool
	}{
		{
			name: "Update chapter",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.UpdateChapterRequest{
					ChapterId: "1",
					Number:    utils.Ptr[float32](1),
					Title:     utils.Ptr("test"),
				},
			},
			want: &pb.ChapterResponse{
				ChapterId: "1",
				MangaId:   "1",
				Number:    1,
				Title:     "test",
				Pages:     []*pb.ChapterResponse_PageHead{{PageId: "1", Number: 1}},
			},
			wantErr: false,
		},
		{
			name: "Chapter not found",
			fields: fields{
				ChapterCase: ChapterCaseMock,
				PageCase:    PageCaseMock,
			},
			args: args{
				in0: context.Background(),
				req: &pb.UpdateChapterRequest{
					ChapterId: "5",
					Number:    utils.Ptr[float32](1),
					Title:     utils.Ptr("test"),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				ChapterCase:                tt.fields.ChapterCase,
				PageCase:                   tt.fields.PageCase,
				UnimplementedReadingServer: tt.fields.UnimplementedReadingServer,
			}
			got, err := s.UpdateChapter(tt.args.in0, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !proto.Equal(got, tt.want) {
				t.Errorf("UpdateChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
