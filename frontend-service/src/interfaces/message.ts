import { AIResponse } from "./aiResp";

export interface Message {
  id: number;
  text: string | AIResponse;
  sender: "user" | "ai";
  timestamp: Date;
}