import { TerminalEntryComponent } from "./TerminalEntryComponent";
import styles from "./TerminalContents.module.css";
import {
  TerminalTooltip,
  TerminalTooltipProps,
} from "./tooltip/TerminalTooltip";
import { TerminalEntry } from "./definitions";
import { TerminalScrollIntoView } from "./TerminalScrollIntoView";

type Props = {
  entries: TerminalEntry[];
  tooltip?: TerminalTooltipProps;
  isAnimate?: boolean;
};

export type TerminalContentsProps = Props;

export function TerminalContents(props: Props) {
  console.log("TerminalContents:", props.isAnimate);
  return (
    <div className={styles.component}>
      {props.entries.map((e, i) => (
        <TerminalScrollIntoView
          key={e.id}
          doScroll={props.isAnimate && i === props.entries.length - 1}
        >
          <TerminalEntryComponent entry={e} />
        </TerminalScrollIntoView>
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
