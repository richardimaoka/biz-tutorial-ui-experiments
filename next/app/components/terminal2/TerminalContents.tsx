import { TerminalEntryComponent } from "./TerminalEntryComponent";
import styles from "./TerminalContents.module.css";
import {
  TerminalTooltip,
  TerminalTooltipProps,
} from "./tooltip/TerminalTooltip";
import { TerminalEntry } from "./definitions";

type Props = {
  entries: TerminalEntry[];
  tooltip?: TerminalTooltipProps;
};

export type TerminalContentsProps = Props;

export function TerminalContents(props: Props) {
  return (
    <div className={styles.component}>
      {props.entries.map((e) => (
        <TerminalEntryComponent key={e.id} entry={e} />
      ))}
      {
        // Terminal tooltip can be shown only at the bottom
        props.tooltip && (
          <TerminalTooltip
            markdownBody={props.tooltip.markdownBody}
            hidden={props.tooltip.hidden}
          />
        )
      }
    </div>
  );
}
