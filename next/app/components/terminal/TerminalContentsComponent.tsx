import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { TerminalNodeComponent } from "./TerminalNodeComponent";

const fragmentDefinition = graphql(`
  fragment TerminalContentsComponent_Fragment on Terminal {
    nodes {
      ...TerminalNodeComponent_Fragment
    }
  }
`);

interface TerminalContentsComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const TerminalContentsComponent = (
  props: TerminalContentsComponentProps
) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.nodes) {
    return <></>;
  }
  const nodes = nonNullArray(fragment.nodes);

  return (
    <div className={styles.contents}>
      {nodes.map((node, index) => (
        <TerminalNodeComponent key={index} fragment={node} />
      ))}
    </div>
  );
};
