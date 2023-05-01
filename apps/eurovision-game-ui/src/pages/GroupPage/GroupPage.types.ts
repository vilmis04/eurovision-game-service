export enum SubmitTypes {
	CREATE = "CREATE",
	UPDATE = "UPDATE",
	DELETE = "DELETE",
	INVITE = "INVITE",
	LEAVE = "LEAVE",
}

export interface IGroupForAction {
	name: string;
	_id: string;
}
