import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgEndGame } from "./types/checkers/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/jsc.checkers.MsgCreateGame", MsgCreateGame],
    ["/jsc.checkers.MsgEndGame", MsgEndGame],
    
];

export { msgTypes }