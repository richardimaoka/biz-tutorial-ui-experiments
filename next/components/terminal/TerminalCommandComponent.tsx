import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const TerminalCommand_Fragment = graphql(`
  fragment TerminalCommand_Fragment on TerminalCommand {
    command
  }
`);

export interface TerminalCommandComponentProps {
  fragment: FragmentType<typeof TerminalCommand_Fragment>;
}

interface CodeComponentProps {
  command: string | null | undefined;
}

const CodeComponent = ({ command }: CodeComponentProps) => {
  return <code>{command}</code>;
};

export const TerminalCommandComponent = (
  props: TerminalCommandComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalCommand_Fragment, props.fragment);

  return (
    <pre
      css={css`
        margin: 1px 0px;
        padding: 4px;
        background-color: #1e1e1e;
        color: #f1f1f1;
        border-bottom: 1px solid #333333;
      `}
    >
      <CodeComponent command={fragment.command} />
    </pre>
  );
};
