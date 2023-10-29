import styles from "./TerminalContents.module.css";
import { TerminalScrollIntoView } from "./TerminalScrollIntoView";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { TerminalTooltipGql } from "../tooltip/TerminalTooltipGql";
import { TerminalEntryComponentGql } from "./TerminalEntryComponentGql";

const fragmentDefinition = graphql(`
  fragment TerminalContentsGql on Terminal2 {
    nodes {
      id
      ...TerminalEntryComponentGql
    }
    tooltip {
      ...TerminalTooltipGql
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  isAnimate: boolean;
}

export function TerminalContentsGql(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  function isLastEntry(i: number) {
    return i === fragment.nodes.length - 1;
  }

  return (
    <div className={styles.component}>
      {fragment.nodes.map((n, i) => (
        <TerminalScrollIntoView
          key={n.id}
          doScroll={props.isAnimate && isLastEntry(i)}
        >
          <TerminalEntryComponentGql
            fragment={n}
            animate={props.isAnimate && isLastEntry(i)}
          />
          {
            // Terminal tooltip can be shown only at the bottom
            fragment.tooltip && isLastEntry(i) && (
              <TerminalTooltipGql
                fragment={fragment.tooltip}
                // hidden={props.tooltip.hidden}
              />
            )
          }
        </TerminalScrollIntoView>
      ))}
    </div>
  );
}
