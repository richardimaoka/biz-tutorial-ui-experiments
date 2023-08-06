import { source_code_pro } from "../fonts/fonts";
import { FlickeringTrail } from "./FlickeringTrail";
import { TerminalPrompt } from "./TerminalPrompt";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalCommand_Fragment on TerminalCommand {
    command
    beforeExecution
  }
`);

export interface TerminalCommandComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const TerminalCommandComponent = (
  props: TerminalCommandComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <pre className={styles.command}>
      <TerminalPrompt />
      <code
        // needs to specify font here, as <code> has its own font
        className={source_code_pro.className}
      >
        {fragment.command}
      </code>
      {fragment.beforeExecution && <FlickeringTrail />}
    </pre>
  );
};
