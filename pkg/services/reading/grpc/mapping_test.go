package grpc

import (
	pb "micromango/pkg/grpc/reading"
	"micromango/pkg/services/reading/entity"
	"reflect"
	"testing"
)

func TestChapterToHead(t *testing.T) {
	type args struct {
		c entity.ChapterHead
	}
	tests := []struct {
		name string
		args args
		want *pb.MangaContentResponse_ChapterHead
	}{
		{
			name: "Chapter to head",
			args: args{
				c: entity.ChapterHead{
					ChapterId: "1",
					Number:    1,
					Title:     "test",
					Pages:     1,
					Read:      false,
				},
			},
			want: &pb.MangaContentResponse_ChapterHead{
				ChapterId: "1",
				Number:    1,
				Title:     "test",
				Pages:     1,
				Read:      false,
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChapterToHead(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChapterToHead() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapterToPb(t *testing.T) {
	type args struct {
		c entity.Chapter
	}
	tests := []struct {
		name string
		args args
		want *pb.ChapterResponse
	}{
		{
			name: "Chapter to pb",
			args: args{
				c: entity.Chapter{
					MangaId:   "1",
					ChapterId: "1",
					Number:    1,
					Title:     "test",
					Pages: []entity.Page{{
						PageId:    "1",
						MangaId:   "1",
						ChapterId: "1",
						Number:    1,
						Image:     "1",
					}},
				},
			},
			want: &pb.ChapterResponse{
				ChapterId: "1",
				MangaId:   "1",
				Number:    1,
				Title:     "test",
				Pages: []*pb.ChapterResponse_PageHead{{
					PageId: "1",
					Number: 1,
					Image:  "1",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChapterToPb(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChapterToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageToPb(t *testing.T) {
	type args struct {
		p entity.Page
	}
	tests := []struct {
		name string
		args args
		want *pb.PageResponse
	}{
		{
			name: "Page to pb",
			args: args{
				p: entity.Page{
					PageId:    "1",
					MangaId:   "1",
					ChapterId: "1",
					Number:    1,
					Image:     "1",
				},
			},
			want: &pb.PageResponse{
				PageId:    "1",
				Number:    1,
				Image:     "1",
				ChapterId: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PageToPb(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PageToPb() = %v, want %v", got, tt.want)
			}
		})
	}
}
