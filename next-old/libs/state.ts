export interface TerminalCommand {
  command: String;
  commandWrittenLength: number;
  state: "writing" | "ready to run";
  extraAction: String;
}

export interface TerminalOutput {
  command: String;
  // writing: auto transition to "ready to run"
  // ready to run: upon user action, transition to next state
  state: "writing" | "ready to run";
}

export interface TerminalState {
  elements: (TerminalCommand | TerminalOutput)[];
  stepAt: number;
}
