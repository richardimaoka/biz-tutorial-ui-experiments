import { CommandComponent } from "./CommandComponent";
import { OutputComponent } from "./OutputComponent";

interface Props {
  entry: TerminalEntry;
}

// Specify the return type JSX.Element for switch case's comprehensiveness check
export function TerminalEntryComponent(props: Props): JSX.Element {
  switch (props.entry.kind) {
    case "command":
      return <CommandComponent command={props.entry.command} />;
    case "output":
      return <OutputComponent output={props.entry.output} />;
  }
  return <div></div>;
}
