package observer


type newsListener interface{
	Update (news string)
}