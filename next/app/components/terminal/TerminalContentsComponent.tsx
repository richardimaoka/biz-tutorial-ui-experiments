import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { TerminalNodeComponent } from "./TerminalNodeComponent";
import { TerminalScrollInto } from "./TerminalScrollInto";

const fragmentDefinition = graphql(`
  fragment TerminalContentsComponent_Fragment on Terminal {
    nodes {
      ...TerminalNodeComponent_Fragment
    }
  }
`);

interface TerminalContentsComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
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
      {nodes.map((node, index) =>
        index < nodes.length - 1 ? (
          <div key={index}>
            <TerminalNodeComponent
              fragment={node}
              skipAnimation={props.skipAnimation}
            />
          </div>
        ) : (
          <TerminalScrollInto key={index}>
            <TerminalNodeComponent
              fragment={node}
              skipAnimation={props.skipAnimation}
            />
          </TerminalScrollInto>
        )
      )}
    </div>
  );
};
