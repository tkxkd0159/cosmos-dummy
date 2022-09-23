import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgEndGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jsc.checkers.MsgEndGame", MsgEndGame],
    ["/jsc.checkers.MsgPlayMove", MsgPlayMove],
    ["/jsc.checkers.MsgRejectGame", MsgRejectGame],
    ["/jsc.checkers.MsgCreateGame", MsgCreateGame],
    
];

export { msgTypes }