import { CurrentDirectory } from "./CurrentDirectory";
import { TerminalContentsComponent } from "./TerminalContentsComponent";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalComponent_Fragment on Terminal {
    ...TerminalCurrentDirectory_Fragment
    ...TerminalContentsComponent_Fragment
  }
`);

interface TerminalComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
}

export const TerminalComponent = (props: TerminalComponentProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.terminal}>
      <CurrentDirectory fragment={fragment} />
      <TerminalContentsComponent
        fragment={fragment}
        skipAnimation={props.skipAnimation}
      />
    </div>
  );
};
