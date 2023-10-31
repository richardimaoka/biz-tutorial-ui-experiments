import { GqlTerminalContents } from "./contents/GqlTerminalContents";
import { GqlTerminalHeader } from "./header/GqlTerminalHeader";
import styles from "./TerminalComponent.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlTerminalComponent on Terminal2 {
    ...GqlTerminalHeader
    ...GqlTerminalContents
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlTerminalComponent(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return (
    <div className={styles.component}>
      <div className={styles.header}>
        <GqlTerminalHeader fragment={fragment} />
      </div>
      <div className={styles.contents}>
        <GqlTerminalContents fragment={fragment} isAnimate={true} />
      </div>
    </div>
  );
}
