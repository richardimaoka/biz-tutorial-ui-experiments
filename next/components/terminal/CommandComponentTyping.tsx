import { css } from "@emotion/react";
import { useEffect, useState } from "react";

interface CommandComponentTypingProps {
  command: string;
}

export const CommandComponentTyping = ({
  command,
}: CommandComponentTypingProps): JSX.Element => {
  const [writtenLength, setWrittenLength] = useState(0);
  useEffect(() => {
    const nextWrittenLength = Math.min(writtenLength + 6, command.length);
    if (nextWrittenLength > writtenLength) {
      setTimeout(() => {
        setWrittenLength(nextWrittenLength);
      }, 30);
    }
  });

  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #1e1e1e;
    color: #f1f1f1;
    border-bottom: 1px solid #333333;
  `;
  return (
    <pre css={cssTerminalCommand}>
      <code>{command.substring(0, writtenLength)}</code>
    </pre>
  );
};
