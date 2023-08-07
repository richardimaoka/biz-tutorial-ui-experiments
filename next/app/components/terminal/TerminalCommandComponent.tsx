import { TerminalCommandTypeIn } from "./TerminalCommandTypeIn";
import { TerminalPrompt } from "./TerminalPrompt";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalCommand_Fragment on TerminalCommand {
    command
    beforeExecution
  }
`);

interface TerminalCommandStaticProps {
  command: string;
}

const TerminalCommandStatic = ({ command }: TerminalCommandStaticProps) => (
  <code>{command}</code>
);

interface TerminalCommandComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  skipAnimation?: boolean;
}

export const TerminalCommandComponent = (
  props: TerminalCommandComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.command) {
    return <></>;
  }

  return (
    <pre className={styles.command}>
      <TerminalPrompt />
      {fragment.beforeExecution && !props.skipAnimation ? (
        <TerminalCommandTypeIn command={fragment.command} />
      ) : (
        <TerminalCommandStatic command={fragment.command} />
      )}
    </pre>
  );
};
