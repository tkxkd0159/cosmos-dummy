import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateGame } from "./types/checkers/checkers/tx";
import { MsgRejectGame } from "./types/checkers/checkers/tx";
import { MsgPlayMove } from "./types/checkers/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/checkers.checkers.MsgCreateGame", MsgCreateGame],
    ["/checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/checkers.checkers.MsgPlayMove", MsgPlayMove],
    
];

export { msgTypes }