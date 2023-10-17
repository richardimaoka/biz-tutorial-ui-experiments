import { TerminalEntryComponent } from "./TerminalEntryComponent";
import styles from "./TerminalContents.module.css";

type Props = {
  entries: TerminalEntry[];
};

export type TerminalContentsProps = Props;

export function TerminalContents(props: Props) {
  return (
    <div className={styles.component}>
      {props.entries.map((e) => (
        <TerminalEntryComponent key={e.id} entry={e} />
      ))}
    </div>
  );
}
