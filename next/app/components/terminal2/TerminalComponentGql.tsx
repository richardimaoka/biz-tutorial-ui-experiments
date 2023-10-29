import { TerminalContentsGql } from "./contents/TerminalContentsGql";
import { TerminalHeaderGql } from "./header/TerminalHeaderGql";
import styles from "./TerminalComponent.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalComponentGql on Terminal2 {
    ...TerminalHeaderGql
    ...TerminalContentsGql
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function TerminalComponentGql(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return (
    <div className={styles.component}>
      <div className={styles.header}>
        <TerminalHeaderGql fragment={fragment} />
      </div>
      <div className={styles.contents}>
        <TerminalContentsGql fragment={fragment} isAnimate={true} />
      </div>
    </div>
  );
}
