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
  function isLastEntry(i: number) {
    return i === props.entries.length - 1;
  }

  return (
    <div className={styles.component}>
      {props.entries.map((e, i) => (
        <TerminalScrollIntoView
          key={e.id}
          doScroll={props.isAnimate && isLastEntry(i)}
        >
          <TerminalEntryComponent entry={e} />
          {
            // Terminal tooltip can be shown only at the bottom
            props.tooltip && isLastEntry(i) && (
              <TerminalTooltip
                markdownBody={props.tooltip.markdownBody}
                hidden={props.tooltip.hidden}
              />
            )
          }
        </TerminalScrollIntoView>
      ))}
    </div>
  );
}
