package posts

import "github.com/jchavannes/jgo/web"

func GetRoutes() []web.Route {
	return []web.Route{
		newRoute,
		topRoute,
		archiveRoute,
	}
}