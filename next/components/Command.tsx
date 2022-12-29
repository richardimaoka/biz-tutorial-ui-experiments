import { css } from "@emotion/react";
import { useEffect, useState } from "react";

interface CommandProps {
  command: string;
}

export const Command = ({ command }: CommandProps): JSX.Element => {
  const [writtenLength, setWrittenLength] = useState(0);
  useEffect(() => {
    if (writtenLength < command.length) {
      setTimeout(() => {
        setWrittenLength(writtenLength + 6);
      }, 30);
    }
  });

  const cssTerminalCommand = css`
    margin: 1px 0px;
    padding: 4px;
    background-color: #3a3a3a;
    color: white;
    border-bottom: 1px solid white;
  `;
  return (
    <pre css={cssTerminalCommand}>
      <code>{command.substring(0, writtenLength)}</code>
    </pre>
  );
};
