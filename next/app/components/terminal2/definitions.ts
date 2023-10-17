type TerminalEntryCommand = {
  kind: "command";
  id: string;
  command: string;
  isExecuted: boolean;
};

type TerminalEntryOutput = {
  kind: "output";
  id: string;
  output: string;
};

type TerminalEntry = TerminalEntryCommand | TerminalEntryOutput;
