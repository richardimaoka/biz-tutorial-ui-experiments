import styles from "./GqlTerminalContents.module.css";
import { TerminalScrollIntoView } from "./TerminalScrollIntoView";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlTerminalTooltip } from "../tooltip/GqlTerminalTooltip";
import { GqlTerminalEntryComponent } from "../entry/GqlTerminalEntryComponent";

const fragmentDefinition = graphql(`
  fragment GqlTerminalContents on Terminal {
    entries {
      id
      ...GqlTerminalEntryComponent
    }
    tooltip {
      ...GqlTerminalTooltip
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTerminalContents(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  function isLastEntry(i: number) {
    return i === fragment.entries.length - 1;
  }

  return (
    <div className={styles.component}>
      {fragment.entries.map((n, i) => (
        <TerminalScrollIntoView key={n.id} doScroll={isLastEntry(i)}>
          <GqlTerminalEntryComponent
            fragment={n}
            isLastEntry={isLastEntry(i)}
          />
          {
            // Terminal tooltip can be shown only at the bottom
            fragment.tooltip && isLastEntry(i) && (
              <GqlTerminalTooltip
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
