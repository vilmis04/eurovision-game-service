import { ICountry } from "../country";
import { VoteTypes } from "../votes";

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

export interface IFinalsList extends Pick<ICountry, "name"> {
	position: VoteTypes;
}

export type TUserVoteResponse = {
	member: string;
	semi1: string[];
	semi2: string[];
	finals: IFinalsList[];
};
