import { GqlTerminalContents } from "./contents/GqlTerminalContents";
import { GqlTerminalHeader } from "./header/GqlTerminalHeader";
import styles from "./TerminalComponent.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlTerminalColumn on TerminalColumn2 {
    ...GqlTerminalHeader
    terminals {
      ...GqlTerminalContents
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  selectIndex: number;
}

export function GqlTerminalColumn(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const selectTerminal = fragment.terminals[props.selectIndex];

  return (
    <div className={styles.component}>
      <div className={styles.header}>
        <GqlTerminalHeader
          fragment={fragment}
          selectIndex={props.selectIndex}
        />
      </div>
      <div className={styles.contents}>
        <GqlTerminalContents fragment={selectTerminal} isAnimate={true} />
      </div>
    </div>
  );
}
