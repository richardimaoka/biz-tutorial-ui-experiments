import { css } from "@emotion/react";
import { useEffect, useState } from "react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const TerminalCommandWriting_Fragment = graphql(`
  fragment TerminalCommandWriting_Fragment on TerminalCommand {
    command
  }
`);

export interface TerminalCommandWritingComponentProps {
  fragment: FragmentType<typeof TerminalCommandWriting_Fragment>;
}

export const TerminalCommandWritingComponent = (
  props: TerminalCommandWritingComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalCommandWriting_Fragment, props.fragment);
  const command = fragment.command;
  const [writtenLength, setWrittenLength] = useState(0);

  //TODO: can avoid re-render with React.memo (+ useMemo hook?)
  useEffect(() => {
    if (command && writtenLength < command.length) {
      const incrementStep = command.length / 10;
      const nextLength = Math.min(
        writtenLength + incrementStep,
        command.length
      );
      setTimeout(() => {
        setWrittenLength(nextLength);
      }, 20);
    }
  });

  return (
    <pre
      css={css`
        margin: 1px 0px;
        padding: 4px;
        background-color: #1e1e1e;
        color: #f1f1f1;
        /* border-bottom: 1px solid #333333; */
      `}
    >
      <code>{command && command.substring(0, writtenLength)}</code>
    </pre>
  );
};
