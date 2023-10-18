import { TerminalContents, TerminalContentsProps } from "./TerminalContents";
import { TerminalHeader, TerminalHeaderProps } from "./header/TerminalHeader";
import styles from "./TerminalComponent.module.css";

type Props = TerminalContentsProps & TerminalHeaderProps;

export type TerminalComponentProps = Props;

export function TerminalComponent(props: Props) {
  return (
    <div className={styles.component}>
      <div className={styles.header}>
        <TerminalHeader
          currentDirectory={props.currentDirectory}
          selectTab={props.selectTab}
          tabs={props.tabs}
        />
      </div>
      <div className={styles.contents}>
        <TerminalContents
          entries={props.entries}
          tooltip={props.tooltip}
          isAnimate={props.isAnimate}
        />
      </div>
    </div>
  );
}
