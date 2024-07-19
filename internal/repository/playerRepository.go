package repository

// make interface for player repository
type PlayerRepoInterface interface{
	PlayerAdd
	PlayerUpdate
	PlayerDelete
	PlayerView
}
type PlayerAdd interface{
	PlayerAdd()
}
type PlayerUpdate interface{
	PlayerUpdate()
}
type PlayerDelete interface{
	PlayerDelete()
}
type PlayerView interface{
	PlayerView()
}