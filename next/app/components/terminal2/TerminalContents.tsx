import { TerminalEntryComponent } from "./TerminalEntryComponent";

interface Props {
  entries: TerminalEntry[];
}

export function TerminalContents(props: Props) {
  return (
    <div>
      {props.entries.map((e) => (
        <TerminalEntryComponent key={e.id} entry={e} />
      ))}
    </div>
  );
}
