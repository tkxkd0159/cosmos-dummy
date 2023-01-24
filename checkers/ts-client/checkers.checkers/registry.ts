import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRejectGame } from "./types/checkers/checkers/tx";
import { MsgPlayMove } from "./types/checkers/checkers/tx";
import { MsgCreateGame } from "./types/checkers/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/checkers.checkers.MsgPlayMove", MsgPlayMove],
    ["/checkers.checkers.MsgCreateGame", MsgCreateGame],
    
];

export { msgTypes }