export const initialValues = {
	name: "",
	_id: "",
};

export interface ICreateGroupFormData {
	name: string;
	_id: string;
}

export interface IGroup {
	name: string;
	members: string[];
	owner: string;
	yearCreated: string;
}
