package handlers

import (
	"SocialNetworkRestApi/api/internal/server/websocket"
	"SocialNetworkRestApi/api/pkg/models"
	"SocialNetworkRestApi/api/pkg/services"
	"log"
)

type Application struct {
	Logger              *log.Logger
	WS                  *websocket.WebsocketServer
	UserService         services.IUserService
	NotificationService services.INotificationService
	PostService         services.IPostService
	CommentService      services.ICommentService
	ChatService         services.IChatService
	GroupService        services.IGroupService
	GroupMemberService  services.IGroupMemberService
	GroupEventService   services.IGroupEventService
}

func InitApp(repositories *models.Repositories, logger *log.Logger) *Application {

	userServices := services.InitUserService(
		logger,
		repositories.UserRepo,
		repositories.SessionRepo,
		repositories.FollowerRepo,
		repositories.NotificationRepo,
	)

	notificationServices := services.InitNotificationService(
		logger,
		repositories.UserRepo,
		repositories.FollowerRepo,
		repositories.NotificationRepo,
		repositories.GroupRepo,
		repositories.GroupMemberRepo,
	)

	chatServices := services.InitChatService(
		logger,
		repositories.UserRepo,
		repositories.MessageRepo,
		repositories.GroupRepo,
	)

	groupEventServices := services.InitGroupEventService(
		logger,
		repositories.GroupEventAttendance,
		repositories.EventRepo,
		repositories.GroupRepo,
		repositories.GroupMemberRepo,
		repositories.UserRepo,
		repositories.NotificationRepo,
	)

	return &Application{
		Logger: logger,
		WS: websocket.InitWebsocket(
			logger,
			userServices,
			notificationServices,
			chatServices,
			services.InitGroupService(
				logger,
				repositories.GroupRepo,
				repositories.GroupMemberRepo,
				repositories.UserRepo,
			),
			services.InitGroupMemberService(logger, repositories.GroupMemberRepo),
		),
		UserService:         userServices,
		NotificationService: notificationServices,
		PostService:         services.InitPostService(logger, repositories.PostRepo, repositories.AllowedPostRepo),
		CommentService:      services.InitCommentService(logger, repositories.CommentRepo),
		ChatService:         chatServices,
		GroupService: services.InitGroupService(
			logger,
			repositories.GroupRepo,
			repositories.GroupMemberRepo,
			repositories.UserRepo,
		),
		GroupMemberService: services.InitGroupMemberService(logger, repositories.GroupMemberRepo),
		GroupEventService:  groupEventServices,
	}
}
