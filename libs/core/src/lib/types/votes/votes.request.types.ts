import { IVotes } from "./votes.types";

export interface IGetVotesRequest extends Omit<IVotes, "votes"> {}

export interface IUpdateVotesRequest extends Omit<IVotes, "username"> {}
