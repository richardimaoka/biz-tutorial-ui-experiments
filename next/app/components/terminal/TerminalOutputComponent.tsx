import { source_code_pro } from "../fonts/fonts";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalOutput_Fragment on TerminalOutput {
    output
  }
`);

interface TerminalOutputComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const TerminalOutputComponent = (
  props: TerminalOutputComponentProps
) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <pre className={styles.output}>
      <code
        // needs to specify font here, as <code> has its own font
        className={source_code_pro.className}
      >
        {fragment.output}
      </code>
    </pre>
  );
};
