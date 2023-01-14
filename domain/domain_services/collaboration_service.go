package domain_services

type CollaborationService struct {
}

func NewCollaborationService() *CollaborationService {
	return &CollaborationService{}
}

// wire

var singletonCollaborationService *CollaborationService = initSingletonCollaborationService()

func GetSingletonCollaborationService() *CollaborationService {
	return singletonCollaborationService
}

func initSingletonCollaborationService() *CollaborationService {
	return NewCollaborationService()
}

// methods
