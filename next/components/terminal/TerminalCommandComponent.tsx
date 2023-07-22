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
  scrollIntoView: () => void;
}

interface CodeComponentProps {
  command: string | null | undefined;
  scrollIntoView: () => void;
}

const TypeInCodeComponent = ({
  command,
  scrollIntoView,
}: CodeComponentProps) => {
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
    scrollIntoView();
  });
  return <code>&gt; {command?.substring(0, writtenLength)}</code>;
};

export const TerminalCommandComponent = (
  props: TerminalCommandComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalCommand_Fragment, props.fragment);
  const scrollIntoView = props.scrollIntoView;

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
      `}
    >
      {fragment.beforeExecution && animate ? (
        <TypeInCodeComponent
          command={fragment.command}
          scrollIntoView={scrollIntoView}
        />
      ) : (
        <code>&gt; {fragment.command}</code>
      )}
    </pre>
  );
};
