import { css } from "@emotion/react";

interface CommandProps {
  command: string;
  writtenLength: number;
}

export const Command = ({
  command,
  writtenLength,
}: CommandProps): JSX.Element => {
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
