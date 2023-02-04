import { css } from "@emotion/react";
import { useRouter } from "next/router";
import { memo, useEffect, useState } from "react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";

const TerminalCommand_Fragment = graphql(`
  fragment TerminalCommand_Fragment on TerminalCommand {
    command
    beforeExecution
  }
`);

export interface TerminalCommandComponentProps {
  fragment: FragmentType<typeof TerminalCommand_Fragment>;
}

interface CodeComponentProps {
  command: string | null | undefined;
}

const TypeInCodeComponent = ({ command }: CodeComponentProps) => {
  const [writtenLength, setWrittenLength] = useState(0);

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
  return <code>{command?.substring(0, writtenLength)}</code>;
};

export const TerminalCommandComponent = (
  props: TerminalCommandComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalCommand_Fragment, props.fragment);

  const router = useRouter();
  const { skipAnimation } = router.query;
  const animate = skipAnimation !== "true";

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
      {fragment.beforeExecution && animate ? (
        <TypeInCodeComponent command={fragment.command} />
      ) : (
        <code>{fragment.command}</code>
      )}
    </pre>
  );
};
