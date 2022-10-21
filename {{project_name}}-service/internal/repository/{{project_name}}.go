package repository

import (
	pb "github.com/salazarhugo/cheers1/genproto/cheers/chat/v1"
	"github.com/salazarhugo/cheers1/genproto/cheers/type/user"
	"github.com/salazarhugo/cheers1/services/chat-service/cache"
)

type {{project_name}}Repository interface {
	Create{{project_name}} () error
	Get{{project_name}} () error
	Update{{project_name}} () error
	Delete{{project_name}} () error
}

type {{project_name}}Repository struct {
	cache cache.RoomCache
}

func New{{project_name}}Repository(cache cache.RoomCache) {{project_name}}Repository {
	return &{{project_name}}Repository{
		cache: cache,
	}
}
