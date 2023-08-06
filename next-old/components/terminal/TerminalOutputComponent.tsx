import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const TerminalOutput_Fragment = graphql(`
  fragment TerminalOutput_Fragment on TerminalOutput {
    output
  }
`);

export interface TerminalOutputComponentProps {
  fragment: FragmentType<typeof TerminalOutput_Fragment>;
}

export const TerminalOutputComponent = (
  props: TerminalOutputComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalOutput_Fragment, props.fragment);

  return (
    <pre
      css={css`
        margin: 1px 0px;
        padding: 4px;
        background-color: #1e1e1e;
        color: #979797;
      `}
    >
      <code>{fragment.output}</code>
    </pre>
  );
};
