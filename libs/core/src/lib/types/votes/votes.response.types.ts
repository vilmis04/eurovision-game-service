import { CountryTypes } from "../country.types";
import { VoteTypes } from "./votes.types";

export interface IGetVotesResponse {
	username?: string;
	[CountryTypes.ARMENIA]: VoteTypes | null;
	[CountryTypes.AUSTRALIA]: VoteTypes | null;
	[CountryTypes.AZERBAIJAN]: VoteTypes | null;
	[CountryTypes.BELGIUM]: VoteTypes | null;
	[CountryTypes.CZECH_REPUBLIC]: VoteTypes | null;
	[CountryTypes.ESTONIA]: VoteTypes | null;
	[CountryTypes.FINLAND]: VoteTypes | null;
	[CountryTypes.FRANCE]: VoteTypes | null;
	[CountryTypes.GERMANY]: VoteTypes | null;
	[CountryTypes.GREECE]: VoteTypes | null;
	[CountryTypes.ICELAND]: VoteTypes | null;
	[CountryTypes.ITALY]: VoteTypes | null;
	[CountryTypes.LITHUANIA]: VoteTypes | null;
	[CountryTypes.MOLDOVA]: VoteTypes | null;
	[CountryTypes.NETHERLANDS]: VoteTypes | null;
	[CountryTypes.NORWAY]: VoteTypes | null;
	[CountryTypes.POLAND]: VoteTypes | null;
	[CountryTypes.PORTUGAL]: VoteTypes | null;
	[CountryTypes.ROMANIA]: VoteTypes | null;
	[CountryTypes.SERBIA]: VoteTypes | null;
	[CountryTypes.SPAIN]: VoteTypes | null;
	[CountryTypes.SWEDEN]: VoteTypes | null;
	[CountryTypes.SWITZERLAND]: VoteTypes | null;
	[CountryTypes.UKRAINE]: VoteTypes | null;
	[CountryTypes.UNITED_KINGDOM]: VoteTypes | null;
}
