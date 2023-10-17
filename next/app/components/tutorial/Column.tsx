import { TerminalComponent } from "../terminal2/TerminalComponent";
import styles from "./Column.module.css";
import { TutorialColumnProps } from "./definitions";

interface Props {
  // Don't pass in children. Use TutorialColumnProps instead of children, as we want to restrict what can be inside a Column.
  column: TutorialColumnProps;
}

// Explicit return type for switch statement comprehensiveness check
export function Column(props: Props): JSX.Element {
  switch (props.column.kind) {
    case "Terminal":
      const terminal = props.column;
      return (
        <div className={styles.component}>
          <TerminalComponent
            currentDirectory={terminal.currentDirectory}
            entries={terminal.entries}
            selectTab={terminal.selectTab}
            tabs={terminal.tabs}
          />
        </div>
      );
  }
}
