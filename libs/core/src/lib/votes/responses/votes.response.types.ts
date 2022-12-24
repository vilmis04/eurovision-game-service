import { CountryTypes } from "../../types/country.types";
import { VoteTypes } from "../../types/vote.types";

export interface IGetVotesResponse {
	// _id: string;
	// username: string;
	[CountryTypes.ARMENIA]: VoteTypes;
	[CountryTypes.AUSTRALIA]: VoteTypes;
	[CountryTypes.AZERBAIJAN]: VoteTypes;
	[CountryTypes.BELGIUM]: VoteTypes;
	[CountryTypes.CZECH_REPUBLIC]: VoteTypes;
	[CountryTypes.ESTONIA]: VoteTypes;
	[CountryTypes.FINLAND]: VoteTypes;
	[CountryTypes.FRANCE]: VoteTypes;
	[CountryTypes.GERMANY]: VoteTypes;
	[CountryTypes.GREECE]: VoteTypes;
	[CountryTypes.ICELAND]: VoteTypes;
	[CountryTypes.ITALY]: VoteTypes;
	[CountryTypes.LITHUANIA]: VoteTypes;
	[CountryTypes.MOLDOVA]: VoteTypes;
	[CountryTypes.NETHERLANDS]: VoteTypes;
	[CountryTypes.NORWAY]: VoteTypes;
	[CountryTypes.POLAND]: VoteTypes;
	[CountryTypes.PORTUGAL]: VoteTypes;
	[CountryTypes.ROMANIA]: VoteTypes;
	[CountryTypes.SERBIA]: VoteTypes;
	[CountryTypes.SPAIN]: VoteTypes;
	[CountryTypes.SWEDEN]: VoteTypes;
	[CountryTypes.SWITZERLAND]: VoteTypes;
	[CountryTypes.UKRAINE]: VoteTypes;
	[CountryTypes.UNITED_KINGDOM]: VoteTypes;
}
